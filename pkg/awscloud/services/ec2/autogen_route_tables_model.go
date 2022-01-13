//AUTOGENERATED CODE DO NOT EDIT
package ec2

type RouteTable struct {
	Associations    []*RouteTableAssociation `parquet:"name=associations,type=MAP,convertedtype=LIST"`
	OwnerId         string                   `parquet:"name=owner_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	PropagatingVgws []*PropagatingVgw        `parquet:"name=propagating_vgws,type=MAP,convertedtype=LIST"`
	RouteTableId    string                   `parquet:"name=route_table_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	Routes          []*Route                 `parquet:"name=routes,type=MAP,convertedtype=LIST"`
	TagsOld         []*Tag                   `parquet:"name=tags_old,type=MAP,convertedtype=LIST"`
	VpcId           string                   `parquet:"name=vpc_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	AccountId       string                   `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region          string                   `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime      int64                    `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	Tags            map[string]string        `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
}
