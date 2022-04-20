//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"time"
)

type NetworkAcl struct {
	Associations  []*NetworkAclAssociation `bson:"associations,omitempty" dynamodbav:"associations,omitempty" json:"associations,omitempty" diff:"associations"`
	Entries       []*NetworkAclEntry       `bson:"entries,omitempty" dynamodbav:"entries,omitempty" json:"entries,omitempty" diff:"entries"`
	IsDefault     bool                     `bson:"is_default,omitempty" dynamodbav:"is_default" json:"is_default,omitempty" diff:"is_default"`
	NetworkAclId  string                   `bson:"network_acl_id,omitempty" dynamodbav:"network_acl_id,omitempty" inventory_primary_key:"true" json:"network_acl_id,omitempty" diff:"network_acl_id,identifier"`
	OwnerId       string                   `bson:"owner_id,omitempty" dynamodbav:"owner_id,omitempty" json:"owner_id,omitempty" diff:"owner_id"`
	Tags          map[string]string        `bson:"tags,omitempty" dynamodbav:"tags,omitempty" json:"tags,omitempty" diff:"tags"`
	VpcId         string                   `bson:"vpc_id,omitempty" dynamodbav:"vpc_id,omitempty" json:"vpc_id,omitempty" diff:"vpc_id"`
	AccountId     string                   `bson:"account_id,omitempty" dynamodbav:"account_id,omitempty" json:"account_id,omitempty" diff:"account_id"`
	Region        string                   `bson:"region,omitempty" dynamodbav:"region,omitempty" json:"region,omitempty" diff:"region"`
	ReportTime    time.Time                `bson:"report_time,omitempty" dynamodbav:"report_time,unixtime,omitempty" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID string                   `bson:"_id,omitempty" dynamodbav:"_id,omitempty" json:"_id,omitempty" diff:"-"`
}