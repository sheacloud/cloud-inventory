package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DistinctReportTimes(ctx context.Context, coll *mongo.Collection, reportDateUnixMilli int64) ([]int64, error) {
	reportDate := time.UnixMilli(reportDateUnixMilli)
	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"report_time", bson.D{{"$gte", reportDateUnixMilli}}}},
				bson.D{{"report_time", bson.D{{"$lt", reportDate.AddDate(0, 0, 1).UnixMilli()}}}},
			},
		},
	}
	results, err := coll.Distinct(ctx, "report_time", filter)
	if err != nil {
		return nil, err
	}
	reportTimes := make([]int64, len(results))
	for i, result := range results {
		reportTimes[i] = result.(int64)
	}

	return reportTimes, nil
}

func GetReportTime(ctx context.Context, coll *mongo.Collection, reportDateUnixMilli int64, timeSelection db.TimeSelection, timeReferenceUnixMilli int64) (*int64, error) {
	reportDate := time.UnixMilli(reportDateUnixMilli)
	var filter bson.D
	switch timeSelection {
	case db.TimeSelectionLatest:
		filter = bson.D{
			{"$and",
				bson.A{
					bson.D{{"report_time", bson.D{{"$gte", reportDateUnixMilli}}}},
					bson.D{{"report_time", bson.D{{"$lt", reportDate.AddDate(0, 0, 1).UnixMilli()}}}},
				},
			},
		}
	case db.TimeSelectionBefore:
		filter = bson.D{{"report_time", bson.D{{"$lt", timeReferenceUnixMilli}}}}
	case db.TimeSelectionAfter:
		filter = bson.D{{"report_time", bson.D{{"$gt", timeReferenceUnixMilli}}}}
	case db.TimeSelectionAt:
		return &timeReferenceUnixMilli, nil
	}

	opts := options.FindOne().SetProjection(bson.D{{"report_time", 1}, {"_id", 0}}).SetSort(bson.D{{"report_time", -1}})

	var result struct {
		ReportTime int64 `bson:"report_time"`
	}
	err := coll.FindOne(ctx, filter, opts).Decode(&result)
	if err != nil {
		fmt.Println("failed to run query")
		return nil, err
	}

	return &result.ReportTime, nil
}
