package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

func FetchInstances(ctx context.Context, params *aws.AwsFetchInput) ([]*Instance, *aws.AwsFetchOutputMetadata) {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "ec2",
		Resource:   "vpcs",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime,
	}
	resources := []*Instance{}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.EC2()

	paginator := ec2.NewDescribeInstancesPaginator(client, &ec2.DescribeInstancesInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeInstances in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, reservation := range output.Reservations {
			for _, object := range reservation.Instances {

				model := new(Instance)
				copier.Copy(&model, &object)

				model.Tags = ConvertTags(object.Tags)
				model.AccountId = params.AccountId
				model.Region = params.Region
				model.ReportTime = params.ReportTime
				model.InventoryUUID = uuid.New().String()

				resources = append(resources, model)
				fetchedResources++
			}
		}

	}

	inventoryResults.FetchedResources = fetchedResources
	inventoryResults.FailedResources = failedResources
	inventoryResults.HadErrors = len(fetchingErrors) > 0

	return resources, &aws.AwsFetchOutputMetadata{
		FetchingErrors:   fetchingErrors,
		InventoryResults: inventoryResults,
	}
}
