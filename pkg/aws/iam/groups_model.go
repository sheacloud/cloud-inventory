// AUTOGENERATED, DO NOT EDIT
package iam

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
	"time"

	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"sync"
)

var customGroupModelPostprocessingFuncs []func(ctx context.Context, client *iam.Client, cfg aws.Config, x *GroupModel) = []func(ctx context.Context, client *iam.Client, cfg aws.Config, x *GroupModel){}
var customGroupModelFuncsLock sync.Mutex

func registerCustomGroupModelPostprocessingFunc(f func(ctx context.Context, client *iam.Client, cfg aws.Config, x *GroupModel)) {
	customGroupModelFuncsLock.Lock()
	defer customGroupModelFuncsLock.Unlock()

	customGroupModelPostprocessingFuncs = append(customGroupModelPostprocessingFuncs, f)
}

func init() {
	Controller.RegisterDataSource("groups", GroupDataSource)
}

type GroupModel struct {
	Arn              string `parquet:"name=arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	CreateDate       *time.Time
	GroupId          string                      `parquet:"name=group_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	GroupName        string                      `parquet:"name=group_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	Path             string                      `parquet:"name=path,type=BYTE_ARRAY,convertedtype=UTF8"`
	CreateDateMilli  int64                       `parquet:"name=create_date_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	AccountId        string                      `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region           string                      `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime       int64                       `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	InlinePolicies   []string                    `parquet:"name=inline_policies,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	UserIds          []string                    `parquet:"name=user_ids,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	AttachedPolicies []*AttachedPolicyGroupModel `parquet:"name=attached_policies,type=MAP,convertedtype=LIST"`
}

type AttachedPolicyGroupModel struct {
	PolicyArn  string `parquet:"name=policy_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	PolicyName string `parquet:"name=policy_name,type=BYTE_ARRAY,convertedtype=UTF8"`
}

func GroupDataSource(ctx context.Context, client *iam.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(GroupModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := iam.NewListGroupsPaginator(client, &iam.ListGroupsInput{})

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
			}).Error("error calling ListGroups")
			return err
		}

		for _, var0 := range output.Groups {

			model := new(GroupModel)
			copier.Copy(&model, &var0)

			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region
			model.ReportTime = reportTime.UTC().UnixMilli()

			for _, f := range customGroupModelPostprocessingFuncs {
				f(ctx, client, cfg, model)
			}

			errors := storageContextSet.Store(ctx, model)
			for storageContext, err := range errors {
				storage.LogContextError(storageContext, fmt.Sprintf("Error storing GroupModel: %v", err))
			}
		}

	}

	return nil
}
