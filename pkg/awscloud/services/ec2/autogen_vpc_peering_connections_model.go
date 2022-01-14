//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"time"
)

type VpcPeeringConnection struct {
	AccepterVpcInfo *VpcPeeringConnectionVpcInfo `parquet:"name=accepter_vpc_info"`
	ExpirationTime *time.Time 
	RequesterVpcInfo *VpcPeeringConnectionVpcInfo `parquet:"name=requester_vpc_info"`
	Status *VpcPeeringConnectionStateReason `parquet:"name=status"`
	TagsOld []*Tag `parquet:"name=tags_old,type=MAP,convertedtype=LIST"`
	VpcPeeringConnectionId string `parquet:"name=vpc_peering_connection_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	AccountId string `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region string `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime int64 `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	ExpirationTimeMilli int64 `parquet:"name=expiration_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	Tags map[string]string `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
}

