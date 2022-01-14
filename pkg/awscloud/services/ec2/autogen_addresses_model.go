//AUTOGENERATED CODE DO NOT EDIT
package ec2



type Address struct {
	AllocationId string `parquet:"name=allocation_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	AssociationId string `parquet:"name=association_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	CarrierIp string `parquet:"name=carrier_ip,type=BYTE_ARRAY,convertedtype=UTF8"`
	CustomerOwnedIp string `parquet:"name=customer_owned_ip,type=BYTE_ARRAY,convertedtype=UTF8"`
	CustomerOwnedIpv4Pool string `parquet:"name=customer_owned_ipv4_pool,type=BYTE_ARRAY,convertedtype=UTF8"`
	Domain string `parquet:"name=domain,type=BYTE_ARRAY,convertedtype=UTF8"`
	InstanceId string `parquet:"name=instance_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	NetworkBorderGroup string `parquet:"name=network_border_group,type=BYTE_ARRAY,convertedtype=UTF8"`
	NetworkInterfaceId string `parquet:"name=network_interface_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	NetworkInterfaceOwnerId string `parquet:"name=network_interface_owner_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	PrivateIpAddress string `parquet:"name=private_ip_address,type=BYTE_ARRAY,convertedtype=UTF8"`
	PublicIp string `parquet:"name=public_ip,type=BYTE_ARRAY,convertedtype=UTF8"`
	PublicIpv4Pool string `parquet:"name=public_ipv4_pool,type=BYTE_ARRAY,convertedtype=UTF8"`
	TagsOld []*Tag `parquet:"name=tags_old,type=MAP,convertedtype=LIST"`
	AccountId string `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region string `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime int64 `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	Tags map[string]string `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
}

