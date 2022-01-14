//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

func FetchInstance(ctx context.Context, params *awscloud.AwsFetchInput) *awscloud.AwsFetchOutput {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "ec2",
		Resource:   "vpcs",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime.UTC().UnixMilli(),
	}

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
				model.ReportTime = params.ReportTime.UTC().UnixMilli()

				if err = PostProcessInstance(ctx, params, model); err != nil {
					fetchingErrors = append(fetchingErrors, fmt.Errorf("error post-processing Instance %s %s/%s: %w", model.InstanceId, params.AccountId, params.Region, err))
					failedResources++
				}

				err = params.OutputFile.Write(ctx, model)
				if err != nil {
					fetchingErrors = append(fetchingErrors, fmt.Errorf("error storing Instance model in %s/%s: %w", params.AccountId, params.Region, err))
				}
				fetchedResources++
			}
		}

	}

	inventoryResults.FetchedResources = fetchedResources
	inventoryResults.FailedResources = failedResources
	inventoryResults.HadErrors = len(fetchingErrors) > 0

	return &awscloud.AwsFetchOutput{
		FetchingErrors:   fetchingErrors,
		InventoryResults: inventoryResults,
		ResourceName:     "Instances",
		AccountId:        params.AccountId,
		Region:           params.Region,
	}
}