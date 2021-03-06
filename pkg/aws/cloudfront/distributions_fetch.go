//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_fetching_file.tmpl
package cloudfront

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

func FetchDistributions(ctx context.Context, params *aws.AwsFetchInput) ([]*Distribution, *aws.AwsFetchOutputMetadata) {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "cloudfront",
		Resource:   "distributions",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime,
	}
	resources := []*Distribution{}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.CloudFront()

	paginator := cloudfront.NewListDistributionsPaginator(client, &cloudfront.ListDistributionsInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling ListDistributions in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.DistributionList.Items {

			result, err := client.GetDistribution(ctx, &cloudfront.GetDistributionInput{
				Id: object.Id,
			})
			if err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling GetDistribution in %s/%s: %w", params.AccountId, params.Region, err))
				failedResources++
				continue
			}

			model := new(Distribution)
			copier.CopyWithOption(&model, &result.Distribution, aws.CopyOption)

			tagsResult, err := client.ListTagsForResource(ctx, &cloudfront.ListTagsForResourceInput{
				Resource: result.Distribution.ARN,
			})
			if err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling ListTagsForResource in %s/%s: %w", params.AccountId, params.Region, err))
				failedResources++
				continue
			}

			model.Tags = ConvertTags(tagsResult.Tags.Items)
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
