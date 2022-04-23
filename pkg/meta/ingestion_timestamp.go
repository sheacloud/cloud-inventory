package meta

type IngestionTimestamp struct {
	Key        string `bson:"ingestion_key" dynamodbav:"ingestion_key"`
	ReportTime int64  `bson:"report_time" dynamodbav:"report_time"`
}
