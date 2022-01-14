//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"time"
)

type Volume struct {
	Attachments []*VolumeAttachment `parquet:"name=attachments,type=MAP,convertedtype=LIST"`
	AvailabilityZone string `parquet:"name=availability_zone,type=BYTE_ARRAY,convertedtype=UTF8"`
	CreateTime *time.Time 
	Encrypted bool `parquet:"name=encrypted,type=BOOLEAN"`
	FastRestored bool `parquet:"name=fast_restored,type=BOOLEAN"`
	Iops int32 `parquet:"name=iops,type=INT32"`
	KmsKeyId string `parquet:"name=kms_key_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	MultiAttachEnabled bool `parquet:"name=multi_attach_enabled,type=BOOLEAN"`
	OutpostArn string `parquet:"name=outpost_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	Size int32 `parquet:"name=size,type=INT32"`
	SnapshotId string `parquet:"name=snapshot_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	State string `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
	TagsOld []*Tag `parquet:"name=tags_old,type=MAP,convertedtype=LIST"`
	Throughput int32 `parquet:"name=throughput,type=INT32"`
	VolumeId string `parquet:"name=volume_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	VolumeType string `parquet:"name=volume_type,type=BYTE_ARRAY,convertedtype=UTF8"`
	AccountId string `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region string `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime int64 `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	CreateTimeMilli int64 `parquet:"name=create_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	Tags map[string]string `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
}
