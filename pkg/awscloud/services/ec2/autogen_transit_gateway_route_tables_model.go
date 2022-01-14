//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"time"
)

type TransitGatewayRouteTable struct {
	CreationTime *time.Time 
	DefaultAssociationRouteTable bool `parquet:"name=default_association_route_table,type=BOOLEAN"`
	DefaultPropagationRouteTable bool `parquet:"name=default_propagation_route_table,type=BOOLEAN"`
	State string `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
	TagsOld []*Tag `parquet:"name=tags_old,type=MAP,convertedtype=LIST"`
	TransitGatewayId string `parquet:"name=transit_gateway_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	TransitGatewayRouteTableId string `parquet:"name=transit_gateway_route_table_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	AccountId string `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region string `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime int64 `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	CreationTimeMilli int64 `parquet:"name=creation_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	Tags map[string]string `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
}
