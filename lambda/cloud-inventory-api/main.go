package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
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

	router    = gin.Default()
	ginLambda *ginadapter.GinLambdaV2
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

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	dynamoClient := dynamodb.NewFromConfig(cfg)

	dao := dynamoDAO.NewDynamoDBDAO(dynamoClient, 3)

	router = api.GetRouter(dao)

	ginLambda = ginadapter.NewV2(router)
}

func Handler(ctx context.Context, event events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, event)
}

func main() {
	lambda.Start(Handler)
}
