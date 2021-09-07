package iam

import (
	"context"
	"fmt"
	"time"

	awsiam "github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

func init() {
	Controller.RegisterDataSource("roles", RoleDataSource)
}

type RoleModel struct {
	Arn                      string `parquet:"name=arn, type=BYTE_ARRAY, convertedtype=UTF8"`
	CreateDate               time.Time
	CreateDateMilli          int64                             `parquet:"name=create_date, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	Path                     string                            `parquet:"name=path, type=BYTE_ARRAY, convertedtype=UTF8"`
	RoleId                   string                            `parquet:"name=role_id, type=BYTE_ARRAY, convertedtype=UTF8" inventory_primary_key:"true"`
	RoleName                 string                            `parquet:"name=role_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	AssumeRolePolicyDocument string                            `parquet:"name=assume_role_policy_document, type=BYTE_ARRAY, convertedtype=UTF8"`
	Description              string                            `parquet:"name=description, type=BYTE_ARRAY, convertedtype=UTF8"`
	MaxSessionDuration       int32                             `parquet:"name=max_session_duration, type=INT32"`
	PermissionsBoundary      *AttachedPermissionsBoundaryModel `parquet:"name=permissions_boundary"`
	RoleLastUsed             *RoleLastUsedModel                `parquet:"name=role_last_used"`
	Tags                     map[string]string                 `parquet:"name=tags, type=MAP, keytype=BYTE_ARRAY, keyconvertedtype=UTF8, valuetype=BYTE_ARRAY, valueconvertedtype=UTF8"`
	AccountId                string                            `parquet:"name=account_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Region                   string                            `parquet:"name=region, type=BYTE_ARRAY, convertedtype=UTF8"`
	ReportTime               int64                             `parquet:"name=report_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
}

type AttachedPermissionsBoundaryModel struct {
	PermissionsBoundaryArn  string `parquet:"name=permissions_boundary_arn, type=BYTE_ARRAY, convertedtype=UTF8"`
	PermissionsBoundaryType string `parquet:"name=permissions_boundary_type, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type RoleLastUsedModel struct {
	LastUsedDate      time.Time
	LastUsedDateMilli int64  `parquet:"name=last_used_date, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	Region            string `parquet:"name=region, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type RoleDataSourceClient interface {
	ListRoles(context.Context, *awsiam.ListRolesInput, ...func(*awsiam.Options)) (*awsiam.ListRolesOutput, error)
}

func RoleDataSource(ctx context.Context, client *awsiam.Client, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	return roleDataSource(ctx, client, reportTime, storageConfig, storageManager)
}

// function with client as a specific interface, allowing mocking/testing
func roleDataSource(ctx context.Context, client RoleDataSourceClient, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(RoleModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := awsiam.NewListRolesPaginator(client, &awsiam.ListRolesInput{})

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
			}).Error("error calling ListRoles")
			return err
		}

		for _, role := range output.Roles {
			model := new(RoleModel)
			copier.Copy(&model, &role)

			model.Tags = GetTagMap(role.Tags)
			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region
			model.ReportTime = reportTime.UnixMilli()
			model.CreateDateMilli = model.CreateDate.UnixMilli()

			if model.RoleLastUsed != nil {
				model.RoleLastUsed.LastUsedDateMilli = model.RoleLastUsed.LastUsedDate.UnixMilli()
			}

			errors := storageContextSet.Store(ctx, model)
			for storageContext, err := range errors {
				storage.LogContextError(storageContext, fmt.Sprintf("Error storing RoleModel: %v", err))
			}
		}
	}

	return nil
}
