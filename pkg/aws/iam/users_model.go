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

var customUserModelPostprocessingFuncs []func(ctx context.Context, client *iam.Client, cfg aws.Config, x *UserModel) = []func(ctx context.Context, client *iam.Client, cfg aws.Config, x *UserModel){}
var customUserModelFuncsLock sync.Mutex

func registerCustomUserModelPostprocessingFunc(f func(ctx context.Context, client *iam.Client, cfg aws.Config, x *UserModel)) {
	customUserModelFuncsLock.Lock()
	defer customUserModelFuncsLock.Unlock()

	customUserModelPostprocessingFuncs = append(customUserModelPostprocessingFuncs, f)
}

func init() {
	Controller.RegisterDataSource("users", UserDataSource)
}

type UserModel struct {
	Arn                   string `parquet:"name=arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	CreateDate            *time.Time
	CreateDateMilli       int64  `parquet:"name=create_date, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	Path                  string `parquet:"name=path,type=BYTE_ARRAY,convertedtype=UTF8"`
	UserId                string `parquet:"name=user_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	UserName              string `parquet:"name=user_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	PasswordLastUsed      *time.Time
	PasswordLastUsedMilli int64                                 `parquet:"name=password_last_used, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	PermissionsBoundary   *AttachedPermissionsBoundaryUserModel `parquet:"name=permissions_boundary"`
	Tags                  map[string]string                     `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
	AccountId             string                                `parquet:"name=account_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Region                string                                `parquet:"name=region, type=BYTE_ARRAY, convertedtype=UTF8"`
	ReportTime            int64                                 `parquet:"name=report_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	AccessKeys            []*AccessKeyMetadataUserModel         `parquet:"name=access_keys,type=LIST"`
	LoginProfile          *LoginProfileUserModel                `parquet:"name=login_profile"`
	AttachedPolicies      []*AttachedPolicyUserModel            `parquet:"name=attached_policies,type=LIST"`
	InlinePolicies        []string                              `parquet:"name=inline_policies,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	GroupIds              []string                              `parquet:"name=group_ids,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
}

type AttachedPermissionsBoundaryUserModel struct {
	PermissionsBoundaryArn  string `parquet:"name=permissions_boundary_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	PermissionsBoundaryType string `parquet:"name=permissions_boundary_type,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type TagUserModel struct {
	Key   string `parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type AccessKeyMetadataUserModel struct {
	AccessKeyId     string `parquet:"name=access_key_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	CreateDate      *time.Time
	CreateDateMilli int64  `parquet:"name=create_date, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	Status          string `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type LoginProfileUserModel struct {
	CreateDate            *time.Time
	CreateDateMilli       int64 `parquet:"name=create_date, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	PasswordResetRequired bool  `parquet:"name=password_reset_required,type=BOOLEAN"`
}

type AttachedPolicyUserModel struct {
	PolicyArn  string `parquet:"name=policy_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	PolicyName string `parquet:"name=policy_name,type=BYTE_ARRAY,convertedtype=UTF8"`
}

func UserDataSource(ctx context.Context, client *iam.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(UserModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := iam.NewListUsersPaginator(client, &iam.ListUsersInput{})

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
			}).Error("error calling ListUsers")
			return err
		}

		for _, var0 := range output.Users {

			model := new(UserModel)
			copier.Copy(&model, &var0)

			model.Tags = GetTagMap(var0.Tags)
			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region
			model.ReportTime = reportTime.UTC().UnixMilli()

			for _, f := range customUserModelPostprocessingFuncs {
				f(ctx, client, cfg, model)
			}

			errors := storageContextSet.Store(ctx, model)
			for storageContext, err := range errors {
				storage.LogContextError(storageContext, fmt.Sprintf("Error storing UserModel: %v", err))
			}
		}

	}

	return nil
}
