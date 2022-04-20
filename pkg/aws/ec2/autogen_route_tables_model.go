//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"time"
)

type RouteTable struct {
	Associations    []*RouteTableAssociation `bson:"associations,omitempty" dynamodbav:"associations,omitempty" json:"associations,omitempty" diff:"associations"`
	OwnerId         string                   `bson:"owner_id,omitempty" dynamodbav:"owner_id,omitempty" json:"owner_id,omitempty" diff:"owner_id"`
	PropagatingVgws []*PropagatingVgw        `bson:"propagating_vgws,omitempty" dynamodbav:"propagating_vgws,omitempty" json:"propagating_vgws,omitempty" diff:"propagating_vgws"`
	RouteTableId    string                   `bson:"route_table_id,omitempty" dynamodbav:"route_table_id,omitempty" inventory_primary_key:"true" json:"route_table_id,omitempty" diff:"route_table_id,identifier"`
	Routes          []*Route                 `bson:"routes,omitempty" dynamodbav:"routes,omitempty" json:"routes,omitempty" diff:"routes"`
	Tags            map[string]string        `bson:"tags,omitempty" dynamodbav:"tags,omitempty" json:"tags,omitempty" diff:"tags"`
	VpcId           string                   `bson:"vpc_id,omitempty" dynamodbav:"vpc_id,omitempty" json:"vpc_id,omitempty" diff:"vpc_id"`
	AccountId       string                   `bson:"account_id,omitempty" dynamodbav:"account_id,omitempty" json:"account_id,omitempty" diff:"account_id"`
	Region          string                   `bson:"region,omitempty" dynamodbav:"region,omitempty" json:"region,omitempty" diff:"region"`
	ReportTime      time.Time                `bson:"report_time,omitempty" dynamodbav:"report_time,unixtime,omitempty" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID   string                   `bson:"_id,omitempty" dynamodbav:"_id,omitempty" json:"_id,omitempty" diff:"-"`
}
