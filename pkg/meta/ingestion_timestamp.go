package meta

import "time"

type IngestionTimestamp struct {
	Key        string    `bson:"ingestion_key" dynamodbav:"ingestion_key"`
	ReportTime time.Time `bson:"report_time" dynamodbav:"report_time,unixtime"`
}
