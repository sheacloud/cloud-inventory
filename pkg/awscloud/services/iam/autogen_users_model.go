//AUTOGENERATED CODE DO NOT EDIT
package iam

import (
	"time"
)

type User struct {
	Arn                   string `parquet:"name=arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"arn" diff:"arn"`
	CreateDate            *time.Time
	Path                  string `parquet:"name=path,type=BYTE_ARRAY,convertedtype=UTF8" json:"path" diff:"path"`
	UserId                string `parquet:"name=user_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"user_id" diff:"user_id,identifier"`
	UserName              string `parquet:"name=user_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"user_name" diff:"user_name"`
	PasswordLastUsed      *time.Time
	PermissionsBoundary   *AttachedPermissionsBoundary `parquet:"name=permissions_boundary" json:"permissions_boundary" diff:"permissions_boundary"`
	Tags                  map[string]string            `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"tags" diff:"tags"`
	AccountId             string                       `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"account_id" diff:"account_id"`
	Region                string                       `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8" json:"region" diff:"region"`
	ReportTime            int64                        `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"report_time" diff:"report_time,immutable"`
	InlinePolicies        []string                     `parquet:"name=inline_policies,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8" json:"inline_policies" diff:"inline_policies"`
	GroupIds              []string                     `parquet:"name=group_ids,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8" json:"group_ids" diff:"group_ids"`
	AccessKeys            []*AccessKeyMetadata         `parquet:"name=access_keys,type=MAP,convertedtype=LIST" json:"access_keys" diff:"access_keys"`
	LoginProfile          *LoginProfile                `parquet:"name=login_profile" json:"login_profile" diff:"login_profile"`
	AttachedPolicies      []*AttachedPolicy            `parquet:"name=attached_policies,type=MAP,convertedtype=LIST" json:"attached_policies" diff:"attached_policies"`
	CreateDateMilli       int64                        `parquet:"name=create_date,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"create_date" diff:"create_date"`
	PasswordLastUsedMilli int64                        `parquet:"name=password_last_used,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"password_last_used" diff:"password_last_used"`
}

func (x *User) GetReportTime() int64 {
	return x.ReportTime
}
