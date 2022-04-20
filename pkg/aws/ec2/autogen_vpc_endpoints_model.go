//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"time"
)

type VpcEndpoint struct {
	CreationTimestamp   *time.Time                 `bson:"creation_timestamp,omitempty" dynamodbav:"creation_timestamp,unixtime,omitempty" json:"creation_timestamp,omitempty" diff:"creation_timestamp"`
	DnsEntries          []*DnsEntry                `bson:"dns_entries,omitempty" dynamodbav:"dns_entries,omitempty" json:"dns_entries,omitempty" diff:"dns_entries"`
	Groups              []*SecurityGroupIdentifier `bson:"groups,omitempty" dynamodbav:"groups,omitempty" json:"groups,omitempty" diff:"groups"`
	LastError           *LastError                 `bson:"last_error,omitempty" dynamodbav:"last_error,omitempty" json:"last_error,omitempty" diff:"last_error"`
	NetworkInterfaceIds []string                   `bson:"network_interface_ids,omitempty" dynamodbav:"network_interface_ids,omitempty" json:"network_interface_ids,omitempty" diff:"network_interface_ids"`
	OwnerId             string                     `bson:"owner_id,omitempty" dynamodbav:"owner_id,omitempty" json:"owner_id,omitempty" diff:"owner_id"`
	PolicyDocument      string                     `bson:"policy_document,omitempty" dynamodbav:"policy_document,omitempty" json:"policy_document,omitempty" diff:"policy_document"`
	PrivateDnsEnabled   bool                       `bson:"private_dns_enabled,omitempty" dynamodbav:"private_dns_enabled" json:"private_dns_enabled,omitempty" diff:"private_dns_enabled"`
	RequesterManaged    bool                       `bson:"requester_managed,omitempty" dynamodbav:"requester_managed" json:"requester_managed,omitempty" diff:"requester_managed"`
	RouteTableIds       []string                   `bson:"route_table_ids,omitempty" dynamodbav:"route_table_ids,omitempty" json:"route_table_ids,omitempty" diff:"route_table_ids"`
	ServiceName         string                     `bson:"service_name,omitempty" dynamodbav:"service_name,omitempty" json:"service_name,omitempty" diff:"service_name"`
	State               string                     `bson:"state,omitempty" dynamodbav:"state,omitempty" json:"state,omitempty" diff:"state"`
	SubnetIds           []string                   `bson:"subnet_ids,omitempty" dynamodbav:"subnet_ids,omitempty" json:"subnet_ids,omitempty" diff:"subnet_ids"`
	Tags                map[string]string          `bson:"tags,omitempty" dynamodbav:"tags,omitempty" json:"tags,omitempty" diff:"tags"`
	VpcEndpointId       string                     `bson:"vpc_endpoint_id,omitempty" dynamodbav:"vpc_endpoint_id,omitempty" inventory_primary_key:"true" json:"vpc_endpoint_id,omitempty" diff:"vpc_endpoint_id,identifier"`
	VpcEndpointType     string                     `bson:"vpc_endpoint_type,omitempty" dynamodbav:"vpc_endpoint_type,omitempty" json:"vpc_endpoint_type,omitempty" diff:"vpc_endpoint_type"`
	VpcId               string                     `bson:"vpc_id,omitempty" dynamodbav:"vpc_id,omitempty" json:"vpc_id,omitempty" diff:"vpc_id"`
	AccountId           string                     `bson:"account_id,omitempty" dynamodbav:"account_id,omitempty" json:"account_id,omitempty" diff:"account_id"`
	Region              string                     `bson:"region,omitempty" dynamodbav:"region,omitempty" json:"region,omitempty" diff:"region"`
	ReportTime          time.Time                  `bson:"report_time,omitempty" dynamodbav:"report_time,unixtime,omitempty" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID       string                     `bson:"_id,omitempty" dynamodbav:"_id,omitempty" json:"_id,omitempty" diff:"-"`
}