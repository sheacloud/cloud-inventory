//AUTOGENERATED CODE DO NOT EDIT
package ec2



type NetworkInterface struct {
	Association *NetworkInterfaceAssociation `parquet:"name=association"`
	Attachment *NetworkInterfaceAttachment `parquet:"name=attachment"`
	AvailabilityZone string `parquet:"name=availability_zone,type=BYTE_ARRAY,convertedtype=UTF8"`
	DenyAllIgwTraffic bool `parquet:"name=deny_all_igw_traffic,type=BOOLEAN"`
	Description string `parquet:"name=description,type=BYTE_ARRAY,convertedtype=UTF8"`
	Groups []*GroupIdentifier `parquet:"name=groups,type=MAP,convertedtype=LIST"`
	InterfaceType string `parquet:"name=interface_type,type=BYTE_ARRAY,convertedtype=UTF8"`
	Ipv4Prefixes []*Ipv4PrefixSpecification `parquet:"name=ipv4_prefixes,type=MAP,convertedtype=LIST"`
	Ipv6Address string `parquet:"name=ipv6_address,type=BYTE_ARRAY,convertedtype=UTF8"`
	Ipv6Addresses []*NetworkInterfaceIpv6Address `parquet:"name=ipv6_addresses,type=MAP,convertedtype=LIST"`
	Ipv6Native bool `parquet:"name=ipv6_native,type=BOOLEAN"`
	Ipv6Prefixes []*Ipv6PrefixSpecification `parquet:"name=ipv6_prefixes,type=MAP,convertedtype=LIST"`
	MacAddress string `parquet:"name=mac_address,type=BYTE_ARRAY,convertedtype=UTF8"`
	NetworkInterfaceId string `parquet:"name=network_interface_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	OutpostArn string `parquet:"name=outpost_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	OwnerId string `parquet:"name=owner_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	PrivateDnsName string `parquet:"name=private_dns_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	PrivateIpAddress string `parquet:"name=private_ip_address,type=BYTE_ARRAY,convertedtype=UTF8"`
	PrivateIpAddresses []*NetworkInterfacePrivateIpAddress `parquet:"name=private_ip_addresses,type=MAP,convertedtype=LIST"`
	RequesterId string `parquet:"name=requester_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	RequesterManaged bool `parquet:"name=requester_managed,type=BOOLEAN"`
	SourceDestCheck bool `parquet:"name=source_dest_check,type=BOOLEAN"`
	Status string `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
	SubnetId string `parquet:"name=subnet_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	TagSetOld []*Tag `parquet:"name=tag_set_old,type=MAP,convertedtype=LIST"`
	VpcId string `parquet:"name=vpc_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	AccountId string `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region string `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime int64 `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	Tags map[string]string `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
}
