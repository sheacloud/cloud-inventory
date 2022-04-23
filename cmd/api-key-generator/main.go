package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	_ "github.com/sheacloud/cloud-inventory/docs"
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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: api-key-generator <user_id>")
		os.Exit(1)
	}
	userName := os.Args[1]
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	dynamodbClient := dynamodb.NewFromConfig(cfg)

	// generate random api key
	apiKeyBytes := make([]byte, 32)
	_, err = rand.Read(apiKeyBytes)
	if err != nil {
		panic(err)
	}

	// hash the api key
	h := sha256.New()
	h.Write(apiKeyBytes)
	apiKeyHash := h.Sum(nil)

	// check if the api key is already in the table
	getResponse, err := dynamodbClient.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(viper.GetString("api_keys_table")),
		Key: map[string]types.AttributeValue{
			"api_key_hash": &types.AttributeValueMemberB{
				Value: apiKeyHash,
			},
		},
	})
	if err != nil {
		panic(err)
	}
	if getResponse.Item == nil {
		// api key doesn't exist, add it to dynamodb
		_, err := dynamodbClient.PutItem(context.TODO(), &dynamodb.PutItemInput{
			TableName: aws.String(viper.GetString("api_keys_table")),
			Item: map[string]types.AttributeValue{
				"api_key_hash": &types.AttributeValueMemberB{
					Value: apiKeyHash,
				},
				"user_id": &types.AttributeValueMemberS{
					Value: userName,
				},
			},
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("API key:", base64.StdEncoding.EncodeToString(apiKeyBytes))
	} else {
		panic("api key already exists, please try again")
	}

}
