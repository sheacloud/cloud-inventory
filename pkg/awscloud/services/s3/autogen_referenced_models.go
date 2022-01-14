//AUTOGENERATED CODE DO NOT EDIT
package s3

import (
	"time"
)

type ReplicationConfiguration struct {
	Role string `parquet:"name=role,type=BYTE_ARRAY,convertedtype=UTF8"`
	Rules []*ReplicationRule `parquet:"name=rules,type=MAP,convertedtype=LIST"`
}

type ReplicationRule struct {
	Destination *Destination `parquet:"name=destination"`
	Status string `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
	DeleteMarkerReplication *DeleteMarkerReplication `parquet:"name=delete_marker_replication"`
	ExistingObjectReplication *ExistingObjectReplication `parquet:"name=existing_object_replication"`
	ID string `parquet:"name=id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Prefix string `parquet:"name=prefix,type=BYTE_ARRAY,convertedtype=UTF8"`
	Priority int32 `parquet:"name=priority,type=INT32"`
	SourceSelectionCriteria *SourceSelectionCriteria `parquet:"name=source_selection_criteria"`
}

type Destination struct {
	Bucket string `parquet:"name=bucket,type=BYTE_ARRAY,convertedtype=UTF8"`
	AccessControlTranslation *AccessControlTranslation `parquet:"name=access_control_translation"`
	Account string `parquet:"name=account,type=BYTE_ARRAY,convertedtype=UTF8"`
	EncryptionConfiguration *EncryptionConfiguration `parquet:"name=encryption_configuration"`
	Metrics *Metrics `parquet:"name=metrics"`
	ReplicationTime *ReplicationTime `parquet:"name=replication_time"`
	StorageClass string `parquet:"name=storage_class,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type AccessControlTranslation struct {
	Owner string `parquet:"name=owner,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type EncryptionConfiguration struct {
	ReplicaKmsKeyID string `parquet:"name=replica_kms_key_id,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type Metrics struct {
	Status string `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
	EventThreshold *ReplicationTimeValue `parquet:"name=event_threshold"`
}

type ReplicationTimeValue struct {
	Minutes int32 `parquet:"name=minutes,type=INT32"`
}

type ReplicationTime struct {
	Status string `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
	Time *ReplicationTimeValue `parquet:"name=time"`
}

type DeleteMarkerReplication struct {
	Status string `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type ExistingObjectReplication struct {
	Status string `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type SourceSelectionCriteria struct {
	ReplicaModifications *ReplicaModifications `parquet:"name=replica_modifications"`
	SseKmsEncryptedObjects *SseKmsEncryptedObjects `parquet:"name=sse_kms_encrypted_objects"`
}

type ReplicaModifications struct {
	Status string `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type SseKmsEncryptedObjects struct {
	Status string `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type Grant struct {
	Grantee *Grantee `parquet:"name=grantee"`
	Permission string `parquet:"name=permission,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type Grantee struct {
	Type string `parquet:"name=type,type=BYTE_ARRAY,convertedtype=UTF8"`
	DisplayName string `parquet:"name=display_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	EmailAddress string `parquet:"name=email_address,type=BYTE_ARRAY,convertedtype=UTF8"`
	ID string `parquet:"name=id,type=BYTE_ARRAY,convertedtype=UTF8"`
	URI string `parquet:"name=uri,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type CORSRule struct {
	AllowedMethods []string `parquet:"name=allowed_methods,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	AllowedOrigins []string `parquet:"name=allowed_origins,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	AllowedHeaders []string `parquet:"name=allowed_headers,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	ExposeHeaders []string `parquet:"name=expose_headers,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	ID string `parquet:"name=id,type=BYTE_ARRAY,convertedtype=UTF8"`
	MaxAgeSeconds int32 `parquet:"name=max_age_seconds,type=INT32"`
}

type ServerSideEncryptionConfiguration struct {
	Rules []*ServerSideEncryptionRule `parquet:"name=rules,type=MAP,convertedtype=LIST"`
}

type ServerSideEncryptionRule struct {
	ApplyServerSideEncryptionByDefault *ServerSideEncryptionByDefault `parquet:"name=apply_server_side_encryption_by_default"`
	BucketKeyEnabled bool `parquet:"name=bucket_key_enabled,type=BOOLEAN"`
}

type ServerSideEncryptionByDefault struct {
	SSEAlgorithm string `parquet:"name=sse_algorithm,type=BYTE_ARRAY,convertedtype=UTF8"`
	KMSMasterKeyID string `parquet:"name=kms_master_key_id,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type IntelligentTieringConfiguration struct {
	Id string `parquet:"name=id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Status string `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
	Tierings []*Tiering `parquet:"name=tierings,type=MAP,convertedtype=LIST"`
	Filter *IntelligentTieringFilter `parquet:"name=filter"`
}

type Tiering struct {
	AccessTier string `parquet:"name=access_tier,type=BYTE_ARRAY,convertedtype=UTF8"`
	Days int32 `parquet:"name=days,type=INT32"`
}

type IntelligentTieringFilter struct {
	And *IntelligentTieringAndOperator `parquet:"name=and"`
	Prefix string `parquet:"name=prefix,type=BYTE_ARRAY,convertedtype=UTF8"`
	Tag *Tag `parquet:"name=tag"`
}

type IntelligentTieringAndOperator struct {
	Prefix string `parquet:"name=prefix,type=BYTE_ARRAY,convertedtype=UTF8"`
	Tags []*Tag `parquet:"name=tags,type=MAP,convertedtype=LIST"`
}

type Tag struct {
	Key string `parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type InventoryConfiguration struct {
	Destination *InventoryDestination `parquet:"name=destination"`
	Id string `parquet:"name=id,type=BYTE_ARRAY,convertedtype=UTF8"`
	IncludedObjectVersions string `parquet:"name=included_object_versions,type=BYTE_ARRAY,convertedtype=UTF8"`
	IsEnabled bool `parquet:"name=is_enabled,type=BOOLEAN"`
	Schedule *InventorySchedule `parquet:"name=schedule"`
	Filter *InventoryFilter `parquet:"name=filter"`
	OptionalFields []string `parquet:"name=optional_fields,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
}

type InventoryDestination struct {
	S3BucketDestination *InventoryS3BucketDestination `parquet:"name=s3_bucket_destination"`
}

type InventoryS3BucketDestination struct {
	Bucket string `parquet:"name=bucket,type=BYTE_ARRAY,convertedtype=UTF8"`
	Format string `parquet:"name=format,type=BYTE_ARRAY,convertedtype=UTF8"`
	AccountId string `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Encryption *InventoryEncryption `parquet:"name=encryption"`
	Prefix string `parquet:"name=prefix,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type InventoryEncryption struct {
	SSEKMS *SSEKMS `parquet:"name=ssekms"`
}

type SSEKMS struct {
	KeyId string `parquet:"name=key_id,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type InventorySchedule struct {
	Frequency string `parquet:"name=frequency,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type InventoryFilter struct {
	Prefix string `parquet:"name=prefix,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type LifecycleRule struct {
	Status string `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
	AbortIncompleteMultipartUpload *AbortIncompleteMultipartUpload `parquet:"name=abort_incomplete_multipart_upload"`
	Expiration *LifecycleExpiration `parquet:"name=expiration"`
	ID string `parquet:"name=id,type=BYTE_ARRAY,convertedtype=UTF8"`
	NoncurrentVersionExpiration *NoncurrentVersionExpiration `parquet:"name=noncurrent_version_expiration"`
	NoncurrentVersionTransitions []*NoncurrentVersionTransition `parquet:"name=noncurrent_version_transitions,type=MAP,convertedtype=LIST"`
	Prefix string `parquet:"name=prefix,type=BYTE_ARRAY,convertedtype=UTF8"`
	Transitions []*Transition `parquet:"name=transitions,type=MAP,convertedtype=LIST"`
}

type AbortIncompleteMultipartUpload struct {
	DaysAfterInitiation int32 `parquet:"name=days_after_initiation,type=INT32"`
}

type LifecycleExpiration struct {
	Date *time.Time 
	Days int32 `parquet:"name=days,type=INT32"`
	ExpiredObjectDeleteMarker bool `parquet:"name=expired_object_delete_marker,type=BOOLEAN"`
	DateMilli int64 `parquet:"name=date,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type NoncurrentVersionExpiration struct {
	NoncurrentDays int32 `parquet:"name=noncurrent_days,type=INT32"`
}

type NoncurrentVersionTransition struct {
	NoncurrentDays int32 `parquet:"name=noncurrent_days,type=INT32"`
	StorageClass string `parquet:"name=storage_class,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type Transition struct {
	Date *time.Time 
	Days int32 `parquet:"name=days,type=INT32"`
	StorageClass string `parquet:"name=storage_class,type=BYTE_ARRAY,convertedtype=UTF8"`
	DateMilli int64 `parquet:"name=date,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type LoggingEnabled struct {
	TargetBucket string `parquet:"name=target_bucket,type=BYTE_ARRAY,convertedtype=UTF8"`
	TargetPrefix string `parquet:"name=target_prefix,type=BYTE_ARRAY,convertedtype=UTF8"`
	TargetGrants []*TargetGrant `parquet:"name=target_grants,type=MAP,convertedtype=LIST"`
}

type TargetGrant struct {
	Grantee *Grantee `parquet:"name=grantee"`
	Permission string `parquet:"name=permission,type=BYTE_ARRAY,convertedtype=UTF8"`
}


