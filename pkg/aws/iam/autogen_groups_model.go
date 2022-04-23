//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_resource_file.tmpl
package iam

type Group struct {
	Arn              string            `bson:"arn,omitempty" ion:"arn" dynamodbav:"arn,omitempty" parquet:"name=arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"arn,omitempty" diff:"arn"`
	CreateDate       int64             `bson:"create_date,omitempty" ion:"create_date" dynamodbav:"create_date,omitempty" parquet:"name=create_date,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"create_date,omitempty" diff:"create_date"`
	GroupId          string            `bson:"group_id,omitempty" ion:"group_id" dynamodbav:"group_id,omitempty" parquet:"name=group_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"group_id,omitempty" diff:"group_id,identifier"`
	GroupName        string            `bson:"group_name,omitempty" ion:"group_name" dynamodbav:"group_name,omitempty" parquet:"name=group_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"group_name,omitempty" diff:"group_name"`
	Path             string            `bson:"path,omitempty" ion:"path" dynamodbav:"path,omitempty" parquet:"name=path,type=BYTE_ARRAY,convertedtype=UTF8" json:"path,omitempty" diff:"path"`
	AccountId        string            `bson:"account_id,omitempty" ion:"account_id" dynamodbav:"account_id,omitempty" parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"account_id,omitempty" diff:"account_id"`
	Region           string            `bson:"region,omitempty" ion:"region" dynamodbav:"region,omitempty" parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8" json:"region,omitempty" diff:"region"`
	ReportTime       int64             `bson:"report_time,omitempty" ion:"report_time" dynamodbav:"report_time,omitempty" parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID    string            `bson:"_id,omitempty" ion:"_id" dynamodbav:"_id,omitempty" parquet:"name=inventory_uuid,type=BYTE_ARRAY,convertedtype=UTF8" json:"_id,omitempty" diff:"-"`
	InlinePolicies   []string          `bson:"inline_policies,omitempty" ion:"inline_policies" dynamodbav:"inline_policies,omitempty" parquet:"name=inline_policies,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8" json:"inline_policies,omitempty" diff:"inline_policies"`
	UserIds          []string          `bson:"user_ids,omitempty" ion:"user_ids" dynamodbav:"user_ids,omitempty" parquet:"name=user_ids,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8" json:"user_ids,omitempty" diff:"user_ids"`
	AttachedPolicies []*AttachedPolicy `bson:"attached_policies,omitempty" ion:"attached_policies" dynamodbav:"attached_policies,omitempty" parquet:"name=attached_policies,type=MAP,convertedtype=LIST" json:"attached_policies,omitempty" diff:"attached_policies"`
}
