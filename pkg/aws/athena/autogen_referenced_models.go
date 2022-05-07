//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_referenced_resource_file.tmpl
package athena

type WorkGroupConfiguration struct {
	BytesScannedCutoffPerQuery      int64                `bson:"bytes_scanned_cutoff_per_query,omitempty" ion:"bytes_scanned_cutoff_per_query" dynamodbav:"bytes_scanned_cutoff_per_query,omitempty" parquet:"name=bytes_scanned_cutoff_per_query,type=INT64" json:"bytes_scanned_cutoff_per_query,omitempty" diff:"bytes_scanned_cutoff_per_query"`
	EnforceWorkGroupConfiguration   bool                 `bson:"enforce_work_group_configuration,omitempty" ion:"enforce_work_group_configuration" dynamodbav:"enforce_work_group_configuration" parquet:"name=enforce_work_group_configuration,type=BOOLEAN" json:"enforce_work_group_configuration,omitempty" diff:"enforce_work_group_configuration"`
	EngineVersion                   *EngineVersion       `bson:"engine_version,omitempty" ion:"engine_version" dynamodbav:"engine_version,omitempty" parquet:"name=engine_version" json:"engine_version,omitempty" diff:"engine_version"`
	PublishCloudWatchMetricsEnabled bool                 `bson:"publish_cloud_watch_metrics_enabled,omitempty" ion:"publish_cloud_watch_metrics_enabled" dynamodbav:"publish_cloud_watch_metrics_enabled" parquet:"name=publish_cloud_watch_metrics_enabled,type=BOOLEAN" json:"publish_cloud_watch_metrics_enabled,omitempty" diff:"publish_cloud_watch_metrics_enabled"`
	RequesterPaysEnabled            bool                 `bson:"requester_pays_enabled,omitempty" ion:"requester_pays_enabled" dynamodbav:"requester_pays_enabled" parquet:"name=requester_pays_enabled,type=BOOLEAN" json:"requester_pays_enabled,omitempty" diff:"requester_pays_enabled"`
	ResultConfiguration             *ResultConfiguration `bson:"result_configuration,omitempty" ion:"result_configuration" dynamodbav:"result_configuration,omitempty" parquet:"name=result_configuration" json:"result_configuration,omitempty" diff:"result_configuration"`
}

type EngineVersion struct {
	EffectiveEngineVersion string `bson:"effective_engine_version,omitempty" ion:"effective_engine_version" dynamodbav:"effective_engine_version,omitempty" parquet:"name=effective_engine_version,type=BYTE_ARRAY,convertedtype=UTF8" json:"effective_engine_version,omitempty" diff:"effective_engine_version"`
	SelectedEngineVersion  string `bson:"selected_engine_version,omitempty" ion:"selected_engine_version" dynamodbav:"selected_engine_version,omitempty" parquet:"name=selected_engine_version,type=BYTE_ARRAY,convertedtype=UTF8" json:"selected_engine_version,omitempty" diff:"selected_engine_version"`
}

type ResultConfiguration struct {
	AclConfiguration        *AclConfiguration        `bson:"acl_configuration,omitempty" ion:"acl_configuration" dynamodbav:"acl_configuration,omitempty" parquet:"name=acl_configuration" json:"acl_configuration,omitempty" diff:"acl_configuration"`
	EncryptionConfiguration *EncryptionConfiguration `bson:"encryption_configuration,omitempty" ion:"encryption_configuration" dynamodbav:"encryption_configuration,omitempty" parquet:"name=encryption_configuration" json:"encryption_configuration,omitempty" diff:"encryption_configuration"`
	ExpectedBucketOwner     string                   `bson:"expected_bucket_owner,omitempty" ion:"expected_bucket_owner" dynamodbav:"expected_bucket_owner,omitempty" parquet:"name=expected_bucket_owner,type=BYTE_ARRAY,convertedtype=UTF8" json:"expected_bucket_owner,omitempty" diff:"expected_bucket_owner"`
	OutputLocation          string                   `bson:"output_location,omitempty" ion:"output_location" dynamodbav:"output_location,omitempty" parquet:"name=output_location,type=BYTE_ARRAY,convertedtype=UTF8" json:"output_location,omitempty" diff:"output_location"`
}

type AclConfiguration struct {
	S3AclOption string `bson:"s3_acl_option,omitempty" ion:"s3_acl_option" dynamodbav:"s3_acl_option,omitempty" parquet:"name=s3_acl_option,type=BYTE_ARRAY,convertedtype=UTF8" json:"s3_acl_option,omitempty" diff:"s3_acl_option"`
}

type EncryptionConfiguration struct {
	EncryptionOption string `bson:"encryption_option,omitempty" ion:"encryption_option" dynamodbav:"encryption_option,omitempty" parquet:"name=encryption_option,type=BYTE_ARRAY,convertedtype=UTF8" json:"encryption_option,omitempty" diff:"encryption_option"`
	KmsKey           string `bson:"kms_key,omitempty" ion:"kms_key" dynamodbav:"kms_key,omitempty" parquet:"name=kms_key,type=BYTE_ARRAY,convertedtype=UTF8" json:"kms_key,omitempty" diff:"kms_key"`
}