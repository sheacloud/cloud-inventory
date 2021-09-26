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

var customInstanceProfileModelPostprocessingFuncs []func(ctx context.Context, client *iam.Client, cfg aws.Config, x *InstanceProfileModel) = []func(ctx context.Context, client *iam.Client, cfg aws.Config, x *InstanceProfileModel){}
var customInstanceProfileModelFuncsLock sync.Mutex

func registerCustomInstanceProfileModelPostprocessingFunc(f func(ctx context.Context, client *iam.Client, cfg aws.Config, x *InstanceProfileModel)) {
	customInstanceProfileModelFuncsLock.Lock()
	defer customInstanceProfileModelFuncsLock.Unlock()

	customInstanceProfileModelPostprocessingFuncs = append(customInstanceProfileModelPostprocessingFuncs, f)
}

func init() {
	Controller.RegisterDataSource("instance_profiles", InstanceProfileDataSource)
}

type InstanceProfileModel struct {
	Arn                 string `parquet:"name=arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	CreateDate          *time.Time
	InstanceProfileId   string                      `parquet:"name=instance_profile_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	InstanceProfileName string                      `parquet:"name=instance_profile_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	Path                string                      `parquet:"name=path,type=BYTE_ARRAY,convertedtype=UTF8"`
	Roles               []*RoleInstanceProfileModel `parquet:"name=roles,type=MAP,convertedtype=LIST"`
	TagsOld             []*TagInstanceProfileModel
	CreateDateMilli     int64             `parquet:"name=create_date_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	Tags                map[string]string `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
	AccountId           string            `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region              string            `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime          int64             `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type RoleInstanceProfileModel struct {
	Arn                      string `parquet:"name=arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	CreateDate               *time.Time
	Path                     string                                           `parquet:"name=path,type=BYTE_ARRAY,convertedtype=UTF8"`
	RoleId                   string                                           `parquet:"name=role_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	RoleName                 string                                           `parquet:"name=role_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	AssumeRolePolicyDocument string                                           `parquet:"name=assume_role_policy_document,type=BYTE_ARRAY,convertedtype=UTF8"`
	Description              string                                           `parquet:"name=description,type=BYTE_ARRAY,convertedtype=UTF8"`
	MaxSessionDuration       int32                                            `parquet:"name=max_session_duration,type=INT32"`
	PermissionsBoundary      *AttachedPermissionsBoundaryInstanceProfileModel `parquet:"name=permissions_boundary"`
	RoleLastUsed             *RoleLastUsedInstanceProfileModel                `parquet:"name=role_last_used"`
	TagsOld                  []*TagInstanceProfileModel
	CreateDateMilli          int64             `parquet:"name=create_date_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	Tags                     map[string]string `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
}

type AttachedPermissionsBoundaryInstanceProfileModel struct {
	PermissionsBoundaryArn  string `parquet:"name=permissions_boundary_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	PermissionsBoundaryType string `parquet:"name=permissions_boundary_type,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type RoleLastUsedInstanceProfileModel struct {
	LastUsedDate      *time.Time
	Region            string `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	LastUsedDateMilli int64  `parquet:"name=last_used_date_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type TagInstanceProfileModel struct {
	Key   string `parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

func InstanceProfileDataSource(ctx context.Context, client *iam.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(InstanceProfileModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := iam.NewListInstanceProfilesPaginator(client, &iam.ListInstanceProfilesInput{})

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
			}).Error("error calling ListInstanceProfiles")
			return err
		}

		for _, var0 := range output.InstanceProfiles {

			model := new(InstanceProfileModel)
			copier.Copy(&model, &var0)

			model.Tags = GetTagMap(var0.Tags)
			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region
			model.ReportTime = reportTime.UTC().UnixMilli()

			for _, f := range customInstanceProfileModelPostprocessingFuncs {
				f(ctx, client, cfg, model)
			}

			errors := storageContextSet.Store(ctx, model)
			for storageContext, err := range errors {
				storage.LogContextError(storageContext, fmt.Sprintf("Error storing InstanceProfileModel: %v", err))
			}
		}

	}

	return nil
}
