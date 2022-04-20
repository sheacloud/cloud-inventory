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

func FetchClusters(ctx context.Context, params *localAws.AwsFetchInput) ([]*Cluster, *localAws.AwsFetchOutputMetadata) {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "ecs",
		Resource:   "clusters",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime,
	}
	resources := []*Cluster{}

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

			model := new(Cluster)
			copier.Copy(&model, &cluster)

			model.Tags = ConvertTags(cluster.Tags)
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