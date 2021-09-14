package efs

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

var (
	Controller = EfsController{
		DataSources: map[string]func(ctx context.Context, client *efs.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error{},
	}
)

type EfsController struct {
	DataSources map[string]func(ctx context.Context, client *efs.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error
}

func (e *EfsController) GetRegionOverrides() []string {
	return []string{}
}

func (e *EfsController) GetName() string {
	return "efs"
}

func (e *EfsController) RegisterDataSource(dataSourceName string, dataSourceFunc func(ctx context.Context, client *efs.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error) {
	e.DataSources[dataSourceName] = dataSourceFunc
}

func (e *EfsController) Process(ctx context.Context, accountId, region string, reportTime time.Time, cfg aws.Config, storageManager *storage.StorageManager) map[string]error {
	efsClient := efs.NewFromConfig(cfg)

	errMap := map[string]error{}

	for dataSourceName, dataSourceFunc := range e.DataSources {
		storageConfig := storage.StorageContextConfig{
			Cloud:      "aws",
			Service:    "efs",
			Region:     region,
			AccountId:  accountId,
			DataSource: dataSourceName,
			Date:       reportTime,
		}
		err := dataSourceFunc(ctx, efsClient, cfg, reportTime, storageConfig, storageManager)
		if err != nil {
			errMap[dataSourceName] = err
		}
		logrus.WithFields(logrus.Fields{
			"cloud":       "aws",
			"service":     "efs",
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
