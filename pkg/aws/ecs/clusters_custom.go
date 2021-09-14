package ecs

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

func ClusterDataSource(ctx context.Context, client *ecs.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(ClusterModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := ecs.NewListClustersPaginator(client, &ecs.ListClustersInput{
		MaxResults: aws.Int32(100),
	})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":     storageConfig.Service,
				"data_source": storageConfig.DataSource,
				"account_id":  storageConfig.AccountId,
				"region":      storageConfig.Region,
				"cloud":       storageConfig.Cloud,
				"error":       err,
			}).Error("error calling ListClusters")
			return err
		}

		result, err := client.DescribeClusters(ctx, &ecs.DescribeClustersInput{
			Clusters: output.ClusterArns,
			Include:  []types.ClusterField{types.ClusterFieldAttachments, types.ClusterFieldConfigurations, types.ClusterFieldSettings, types.ClusterFieldStatistics, types.ClusterFieldTags},
		})
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":     storageConfig.Service,
				"data_source": storageConfig.DataSource,
				"account_id":  storageConfig.AccountId,
				"region":      storageConfig.Region,
				"cloud":       storageConfig.Cloud,
				"error":       err,
			}).Error("error calling DescribeClusters")
			return err
		}
		for _, cluster := range result.Clusters {

			model := new(ClusterModel)
			copier.Copy(&model, &cluster)

			model.Tags = GetTagMap(cluster.Tags)
			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region
			model.ReportTime = reportTime.UTC().UnixMilli()

			for _, f := range customClusterModelPostprocessingFuncs {
				f(ctx, client, cfg, model)
			}

			errors := storageContextSet.Store(ctx, model)
			for storageContext, err := range errors {
				storage.LogContextError(storageContext, fmt.Sprintf("Error storing ClusterModel: %v", err))
			}
		}

	}

	return nil
}
