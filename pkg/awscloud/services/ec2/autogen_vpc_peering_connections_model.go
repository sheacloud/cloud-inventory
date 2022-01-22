//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"time"
)

type VpcPeeringConnection struct {
	AccepterVpcInfo        *VpcPeeringConnectionVpcInfo `parquet:"name=accepter_vpc_info" json:"accepter_vpc_info" diff:"accepter_vpc_info"`
	ExpirationTime         *time.Time
	RequesterVpcInfo       *VpcPeeringConnectionVpcInfo     `parquet:"name=requester_vpc_info" json:"requester_vpc_info" diff:"requester_vpc_info"`
	Status                 *VpcPeeringConnectionStateReason `parquet:"name=status" json:"status" diff:"status"`
	Tags                   map[string]string                `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"tags" diff:"tags"`
	VpcPeeringConnectionId string                           `parquet:"name=vpc_peering_connection_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"vpc_peering_connection_id" diff:"vpc_peering_connection_id,identifier"`
	AccountId              string                           `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"account_id" diff:"account_id"`
	Region                 string                           `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8" json:"region" diff:"region"`
	ReportTime             int64                            `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"report_time" diff:"report_time,immutable"`
	ExpirationTimeMilli    int64                            `parquet:"name=expiration_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"expiration_time" diff:"expiration_time"`
}

func (x *VpcPeeringConnection) GetReportTime() int64 {
	return x.ReportTime
}
