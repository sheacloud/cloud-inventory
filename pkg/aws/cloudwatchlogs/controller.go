package cloudwatchlogs

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

var (
	Controller = CloudwatchLogsController{
		DataSources: map[string]func(ctx context.Context, client *cloudwatchlogs.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error{},
	}
)

type CloudwatchLogsController struct {
	DataSources map[string]func(ctx context.Context, client *cloudwatchlogs.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error
}

func (e *CloudwatchLogsController) GetRegionOverrides() []string {
	return []string{}
}

func (e *CloudwatchLogsController) GetName() string {
	return "cloudwatchlogs"
}

func (e *CloudwatchLogsController) RegisterDataSource(dataSourceName string, dataSourceFunc func(ctx context.Context, client *cloudwatchlogs.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error) {
	e.DataSources[dataSourceName] = dataSourceFunc
}

func (e *CloudwatchLogsController) Process(ctx context.Context, accountId, region string, reportTime time.Time, cfg aws.Config, storageManager *storage.StorageManager) map[string]error {
	cloudwatchlogsClient := cloudwatchlogs.NewFromConfig(cfg)

	errMap := map[string]error{}

	for dataSourceName, dataSourceFunc := range e.DataSources {
		storageConfig := storage.StorageContextConfig{
			Cloud:      "aws",
			Service:    "cloudwatchlogs",
			Region:     region,
			AccountId:  accountId,
			DataSource: dataSourceName,
			Date:       reportTime,
		}
		err := dataSourceFunc(ctx, cloudwatchlogsClient, cfg, reportTime, storageConfig, storageManager)
		if err != nil {
			errMap[dataSourceName] = err
		}
		logrus.WithFields(logrus.Fields{
			"cloud":       "aws",
			"service":     "cloudwatchlogs",
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
