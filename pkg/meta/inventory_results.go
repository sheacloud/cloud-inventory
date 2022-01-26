package meta

type InventoryResults struct {
	Cloud            string `parquet:"name=cloud,type=BYTE_ARRAY,convertedtype=UTF8"`
	Service          string `parquet:"name=service,type=BYTE_ARRAY,convertedtype=UTF8"`
	Resource         string `parquet:"name=resource,type=BYTE_ARRAY,convertedtype=UTF8"`
	AccountId        string `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region           string `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	FetchedResources int    `parquet:"name=fetched_resources,type=INT64"`
	FailedResources  int    `parquet:"name=failed_resources,type=INT64"`
	HadErrors        bool   `parquet:"name=had_errors,type=BOOLEAN"`
	ReportTime       int64  `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}
