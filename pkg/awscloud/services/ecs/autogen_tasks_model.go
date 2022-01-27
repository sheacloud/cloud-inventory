//AUTOGENERATED CODE DO NOT EDIT
package ecs

import (
	"time"
)

type Task struct {
	Attachments             []*Attachment           `parquet:"name=attachments,type=MAP,convertedtype=LIST" json:"attachments" diff:"attachments"`
	Attributes              []*Attribute            `parquet:"name=attributes,type=MAP,convertedtype=LIST" json:"attributes" diff:"attributes"`
	AvailabilityZone        string                  `parquet:"name=availability_zone,type=BYTE_ARRAY,convertedtype=UTF8" json:"availability_zone" diff:"availability_zone"`
	CapacityProviderName    string                  `parquet:"name=capacity_provider_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"capacity_provider_name" diff:"capacity_provider_name"`
	ClusterArn              string                  `parquet:"name=cluster_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"cluster_arn" diff:"cluster_arn"`
	Connectivity            string                  `parquet:"name=connectivity,type=BYTE_ARRAY,convertedtype=UTF8" json:"connectivity" diff:"connectivity"`
	ConnectivityAt          *time.Time              `json:"-"`
	ContainerInstanceArn    string                  `parquet:"name=container_instance_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"container_instance_arn" diff:"container_instance_arn"`
	Containers              []*Container            `parquet:"name=containers,type=MAP,convertedtype=LIST" json:"containers" diff:"containers"`
	Cpu                     string                  `parquet:"name=cpu,type=BYTE_ARRAY,convertedtype=UTF8" json:"cpu" diff:"cpu"`
	CreatedAt               *time.Time              `json:"-"`
	DesiredStatus           string                  `parquet:"name=desired_status,type=BYTE_ARRAY,convertedtype=UTF8" json:"desired_status" diff:"desired_status"`
	EnableExecuteCommand    bool                    `parquet:"name=enable_execute_command,type=BOOLEAN" json:"enable_execute_command" diff:"enable_execute_command"`
	EphemeralStorage        *EphemeralStorage       `parquet:"name=ephemeral_storage" json:"ephemeral_storage" diff:"ephemeral_storage"`
	ExecutionStoppedAt      *time.Time              `json:"-"`
	Group                   string                  `parquet:"name=group,type=BYTE_ARRAY,convertedtype=UTF8" json:"group" diff:"group"`
	HealthStatus            string                  `parquet:"name=health_status,type=BYTE_ARRAY,convertedtype=UTF8" json:"health_status" diff:"health_status"`
	InferenceAccelerators   []*InferenceAccelerator `parquet:"name=inference_accelerators,type=MAP,convertedtype=LIST" json:"inference_accelerators" diff:"inference_accelerators"`
	LastStatus              string                  `parquet:"name=last_status,type=BYTE_ARRAY,convertedtype=UTF8" json:"last_status" diff:"last_status"`
	LaunchType              string                  `parquet:"name=launch_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"launch_type" diff:"launch_type"`
	Memory                  string                  `parquet:"name=memory,type=BYTE_ARRAY,convertedtype=UTF8" json:"memory" diff:"memory"`
	Overrides               *TaskOverride           `parquet:"name=overrides" json:"overrides" diff:"overrides"`
	PlatformFamily          string                  `parquet:"name=platform_family,type=BYTE_ARRAY,convertedtype=UTF8" json:"platform_family" diff:"platform_family"`
	PlatformVersion         string                  `parquet:"name=platform_version,type=BYTE_ARRAY,convertedtype=UTF8" json:"platform_version" diff:"platform_version"`
	PullStartedAt           *time.Time              `json:"-"`
	PullStoppedAt           *time.Time              `json:"-"`
	StartedAt               *time.Time              `json:"-"`
	StartedBy               string                  `parquet:"name=started_by,type=BYTE_ARRAY,convertedtype=UTF8" json:"started_by" diff:"started_by"`
	StopCode                string                  `parquet:"name=stop_code,type=BYTE_ARRAY,convertedtype=UTF8" json:"stop_code" diff:"stop_code"`
	StoppedAt               *time.Time              `json:"-"`
	StoppedReason           string                  `parquet:"name=stopped_reason,type=BYTE_ARRAY,convertedtype=UTF8" json:"stopped_reason" diff:"stopped_reason"`
	StoppingAt              *time.Time              `json:"-"`
	Tags                    map[string]string       `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"tags" diff:"tags"`
	TaskArn                 string                  `parquet:"name=task_arn,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"task_arn" diff:"task_arn,identifier"`
	TaskDefinitionArn       string                  `parquet:"name=task_definition_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"task_definition_arn" diff:"task_definition_arn"`
	Version                 int64                   `parquet:"name=version,type=INT64" json:"version" diff:"version"`
	AccountId               string                  `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"account_id" diff:"account_id"`
	Region                  string                  `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8" json:"region" diff:"region"`
	ReportTime              int64                   `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"report_time" diff:"report_time,immutable"`
	ConnectivityAtMilli     int64                   `parquet:"name=connectivity_at,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"connectivity_at" diff:"connectivity_at"`
	CreatedAtMilli          int64                   `parquet:"name=created_at,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"created_at" diff:"created_at"`
	ExecutionStoppedAtMilli int64                   `parquet:"name=execution_stopped_at,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"execution_stopped_at" diff:"execution_stopped_at"`
	PullStartedAtMilli      int64                   `parquet:"name=pull_started_at,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"pull_started_at" diff:"pull_started_at"`
	PullStoppedAtMilli      int64                   `parquet:"name=pull_stopped_at,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"pull_stopped_at" diff:"pull_stopped_at"`
	StartedAtMilli          int64                   `parquet:"name=started_at,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"started_at" diff:"started_at"`
	StoppedAtMilli          int64                   `parquet:"name=stopped_at,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"stopped_at" diff:"stopped_at"`
	StoppingAtMilli         int64                   `parquet:"name=stopping_at,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"stopping_at" diff:"stopping_at"`
}

func (x *Task) GetReportTime() int64 {
	return x.ReportTime
}
