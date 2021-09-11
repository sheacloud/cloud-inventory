package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/sheacloud/cloud-inventory/internal/catalog"
	"github.com/sheacloud/cloud-inventory/internal/processor"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/spf13/viper"
)

var (
	parquetS3Viper = viper.New()
	logViper       = viper.New()
	awsViper       = viper.New()
)

func initLogOptions() {
	logViper.SetEnvPrefix("log")
	logViper.AutomaticEnv()

	logViper.BindEnv("level")
	logViper.SetDefault("level", "info")

	logViper.BindEnv("caller")
	logViper.SetDefault("caller", false)
}

func initParquetS3Options() {
	parquetS3Viper.SetEnvPrefix("parquet_s3")
	parquetS3Viper.AutomaticEnv()

	parquetS3Viper.BindEnv("bucket")

	parquetS3Viper.BindEnv("path_prefix")
	parquetS3Viper.SetDefault("path_prefix", "parquet/")

	parquetS3Viper.BindEnv("file_extension")
	parquetS3Viper.SetDefault("file_extension", "parquet")

	parquetS3Viper.BindEnv("parquet_pages")
	parquetS3Viper.SetDefault("parquet_pages", 4)
}

func initAwsOptions() {
	awsViper.SetEnvPrefix("aws")
	awsViper.AutomaticEnv()

	awsViper.BindEnv("accounts")

	awsViper.BindEnv("regions")
	awsViper.SetDefault("regions", "us-east-1,us-west-2")

	awsViper.BindEnv("assume_role_name")
}

func init() {
	initLogOptions()
	initParquetS3Options()
	initAwsOptions()
}

func stringInList(s string, sList []string) bool {
	for _, s2 := range sList {
		if s == s2 {
			return true
		}
	}
	return false
}

func main() {
	baseAwsConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	reportTime := time.Now().UTC()
	parquetBackend := storage.NewParquetS3Backend(parquetS3Viper.GetString("bucket"), parquetS3Viper.GetString("path_prefix"), parquetS3Viper.GetString("file_extension"), s3.NewFromConfig(baseAwsConfig), int64(parquetS3Viper.GetInt("parquet_pages")))

	storageManager := storage.StorageManager{
		StorageBackends: []storage.StorageBackend{
			parquetBackend,
		},
	}

	awsProcessor := processor.NewAwsJobProcessor(16, 1000)
	awsProcessor.Start()

	accounts := strings.Split(awsViper.GetString("accounts"), ",")
	regions := strings.Split(awsViper.GetString("regions"), ",")
	assumeRoleName := awsViper.GetString("assume_role_name")

	stsSvc := sts.NewFromConfig(baseAwsConfig)

	for _, account := range accounts {
		accountAwsConfig := baseAwsConfig.Copy()
		creds := stscreds.NewAssumeRoleProvider(stsSvc, fmt.Sprintf("arn:aws:iam::%s:role/%s", account, assumeRoleName))
		accountAwsConfig.Credentials = aws.NewCredentialsCache(creds)
		for _, region := range regions {
			regionConfig := accountAwsConfig.Copy()
			regionConfig.Region = region
			for _, serviceController := range catalog.AwsServiceControllers {
				regionOverrides := serviceController.GetRegionOverrides()
				if len(regionOverrides) > 0 && !stringInList(region, regionOverrides) {
					//skip running in this region
					continue
				}

				awsProcessor.AddJob(processor.NewAwsJob(
					serviceController,
					context.TODO(),
					account,
					region,
					reportTime,
					regionConfig,
					&storageManager,
				))
			}
		}
	}

	awsProcessor.WaitForCompletion()
	parquetBackend.CloseStorageContexts(context.TODO())
}
