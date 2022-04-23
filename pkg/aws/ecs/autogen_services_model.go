//AUTOGENERATED CODE DO NOT EDIT
package ecs

type Service struct {
	CapacityProviderStrategy      []*CapacityProviderStrategyItem `bson:"capacity_provider_strategy,omitempty" ion:"capacity_provider_strategy" dynamodbav:"capacity_provider_strategy,omitempty" parquet:"name=capacity_provider_strategy,type=MAP,convertedtype=LIST" json:"capacity_provider_strategy,omitempty" diff:"capacity_provider_strategy"`
	ClusterArn                    string                          `bson:"cluster_arn,omitempty" ion:"cluster_arn" dynamodbav:"cluster_arn,omitempty" parquet:"name=cluster_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"cluster_arn,omitempty" diff:"cluster_arn"`
	CreatedAt                     int64                           `bson:"created_at,omitempty" ion:"created_at" dynamodbav:"created_at,omitempty" parquet:"name=created_at,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"created_at,omitempty" diff:"created_at"`
	CreatedBy                     string                          `bson:"created_by,omitempty" ion:"created_by" dynamodbav:"created_by,omitempty" parquet:"name=created_by,type=BYTE_ARRAY,convertedtype=UTF8" json:"created_by,omitempty" diff:"created_by"`
	DeploymentConfiguration       *DeploymentConfiguration        `bson:"deployment_configuration,omitempty" ion:"deployment_configuration" dynamodbav:"deployment_configuration,omitempty" parquet:"name=deployment_configuration" json:"deployment_configuration,omitempty" diff:"deployment_configuration"`
	DeploymentController          *DeploymentController           `bson:"deployment_controller,omitempty" ion:"deployment_controller" dynamodbav:"deployment_controller,omitempty" parquet:"name=deployment_controller" json:"deployment_controller,omitempty" diff:"deployment_controller"`
	Deployments                   []*Deployment                   `bson:"deployments,omitempty" ion:"deployments" dynamodbav:"deployments,omitempty" parquet:"name=deployments,type=MAP,convertedtype=LIST" json:"deployments,omitempty" diff:"deployments"`
	DesiredCount                  int32                           `bson:"desired_count,omitempty" ion:"desired_count" dynamodbav:"desired_count,omitempty" parquet:"name=desired_count,type=INT32" json:"desired_count,omitempty" diff:"desired_count"`
	EnableECSManagedTags          bool                            `bson:"enable_ecs_managed_tags,omitempty" ion:"enable_ecs_managed_tags" dynamodbav:"enable_ecs_managed_tags" parquet:"name=enable_ecs_managed_tags,type=BOOLEAN" json:"enable_ecs_managed_tags,omitempty" diff:"enable_ecs_managed_tags"`
	EnableExecuteCommand          bool                            `bson:"enable_execute_command,omitempty" ion:"enable_execute_command" dynamodbav:"enable_execute_command" parquet:"name=enable_execute_command,type=BOOLEAN" json:"enable_execute_command,omitempty" diff:"enable_execute_command"`
	HealthCheckGracePeriodSeconds int32                           `bson:"health_check_grace_period_seconds,omitempty" ion:"health_check_grace_period_seconds" dynamodbav:"health_check_grace_period_seconds,omitempty" parquet:"name=health_check_grace_period_seconds,type=INT32" json:"health_check_grace_period_seconds,omitempty" diff:"health_check_grace_period_seconds"`
	LaunchType                    string                          `bson:"launch_type,omitempty" ion:"launch_type" dynamodbav:"launch_type,omitempty" parquet:"name=launch_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"launch_type,omitempty" diff:"launch_type"`
	LoadBalancers                 []*LoadBalancer                 `bson:"load_balancers,omitempty" ion:"load_balancers" dynamodbav:"load_balancers,omitempty" parquet:"name=load_balancers,type=MAP,convertedtype=LIST" json:"load_balancers,omitempty" diff:"load_balancers"`
	NetworkConfiguration          *NetworkConfiguration           `bson:"network_configuration,omitempty" ion:"network_configuration" dynamodbav:"network_configuration,omitempty" parquet:"name=network_configuration" json:"network_configuration,omitempty" diff:"network_configuration"`
	PendingCount                  int32                           `bson:"pending_count,omitempty" ion:"pending_count" dynamodbav:"pending_count,omitempty" parquet:"name=pending_count,type=INT32" json:"pending_count,omitempty" diff:"pending_count"`
	PlacementConstraints          []*PlacementConstraint          `bson:"placement_constraints,omitempty" ion:"placement_constraints" dynamodbav:"placement_constraints,omitempty" parquet:"name=placement_constraints,type=MAP,convertedtype=LIST" json:"placement_constraints,omitempty" diff:"placement_constraints"`
	PlacementStrategy             []*PlacementStrategy            `bson:"placement_strategy,omitempty" ion:"placement_strategy" dynamodbav:"placement_strategy,omitempty" parquet:"name=placement_strategy,type=MAP,convertedtype=LIST" json:"placement_strategy,omitempty" diff:"placement_strategy"`
	PlatformFamily                string                          `bson:"platform_family,omitempty" ion:"platform_family" dynamodbav:"platform_family,omitempty" parquet:"name=platform_family,type=BYTE_ARRAY,convertedtype=UTF8" json:"platform_family,omitempty" diff:"platform_family"`
	PlatformVersion               string                          `bson:"platform_version,omitempty" ion:"platform_version" dynamodbav:"platform_version,omitempty" parquet:"name=platform_version,type=BYTE_ARRAY,convertedtype=UTF8" json:"platform_version,omitempty" diff:"platform_version"`
	PropagateTags                 string                          `bson:"propagate_tags,omitempty" ion:"propagate_tags" dynamodbav:"propagate_tags,omitempty" parquet:"name=propagate_tags,type=BYTE_ARRAY,convertedtype=UTF8" json:"propagate_tags,omitempty" diff:"propagate_tags"`
	RoleArn                       string                          `bson:"role_arn,omitempty" ion:"role_arn" dynamodbav:"role_arn,omitempty" parquet:"name=role_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"role_arn,omitempty" diff:"role_arn"`
	RunningCount                  int32                           `bson:"running_count,omitempty" ion:"running_count" dynamodbav:"running_count,omitempty" parquet:"name=running_count,type=INT32" json:"running_count,omitempty" diff:"running_count"`
	SchedulingStrategy            string                          `bson:"scheduling_strategy,omitempty" ion:"scheduling_strategy" dynamodbav:"scheduling_strategy,omitempty" parquet:"name=scheduling_strategy,type=BYTE_ARRAY,convertedtype=UTF8" json:"scheduling_strategy,omitempty" diff:"scheduling_strategy"`
	ServiceArn                    string                          `bson:"service_arn,omitempty" ion:"service_arn" dynamodbav:"service_arn,omitempty" parquet:"name=service_arn,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"service_arn,omitempty" diff:"service_arn,identifier"`
	ServiceName                   string                          `bson:"service_name,omitempty" ion:"service_name" dynamodbav:"service_name,omitempty" parquet:"name=service_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"service_name,omitempty" diff:"service_name"`
	ServiceRegistries             []*ServiceRegistry              `bson:"service_registries,omitempty" ion:"service_registries" dynamodbav:"service_registries,omitempty" parquet:"name=service_registries,type=MAP,convertedtype=LIST" json:"service_registries,omitempty" diff:"service_registries"`
	Status                        string                          `bson:"status,omitempty" ion:"status" dynamodbav:"status,omitempty" parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8" json:"status,omitempty" diff:"status"`
	Tags                          map[string]string               `bson:"tags,omitempty" ion:"tags" dynamodbav:"tags,omitempty" parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"tags,omitempty" diff:"tags"`
	TaskDefinition                string                          `bson:"task_definition,omitempty" ion:"task_definition" dynamodbav:"task_definition,omitempty" parquet:"name=task_definition,type=BYTE_ARRAY,convertedtype=UTF8" json:"task_definition,omitempty" diff:"task_definition"`
	TaskSets                      []*TaskSet                      `bson:"task_sets,omitempty" ion:"task_sets" dynamodbav:"task_sets,omitempty" parquet:"name=task_sets,type=MAP,convertedtype=LIST" json:"task_sets,omitempty" diff:"task_sets"`
	AccountId                     string                          `bson:"account_id,omitempty" ion:"account_id" dynamodbav:"account_id,omitempty" parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"account_id,omitempty" diff:"account_id"`
	Region                        string                          `bson:"region,omitempty" ion:"region" dynamodbav:"region,omitempty" parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8" json:"region,omitempty" diff:"region"`
	ReportTime                    int64                           `bson:"report_time,omitempty" ion:"report_time" dynamodbav:"report_time,omitempty" parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID                 string                          `bson:"_id,omitempty" ion:"_id" dynamodbav:"_id,omitempty" parquet:"name=inventory_uuid,type=BYTE_ARRAY,convertedtype=UTF8" json:"_id,omitempty" diff:"-"`
}
