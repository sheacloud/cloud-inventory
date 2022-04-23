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

func FetchImages(ctx context.Context, params *aws.AwsFetchInput) ([]*Image, *aws.AwsFetchOutputMetadata) {
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
	resources := []*Image{}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.EC2()

	result, err := client.DescribeImages(ctx, &ec2.DescribeImagesInput{
		Owners: []string{"self"},
	})
	if err != nil {
		fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeImages in %s/%s: %w", params.AccountId, params.Region, err))
		inventoryResults.FetchedResources = 0
		inventoryResults.FailedResources = 0
		inventoryResults.HadErrors = true
		return nil, &aws.AwsFetchOutputMetadata{
			FetchingErrors:   fetchingErrors,
			InventoryResults: inventoryResults,
		}
	}

	results := []*ec2.DescribeImagesOutput{result}
	for _, output := range results {

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeImages in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.Images {

			model := new(Image)
			copier.CopyWithOption(&model, &object, aws.CopyOption)

			model.Tags = ConvertTags(object.Tags)
			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime
			model.InventoryUUID = uuid.New().String()

			resources = append(resources, model)
			fetchedResources++
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
