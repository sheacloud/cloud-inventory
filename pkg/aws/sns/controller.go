package sns

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

var (
	Controller = SnsController{
		DataSources: map[string]func(ctx context.Context, client *sns.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error{},
	}
)

type SnsController struct {
	DataSources map[string]func(ctx context.Context, client *sns.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error
}

func (e *SnsController) GetRegionOverrides() []string {
	return []string{}
}

func (e *SnsController) GetName() string {
	return "sns"
}

func (e *SnsController) RegisterDataSource(dataSourceName string, dataSourceFunc func(ctx context.Context, client *sns.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error) {
	e.DataSources[dataSourceName] = dataSourceFunc
}

func (e *SnsController) Process(ctx context.Context, accountId, region string, reportTime time.Time, cfg aws.Config, storageManager *storage.StorageManager) map[string]error {
	snsClient := sns.NewFromConfig(cfg)

	errMap := map[string]error{}

	for dataSourceName, dataSourceFunc := range e.DataSources {
		storageConfig := storage.StorageContextConfig{
			Cloud:      "aws",
			Service:    "sns",
			Region:     region,
			AccountId:  accountId,
			DataSource: dataSourceName,
			Date:       reportTime,
		}
		err := dataSourceFunc(ctx, snsClient, cfg, reportTime, storageConfig, storageManager)
		if err != nil {
			errMap[dataSourceName] = err
		}
		logrus.WithFields(logrus.Fields{
			"cloud":       "aws",
			"service":     "sns",
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
