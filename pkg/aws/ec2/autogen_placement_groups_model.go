//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"time"
)

type PlacementGroup struct {
	GroupId        string            `bson:"group_id,omitempty" dynamodbav:"group_id,omitempty" inventory_primary_key:"true" json:"group_id,omitempty" diff:"group_id,identifier"`
	GroupName      string            `bson:"group_name,omitempty" dynamodbav:"group_name,omitempty" json:"group_name,omitempty" diff:"group_name"`
	PartitionCount int32             `bson:"partition_count,omitempty" dynamodbav:"partition_count,omitempty" json:"partition_count,omitempty" diff:"partition_count"`
	State          string            `bson:"state,omitempty" dynamodbav:"state,omitempty" json:"state,omitempty" diff:"state"`
	Strategy       string            `bson:"strategy,omitempty" dynamodbav:"strategy,omitempty" json:"strategy,omitempty" diff:"strategy"`
	Tags           map[string]string `bson:"tags,omitempty" dynamodbav:"tags,omitempty" json:"tags,omitempty" diff:"tags"`
	AccountId      string            `bson:"account_id,omitempty" dynamodbav:"account_id,omitempty" json:"account_id,omitempty" diff:"account_id"`
	Region         string            `bson:"region,omitempty" dynamodbav:"region,omitempty" json:"region,omitempty" diff:"region"`
	ReportTime     time.Time         `bson:"report_time,omitempty" dynamodbav:"report_time,unixtime,omitempty" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID  string            `bson:"_id,omitempty" dynamodbav:"_id,omitempty" json:"_id,omitempty" diff:"-"`
}