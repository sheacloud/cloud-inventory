package main

import (
	"bytes"
	"context"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/golang-jwt/jwt"
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

	jwks JWKS
)

func initOptions() {
	viper.SetEnvPrefix("cloud_inventory")
	viper.AutomaticEnv()

	viper.BindEnv("log_level")
	viper.SetDefault("log_level", "info")

	viper.BindEnv("log_caller")
	viper.SetDefault("log_caller", false)

	viper.BindEnv("api_keys_table")

	viper.BindEnv("cognito_user_pool_id")

	viper.BindEnv("cognito_user_pool_region")
	viper.SetDefault("cognito_user_pool_region", "us-east-1")
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
	if viper.GetString("cognito_user_pool_id") == "" {
		panic("cognito_user_pool_id is required")
	}
}

func downloadJWTKeys() {
	resp, err := http.Get(fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", viper.GetString("cognito_user_pool_region"), viper.GetString("cognito_user_pool_id")))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &jwks)
	if err != nil {
		panic(err)
	}
}

func init() {
	initOptions()
	initLogging()
	validateOptions()
	downloadJWTKeys()
}

type JWKS struct {
	Keys []JWK `json:"keys"`
}

type JWK struct {
	Alg string `json:"alg"`
	E   string `json:"e"`
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	N   string `json:"n"`
	Use string `json:"use"`
}

func handleApiKey(ctx context.Context, apiKey string) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {
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

func handleAuthorizerToken(ctx context.Context, authorizerToken string) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {

	token, err := jwt.Parse(authorizerToken, func(token *jwt.Token) (interface{}, error) {
		// return the public key for the jwt
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		for _, key := range jwks.Keys {
			if key.Kid == token.Header["kid"] {
				nBytes, err := base64.RawURLEncoding.DecodeString(key.N)
				if err != nil {
					return nil, err
				}
				n := big.NewInt(0)
				n.SetBytes(nBytes)

				eBytes, err := base64.RawURLEncoding.DecodeString(key.E)
				if err != nil {
					return nil, err
				}
				var eBytesPadded []byte
				if len(eBytes) < 8 {
					eBytesPadded = make([]byte, 8-len(eBytes), 8)
					eBytesPadded = append(eBytesPadded, eBytes...)
				} else {
					eBytesPadded = eBytes
				}
				eReader := bytes.NewReader(eBytesPadded)
				var e uint64
				err = binary.Read(eReader, binary.BigEndian, &e)
				if err != nil {
					return nil, err
				}

				rsaPublicKey := rsa.PublicKey{N: n, E: int(e)}
				return &rsaPublicKey, nil
			}
		}

		return nil, fmt.Errorf("Unable to find key for kid: %v", token.Header["kid"])
	})
	if err != nil {
		logrus.Error("failed to parse token", err)
		return events.APIGatewayV2CustomAuthorizerSimpleResponse{
			IsAuthorized: false,
			Context: map[string]interface{}{
				"message": "Failed to parse token",
			},
		}, nil
	}

	if token.Valid {
		logrus.WithFields(logrus.Fields{
			"claims": token.Claims,
		}).Info("valid token")
		return events.APIGatewayV2CustomAuthorizerSimpleResponse{
			IsAuthorized: true,
			Context:      map[string]interface{}{},
		}, nil
	} else {
		logrus.Error("invalid token")
		return events.APIGatewayV2CustomAuthorizerSimpleResponse{
			IsAuthorized: false,
			Context: map[string]interface{}{
				"message": "Invalid token",
			},
		}, nil
	}
}

func handleRequest(ctx context.Context, event events.APIGatewayCustomAuthorizerRequestTypeRequest) (events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {
	authorizerToken, ok := event.Headers["authorization"]
	if !ok {
		return events.APIGatewayV2CustomAuthorizerSimpleResponse{
			IsAuthorized: false,
			Context: map[string]interface{}{
				"message": "Authorization header not found",
			},
		}, nil
	}

	if strings.HasPrefix(authorizerToken, "Bearer ") {
		// remove the bearer prefix
		authorizerToken = strings.TrimPrefix(authorizerToken, "Bearer ")
		return handleAuthorizerToken(ctx, authorizerToken)
	} else {
		return handleApiKey(ctx, authorizerToken)
	}
}

func main() {
	lambda.Start(handleRequest)
}
