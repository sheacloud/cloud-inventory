//AUTOGENERATED CODE DO NOT EDIT
package sns

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

func FetchTopics(ctx context.Context, params *aws.AwsFetchInput) ([]*Topic, *aws.AwsFetchOutputMetadata) {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "sns",
		Resource:   "topics",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime,
	}
	resources := []*Topic{}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.SNS()

	paginator := sns.NewListTopicsPaginator(client, &sns.ListTopicsInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling ListTopics in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.Topics {

			model := new(Topic)
			copier.CopyWithOption(&model, &object, aws.CopyOption)

			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime
			model.InventoryUUID = uuid.New().String()

			if err = PostProcessTopic(ctx, params, model); err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error post-processing Topic %s %s/%s: %w", model.TopicArn, params.AccountId, params.Region, err))
				failedResources++
			}

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
