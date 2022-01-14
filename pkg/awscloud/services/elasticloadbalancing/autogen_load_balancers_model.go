//AUTOGENERATED CODE DO NOT EDIT
package elasticloadbalancing

import (
	"time"
)

type LoadBalancerDescription struct {
	AvailabilityZones []string `parquet:"name=availability_zones,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	BackendServerDescriptions []*BackendServerDescription `parquet:"name=backend_server_descriptions,type=MAP,convertedtype=LIST"`
	CanonicalHostedZoneName string `parquet:"name=canonical_hosted_zone_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	CanonicalHostedZoneNameID string `parquet:"name=canonical_hosted_zone_name_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	CreatedTime *time.Time 
	DNSName string `parquet:"name=dns_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	HealthCheck *HealthCheck `parquet:"name=health_check"`
	Instances []*Instance `parquet:"name=instances,type=MAP,convertedtype=LIST"`
	ListenerDescriptions []*ListenerDescription `parquet:"name=listener_descriptions,type=MAP,convertedtype=LIST"`
	LoadBalancerName string `parquet:"name=load_balancer_name,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	Policies *Policies `parquet:"name=policies"`
	Scheme string `parquet:"name=scheme,type=BYTE_ARRAY,convertedtype=UTF8"`
	SecurityGroups []string `parquet:"name=security_groups,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	SourceSecurityGroup *SourceSecurityGroup `parquet:"name=source_security_group"`
	Subnets []string `parquet:"name=subnets,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	VPCId string `parquet:"name=vpc_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	AccountId string `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region string `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime int64 `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	CreatedTimeMilli int64 `parquet:"name=created_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

