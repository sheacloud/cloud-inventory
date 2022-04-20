//AUTOGENERATED CODE DO NOT EDIT
package sqs

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/google/uuid"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

func FetchQueues(ctx context.Context, params *localAws.AwsFetchInput) ([]*Queue, *localAws.AwsFetchOutputMetadata) {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "sqs",
		Resource:   "queues",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime,
	}
	resources := []*Queue{}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.SQS()

	paginator := sqs.NewListQueuesPaginator(client, &sqs.ListQueuesInput{
		MaxResults: aws.Int32(100),
	})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling ListGateways in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, queueURL := range output.QueueUrls {
			result, err := client.GetQueueAttributes(ctx, &sqs.GetQueueAttributesInput{
				QueueUrl: aws.String(queueURL),
				AttributeNames: []types.QueueAttributeName{
					types.QueueAttributeNameAll,
				},
			})
			if err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling GetQueueAttributes on %s in %s/%s: %w", queueURL, params.AccountId, params.Region, err))
				continue
			}
			model := new(Queue)
			model.QueueURL = queueURL
			for key, value := range result.Attributes {
				err = nil
				switch key {
				case "Policy":
					model.Policy = value
				case "VisibilityTimeout":
					model.VisibilityTimeout, err = strconv.ParseInt(value, 10, 64)
				case "MaximumMessageSize":
					model.MaximumMessageSize, err = strconv.ParseInt(value, 10, 64)
				case "MessageRetentionPeriod":
					model.MessageRetentionPeriod, err = strconv.ParseInt(value, 10, 64)
				case "ApproximateNumberOfMessages":
					model.ApproximateNumberOfMessages, err = strconv.ParseInt(value, 10, 64)
				case "ApproximateNumberOfMessagesNotVisible":
					model.ApproximateNumberOfMessagesNotVisible, err = strconv.ParseInt(value, 10, 64)
				case "CreatedTimestamp":
					model.CreatedTimestamp = value
				case "LastModifiedTimestamp":
					model.LastModifiedTimestamp = value
				case "QueueArn":
					model.QueueArn = value
				case "ApproximateNumberOfMessagesDelayed":
					model.ApproximateNumberOfMessagesDelayed, err = strconv.ParseInt(value, 10, 64)
				case "DelaySeconds":
					model.DelaySeconds, err = strconv.ParseInt(value, 10, 64)
				case "ReceiveMessageWaitTimeSeconds":
					model.ReceiveMessageWaitTimeSeconds, err = strconv.ParseInt(value, 10, 64)
				case "RedrivePolicy":
					model.RedrivePolicy = value
				case "FifoQueue":
					if value == "true" {
						model.FifoQueue = true
					}
				case "ContentBasedDeduplication":
					if value == "true" {
						model.ContentBasedDeduplication = true
					}
				case "KmsMasterKeyId":
					model.KmsMasterKeyId = value
				case "KmsDataKeyReusePeriodSeconds":
					model.KmsDataKeyReusePeriodSeconds, err = strconv.ParseInt(value, 10, 64)
				case "DeduplicationScope":
					model.DeduplicationScope = value
				case "FifoThroughputLimit":
					model.FifoThroughputLimit = value
				}
				if err != nil {
					continue
				}
			}

			tagsResult, err := client.ListQueueTags(ctx, &sqs.ListQueueTagsInput{
				QueueUrl: aws.String(model.QueueURL),
			})
			if err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling ListQueueTags on %s in %s/%s: %w", queueURL, params.AccountId, params.Region, err))
				continue
			}
			model.Tags = tagsResult.Tags

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

	return resources, &localAws.AwsFetchOutputMetadata{
		FetchingErrors:   fetchingErrors,
		InventoryResults: inventoryResults,
	}
}
