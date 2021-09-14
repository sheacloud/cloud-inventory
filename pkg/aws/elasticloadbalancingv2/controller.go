package elasticloadbalancingv2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

var (
	Controller = ElasticLoadBalancingV2Controller{
		DataSources: map[string]func(ctx context.Context, client *elasticloadbalancingv2.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error{},
	}
)

type ElasticLoadBalancingV2Controller struct {
	DataSources map[string]func(ctx context.Context, client *elasticloadbalancingv2.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error
}

func (e *ElasticLoadBalancingV2Controller) GetRegionOverrides() []string {
	return []string{}
}

func (e *ElasticLoadBalancingV2Controller) GetName() string {
	return "elasticloadbalancingv2"
}

func (e *ElasticLoadBalancingV2Controller) RegisterDataSource(dataSourceName string, dataSourceFunc func(ctx context.Context, client *elasticloadbalancingv2.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error) {
	e.DataSources[dataSourceName] = dataSourceFunc
}

func (e *ElasticLoadBalancingV2Controller) Process(ctx context.Context, accountId, region string, reportTime time.Time, cfg aws.Config, storageManager *storage.StorageManager) map[string]error {
	elasticloadbalancingv2Client := elasticloadbalancingv2.NewFromConfig(cfg)

	errMap := map[string]error{}

	for dataSourceName, dataSourceFunc := range e.DataSources {
		storageConfig := storage.StorageContextConfig{
			Cloud:      "aws",
			Service:    "elasticloadbalancingv2",
			Region:     region,
			AccountId:  accountId,
			DataSource: dataSourceName,
			Date:       reportTime,
		}
		err := dataSourceFunc(ctx, elasticloadbalancingv2Client, cfg, reportTime, storageConfig, storageManager)
		if err != nil {
			errMap[dataSourceName] = err
		}
		logrus.WithFields(logrus.Fields{
			"cloud":       "aws",
			"service":     "elasticloadbalancingv2",
			"region":      region,
			"account_id":  accountId,
			"datasource":  dataSourceName,
			"report_time": reportTime,
		}).Info("processed data source")
	}

	if len(errMap) == 0 {
		return nil
	} else {
		return errMap
	}
}

func GetTagMap(tags []types.Tag) map[string]string {
	tagMap := make(map[string]string)
	for _, tag := range tags {
		tagMap[*tag.Key] = *tag.Value
	}
	return tagMap
}
