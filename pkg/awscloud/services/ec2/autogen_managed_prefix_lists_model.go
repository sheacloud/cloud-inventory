//AUTOGENERATED CODE DO NOT EDIT
package ec2

type ManagedPrefixList struct {
	AddressFamily  string            `parquet:"name=address_family,type=BYTE_ARRAY,convertedtype=UTF8"`
	MaxEntries     int32             `parquet:"name=max_entries,type=INT32"`
	OwnerId        string            `parquet:"name=owner_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	PrefixListArn  string            `parquet:"name=prefix_list_arn,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	PrefixListId   string            `parquet:"name=prefix_list_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	PrefixListName string            `parquet:"name=prefix_list_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	State          string            `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
	StateMessage   string            `parquet:"name=state_message,type=BYTE_ARRAY,convertedtype=UTF8"`
	TagsOld        []*Tag            `parquet:"name=tags_old,type=MAP,convertedtype=LIST"`
	Version        int64             `parquet:"name=version,type=INT64"`
	AccountId      string            `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region         string            `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime     int64             `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	Tags           map[string]string `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
}
