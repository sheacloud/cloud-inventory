package main

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
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

	viper.BindEnv("api_keys_table")
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
	if viper.GetString("api_keys_table") == "" {
		panic("api_keys_table is required")
	}
}

func init() {
	initOptions()
	initLogging()
	validateOptions()
}

func handleRequest(ctx context.Context, event events.APIGatewayCustomAuthorizerRequestTypeRequest) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {
	apiKey, ok := event.Headers["x-api-key"]
	if !ok {
		logrus.Info("missing auth header")
		return events.APIGatewayV2CustomAuthorizerSimpleResponse{
			IsAuthorized: false,
			Context: map[string]interface{}{
				"message": "X-API-Key header is missing",
			},
		}, nil
	}

	// convert token from base64 to bytes
	apiKeyBytes, err := base64.StdEncoding.DecodeString(apiKey)
	if err != nil {
		logrus.Error("failed to decode token")
		return events.APIGatewayV2CustomAuthorizerSimpleResponse{
			IsAuthorized: false,
			Context: map[string]interface{}{
				"message": "Failed to decode token from base64",
			},
		}, nil
	}

	// hash the token
	h := sha256.New()
	h.Write(apiKeyBytes)
	apiKeyHash := h.Sum(nil)

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	dynamodbClient := dynamodb.NewFromConfig(cfg)

	// check the database for the token
	getResponse, err := dynamodbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(viper.GetString("api_keys_table")),
		Key: map[string]types.AttributeValue{
			"api_key_hash": &types.AttributeValueMemberB{
				Value: apiKeyHash,
			},
		},
	})
	if err != nil {
		logrus.Fatal("failed to get item from dynamodb", err)
	}

	if getResponse.Item == nil {
		// token not found, unauthorized request
		return events.APIGatewayV2CustomAuthorizerSimpleResponse{
			IsAuthorized: false,
			Context: map[string]interface{}{
				"message": "Unauthorized",
			},
		}, nil
	}
	userId := ""
	if getResponse.Item["user_id"] != nil {
		userId = getResponse.Item["user_id"].(*types.AttributeValueMemberS).Value
	}

	// token found, authorized request
	// update the last-accessed field in the database
	_, err = dynamodbClient.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: aws.String(viper.GetString("api_key_table")),
		Key: map[string]types.AttributeValue{
			"api_key_hash": &types.AttributeValueMemberB{
				Value: apiKeyHash,
			},
		},
		UpdateExpression: aws.String("set last_accessed = :last_accessed"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":last_accessed": &types.AttributeValueMemberS{
				Value: time.Now().UTC().Format(time.RFC3339),
			},
		},
	})
	if err != nil {
		logrus.Error("failed to update item in dynamodb", err)
	}
	return events.APIGatewayV2CustomAuthorizerSimpleResponse{
		IsAuthorized: true,
		Context: map[string]interface{}{
			"user_id": userId,
		},
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
