//AUTOGENERATED CODE DO NOT EDIT
package ecs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

func FetchTasks(ctx context.Context, params *localAws.AwsFetchInput) ([]*Task, *localAws.AwsFetchOutputMetadata) {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "ecs",
		Resource:   "tasks",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime,
	}
	resources := []*Task{}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.ECS()

	paginator := ecs.NewListClustersPaginator(client, &ecs.ListClustersInput{
		MaxResults: aws.Int32(100),
	})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling ListClusters in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		result, err := client.DescribeClusters(ctx, &ecs.DescribeClustersInput{
			Clusters: output.ClusterArns,
			Include:  []types.ClusterField{types.ClusterFieldAttachments, types.ClusterFieldConfigurations, types.ClusterFieldSettings, types.ClusterFieldStatistics, types.ClusterFieldTags},
		})
		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeClusters in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, cluster := range result.Clusters {
			paginator := ecs.NewListTasksPaginator(client, &ecs.ListTasksInput{
				MaxResults: aws.Int32(10),
				Cluster:    cluster.ClusterName,
			})

			for paginator.HasMorePages() {
				output, err := paginator.NextPage(ctx)
				if err != nil {
					fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling ListTasks on %s in %s/%s: %w", *cluster.ClusterArn, params.AccountId, params.Region, err))
					break
				}

				if len(output.TaskArns) == 0 {
					continue
				}

				result, err := client.DescribeTasks(ctx, &ecs.DescribeTasksInput{
					Tasks:   output.TaskArns,
					Cluster: cluster.ClusterName,
					Include: []types.TaskField{
						types.TaskFieldTags,
					},
				})
				if err != nil {
					fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeTasks on %s in %s/%s: %w", *cluster.ClusterArn, params.AccountId, params.Region, err))
					break
				}
				for _, task := range result.Tasks {
					model := new(Task)
					copier.CopyWithOption(&model, &task, localAws.CopyOption)

					model.Tags = ConvertTags(task.Tags)
					model.AccountId = params.AccountId
					model.Region = params.Region
					model.ReportTime = params.ReportTime
					model.InventoryUUID = uuid.New().String()

					resources = append(resources, model)
					fetchedResources++
				}

			}
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
