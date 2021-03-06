//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_resource_file.tmpl
package iam

type User struct {
	Arn                 string                       `bson:"arn,omitempty" ion:"arn" dynamodbav:"arn,omitempty" parquet:"name=arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"arn,omitempty" diff:"arn"`
	CreateDate          int64                        `bson:"create_date,omitempty" ion:"create_date" dynamodbav:"create_date,omitempty" parquet:"name=create_date,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"create_date,omitempty" diff:"create_date"`
	Path                string                       `bson:"path,omitempty" ion:"path" dynamodbav:"path,omitempty" parquet:"name=path,type=BYTE_ARRAY,convertedtype=UTF8" json:"path,omitempty" diff:"path"`
	UserId              string                       `bson:"user_id,omitempty" ion:"user_id" dynamodbav:"user_id,omitempty" parquet:"name=user_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"user_id,omitempty" diff:"user_id,identifier"`
	UserName            string                       `bson:"user_name,omitempty" ion:"user_name" dynamodbav:"user_name,omitempty" parquet:"name=user_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"user_name,omitempty" diff:"user_name"`
	PasswordLastUsed    int64                        `bson:"password_last_used,omitempty" ion:"password_last_used" dynamodbav:"password_last_used,omitempty" parquet:"name=password_last_used,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"password_last_used,omitempty" diff:"password_last_used"`
	PermissionsBoundary *AttachedPermissionsBoundary `bson:"permissions_boundary,omitempty" ion:"permissions_boundary" dynamodbav:"permissions_boundary,omitempty" parquet:"name=permissions_boundary" json:"permissions_boundary,omitempty" diff:"permissions_boundary"`
	Tags                map[string]string            `bson:"tags,omitempty" ion:"tags" dynamodbav:"tags,omitempty" parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"tags,omitempty" diff:"tags"`
	AccountId           string                       `bson:"account_id,omitempty" ion:"account_id" dynamodbav:"account_id,omitempty" parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"account_id,omitempty" diff:"account_id"`
	Region              string                       `bson:"region,omitempty" ion:"region" dynamodbav:"region,omitempty" parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8" json:"region,omitempty" diff:"region"`
	ReportTime          int64                        `bson:"report_time,omitempty" ion:"report_time" dynamodbav:"report_time,omitempty" parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID       string                       `bson:"_id,omitempty" ion:"_id" dynamodbav:"_id,omitempty" parquet:"name=inventory_uuid,type=BYTE_ARRAY,convertedtype=UTF8" json:"_id,omitempty" diff:"-"`
	InlinePolicies      []string                     `bson:"inline_policies,omitempty" ion:"inline_policies" dynamodbav:"inline_policies,omitempty" parquet:"name=inline_policies,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8" json:"inline_policies,omitempty" diff:"inline_policies"`
	GroupIds            []string                     `bson:"group_ids,omitempty" ion:"group_ids" dynamodbav:"group_ids,omitempty" parquet:"name=group_ids,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8" json:"group_ids,omitempty" diff:"group_ids"`
	AccessKeys          []*AccessKeyMetadata         `bson:"access_keys,omitempty" ion:"access_keys" dynamodbav:"access_keys,omitempty" parquet:"name=access_keys,type=MAP,convertedtype=LIST" json:"access_keys,omitempty" diff:"access_keys"`
	LoginProfile        *LoginProfile                `bson:"login_profile,omitempty" ion:"login_profile" dynamodbav:"login_profile,omitempty" parquet:"name=login_profile" json:"login_profile,omitempty" diff:"login_profile"`
	AttachedPolicies    []*AttachedPolicy            `bson:"attached_policies,omitempty" ion:"attached_policies" dynamodbav:"attached_policies,omitempty" parquet:"name=attached_policies,type=MAP,convertedtype=LIST" json:"attached_policies,omitempty" diff:"attached_policies"`
}
