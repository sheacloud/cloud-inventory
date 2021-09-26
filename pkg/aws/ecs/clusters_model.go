// AUTOGENERATED, DO NOT EDIT
package ecs

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"sync"
	"time"
)

var customClusterModelPostprocessingFuncs []func(ctx context.Context, client *ecs.Client, cfg aws.Config, x *ClusterModel) = []func(ctx context.Context, client *ecs.Client, cfg aws.Config, x *ClusterModel){}
var customClusterModelFuncsLock sync.Mutex

func registerCustomClusterModelPostprocessingFunc(f func(ctx context.Context, client *ecs.Client, cfg aws.Config, x *ClusterModel)) {
	customClusterModelFuncsLock.Lock()
	defer customClusterModelFuncsLock.Unlock()

	customClusterModelPostprocessingFuncs = append(customClusterModelPostprocessingFuncs, f)
}

func init() {
	Controller.RegisterDataSource("clusters", ClusterDataSource)
}

type ClusterModel struct {
	ActiveServicesCount               int32                                       `parquet:"name=active_services_count,type=INT32"`
	Attachments                       []*AttachmentClusterModel                   `parquet:"name=attachments,type=MAP,convertedtype=LIST"`
	AttachmentsStatus                 string                                      `parquet:"name=attachments_status,type=BYTE_ARRAY,convertedtype=UTF8"`
	CapacityProviders                 []string                                    `parquet:"name=capacity_providers,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	ClusterArn                        string                                      `parquet:"name=cluster_arn,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	ClusterName                       string                                      `parquet:"name=cluster_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	Configuration                     *ClusterConfigurationClusterModel           `parquet:"name=configuration"`
	DefaultCapacityProviderStrategy   []*CapacityProviderStrategyItemClusterModel `parquet:"name=default_capacity_provider_strategy,type=MAP,convertedtype=LIST"`
	PendingTasksCount                 int32                                       `parquet:"name=pending_tasks_count,type=INT32"`
	RegisteredContainerInstancesCount int32                                       `parquet:"name=registered_container_instances_count,type=INT32"`
	RunningTasksCount                 int32                                       `parquet:"name=running_tasks_count,type=INT32"`
	Settings                          []*ClusterSettingClusterModel               `parquet:"name=settings,type=MAP,convertedtype=LIST"`
	Statistics                        []*KeyValuePairClusterModel                 `parquet:"name=statistics,type=MAP,convertedtype=LIST"`
	Status                            string                                      `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
	TagsOld                           []*TagClusterModel
	Tags                              map[string]string      `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
	AccountId                         string                 `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region                            string                 `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime                        int64                  `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	Services                          []*ServiceClusterModel `parquet:"name=services,type=MAP,convertedtype=LIST"`
	Tasks                             []*TaskClusterModel    `parquet:"name=tasks,type=MAP,convertedtype=LIST"`
}

type AttachmentClusterModel struct {
	Details []*KeyValuePairClusterModel `parquet:"name=details,type=MAP,convertedtype=LIST"`
	Id      string                      `parquet:"name=id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Status  string                      `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
	Type    string                      `parquet:"name=type,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type KeyValuePairClusterModel struct {
	Name  string `parquet:"name=name,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type ClusterConfigurationClusterModel struct {
	ExecuteCommandConfiguration *ExecuteCommandConfigurationClusterModel `parquet:"name=execute_command_configuration"`
}

type ExecuteCommandConfigurationClusterModel struct {
	KmsKeyId         string                                      `parquet:"name=kms_key_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	LogConfiguration *ExecuteCommandLogConfigurationClusterModel `parquet:"name=log_configuration"`
	Logging          string                                      `parquet:"name=logging,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type ExecuteCommandLogConfigurationClusterModel struct {
	CloudWatchEncryptionEnabled bool   `parquet:"name=cloud_watch_encryption_enabled,type=BOOLEAN"`
	CloudWatchLogGroupName      string `parquet:"name=cloud_watch_log_group_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	S3BucketName                string `parquet:"name=s3_bucket_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	S3EncryptionEnabled         bool   `parquet:"name=s3_encryption_enabled,type=BOOLEAN"`
	S3KeyPrefix                 string `parquet:"name=s3_key_prefix,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type CapacityProviderStrategyItemClusterModel struct {
	CapacityProvider string `parquet:"name=capacity_provider,type=BYTE_ARRAY,convertedtype=UTF8"`
	Base             int32  `parquet:"name=base,type=INT32"`
	Weight           int32  `parquet:"name=weight,type=INT32"`
}

type ClusterSettingClusterModel struct {
	Name  string `parquet:"name=name,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type TagClusterModel struct {
	Key   string `parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type ServiceClusterModel struct {
	CapacityProviderStrategy      []*CapacityProviderStrategyItemClusterModel `parquet:"name=capacity_provider_strategy,type=MAP,convertedtype=LIST"`
	ClusterArn                    string                                      `parquet:"name=cluster_arn,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	CreatedAt                     *time.Time
	CreatedBy                     string                               `parquet:"name=created_by,type=BYTE_ARRAY,convertedtype=UTF8"`
	DeploymentConfiguration       *DeploymentConfigurationClusterModel `parquet:"name=deployment_configuration"`
	DeploymentController          *DeploymentControllerClusterModel    `parquet:"name=deployment_controller"`
	Deployments                   []*DeploymentClusterModel            `parquet:"name=deployments,type=MAP,convertedtype=LIST"`
	DesiredCount                  int32                                `parquet:"name=desired_count,type=INT32"`
	EnableECSManagedTags          bool                                 `parquet:"name=enable_ecs_managed_tags,type=BOOLEAN"`
	EnableExecuteCommand          bool                                 `parquet:"name=enable_execute_command,type=BOOLEAN"`
	Events                        []*ServiceEventClusterModel          `parquet:"name=events,type=MAP,convertedtype=LIST"`
	HealthCheckGracePeriodSeconds int32                                `parquet:"name=health_check_grace_period_seconds,type=INT32"`
	LaunchType                    string                               `parquet:"name=launch_type,type=BYTE_ARRAY,convertedtype=UTF8"`
	LoadBalancers                 []*LoadBalancerClusterModel          `parquet:"name=load_balancers,type=MAP,convertedtype=LIST"`
	NetworkConfiguration          *NetworkConfigurationClusterModel    `parquet:"name=network_configuration"`
	PendingCount                  int32                                `parquet:"name=pending_count,type=INT32"`
	PlacementConstraints          []*PlacementConstraintClusterModel   `parquet:"name=placement_constraints,type=MAP,convertedtype=LIST"`
	PlacementStrategy             []*PlacementStrategyClusterModel     `parquet:"name=placement_strategy,type=MAP,convertedtype=LIST"`
	PlatformVersion               string                               `parquet:"name=platform_version,type=BYTE_ARRAY,convertedtype=UTF8"`
	PropagateTags                 string                               `parquet:"name=propagate_tags,type=BYTE_ARRAY,convertedtype=UTF8"`
	RoleArn                       string                               `parquet:"name=role_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	RunningCount                  int32                                `parquet:"name=running_count,type=INT32"`
	SchedulingStrategy            string                               `parquet:"name=scheduling_strategy,type=BYTE_ARRAY,convertedtype=UTF8"`
	ServiceArn                    string                               `parquet:"name=service_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	ServiceName                   string                               `parquet:"name=service_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	ServiceRegistries             []*ServiceRegistryClusterModel       `parquet:"name=service_registries,type=MAP,convertedtype=LIST"`
	Status                        string                               `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
	Tags                          []*TagClusterModel                   `parquet:"name=tags,type=MAP,convertedtype=LIST"`
	TaskDefinition                string                               `parquet:"name=task_definition,type=BYTE_ARRAY,convertedtype=UTF8"`
	TaskSets                      []*TaskSetClusterModel               `parquet:"name=task_sets,type=MAP,convertedtype=LIST"`
	CreatedAtMilli                int64                                `parquet:"name=created_at_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type DeploymentConfigurationClusterModel struct {
	DeploymentCircuitBreaker *DeploymentCircuitBreakerClusterModel `parquet:"name=deployment_circuit_breaker"`
	MaximumPercent           int32                                 `parquet:"name=maximum_percent,type=INT32"`
	MinimumHealthyPercent    int32                                 `parquet:"name=minimum_healthy_percent,type=INT32"`
}

type DeploymentCircuitBreakerClusterModel struct {
	Enable   bool `parquet:"name=enable,type=BOOLEAN"`
	Rollback bool `parquet:"name=rollback,type=BOOLEAN"`
}

type DeploymentControllerClusterModel struct {
	Type string `parquet:"name=type,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type DeploymentClusterModel struct {
	CapacityProviderStrategy []*CapacityProviderStrategyItemClusterModel `parquet:"name=capacity_provider_strategy,type=MAP,convertedtype=LIST"`
	CreatedAt                *time.Time
	DesiredCount             int32                             `parquet:"name=desired_count,type=INT32"`
	FailedTasks              int32                             `parquet:"name=failed_tasks,type=INT32"`
	Id                       string                            `parquet:"name=id,type=BYTE_ARRAY,convertedtype=UTF8"`
	LaunchType               string                            `parquet:"name=launch_type,type=BYTE_ARRAY,convertedtype=UTF8"`
	NetworkConfiguration     *NetworkConfigurationClusterModel `parquet:"name=network_configuration"`
	PendingCount             int32                             `parquet:"name=pending_count,type=INT32"`
	PlatformVersion          string                            `parquet:"name=platform_version,type=BYTE_ARRAY,convertedtype=UTF8"`
	RolloutState             string                            `parquet:"name=rollout_state,type=BYTE_ARRAY,convertedtype=UTF8"`
	RolloutStateReason       string                            `parquet:"name=rollout_state_reason,type=BYTE_ARRAY,convertedtype=UTF8"`
	RunningCount             int32                             `parquet:"name=running_count,type=INT32"`
	Status                   string                            `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
	TaskDefinition           string                            `parquet:"name=task_definition,type=BYTE_ARRAY,convertedtype=UTF8"`
	UpdatedAt                *time.Time
	CreatedAtMilli           int64 `parquet:"name=created_at_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	UpdatedAtMilli           int64 `parquet:"name=updated_at_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type NetworkConfigurationClusterModel struct {
	AwsvpcConfiguration *AwsVpcConfigurationClusterModel `parquet:"name=awsvpc_configuration"`
}

type AwsVpcConfigurationClusterModel struct {
	Subnets        []string `parquet:"name=subnets,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	AssignPublicIp string   `parquet:"name=assign_public_ip,type=BYTE_ARRAY,convertedtype=UTF8"`
	SecurityGroups []string `parquet:"name=security_groups,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
}

type ServiceEventClusterModel struct {
	CreatedAt      *time.Time
	Id             string `parquet:"name=id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Message        string `parquet:"name=message,type=BYTE_ARRAY,convertedtype=UTF8"`
	CreatedAtMilli int64  `parquet:"name=created_at_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type LoadBalancerClusterModel struct {
	ContainerName    string `parquet:"name=container_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	ContainerPort    int32  `parquet:"name=container_port,type=INT32"`
	LoadBalancerName string `parquet:"name=load_balancer_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	TargetGroupArn   string `parquet:"name=target_group_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type PlacementConstraintClusterModel struct {
	Expression string `parquet:"name=expression,type=BYTE_ARRAY,convertedtype=UTF8"`
	Type       string `parquet:"name=type,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type PlacementStrategyClusterModel struct {
	Field string `parquet:"name=field,type=BYTE_ARRAY,convertedtype=UTF8"`
	Type  string `parquet:"name=type,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type ServiceRegistryClusterModel struct {
	ContainerName string `parquet:"name=container_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	ContainerPort int32  `parquet:"name=container_port,type=INT32"`
	Port          int32  `parquet:"name=port,type=INT32"`
	RegistryArn   string `parquet:"name=registry_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type TaskSetClusterModel struct {
	CapacityProviderStrategy []*CapacityProviderStrategyItemClusterModel `parquet:"name=capacity_provider_strategy,type=MAP,convertedtype=LIST"`
	ClusterArn               string                                      `parquet:"name=cluster_arn,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	ComputedDesiredCount     int32                                       `parquet:"name=computed_desired_count,type=INT32"`
	CreatedAt                *time.Time
	ExternalId               string                            `parquet:"name=external_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Id                       string                            `parquet:"name=id,type=BYTE_ARRAY,convertedtype=UTF8"`
	LaunchType               string                            `parquet:"name=launch_type,type=BYTE_ARRAY,convertedtype=UTF8"`
	LoadBalancers            []*LoadBalancerClusterModel       `parquet:"name=load_balancers,type=MAP,convertedtype=LIST"`
	NetworkConfiguration     *NetworkConfigurationClusterModel `parquet:"name=network_configuration"`
	PendingCount             int32                             `parquet:"name=pending_count,type=INT32"`
	PlatformVersion          string                            `parquet:"name=platform_version,type=BYTE_ARRAY,convertedtype=UTF8"`
	RunningCount             int32                             `parquet:"name=running_count,type=INT32"`
	Scale                    *ScaleClusterModel                `parquet:"name=scale"`
	ServiceArn               string                            `parquet:"name=service_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	ServiceRegistries        []*ServiceRegistryClusterModel    `parquet:"name=service_registries,type=MAP,convertedtype=LIST"`
	StabilityStatus          string                            `parquet:"name=stability_status,type=BYTE_ARRAY,convertedtype=UTF8"`
	StabilityStatusAt        *time.Time
	StartedBy                string             `parquet:"name=started_by,type=BYTE_ARRAY,convertedtype=UTF8"`
	Status                   string             `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
	Tags                     []*TagClusterModel `parquet:"name=tags,type=MAP,convertedtype=LIST"`
	TaskDefinition           string             `parquet:"name=task_definition,type=BYTE_ARRAY,convertedtype=UTF8"`
	TaskSetArn               string             `parquet:"name=task_set_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	UpdatedAt                *time.Time
	CreatedAtMilli           int64 `parquet:"name=created_at_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	StabilityStatusAtMilli   int64 `parquet:"name=stability_status_at_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	UpdatedAtMilli           int64 `parquet:"name=updated_at_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type ScaleClusterModel struct {
	Unit  string  `parquet:"name=unit,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value float64 `parquet:"name=value,type=DOUBLE"`
}

type TaskClusterModel struct {
	Attachments             []*AttachmentClusterModel `parquet:"name=attachments,type=MAP,convertedtype=LIST"`
	Attributes              []*AttributeClusterModel  `parquet:"name=attributes,type=MAP,convertedtype=LIST"`
	AvailabilityZone        string                    `parquet:"name=availability_zone,type=BYTE_ARRAY,convertedtype=UTF8"`
	CapacityProviderName    string                    `parquet:"name=capacity_provider_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	ClusterArn              string                    `parquet:"name=cluster_arn,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	Connectivity            string                    `parquet:"name=connectivity,type=BYTE_ARRAY,convertedtype=UTF8"`
	ConnectivityAt          *time.Time
	ContainerInstanceArn    string                   `parquet:"name=container_instance_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	Containers              []*ContainerClusterModel `parquet:"name=containers,type=MAP,convertedtype=LIST"`
	Cpu                     string                   `parquet:"name=cpu,type=BYTE_ARRAY,convertedtype=UTF8"`
	CreatedAt               *time.Time
	DesiredStatus           string                        `parquet:"name=desired_status,type=BYTE_ARRAY,convertedtype=UTF8"`
	EnableExecuteCommand    bool                          `parquet:"name=enable_execute_command,type=BOOLEAN"`
	EphemeralStorage        *EphemeralStorageClusterModel `parquet:"name=ephemeral_storage"`
	ExecutionStoppedAt      *time.Time
	Group                   string                              `parquet:"name=group,type=BYTE_ARRAY,convertedtype=UTF8"`
	HealthStatus            string                              `parquet:"name=health_status,type=BYTE_ARRAY,convertedtype=UTF8"`
	InferenceAccelerators   []*InferenceAcceleratorClusterModel `parquet:"name=inference_accelerators,type=MAP,convertedtype=LIST"`
	LastStatus              string                              `parquet:"name=last_status,type=BYTE_ARRAY,convertedtype=UTF8"`
	LaunchType              string                              `parquet:"name=launch_type,type=BYTE_ARRAY,convertedtype=UTF8"`
	Memory                  string                              `parquet:"name=memory,type=BYTE_ARRAY,convertedtype=UTF8"`
	Overrides               *TaskOverrideClusterModel           `parquet:"name=overrides"`
	PlatformVersion         string                              `parquet:"name=platform_version,type=BYTE_ARRAY,convertedtype=UTF8"`
	PullStartedAt           *time.Time
	PullStoppedAt           *time.Time
	StartedAt               *time.Time
	StartedBy               string `parquet:"name=started_by,type=BYTE_ARRAY,convertedtype=UTF8"`
	StopCode                string `parquet:"name=stop_code,type=BYTE_ARRAY,convertedtype=UTF8"`
	StoppedAt               *time.Time
	StoppedReason           string `parquet:"name=stopped_reason,type=BYTE_ARRAY,convertedtype=UTF8"`
	StoppingAt              *time.Time
	Tags                    []*TagClusterModel `parquet:"name=tags,type=MAP,convertedtype=LIST"`
	TaskArn                 string             `parquet:"name=task_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	TaskDefinitionArn       string             `parquet:"name=task_definition_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	Version                 int64              `parquet:"name=version,type=INT64"`
	ConnectivityAtMilli     int64              `parquet:"name=connectivity_at_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	CreatedAtMilli          int64              `parquet:"name=created_at_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	ExecutionStoppedAtMilli int64              `parquet:"name=execution_stopped_at_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	PullStartedAtMilli      int64              `parquet:"name=pull_started_at_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	PullStoppedAtMilli      int64              `parquet:"name=pull_stopped_at_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	StartedAtMilli          int64              `parquet:"name=started_at_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	StoppedAtMilli          int64              `parquet:"name=stopped_at_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	StoppingAtMilli         int64              `parquet:"name=stopping_at_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type AttributeClusterModel struct {
	Name       string `parquet:"name=name,type=BYTE_ARRAY,convertedtype=UTF8"`
	TargetId   string `parquet:"name=target_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	TargetType string `parquet:"name=target_type,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value      string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type ContainerClusterModel struct {
	ContainerArn      string                          `parquet:"name=container_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	Cpu               string                          `parquet:"name=cpu,type=BYTE_ARRAY,convertedtype=UTF8"`
	ExitCode          int32                           `parquet:"name=exit_code,type=INT32"`
	GpuIds            []string                        `parquet:"name=gpu_ids,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	HealthStatus      string                          `parquet:"name=health_status,type=BYTE_ARRAY,convertedtype=UTF8"`
	Image             string                          `parquet:"name=image,type=BYTE_ARRAY,convertedtype=UTF8"`
	ImageDigest       string                          `parquet:"name=image_digest,type=BYTE_ARRAY,convertedtype=UTF8"`
	LastStatus        string                          `parquet:"name=last_status,type=BYTE_ARRAY,convertedtype=UTF8"`
	ManagedAgents     []*ManagedAgentClusterModel     `parquet:"name=managed_agents,type=MAP,convertedtype=LIST"`
	Memory            string                          `parquet:"name=memory,type=BYTE_ARRAY,convertedtype=UTF8"`
	MemoryReservation string                          `parquet:"name=memory_reservation,type=BYTE_ARRAY,convertedtype=UTF8"`
	Name              string                          `parquet:"name=name,type=BYTE_ARRAY,convertedtype=UTF8"`
	NetworkBindings   []*NetworkBindingClusterModel   `parquet:"name=network_bindings,type=MAP,convertedtype=LIST"`
	NetworkInterfaces []*NetworkInterfaceClusterModel `parquet:"name=network_interfaces,type=MAP,convertedtype=LIST"`
	Reason            string                          `parquet:"name=reason,type=BYTE_ARRAY,convertedtype=UTF8"`
	RuntimeId         string                          `parquet:"name=runtime_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	TaskArn           string                          `parquet:"name=task_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type ManagedAgentClusterModel struct {
	LastStartedAt      *time.Time
	LastStatus         string `parquet:"name=last_status,type=BYTE_ARRAY,convertedtype=UTF8"`
	Name               string `parquet:"name=name,type=BYTE_ARRAY,convertedtype=UTF8"`
	Reason             string `parquet:"name=reason,type=BYTE_ARRAY,convertedtype=UTF8"`
	LastStartedAtMilli int64  `parquet:"name=last_started_at_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type NetworkBindingClusterModel struct {
	BindIP        string `parquet:"name=bind_ip,type=BYTE_ARRAY,convertedtype=UTF8"`
	ContainerPort int32  `parquet:"name=container_port,type=INT32"`
	HostPort      int32  `parquet:"name=host_port,type=INT32"`
	Protocol      string `parquet:"name=protocol,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type NetworkInterfaceClusterModel struct {
	AttachmentId       string `parquet:"name=attachment_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Ipv6Address        string `parquet:"name=ipv6_address,type=BYTE_ARRAY,convertedtype=UTF8"`
	PrivateIpv4Address string `parquet:"name=private_ipv4_address,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type EphemeralStorageClusterModel struct {
	SizeInGiB int32 `parquet:"name=size_in_gi_b,type=INT32"`
}

type InferenceAcceleratorClusterModel struct {
	DeviceName string `parquet:"name=device_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	DeviceType string `parquet:"name=device_type,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type TaskOverrideClusterModel struct {
	ContainerOverrides            []*ContainerOverrideClusterModel            `parquet:"name=container_overrides,type=MAP,convertedtype=LIST"`
	Cpu                           string                                      `parquet:"name=cpu,type=BYTE_ARRAY,convertedtype=UTF8"`
	EphemeralStorage              *EphemeralStorageClusterModel               `parquet:"name=ephemeral_storage"`
	ExecutionRoleArn              string                                      `parquet:"name=execution_role_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	InferenceAcceleratorOverrides []*InferenceAcceleratorOverrideClusterModel `parquet:"name=inference_accelerator_overrides,type=MAP,convertedtype=LIST"`
	Memory                        string                                      `parquet:"name=memory,type=BYTE_ARRAY,convertedtype=UTF8"`
	TaskRoleArn                   string                                      `parquet:"name=task_role_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type ContainerOverrideClusterModel struct {
	Command              []string                           `parquet:"name=command,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	Cpu                  int32                              `parquet:"name=cpu,type=INT32"`
	Environment          []*KeyValuePairClusterModel        `parquet:"name=environment,type=MAP,convertedtype=LIST"`
	EnvironmentFiles     []*EnvironmentFileClusterModel     `parquet:"name=environment_files,type=MAP,convertedtype=LIST"`
	Memory               int32                              `parquet:"name=memory,type=INT32"`
	MemoryReservation    int32                              `parquet:"name=memory_reservation,type=INT32"`
	Name                 string                             `parquet:"name=name,type=BYTE_ARRAY,convertedtype=UTF8"`
	ResourceRequirements []*ResourceRequirementClusterModel `parquet:"name=resource_requirements,type=MAP,convertedtype=LIST"`
}

type EnvironmentFileClusterModel struct {
	Type  string `parquet:"name=type,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type ResourceRequirementClusterModel struct {
	Type  string `parquet:"name=type,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type InferenceAcceleratorOverrideClusterModel struct {
	DeviceName string `parquet:"name=device_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	DeviceType string `parquet:"name=device_type,type=BYTE_ARRAY,convertedtype=UTF8"`
}
