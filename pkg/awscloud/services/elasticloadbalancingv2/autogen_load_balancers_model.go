//AUTOGENERATED CODE DO NOT EDIT
package elasticloadbalancingv2

import (
	"time"
)

type LoadBalancer struct {
	AvailabilityZones []*AvailabilityZone `parquet:"name=availability_zones,type=MAP,convertedtype=LIST"`
	CanonicalHostedZoneId string `parquet:"name=canonical_hosted_zone_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	CreatedTime *time.Time 
	CustomerOwnedIpv4Pool string `parquet:"name=customer_owned_ipv4_pool,type=BYTE_ARRAY,convertedtype=UTF8"`
	DNSName string `parquet:"name=dns_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	IpAddressType string `parquet:"name=ip_address_type,type=BYTE_ARRAY,convertedtype=UTF8"`
	LoadBalancerArn string `parquet:"name=load_balancer_arn,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	LoadBalancerName string `parquet:"name=load_balancer_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	Scheme string `parquet:"name=scheme,type=BYTE_ARRAY,convertedtype=UTF8"`
	SecurityGroups []string `parquet:"name=security_groups,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	State *LoadBalancerState `parquet:"name=state"`
	Type string `parquet:"name=type,type=BYTE_ARRAY,convertedtype=UTF8"`
	VpcId string `parquet:"name=vpc_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	AccountId string `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region string `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime int64 `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	Tags map[string]string `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
	Listeners []*Listener `parquet:"name=listeners,type=MAP,convertedtype=LIST"`
	CreatedTimeMilli int64 `parquet:"name=created_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

