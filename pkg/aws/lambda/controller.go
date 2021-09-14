package lambda

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

var (
	Controller = LambdaController{
		DataSources: map[string]func(ctx context.Context, client *lambda.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error{},
	}
)

type LambdaController struct {
	DataSources map[string]func(ctx context.Context, client *lambda.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error
}

func (e *LambdaController) GetRegionOverrides() []string {
	return []string{}
}

func (e *LambdaController) GetName() string {
	return "lambda"
}

func (e *LambdaController) RegisterDataSource(dataSourceName string, dataSourceFunc func(ctx context.Context, client *lambda.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error) {
	e.DataSources[dataSourceName] = dataSourceFunc
}

func (e *LambdaController) Process(ctx context.Context, accountId, region string, reportTime time.Time, cfg aws.Config, storageManager *storage.StorageManager) map[string]error {
	lambdaClient := lambda.NewFromConfig(cfg)

	errMap := map[string]error{}

	for dataSourceName, dataSourceFunc := range e.DataSources {
		storageConfig := storage.StorageContextConfig{
			Cloud:      "aws",
			Service:    "lambda",
			Region:     region,
			AccountId:  accountId,
			DataSource: dataSourceName,
			Date:       reportTime,
		}
		err := dataSourceFunc(ctx, lambdaClient, cfg, reportTime, storageConfig, storageManager)
		if err != nil {
			errMap[dataSourceName] = err
		}
		logrus.WithFields(logrus.Fields{
			"cloud":       "aws",
			"service":     "lambda",
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
