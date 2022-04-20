//AUTOGENERATED CODE DO NOT EDIT
package dynamodb

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	dynamo "github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws/cloudtrail"
	"time"
)

type DynamoDBCloudTrailDAO struct {
	client     *dynamo.Client
	maxRetries int
}

func (dao *DynamoDBCloudTrailDAO) ListTrails(ctx context.Context, reportTime time.Time, accountID, region *string, limit, offset *int64) ([]*cloudtrail.Trail, error) {
	tableName := "cloud-inventory-aws-cloudtrail-trails"
	items, _, err := ListItems(ctx, dao.client, tableName, reportTime, "trail_arn", nil)
	if err != nil {
		return nil, err
	}
	var resources []*cloudtrail.Trail
	err = attributevalue.UnmarshalListOfMaps(items, &resources)
	if err != nil {
		return nil, err
	}
	return resources, nil
}

func (dao *DynamoDBCloudTrailDAO) GetTrail(ctx context.Context, reportTime time.Time, id string) (*cloudtrail.Trail, error) {
	tableName := "cloud-inventory-aws-cloudtrail-trails"
	item, err := GetItem(ctx, dao.client, tableName, reportTime, id, "trail_arn")
	if err != nil {
		return nil, err
	}
	var resource *cloudtrail.Trail
	err = attributevalue.UnmarshalMap(item, &resource)
	if err != nil {
		return nil, err
	}
	return resource, nil
}

func (dao *DynamoDBCloudTrailDAO) PutTrails(ctx context.Context, resources []*cloudtrail.Trail) error {
	items := make([]map[string]types.AttributeValue, len(resources))
	for i, resource := range resources {
		item, err := attributevalue.MarshalMap(resource)
		if err != nil {
			return err
		}
		items[i] = item
	}
	return BatchWriteItems(ctx, dao.client, dao.maxRetries, "cloud-inventory-aws-cloudtrail-trails", items)
}

func (dao *DynamoDBCloudTrailDAO) GetTrailReportTimes(ctx context.Context, reportDate time.Time) ([]string, error) {
	return DistinctReportTimes(ctx, dao.client, reportDate, "aws", "cloudtrail", "trails")
}

func (dao *DynamoDBCloudTrailDAO) GetReferencedTrailReportTime(ctx context.Context, reportDate time.Time, timeSelection db.TimeSelection, timeReference time.Time) (*time.Time, error) {
	return GetReportTime(ctx, dao.client, reportDate, timeSelection, timeReference, "aws", "cloudtrail", "trails")
}
