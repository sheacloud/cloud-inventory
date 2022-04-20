//AUTOGENERATED CODE DO NOT EDIT
package iam

import (
	"time"
)

type User struct {
	Arn                 string                       `bson:"arn,omitempty" dynamodbav:"arn,omitempty" json:"arn,omitempty" diff:"arn"`
	CreateDate          *time.Time                   `bson:"create_date,omitempty" dynamodbav:"create_date,unixtime,omitempty" json:"create_date,omitempty" diff:"create_date"`
	Path                string                       `bson:"path,omitempty" dynamodbav:"path,omitempty" json:"path,omitempty" diff:"path"`
	UserId              string                       `bson:"user_id,omitempty" dynamodbav:"user_id,omitempty" inventory_primary_key:"true" json:"user_id,omitempty" diff:"user_id,identifier"`
	UserName            string                       `bson:"user_name,omitempty" dynamodbav:"user_name,omitempty" json:"user_name,omitempty" diff:"user_name"`
	PasswordLastUsed    *time.Time                   `bson:"password_last_used,omitempty" dynamodbav:"password_last_used,unixtime,omitempty" json:"password_last_used,omitempty" diff:"password_last_used"`
	PermissionsBoundary *AttachedPermissionsBoundary `bson:"permissions_boundary,omitempty" dynamodbav:"permissions_boundary,omitempty" json:"permissions_boundary,omitempty" diff:"permissions_boundary"`
	Tags                map[string]string            `bson:"tags,omitempty" dynamodbav:"tags,omitempty" json:"tags,omitempty" diff:"tags"`
	AccountId           string                       `bson:"account_id,omitempty" dynamodbav:"account_id,omitempty" json:"account_id,omitempty" diff:"account_id"`
	Region              string                       `bson:"region,omitempty" dynamodbav:"region,omitempty" json:"region,omitempty" diff:"region"`
	ReportTime          time.Time                    `bson:"report_time,omitempty" dynamodbav:"report_time,unixtime,omitempty" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID       string                       `bson:"_id,omitempty" dynamodbav:"_id,omitempty" json:"_id,omitempty" diff:"-"`
	InlinePolicies      []string                     `bson:"inline_policies,omitempty" dynamodbav:"inline_policies,omitempty" json:"inline_policies,omitempty" diff:"inline_policies"`
	GroupIds            []string                     `bson:"group_ids,omitempty" dynamodbav:"group_ids,omitempty" json:"group_ids,omitempty" diff:"group_ids"`
	AccessKeys          []*AccessKeyMetadata         `bson:"access_keys,omitempty" dynamodbav:"access_keys,omitempty" json:"access_keys,omitempty" diff:"access_keys"`
	LoginProfile        *LoginProfile                `bson:"login_profile,omitempty" dynamodbav:"login_profile,omitempty" json:"login_profile,omitempty" diff:"login_profile"`
	AttachedPolicies    []*AttachedPolicy            `bson:"attached_policies,omitempty" dynamodbav:"attached_policies,omitempty" json:"attached_policies,omitempty" diff:"attached_policies"`
}
