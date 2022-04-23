package main

import (
	"context"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/sheacloud/cloud-inventory/internal/db"
	dynamoDao "github.com/sheacloud/cloud-inventory/internal/db/dynamodb"
	mongoDao "github.com/sheacloud/cloud-inventory/internal/db/mongo"
	"github.com/sheacloud/cloud-inventory/internal/db/multi"
	s3ParquetDao "github.com/sheacloud/cloud-inventory/internal/db/s3parquet"
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

	viper.BindEnv("aws_accounts_ids")

	viper.BindEnv("aws_regions")

	viper.BindEnv("aws_use_local_credentials")
	viper.SetDefault("aws_use_local_credentials", false)

	viper.BindEnv("aws_assume_role_name")

	viper.BindEnv("aws_processor_routines")
	viper.SetDefault("aws_processor_routines", 32)

	// DAO settings
	viper.BindEnv("mongo_uri")

	viper.BindEnv("dynamodb_table_prefix")

	viper.BindEnv("s3_bucket")
}

func initLogging() {
	logrus.SetLevel(logrusLevels[viper.GetString("log_level")])
	logrus.SetReportCaller(viper.GetBool("log_caller"))
}

func validateOptions() {
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

func initializeDAOs(cfg aws.Config) db.WriterDAO {
	var s3ParquetDAO *s3ParquetDao.S3ParquetWriterDAO
	var mongoDAO *mongoDao.MongoWriterDAO
	var dynamoDAO *dynamoDao.DynamoDBWriterDAO

	selectedDAOs := []db.WriterDAO{}
	if viper.GetString("mongo_uri") != "" {
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(viper.GetString("mongo_uri")))
		if err != nil {
			panic(err)
		}
		mongoDAO = mongoDao.NewMongoWriterDAO(client.Database("cloud-inventory"), 3)
		selectedDAOs = append(selectedDAOs, mongoDAO)
	}
	if viper.GetString("dynamodb_table_prefix") != "" {
		dynamoClient := dynamodb.NewFromConfig(cfg)
		dynamoDAO = dynamoDao.NewDynamoDBWriterDAO(dynamoClient, 3)
		selectedDAOs = append(selectedDAOs, dynamoDAO)
	}
	if viper.GetString("s3_bucket") != "" {
		s3Client := s3.NewFromConfig(cfg)
		s3ParquetDAO = s3ParquetDao.NewS3ParquetWriterDAO(s3Client, viper.GetString("s3_bucket"), 32)
		selectedDAOs = append(selectedDAOs, s3ParquetDAO)
	}

	if len(selectedDAOs) == 0 {
		panic("no DAO configured. Must specify at least one of mongo_uri, dynamodb_table_prefix, s3_bucket")
	}
	if len(selectedDAOs) == 1 {
		return selectedDAOs[0]
	}
	return multi.NewMultiWriterDAO(selectedDAOs)
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	accountIDs := strings.Split(viper.GetString("aws_account_ids"), ",")
	regions := strings.Split(viper.GetString("aws_regions"), ",")

	dao := initializeDAOs(cfg)

	inventory.FetchAwsInventory(context.TODO(), accountIDs, regions, cfg, viper.GetBool("aws_use_local_credentials"), viper.GetString("aws_assume_role_name"), time.Now().UTC().UnixMilli(), dao, viper.GetInt("aws_processor_routines"))

	dao.Finish(context.TODO())
}
