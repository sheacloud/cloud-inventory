package ec2

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

var (
	Controller = Ec2Controller{
		DataSources: map[string]func(ctx context.Context, client *awsec2.Client, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error{},
	}
)

type Ec2Controller struct {
	DataSources map[string]func(ctx context.Context, client *awsec2.Client, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error
}

func (e *Ec2Controller) GetName() string {
	return "ec2"
}

func (e *Ec2Controller) RegisterDataSource(dataSourceName string, dataSourceFunc func(ctx context.Context, client *awsec2.Client, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error) {
	e.DataSources[dataSourceName] = dataSourceFunc
}

func (e *Ec2Controller) Process(ctx context.Context, accountId, region string, date time.Time, cfg aws.Config, storageManager *storage.StorageManager) map[string]error {
	ec2Client := awsec2.NewFromConfig(cfg)

	errMap := map[string]error{}

	for dataSourceName, dataSourceFunc := range e.DataSources {
		storageConfig := storage.StorageContextConfig{
			Cloud:      "aws",
			Service:    "ec2",
			Region:     region,
			AccountId:  accountId,
			DataSource: dataSourceName,
			Date:       date,
		}
		err := dataSourceFunc(ctx, ec2Client, storageConfig, storageManager)
		if err != nil {
			errMap[dataSourceName] = err
		}
		logrus.WithFields(logrus.Fields{
			"cloud":      "aws",
			"service":    "ec2",
			"region":     region,
			"account_id": accountId,
			"datasource": dataSourceName,
			"date":       date,
		}).Info("processed data source")
	}

	if len(errMap) == 0 {
		return nil
	} else {
		return errMap
	}
}

func GetTagMap(tags []ec2types.Tag) map[string]string {
	tagMap := make(map[string]string)
	for _, tag := range tags {
		tagMap[*tag.Key] = *tag.Value
	}
	return tagMap
}
