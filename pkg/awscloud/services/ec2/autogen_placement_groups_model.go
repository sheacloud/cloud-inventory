//AUTOGENERATED CODE DO NOT EDIT
package ec2

type PlacementGroup struct {
	GroupId        string            `parquet:"name=group_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"group_id" diff:"group_id,identifier"`
	GroupName      string            `parquet:"name=group_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"group_name" diff:"group_name"`
	PartitionCount int32             `parquet:"name=partition_count,type=INT32" json:"partition_count" diff:"partition_count"`
	State          string            `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8" json:"state" diff:"state"`
	Strategy       string            `parquet:"name=strategy,type=BYTE_ARRAY,convertedtype=UTF8" json:"strategy" diff:"strategy"`
	Tags           map[string]string `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"tags" diff:"tags"`
	AccountId      string            `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"account_id" diff:"account_id"`
	Region         string            `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8" json:"region" diff:"region"`
	ReportTime     int64             `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"report_time" diff:"report_time,immutable"`
}

func (x *PlacementGroup) GetReportTime() int64 {
	return x.ReportTime
}
