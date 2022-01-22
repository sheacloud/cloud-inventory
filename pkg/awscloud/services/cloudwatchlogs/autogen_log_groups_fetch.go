//AUTOGENERATED CODE DO NOT EDIT
package cloudwatchlogs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

func FetchLogGroup(ctx context.Context, params *awscloud.AwsFetchInput) *awscloud.AwsFetchOutput {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "cloudwatchlogs",
		Resource:   "log_groups",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime.UTC().UnixMilli(),
	}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.CloudWatchLogs()

	paginator := cloudwatchlogs.NewDescribeLogGroupsPaginator(client, &cloudwatchlogs.DescribeLogGroupsInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeLogGroups in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.LogGroups {

			model := new(LogGroup)
			copier.Copy(&model, &object)

			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime.UTC().UnixMilli()

			err = params.OutputFile.Write(ctx, model)
			if err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error storing LogGroup model in %s/%s: %w", params.AccountId, params.Region, err))
			}
			fetchedResources++
		}

	}

	inventoryResults.FetchedResources = fetchedResources
	inventoryResults.FailedResources = failedResources
	inventoryResults.HadErrors = len(fetchingErrors) > 0

	return &awscloud.AwsFetchOutput{
		FetchingErrors:   fetchingErrors,
		InventoryResults: inventoryResults,
		ResourceName:     "log_groups",
		AccountId:        params.AccountId,
		Region:           params.Region,
	}
}
