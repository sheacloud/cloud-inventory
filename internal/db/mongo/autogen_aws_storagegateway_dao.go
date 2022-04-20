//AUTOGENERATED CODE DO NOT EDIT
package mongo

import (
	"context"
	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws/storagegateway"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoStorageGatewayDAO struct {
	db *mongo.Database
}

func (dao *MongoStorageGatewayDAO) ListGateways(ctx context.Context, reportTime time.Time, accountID, region *string, limit, offset *int64) ([]*storagegateway.Gateway, error) {
	filter := bson.D{
		bson.E{"report_time", reportTime},
	}
	if accountID != nil {
		filter = append(filter, bson.E{Key: "account_id", Value: *accountID})
	}
	if region != nil {
		filter = append(filter, bson.E{Key: "region", Value: *region})
	}

	var results []*storagegateway.Gateway
	cursor, err := dao.db.Collection("aws_storagegateway_gateways").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (dao *MongoStorageGatewayDAO) GetGateway(ctx context.Context, reportTime time.Time, id string) (*storagegateway.Gateway, error) {
	filter := bson.D{
		bson.E{"report_time", reportTime},
		bson.E{"gateway_arn", id},
	}

	var result *storagegateway.Gateway
	err := dao.db.Collection("aws_storagegateway_gateways").FindOne(ctx, filter).Decode(&result)

	return result, err
}

func (dao *MongoStorageGatewayDAO) PutGateways(ctx context.Context, resources []*storagegateway.Gateway) error {
	if len(resources) == 0 {
		return nil
	}
	writes := make([]interface{}, len(resources))
	for i, resource := range resources {
		writes[i] = resource
	}
	_, err := dao.db.Collection("aws.storagegateway.gateways").InsertMany(ctx, writes)

	return err
}

func (dao *MongoStorageGatewayDAO) GetGatewayReportTimes(ctx context.Context, reportDate time.Time) ([]string, error) {
	return DistinctReportTimes(ctx, dao.db.Collection("aws_storagegateway_gateways"), reportDate)
}

func (dao *MongoStorageGatewayDAO) GetReferencedGatewayReportTime(ctx context.Context, reportDate time.Time, timeSelection db.TimeSelection, timeReference time.Time) (*time.Time, error) {
	return GetReportTime(ctx, dao.db.Collection("aws_storagegateway_gateways"), reportDate, timeSelection, timeReference)
}