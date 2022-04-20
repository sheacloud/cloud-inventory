package main

import (
	"context"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dynamoDAO "github.com/sheacloud/cloud-inventory/internal/db/dynamodb"
	"github.com/sheacloud/cloud-inventory/internal/inventory"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	logrusLevels = map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
	}
)

func initOptions() {
	viper.SetEnvPrefix("cloud_inventory")
	viper.AutomaticEnv()

	viper.BindEnv("log_level")
	viper.SetDefault("log_level", "info")

	viper.BindEnv("log_caller")
	viper.SetDefault("log_caller", false)

	viper.BindEnv("s3_bucket")

	viper.BindEnv("aws_accounts_ids")

	viper.BindEnv("aws_regions")

	viper.BindEnv("aws_use_local_credentials")
	viper.SetDefault("aws_use_local_credentials", false)

	viper.BindEnv("aws_assume_role_name")

	viper.BindEnv("aws_processor_routines")
	viper.SetDefault("aws_processor_routines", 32)

	viper.BindEnv("mongo_uri")
}

func initLogging() {
	logrus.SetLevel(logrusLevels[viper.GetString("log_level")])
	logrus.SetReportCaller(viper.GetBool("log_caller"))
}

func validateOptions() {
	if viper.GetString("s3_bucket") == "" {
		panic("s3_bucket is required")
	}

	if viper.GetString("aws_account_ids") == "" && !viper.GetBool("aws_use_local_credentials") {
		panic("aws_account_ids is required when aws_use_local_credentials is false")
	}

	if viper.GetString("aws_regions") == "" {
		panic("aws_regions is required")
	}

	if viper.GetString("aws_assume_role_name") == "" && !viper.GetBool("aws_use_local_credentials") {
		panic("aws_assume_role_name is required when aws_use_local_credentials is false")
	}
}

func init() {
	initOptions()
	initLogging()
	validateOptions()
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	// s3Client := s3.NewFromConfig(cfg)
	dynamoClient := dynamodb.NewFromConfig(cfg)

	// fileManager := indexedstorage.NewIndexedFileManager(viper.GetString("s3_bucket"), "inventory/", "parquet", s3Client)

	accountIDs := strings.Split(viper.GetString("aws_account_ids"), ",")
	regions := strings.Split(viper.GetString("aws_regions"), ",")

	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(viper.GetString("mongo_uri")))
	// if err != nil {
	// 	panic(err)
	// }

	// dao := mongoDAO.NewMongoDAO(client.Database("cloud-inventory"))
	dao := dynamoDAO.NewDynamoDBDAO(dynamoClient, 3)

	inventory.FetchAwsInventory(context.TODO(), accountIDs, regions, cfg, viper.GetBool("aws_use_local_credentials"), viper.GetString("aws_assume_role_name"), time.Now().UTC(), dao, viper.GetInt("aws_processor_routines"))
}
