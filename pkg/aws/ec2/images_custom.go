package ec2

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

func ImageDataSource(ctx context.Context, client *ec2.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(ImageModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	result, err := client.DescribeImages(ctx, &ec2.DescribeImagesInput{
		Owners: []string{"self"},
	})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     storageConfig.Service,
			"data_source": storageConfig.DataSource,
			"account_id":  storageConfig.AccountId,
			"region":      storageConfig.Region,
			"cloud":       storageConfig.Cloud,
			"error":       err,
		}).Error("error calling DescribeImages")
		return err
	}

	for _, image := range result.Images {

		model := new(ImageModel)
		copier.Copy(&model, &image)

		model.Tags = GetTagMap(image.Tags)
		model.AccountId = storageConfig.AccountId
		model.Region = storageConfig.Region
		model.ReportTime = reportTime.UTC().UnixMilli()

		for _, f := range customImageModelPostprocessingFuncs {
			f(ctx, client, cfg, model)
		}

		errors := storageContextSet.Store(ctx, model)
		for storageContext, err := range errors {
			storage.LogContextError(storageContext, fmt.Sprintf("Error storing ImageModel: %v", err))
		}
	}

	return nil
}
