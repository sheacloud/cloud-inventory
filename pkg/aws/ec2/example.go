package ec2

import (
	"context"
	"fmt"
	"time"

	awsec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

// func init() {
// 	Controller.RegisterDataSource("example", ExampleDataSource)
// }

type ExampleModel struct {
	Field      string            `parquet:"name=field, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Tags       map[string]string `parquet:"name=tags, type=MAP, keytype=BYTE_ARRAY, keyconvertedtype=UTF8, valuetype=BYTE_ARRAY, valueconvertedtype=UTF8"`
	AccountId  string            `parquet:"name=account_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Region     string            `parquet:"name=region, type=BYTE_ARRAY, convertedtype=UTF8"`
	ReportTime int64             `parquet:"name=report_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
}

type ExampleDataSourceClient interface {
	DescribeVolumes(context.Context, *awsec2.DescribeVolumesInput, ...func(*awsec2.Options)) (*awsec2.DescribeVolumesOutput, error)
}

func ExampleDataSource(ctx context.Context, client *awsec2.Client, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	return exampleDataSource(ctx, client, reportTime, storageConfig, storageManager)
}

// function with client as a specific interface, allowing mocking/testing
func exampleDataSource(ctx context.Context, client ExampleDataSourceClient, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(ExampleModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := awsec2.NewDescribeVolumesPaginator(client, &awsec2.DescribeVolumesInput{})

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
			}).Error("error calling DescribeVolumes")
			return err
		}

		for _, volume := range output.Volumes {
			model := new(ExampleModel)
			copier.Copy(&model, &volume)

			model.Tags = GetTagMap(volume.Tags)
			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region
			model.ReportTime = reportTime.UnixMilli()

			errors := storageContextSet.Store(ctx, model)
			for storageContext, err := range errors {
				storage.LogContextError(storageContext, fmt.Sprintf("Error storing ExampleModel: %v", err))
			}
		}
	}

	return nil
}
