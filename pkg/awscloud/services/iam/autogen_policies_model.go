//AUTOGENERATED CODE DO NOT EDIT
package iam

import (
	"time"
)

type Policy struct {
	Arn                           string `parquet:"name=arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"arn" diff:"arn"`
	AttachmentCount               int32  `parquet:"name=attachment_count,type=INT32" json:"attachment_count" diff:"attachment_count"`
	CreateDate                    *time.Time
	DefaultVersionId              string            `parquet:"name=default_version_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"default_version_id" diff:"default_version_id"`
	Description                   string            `parquet:"name=description,type=BYTE_ARRAY,convertedtype=UTF8" json:"description" diff:"description"`
	IsAttachable                  bool              `parquet:"name=is_attachable,type=BOOLEAN" json:"is_attachable" diff:"is_attachable"`
	Path                          string            `parquet:"name=path,type=BYTE_ARRAY,convertedtype=UTF8" json:"path" diff:"path"`
	PermissionsBoundaryUsageCount int32             `parquet:"name=permissions_boundary_usage_count,type=INT32" json:"permissions_boundary_usage_count" diff:"permissions_boundary_usage_count"`
	PolicyId                      string            `parquet:"name=policy_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"policy_id" diff:"policy_id,identifier"`
	PolicyName                    string            `parquet:"name=policy_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"policy_name" diff:"policy_name"`
	Tags                          map[string]string `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"tags" diff:"tags"`
	UpdateDate                    *time.Time
	AccountId                     string `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"account_id" diff:"account_id"`
	Region                        string `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8" json:"region" diff:"region"`
	ReportTime                    int64  `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"report_time" diff:"report_time,immutable"`
	CreateDateMilli               int64  `parquet:"name=create_date,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"create_date" diff:"create_date"`
	UpdateDateMilli               int64  `parquet:"name=update_date,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"update_date" diff:"update_date"`
}

func (x *Policy) GetReportTime() int64 {
	return x.ReportTime
}
