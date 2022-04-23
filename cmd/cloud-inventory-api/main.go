package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/cloud-inventory/docs"
	"github.com/sheacloud/cloud-inventory/internal/api"
	dynamoDAO "github.com/sheacloud/cloud-inventory/internal/db/dynamodb"
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

	viper.BindEnv("s3_bucket")

	viper.BindEnv("mongo_uri")
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
	if viper.GetString("s3_bucket") == "" {
		panic("s3_bucket is required")
	}
}

func init() {
	initOptions()
	initLogging()
	validateOptions()

	docs.SwaggerInfo.Host = viper.GetString("api_url")

	// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(viper.GetString("mongo_uri")))
	// if err != nil {
	// 	panic(err)
	// }

	// dao := mongoDao.NewMongoDAO(client.Database("cloud-inventory"))

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	dynamoClient := dynamodb.NewFromConfig(cfg)

	dao := dynamoDAO.NewDynamoDBReaderDAO(dynamoClient, 3)

	router = api.GetRouter(dao)
}

func main() {
	router.Run("localhost:3000")
}
