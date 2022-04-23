//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_resource_file.tmpl
package ecs

type Cluster struct {
	ActiveServicesCount               int32                           `bson:"active_services_count,omitempty" ion:"active_services_count" dynamodbav:"active_services_count,omitempty" parquet:"name=active_services_count,type=INT32" json:"active_services_count,omitempty" diff:"active_services_count"`
	Attachments                       []*Attachment                   `bson:"attachments,omitempty" ion:"attachments" dynamodbav:"attachments,omitempty" parquet:"name=attachments,type=MAP,convertedtype=LIST" json:"attachments,omitempty" diff:"attachments"`
	AttachmentsStatus                 string                          `bson:"attachments_status,omitempty" ion:"attachments_status" dynamodbav:"attachments_status,omitempty" parquet:"name=attachments_status,type=BYTE_ARRAY,convertedtype=UTF8" json:"attachments_status,omitempty" diff:"attachments_status"`
	CapacityProviders                 []string                        `bson:"capacity_providers,omitempty" ion:"capacity_providers" dynamodbav:"capacity_providers,omitempty" parquet:"name=capacity_providers,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8" json:"capacity_providers,omitempty" diff:"capacity_providers"`
	ClusterArn                        string                          `bson:"cluster_arn,omitempty" ion:"cluster_arn" dynamodbav:"cluster_arn,omitempty" parquet:"name=cluster_arn,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"cluster_arn,omitempty" diff:"cluster_arn,identifier"`
	ClusterName                       string                          `bson:"cluster_name,omitempty" ion:"cluster_name" dynamodbav:"cluster_name,omitempty" parquet:"name=cluster_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"cluster_name,omitempty" diff:"cluster_name"`
	Configuration                     *ClusterConfiguration           `bson:"configuration,omitempty" ion:"configuration" dynamodbav:"configuration,omitempty" parquet:"name=configuration" json:"configuration,omitempty" diff:"configuration"`
	DefaultCapacityProviderStrategy   []*CapacityProviderStrategyItem `bson:"default_capacity_provider_strategy,omitempty" ion:"default_capacity_provider_strategy" dynamodbav:"default_capacity_provider_strategy,omitempty" parquet:"name=default_capacity_provider_strategy,type=MAP,convertedtype=LIST" json:"default_capacity_provider_strategy,omitempty" diff:"default_capacity_provider_strategy"`
	PendingTasksCount                 int32                           `bson:"pending_tasks_count,omitempty" ion:"pending_tasks_count" dynamodbav:"pending_tasks_count,omitempty" parquet:"name=pending_tasks_count,type=INT32" json:"pending_tasks_count,omitempty" diff:"pending_tasks_count"`
	RegisteredContainerInstancesCount int32                           `bson:"registered_container_instances_count,omitempty" ion:"registered_container_instances_count" dynamodbav:"registered_container_instances_count,omitempty" parquet:"name=registered_container_instances_count,type=INT32" json:"registered_container_instances_count,omitempty" diff:"registered_container_instances_count"`
	RunningTasksCount                 int32                           `bson:"running_tasks_count,omitempty" ion:"running_tasks_count" dynamodbav:"running_tasks_count,omitempty" parquet:"name=running_tasks_count,type=INT32" json:"running_tasks_count,omitempty" diff:"running_tasks_count"`
	Settings                          []*ClusterSetting               `bson:"settings,omitempty" ion:"settings" dynamodbav:"settings,omitempty" parquet:"name=settings,type=MAP,convertedtype=LIST" json:"settings,omitempty" diff:"settings"`
	Statistics                        []*KeyValuePair                 `bson:"statistics,omitempty" ion:"statistics" dynamodbav:"statistics,omitempty" parquet:"name=statistics,type=MAP,convertedtype=LIST" json:"statistics,omitempty" diff:"statistics"`
	Status                            string                          `bson:"status,omitempty" ion:"status" dynamodbav:"status,omitempty" parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8" json:"status,omitempty" diff:"status"`
	Tags                              map[string]string               `bson:"tags,omitempty" ion:"tags" dynamodbav:"tags,omitempty" parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"tags,omitempty" diff:"tags"`
	AccountId                         string                          `bson:"account_id,omitempty" ion:"account_id" dynamodbav:"account_id,omitempty" parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"account_id,omitempty" diff:"account_id"`
	Region                            string                          `bson:"region,omitempty" ion:"region" dynamodbav:"region,omitempty" parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8" json:"region,omitempty" diff:"region"`
	ReportTime                        int64                           `bson:"report_time,omitempty" ion:"report_time" dynamodbav:"report_time,omitempty" parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID                     string                          `bson:"_id,omitempty" ion:"_id" dynamodbav:"_id,omitempty" parquet:"name=inventory_uuid,type=BYTE_ARRAY,convertedtype=UTF8" json:"_id,omitempty" diff:"-"`
}
