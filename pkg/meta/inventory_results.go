package meta

type InventoryResults struct {
	InventoryUUID    string `bson:"_id" dynamodbav:"_id" parquet:"name=_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Cloud            string `bson:"cloud" dynamodbav:"cloud" parquet:"name=cloud,type=BYTE_ARRAY,convertedtype=UTF8"`
	Service          string `bson:"service" dynamodbav:"service" parquet:"name=service,type=BYTE_ARRAY,convertedtype=UTF8"`
	Resource         string `bson:"resource" dynamodbav:"resource" parquet:"name=resource,type=BYTE_ARRAY,convertedtype=UTF8"`
	AccountId        string `bson:"account_id" dynamodbav:"account_id" parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region           string `bson:"region" dynamodbav:"region" parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	FetchedResources int    `bson:"fetched_resources" dynamodbav:"fetched_resources" parquet:"name=fetched_resources,type=INT32"`
	FailedResources  int    `bson:"failed_resources" dynamodbav:"failed_resources" parquet:"name=failed_resources,type=INT32"`
	HadErrors        bool   `bson:"had_errors" dynamodbav:"had_errors" parquet:"name=had_errors,type=BOOLEAN"`
	ReportTime       int64  `bson:"report_time" dynamodbav:"report_time" parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}
