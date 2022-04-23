package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/cloud-inventory/docs"
	"github.com/sheacloud/cloud-inventory/internal/api"
	"github.com/sheacloud/cloud-inventory/internal/db"
	dynamoDao "github.com/sheacloud/cloud-inventory/internal/db/dynamodb"
	mongoDao "github.com/sheacloud/cloud-inventory/internal/db/mongo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	logrusLevels = map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
	}

	router = gin.Default()
)

func initOptions() {
	viper.SetEnvPrefix("cloud_inventory")
	viper.AutomaticEnv()

	viper.BindEnv("log_level")
	viper.SetDefault("log_level", "info")

	viper.BindEnv("log_caller")
	viper.SetDefault("log_caller", false)

	viper.BindEnv("api_url")
	viper.SetDefault("api_url", "localhost:3000")

	// DAO settings
	viper.BindEnv("database_type")

	viper.BindEnv("s3_bucket")

	viper.BindEnv("mongo_uri")

	viper.BindEnv("dynamodb_table_prefix")
}

func initLogging() {
	logrus.SetLevel(logrusLevels[viper.GetString("log_level")])
	logrus.SetReportCaller(viper.GetBool("log_caller"))
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02T15:04:05.999999999Z07:00",
	})
}

func validateOptions() {
	if viper.GetString("database_type") == "" {
		panic("database_type is required")
	}
	if viper.GetString("database_type") != "mongo" && viper.GetString("database_type") != "dynamodb" && viper.GetString("database_type") != "s3parquet" {
		panic("database_type must be one of mongo, dynamodb, s3parquet")
	}
}

func initializeDAO(cfg aws.Config) db.ReaderDAO {

	switch viper.GetString("database_type") {
	case "mongo":
		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(viper.GetString("mongo_uri")))
		if err != nil {
			panic(err)
		}
		mongoDAO := mongoDao.NewMongoReaderDAO(client.Database("cloud-inventory"), 3)
		return mongoDAO
	case "dynamodb":
		dynamoClient := dynamodb.NewFromConfig(cfg)
		dynamoDAO := dynamoDao.NewDynamoDBReaderDAO(dynamoClient, 3)
		return dynamoDAO
	}

	panic("No database selected")
}

func init() {
	initOptions()
	initLogging()
	validateOptions()

	docs.SwaggerInfo.Host = viper.GetString("api_url")

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	dao := initializeDAO(cfg)

	router = api.GetRouter(dao)
}

func main() {
	router.Run("localhost:3000")
}
