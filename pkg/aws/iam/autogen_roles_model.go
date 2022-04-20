//AUTOGENERATED CODE DO NOT EDIT
package iam

import (
	"time"
)

type Role struct {
	Arn                 string                       `bson:"arn,omitempty" dynamodbav:"arn,omitempty" json:"arn,omitempty" diff:"arn"`
	CreateDate          *time.Time                   `bson:"create_date,omitempty" dynamodbav:"create_date,unixtime,omitempty" json:"create_date,omitempty" diff:"create_date"`
	Path                string                       `bson:"path,omitempty" dynamodbav:"path,omitempty" json:"path,omitempty" diff:"path"`
	RoleId              string                       `bson:"role_id,omitempty" dynamodbav:"role_id,omitempty" inventory_primary_key:"true" json:"role_id,omitempty" diff:"role_id,identifier"`
	RoleName            string                       `bson:"role_name,omitempty" dynamodbav:"role_name,omitempty" json:"role_name,omitempty" diff:"role_name"`
	Description         string                       `bson:"description,omitempty" dynamodbav:"description,omitempty" json:"description,omitempty" diff:"description"`
	MaxSessionDuration  int32                        `bson:"max_session_duration,omitempty" dynamodbav:"max_session_duration,omitempty" json:"max_session_duration,omitempty" diff:"max_session_duration"`
	PermissionsBoundary *AttachedPermissionsBoundary `bson:"permissions_boundary,omitempty" dynamodbav:"permissions_boundary,omitempty" json:"permissions_boundary,omitempty" diff:"permissions_boundary"`
	RoleLastUsed        *RoleLastUsed                `bson:"role_last_used,omitempty" dynamodbav:"role_last_used,omitempty" json:"role_last_used,omitempty" diff:"role_last_used"`
	Tags                map[string]string            `bson:"tags,omitempty" dynamodbav:"tags,omitempty" json:"tags,omitempty" diff:"tags"`
	AccountId           string                       `bson:"account_id,omitempty" dynamodbav:"account_id,omitempty" json:"account_id,omitempty" diff:"account_id"`
	Region              string                       `bson:"region,omitempty" dynamodbav:"region,omitempty" json:"region,omitempty" diff:"region"`
	ReportTime          time.Time                    `bson:"report_time,omitempty" dynamodbav:"report_time,unixtime,omitempty" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID       string                       `bson:"_id,omitempty" dynamodbav:"_id,omitempty" json:"_id,omitempty" diff:"-"`
	InlinePolicies      []string                     `bson:"inline_policies,omitempty" dynamodbav:"inline_policies,omitempty" json:"inline_policies,omitempty" diff:"inline_policies"`
	AttachedPolicies    []*AttachedPolicy            `bson:"attached_policies,omitempty" dynamodbav:"attached_policies,omitempty" json:"attached_policies,omitempty" diff:"attached_policies"`
}
