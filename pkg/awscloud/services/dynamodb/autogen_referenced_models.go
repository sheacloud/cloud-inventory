//AUTOGENERATED CODE DO NOT EDIT
package dynamodb

import (
	"time"
)

type ArchivalSummary struct {
	ArchivalBackupArn     string     `parquet:"name=archival_backup_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"archival_backup_arn" diff:"archival_backup_arn"`
	ArchivalDateTime      *time.Time `json:"-"`
	ArchivalReason        string     `parquet:"name=archival_reason,type=BYTE_ARRAY,convertedtype=UTF8" json:"archival_reason" diff:"archival_reason"`
	ArchivalDateTimeMilli int64      `parquet:"name=archival_date_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"archival_date_time" diff:"archival_date_time"`
}

type AttributeDefinition struct {
	AttributeName string `parquet:"name=attribute_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"attribute_name" diff:"attribute_name"`
	AttributeType string `parquet:"name=attribute_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"attribute_type" diff:"attribute_type"`
}

type BillingModeSummary struct {
	BillingMode                            string     `parquet:"name=billing_mode,type=BYTE_ARRAY,convertedtype=UTF8" json:"billing_mode" diff:"billing_mode"`
	LastUpdateToPayPerRequestDateTime      *time.Time `json:"-"`
	LastUpdateToPayPerRequestDateTimeMilli int64      `parquet:"name=last_update_to_pay_per_request_date_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"last_update_to_pay_per_request_date_time" diff:"last_update_to_pay_per_request_date_time"`
}

type GlobalSecondaryIndexDescription struct {
	Backfilling           bool                              `parquet:"name=backfilling,type=BOOLEAN" json:"backfilling" diff:"backfilling"`
	IndexArn              string                            `parquet:"name=index_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"index_arn" diff:"index_arn"`
	IndexName             string                            `parquet:"name=index_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"index_name" diff:"index_name"`
	IndexSizeBytes        int64                             `parquet:"name=index_size_bytes,type=INT64" json:"index_size_bytes" diff:"index_size_bytes"`
	IndexStatus           string                            `parquet:"name=index_status,type=BYTE_ARRAY,convertedtype=UTF8" json:"index_status" diff:"index_status"`
	ItemCount             int64                             `parquet:"name=item_count,type=INT64" json:"item_count" diff:"item_count"`
	KeySchema             []*KeySchemaElement               `parquet:"name=key_schema,type=MAP,convertedtype=LIST" json:"key_schema" diff:"key_schema"`
	Projection            *Projection                       `parquet:"name=projection" json:"projection" diff:"projection"`
	ProvisionedThroughput *ProvisionedThroughputDescription `parquet:"name=provisioned_throughput" json:"provisioned_throughput" diff:"provisioned_throughput"`
}

type KeySchemaElement struct {
	AttributeName string `parquet:"name=attribute_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"attribute_name" diff:"attribute_name"`
	KeyType       string `parquet:"name=key_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"key_type" diff:"key_type"`
}

type Projection struct {
	NonKeyAttributes []string `parquet:"name=non_key_attributes,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8" json:"non_key_attributes" diff:"non_key_attributes"`
	ProjectionType   string   `parquet:"name=projection_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"projection_type" diff:"projection_type"`
}

type ProvisionedThroughputDescription struct {
	LastDecreaseDateTime      *time.Time `json:"-"`
	LastIncreaseDateTime      *time.Time `json:"-"`
	NumberOfDecreasesToday    int64      `parquet:"name=number_of_decreases_today,type=INT64" json:"number_of_decreases_today" diff:"number_of_decreases_today"`
	ReadCapacityUnits         int64      `parquet:"name=read_capacity_units,type=INT64" json:"read_capacity_units" diff:"read_capacity_units"`
	WriteCapacityUnits        int64      `parquet:"name=write_capacity_units,type=INT64" json:"write_capacity_units" diff:"write_capacity_units"`
	LastDecreaseDateTimeMilli int64      `parquet:"name=last_decrease_date_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"last_decrease_date_time" diff:"last_decrease_date_time"`
	LastIncreaseDateTimeMilli int64      `parquet:"name=last_increase_date_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"last_increase_date_time" diff:"last_increase_date_time"`
}

type LocalSecondaryIndexDescription struct {
	IndexArn       string              `parquet:"name=index_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"index_arn" diff:"index_arn"`
	IndexName      string              `parquet:"name=index_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"index_name" diff:"index_name"`
	IndexSizeBytes int64               `parquet:"name=index_size_bytes,type=INT64" json:"index_size_bytes" diff:"index_size_bytes"`
	ItemCount      int64               `parquet:"name=item_count,type=INT64" json:"item_count" diff:"item_count"`
	KeySchema      []*KeySchemaElement `parquet:"name=key_schema,type=MAP,convertedtype=LIST" json:"key_schema" diff:"key_schema"`
	Projection     *Projection         `parquet:"name=projection" json:"projection" diff:"projection"`
}

type ReplicaDescription struct {
	GlobalSecondaryIndexes           []*ReplicaGlobalSecondaryIndexDescription `parquet:"name=global_secondary_indexes,type=MAP,convertedtype=LIST" json:"global_secondary_indexes" diff:"global_secondary_indexes"`
	KMSMasterKeyId                   string                                    `parquet:"name=kms_master_key_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"kms_master_key_id" diff:"kms_master_key_id"`
	ProvisionedThroughputOverride    *ProvisionedThroughputOverride            `parquet:"name=provisioned_throughput_override" json:"provisioned_throughput_override" diff:"provisioned_throughput_override"`
	RegionName                       string                                    `parquet:"name=region_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"region_name" diff:"region_name"`
	ReplicaInaccessibleDateTime      *time.Time                                `json:"-"`
	ReplicaStatus                    string                                    `parquet:"name=replica_status,type=BYTE_ARRAY,convertedtype=UTF8" json:"replica_status" diff:"replica_status"`
	ReplicaStatusDescription         string                                    `parquet:"name=replica_status_description,type=BYTE_ARRAY,convertedtype=UTF8" json:"replica_status_description" diff:"replica_status_description"`
	ReplicaStatusPercentProgress     string                                    `parquet:"name=replica_status_percent_progress,type=BYTE_ARRAY,convertedtype=UTF8" json:"replica_status_percent_progress" diff:"replica_status_percent_progress"`
	ReplicaTableClassSummary         *TableClassSummary                        `parquet:"name=replica_table_class_summary" json:"replica_table_class_summary" diff:"replica_table_class_summary"`
	ReplicaInaccessibleDateTimeMilli int64                                     `parquet:"name=replica_inaccessible_date_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"replica_inaccessible_date_time" diff:"replica_inaccessible_date_time"`
}

type ReplicaGlobalSecondaryIndexDescription struct {
	IndexName                     string                         `parquet:"name=index_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"index_name" diff:"index_name"`
	ProvisionedThroughputOverride *ProvisionedThroughputOverride `parquet:"name=provisioned_throughput_override" json:"provisioned_throughput_override" diff:"provisioned_throughput_override"`
}

type ProvisionedThroughputOverride struct {
	ReadCapacityUnits int64 `parquet:"name=read_capacity_units,type=INT64" json:"read_capacity_units" diff:"read_capacity_units"`
}

type TableClassSummary struct {
	LastUpdateDateTime      *time.Time `json:"-"`
	TableClass              string     `parquet:"name=table_class,type=BYTE_ARRAY,convertedtype=UTF8" json:"table_class" diff:"table_class"`
	LastUpdateDateTimeMilli int64      `parquet:"name=last_update_date_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"last_update_date_time" diff:"last_update_date_time"`
}

type RestoreSummary struct {
	RestoreDateTime      *time.Time `json:"-"`
	RestoreInProgress    bool       `parquet:"name=restore_in_progress,type=BOOLEAN" json:"restore_in_progress" diff:"restore_in_progress"`
	SourceBackupArn      string     `parquet:"name=source_backup_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"source_backup_arn" diff:"source_backup_arn"`
	SourceTableArn       string     `parquet:"name=source_table_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"source_table_arn" diff:"source_table_arn"`
	RestoreDateTimeMilli int64      `parquet:"name=restore_date_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"restore_date_time" diff:"restore_date_time"`
}

type SSEDescription struct {
	InaccessibleEncryptionDateTime      *time.Time `json:"-"`
	KMSMasterKeyArn                     string     `parquet:"name=kms_master_key_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"kms_master_key_arn" diff:"kms_master_key_arn"`
	SSEType                             string     `parquet:"name=sse_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"sse_type" diff:"sse_type"`
	Status                              string     `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8" json:"status" diff:"status"`
	InaccessibleEncryptionDateTimeMilli int64      `parquet:"name=inaccessible_encryption_date_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"inaccessible_encryption_date_time" diff:"inaccessible_encryption_date_time"`
}

type StreamSpecification struct {
	StreamEnabled  bool   `parquet:"name=stream_enabled,type=BOOLEAN" json:"stream_enabled" diff:"stream_enabled"`
	StreamViewType string `parquet:"name=stream_view_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"stream_view_type" diff:"stream_view_type"`
}
