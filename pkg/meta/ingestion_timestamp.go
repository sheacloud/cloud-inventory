package meta

type IngestionTimestamp struct {
	Key        string `bson:"ingestion_key" dynamodbav:"ingestion_key" parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime int64  `bson:"report_time" dynamodbav:"report_time" parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}
