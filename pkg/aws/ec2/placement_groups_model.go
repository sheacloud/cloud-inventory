// AUTOGENERATED, DO NOT EDIT
package ec2

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

var customPlacementGroupModelPostprocessingFuncs []func(x *PlacementGroupModel) = []func(x *PlacementGroupModel){}
var customPlacementGroupModelFuncsLock sync.Mutex

func registerCustomPlacementGroupModelPostprocessingFunc(f func(x *PlacementGroupModel)) {
	customPlacementGroupModelFuncsLock.Lock()
	defer customPlacementGroupModelFuncsLock.Unlock()

	customPlacementGroupModelPostprocessingFuncs = append(customPlacementGroupModelPostprocessingFuncs, f)
}

func init() {
	Controller.RegisterDataSource("placement_groups", PlacementGroupDataSource)
}

type PlacementGroupModel struct {
	GroupId        string            `parquet:"name=group_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	GroupName      string            `parquet:"name=group_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	PartitionCount int32             `parquet:"name=partition_count,type=INT32"`
	State          string            `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
	Strategy       string            `parquet:"name=strategy,type=BYTE_ARRAY,convertedtype=UTF8"`
	Tags           map[string]string `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
	AccountId      string            `parquet:"name=account_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Region         string            `parquet:"name=region, type=BYTE_ARRAY, convertedtype=UTF8"`
	ReportTime     int64             `parquet:"name=report_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
}

type TagPlacementGroupModel struct {
	Key   string `parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

func PlacementGroupDataSource(ctx context.Context, client *ec2.Client, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(PlacementGroupModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	params := &ec2.DescribePlacementGroupsInput{}

	result, err := client.DescribePlacementGroups(ctx, params)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     storageConfig.Service,
			"data_source": storageConfig.DataSource,
			"account_id":  storageConfig.AccountId,
			"region":      storageConfig.Region,
			"cloud":       storageConfig.Cloud,
			"error":       err,
		}).Error("error calling DescribePlacementGroups")
		return err
	}

	results := []*ec2.DescribePlacementGroupsOutput{result}
	for _, output := range results {

		for _, var0 := range output.PlacementGroups {

			model := new(PlacementGroupModel)
			copier.Copy(&model, &var0)

			model.Tags = GetTagMap(var0.Tags)
			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region
			model.ReportTime = reportTime.UTC().UnixMilli()

			for _, f := range customPlacementGroupModelPostprocessingFuncs {
				f(model)
			}

			errors := storageContextSet.Store(ctx, model)
			for storageContext, err := range errors {
				storage.LogContextError(storageContext, fmt.Sprintf("Error storing PlacementGroupModel: %v", err))
			}
		}

	}

	return nil
}