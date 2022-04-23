package dynamodb

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/expression"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
	"github.com/sirupsen/logrus"
)

func GetItem(ctx context.Context, client *dynamodb.Client, tableName string, reportTimeUnixMilli int64, id string, idAttributeName string) (map[string]types.AttributeValue, error) {
	keyCondition := expression.Key(idAttributeName).Equal(expression.Value(id))
	keyCondition = keyCondition.And(expression.Key("report_time").Equal(expression.Value(reportTimeUnixMilli)))
	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
	if err != nil {
		return nil, err
	}
	fmt.Println(id)

	resp, err := client.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		IndexName:                 aws.String("id_report_time_index"),
	})

	if err != nil {
		return nil, err
	}

	if len(resp.Items) == 0 {
		return nil, fmt.Errorf("Item not found")
	}
	if len(resp.Items) > 1 {
		return nil, fmt.Errorf("Multiple items found")
	}

	return resp.Items[0], nil
}

func ListItems(ctx context.Context, client *dynamodb.Client, tableName string, reportTimeUnixMilli int64, idAttributeName string, lastKey *string) ([]map[string]types.AttributeValue, *string, error) {
	keyCondition := expression.Key("report_time").Equal(expression.Value(reportTimeUnixMilli))
	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
	if err != nil {
		return nil, nil, err
	}

	var lastEvaluatedKey map[string]types.AttributeValue
	if lastKey != nil {
		lastEvaluatedKey = map[string]types.AttributeValue{
			"report_time": &types.AttributeValueMemberN{
				Value: strconv.FormatInt(reportTimeUnixMilli, 10),
			},
			idAttributeName: &types.AttributeValueMemberS{
				Value: *lastKey,
			},
		}
	}

	resp, err := client.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String(tableName),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		IndexName:                 aws.String("report_time_id_index"),
		ExclusiveStartKey:         lastEvaluatedKey,
	})
	if err != nil {
		return nil, nil, err
	}

	var nextKey *string
	if resp.LastEvaluatedKey != nil {
		tmp := resp.LastEvaluatedKey[idAttributeName].(*types.AttributeValueMemberS).Value
		nextKey = &tmp
	}
	return resp.Items, nextKey, nil
}

func WriteItem(ctx context.Context, client *dynamodb.Client, maxRetries int, tableName string, item interface{}) error {
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	for retry := 0; retry < maxRetries; retry++ {
		_, err = client.PutItem(ctx, &dynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item:      av,
		})
		if err == nil {
			return nil
		}
	}

	return err
}

func BatchWriteItems(ctx context.Context, client *dynamodb.Client, maxRetries int, tableName string, items []map[string]types.AttributeValue) error {
	for i := 0; i < len(items); i += 25 {
		end := i + 25
		if end > len(items) {
			end = len(items)
		}

		batchWriteInput := &dynamodb.BatchWriteItemInput{
			RequestItems: map[string][]types.WriteRequest{
				tableName: make([]types.WriteRequest, end-i),
			},
		}

		for j := i; j < end; j++ {
			item := items[j]
			batchWriteInput.RequestItems[tableName][j-i] = types.WriteRequest{
				PutRequest: &types.PutRequest{
					Item: item,
				},
			}
		}

		for retry := 0; retry < maxRetries; retry++ {
			resp, err := client.BatchWriteItem(ctx, batchWriteInput)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"error":      err,
					"table_name": tableName,
				}).Error("Failed to batch write items")
				return err
			}
			if resp.UnprocessedItems == nil {
				break
			}
			if resp.UnprocessedItems[tableName] == nil {
				break
			}
			batchWriteInput.RequestItems[tableName] = resp.UnprocessedItems[tableName]
		}
	}

	return nil
}

func DistinctReportTimes(ctx context.Context, client *dynamodb.Client, reportDateUnixMilli int64, cloud, service, resource string) ([]int64, error) {
	reportDate := time.UnixMilli(reportDateUnixMilli).UTC()
	reportDate = time.Date(reportDate.Year(), reportDate.Month(), reportDate.Day(), 0, 0, 0, 0, time.UTC)
	lowerTime := expression.Value(reportDate.UTC().UnixMilli())
	upperTime := expression.Value(reportDate.UTC().AddDate(0, 0, 1).UnixMilli())
	keyCondition := expression.Key("ingestion_key").Equal(expression.Value(cloud + ":" + service + ":" + resource))
	keyCondition = keyCondition.And(expression.Key("report_time").Between(lowerTime, upperTime))
	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
	if err != nil {
		return nil, err
	}

	reportTimes := []int64{}
	var lastEvaluatedKey map[string]types.AttributeValue

	for {
		resp, err := client.Query(ctx, &dynamodb.QueryInput{
			TableName:                 aws.String("cloud-inventory-ingestion-timestamps"),
			KeyConditionExpression:    expr.KeyCondition(),
			ProjectionExpression:      expr.Projection(),
			ExpressionAttributeNames:  expr.Names(),
			ExpressionAttributeValues: expr.Values(),
			ScanIndexForward:          aws.Bool(true),
			ExclusiveStartKey:         lastEvaluatedKey,
		})

		if err != nil {
			return nil, err
		}

		var ingestions []meta.IngestionTimestamp
		err = attributevalue.UnmarshalListOfMaps(resp.Items, &ingestions)
		if err != nil {
			return nil, err
		}
		for _, ingestion := range ingestions {
			reportTimes = append(reportTimes, ingestion.ReportTime)
		}

		if resp.LastEvaluatedKey != nil {
			lastEvaluatedKey = resp.LastEvaluatedKey
		} else {
			break
		}
	}

	return reportTimes, nil
}

func GetReportTime(ctx context.Context, client *dynamodb.Client, reportDateUnixMilli int64, timeSelection db.TimeSelection, timeReferenceUnixMilli int64, cloud, service, resource string) (*int64, error) {
	if timeSelection == db.TimeSelectionAt {
		return &timeReferenceUnixMilli, nil
	}
	reportDate := time.UnixMilli(reportDateUnixMilli).UTC()
	reportDate = time.Date(reportDate.Year(), reportDate.Month(), reportDate.Day(), 0, 0, 0, 0, time.UTC)

	keyCondition := expression.Key("ingestion_key").Equal(expression.Value(cloud + ":" + service + ":" + resource))

	switch timeSelection {
	case db.TimeSelectionLatest:
		keyCondition = keyCondition.And(expression.Key("report_time").LessThan(expression.Value(reportDate.AddDate(0, 0, 1).UTC().UnixMilli())))
	case db.TimeSelectionBefore:
		keyCondition = keyCondition.And(expression.Key("report_time").LessThan(expression.Value(timeReferenceUnixMilli)))
	case db.TimeSelectionAfter:
		keyCondition = keyCondition.And(expression.Key("report_time").GreaterThan(expression.Value(timeReferenceUnixMilli)))
	}

	expr, err := expression.NewBuilder().WithKeyCondition(keyCondition).Build()
	if err != nil {
		return nil, err
	}

	resp, err := client.Query(ctx, &dynamodb.QueryInput{
		TableName:                 aws.String("cloud-inventory-ingestion-timestamps"),
		KeyConditionExpression:    expr.KeyCondition(),
		ProjectionExpression:      expr.Projection(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		Limit:                     aws.Int32(1),
		ScanIndexForward:          aws.Bool(false),
	})

	if err != nil {
		return nil, err
	}

	if len(resp.Items) == 0 {
		return nil, fmt.Errorf("No report time found")
	}

	var latestIngestion meta.IngestionTimestamp
	err = attributevalue.UnmarshalMap(resp.Items[0], &latestIngestion)
	if err != nil {
		return nil, err
	}

	return &latestIngestion.ReportTime, nil
}
