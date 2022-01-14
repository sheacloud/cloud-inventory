//AUTOGENERATED CODE DO NOT EDIT
package dynamodb

import (
	"time"
)

type ArchivalSummary struct {
	ArchivalBackupArn string `parquet:"name=archival_backup_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	ArchivalDateTime *time.Time 
	ArchivalReason string `parquet:"name=archival_reason,type=BYTE_ARRAY,convertedtype=UTF8"`
	ArchivalDateTimeMilli int64 `parquet:"name=archival_date_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type AttributeDefinition struct {
	AttributeName string `parquet:"name=attribute_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	AttributeType string `parquet:"name=attribute_type,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type BillingModeSummary struct {
	BillingMode string `parquet:"name=billing_mode,type=BYTE_ARRAY,convertedtype=UTF8"`
	LastUpdateToPayPerRequestDateTime *time.Time 
	LastUpdateToPayPerRequestDateTimeMilli int64 `parquet:"name=last_update_to_pay_per_request_date_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type GlobalSecondaryIndexDescription struct {
	Backfilling bool `parquet:"name=backfilling,type=BOOLEAN"`
	IndexArn string `parquet:"name=index_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	IndexName string `parquet:"name=index_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	IndexSizeBytes int64 `parquet:"name=index_size_bytes,type=INT64"`
	IndexStatus string `parquet:"name=index_status,type=BYTE_ARRAY,convertedtype=UTF8"`
	ItemCount int64 `parquet:"name=item_count,type=INT64"`
	KeySchema []*KeySchemaElement `parquet:"name=key_schema,type=MAP,convertedtype=LIST"`
	Projection *Projection `parquet:"name=projection"`
	ProvisionedThroughput *ProvisionedThroughputDescription `parquet:"name=provisioned_throughput"`
}

type KeySchemaElement struct {
	AttributeName string `parquet:"name=attribute_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	KeyType string `parquet:"name=key_type,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type Projection struct {
	NonKeyAttributes []string `parquet:"name=non_key_attributes,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	ProjectionType string `parquet:"name=projection_type,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type ProvisionedThroughputDescription struct {
	LastDecreaseDateTime *time.Time 
	LastIncreaseDateTime *time.Time 
	NumberOfDecreasesToday int64 `parquet:"name=number_of_decreases_today,type=INT64"`
	ReadCapacityUnits int64 `parquet:"name=read_capacity_units,type=INT64"`
	WriteCapacityUnits int64 `parquet:"name=write_capacity_units,type=INT64"`
	LastDecreaseDateTimeMilli int64 `parquet:"name=last_decrease_date_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	LastIncreaseDateTimeMilli int64 `parquet:"name=last_increase_date_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type LocalSecondaryIndexDescription struct {
	IndexArn string `parquet:"name=index_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	IndexName string `parquet:"name=index_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	IndexSizeBytes int64 `parquet:"name=index_size_bytes,type=INT64"`
	ItemCount int64 `parquet:"name=item_count,type=INT64"`
	KeySchema []*KeySchemaElement `parquet:"name=key_schema,type=MAP,convertedtype=LIST"`
	Projection *Projection `parquet:"name=projection"`
}

type ReplicaDescription struct {
	GlobalSecondaryIndexes []*ReplicaGlobalSecondaryIndexDescription `parquet:"name=global_secondary_indexes,type=MAP,convertedtype=LIST"`
	KMSMasterKeyId string `parquet:"name=kms_master_key_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	ProvisionedThroughputOverride *ProvisionedThroughputOverride `parquet:"name=provisioned_throughput_override"`
	RegionName string `parquet:"name=region_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReplicaInaccessibleDateTime *time.Time 
	ReplicaStatus string `parquet:"name=replica_status,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReplicaStatusDescription string `parquet:"name=replica_status_description,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReplicaStatusPercentProgress string `parquet:"name=replica_status_percent_progress,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReplicaTableClassSummary *TableClassSummary `parquet:"name=replica_table_class_summary"`
	ReplicaInaccessibleDateTimeMilli int64 `parquet:"name=replica_inaccessible_date_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type ReplicaGlobalSecondaryIndexDescription struct {
	IndexName string `parquet:"name=index_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	ProvisionedThroughputOverride *ProvisionedThroughputOverride `parquet:"name=provisioned_throughput_override"`
}

type ProvisionedThroughputOverride struct {
	ReadCapacityUnits int64 `parquet:"name=read_capacity_units,type=INT64"`
}

type TableClassSummary struct {
	LastUpdateDateTime *time.Time 
	TableClass string `parquet:"name=table_class,type=BYTE_ARRAY,convertedtype=UTF8"`
	LastUpdateDateTimeMilli int64 `parquet:"name=last_update_date_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type RestoreSummary struct {
	RestoreDateTime *time.Time 
	RestoreInProgress bool `parquet:"name=restore_in_progress,type=BOOLEAN"`
	SourceBackupArn string `parquet:"name=source_backup_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	SourceTableArn string `parquet:"name=source_table_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	RestoreDateTimeMilli int64 `parquet:"name=restore_date_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type SSEDescription struct {
	InaccessibleEncryptionDateTime *time.Time 
	KMSMasterKeyArn string `parquet:"name=kms_master_key_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	SSEType string `parquet:"name=sse_type,type=BYTE_ARRAY,convertedtype=UTF8"`
	Status string `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
	InaccessibleEncryptionDateTimeMilli int64 `parquet:"name=inaccessible_encryption_date_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type StreamSpecification struct {
	StreamEnabled bool `parquet:"name=stream_enabled,type=BOOLEAN"`
	StreamViewType string `parquet:"name=stream_view_type,type=BYTE_ARRAY,convertedtype=UTF8"`
}

