//AUTOGENERATED CODE DO NOT EDIT
package ec2



type PlacementGroup struct {
	GroupId string `parquet:"name=group_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	GroupName string `parquet:"name=group_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	PartitionCount int32 `parquet:"name=partition_count,type=INT32"`
	State string `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
	Strategy string `parquet:"name=strategy,type=BYTE_ARRAY,convertedtype=UTF8"`
	TagsOld []*Tag `parquet:"name=tags_old,type=MAP,convertedtype=LIST"`
	AccountId string `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region string `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime int64 `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	Tags map[string]string `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
}
