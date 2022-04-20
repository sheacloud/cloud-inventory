//AUTOGENERATED CODE DO NOT EDIT
package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	dynamo "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws/sns"
	"time"
)

type DynamoDBSNSDAO struct {
	client     *dynamo.Client
	maxRetries int
}

func (dao *DynamoDBSNSDAO) ListTopics(ctx context.Context, reportTime time.Time, accountID, region *string, limit, offset *int64) ([]*sns.Topic, error) {
	tableName := "cloud-inventory-aws-sns-topics"
	items, _, err := ListItems(ctx, dao.client, tableName, reportTime, "topic_arn", nil)
	if err != nil {
		return nil, err
	}
	var resources []*sns.Topic
	err = attributevalue.UnmarshalListOfMaps(items, &resources)
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func (dao *DynamoDBSNSDAO) GetTopic(ctx context.Context, reportTime time.Time, id string) (*sns.Topic, error) {
	tableName := "cloud-inventory-aws-sns-topics"
	item, err := GetItem(ctx, dao.client, tableName, reportTime, id, "topic_arn")
	if err != nil {
		return nil, err
	}
	var resource *sns.Topic
	err = attributevalue.UnmarshalMap(item, &resource)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func (dao *DynamoDBSNSDAO) PutTopics(ctx context.Context, resources []*sns.Topic) error {
	items := make([]map[string]types.AttributeValue, len(resources))
	for i, resource := range resources {
		item, err := attributevalue.MarshalMap(resource)
		if err != nil {
			return err
		}
		items[i] = item
	}
	return BatchWriteItems(ctx, dao.client, dao.maxRetries, "cloud-inventory-aws-sns-topics", items)
}

func (dao *DynamoDBSNSDAO) GetTopicReportTimes(ctx context.Context, reportDate time.Time) ([]string, error) {
	return DistinctReportTimes(ctx, dao.client, reportDate, "aws", "sns", "topics")
}

func (dao *DynamoDBSNSDAO) GetReferencedTopicReportTime(ctx context.Context, reportDate time.Time, timeSelection db.TimeSelection, timeReference time.Time) (*time.Time, error) {
	return GetReportTime(ctx, dao.client, reportDate, timeSelection, timeReference, "aws", "sns", "topics")
}

func (dao *DynamoDBSNSDAO) ListSubscriptions(ctx context.Context, reportTime time.Time, accountID, region *string, limit, offset *int64) ([]*sns.Subscription, error) {
	tableName := "cloud-inventory-aws-sns-subscriptions"
	items, _, err := ListItems(ctx, dao.client, tableName, reportTime, "subscription_arn", nil)
	if err != nil {
		return nil, err
	}
	var resources []*sns.Subscription
	err = attributevalue.UnmarshalListOfMaps(items, &resources)
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func (dao *DynamoDBSNSDAO) GetSubscription(ctx context.Context, reportTime time.Time, id string) (*sns.Subscription, error) {
	tableName := "cloud-inventory-aws-sns-subscriptions"
	item, err := GetItem(ctx, dao.client, tableName, reportTime, id, "subscription_arn")
	if err != nil {
		return nil, err
	}
	var resource *sns.Subscription
	err = attributevalue.UnmarshalMap(item, &resource)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func (dao *DynamoDBSNSDAO) PutSubscriptions(ctx context.Context, resources []*sns.Subscription) error {
	items := make([]map[string]types.AttributeValue, len(resources))
	for i, resource := range resources {
		item, err := attributevalue.MarshalMap(resource)
		if err != nil {
			return err
		}
		items[i] = item
	}
	return BatchWriteItems(ctx, dao.client, dao.maxRetries, "cloud-inventory-aws-sns-subscriptions", items)
}

func (dao *DynamoDBSNSDAO) GetSubscriptionReportTimes(ctx context.Context, reportDate time.Time) ([]string, error) {
	return DistinctReportTimes(ctx, dao.client, reportDate, "aws", "sns", "subscriptions")
}

func (dao *DynamoDBSNSDAO) GetReferencedSubscriptionReportTime(ctx context.Context, reportDate time.Time, timeSelection db.TimeSelection, timeReference time.Time) (*time.Time, error) {
	return GetReportTime(ctx, dao.client, reportDate, timeSelection, timeReference, "aws", "sns", "subscriptions")
}