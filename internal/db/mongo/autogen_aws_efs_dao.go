//AUTOGENERATED CODE DO NOT EDIT
package mongo

import (
	"context"
	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws/efs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MongoEFSDAO struct {
	db *mongo.Database
}

func (dao *MongoEFSDAO) ListFileSystems(ctx context.Context, reportTime time.Time, accountID, region *string, limit, offset *int64) ([]*efs.FileSystem, error) {
	filter := bson.D{
		bson.E{"report_time", reportTime},
	}
	if accountID != nil {
		filter = append(filter, bson.E{Key: "account_id", Value: *accountID})
	}
	if region != nil {
		filter = append(filter, bson.E{Key: "region", Value: *region})
	}

	var results []*efs.FileSystem
	cursor, err := dao.db.Collection("aws_efs_file_systems").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (dao *MongoEFSDAO) GetFileSystem(ctx context.Context, reportTime time.Time, id string) (*efs.FileSystem, error) {
	filter := bson.D{
		bson.E{"report_time", reportTime},
		bson.E{"file_system_id", id},
	}

	var result *efs.FileSystem
	err := dao.db.Collection("aws_efs_file_systems").FindOne(ctx, filter).Decode(&result)

	return result, err
}

func (dao *MongoEFSDAO) PutFileSystems(ctx context.Context, resources []*efs.FileSystem) error {
	if len(resources) == 0 {
		return nil
	}
	writes := make([]interface{}, len(resources))
	for i, resource := range resources {
		writes[i] = resource
	}
	_, err := dao.db.Collection("aws.efs.filesystems").InsertMany(ctx, writes)

	return err
}

func (dao *MongoEFSDAO) GetFileSystemReportTimes(ctx context.Context, reportDate time.Time) ([]string, error) {
	return DistinctReportTimes(ctx, dao.db.Collection("aws_efs_file_systems"), reportDate)
}

func (dao *MongoEFSDAO) GetReferencedFileSystemReportTime(ctx context.Context, reportDate time.Time, timeSelection db.TimeSelection, timeReference time.Time) (*time.Time, error) {
	return GetReportTime(ctx, dao.db.Collection("aws_efs_file_systems"), reportDate, timeSelection, timeReference)
}
