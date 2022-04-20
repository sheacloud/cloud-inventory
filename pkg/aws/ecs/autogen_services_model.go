//AUTOGENERATED CODE DO NOT EDIT
package ecs

import (
	"time"
)

type Service struct {
	CapacityProviderStrategy      []*CapacityProviderStrategyItem `bson:"capacity_provider_strategy,omitempty" dynamodbav:"capacity_provider_strategy,omitempty" json:"capacity_provider_strategy,omitempty" diff:"capacity_provider_strategy"`
	ClusterArn                    string                          `bson:"cluster_arn,omitempty" dynamodbav:"cluster_arn,omitempty" json:"cluster_arn,omitempty" diff:"cluster_arn"`
	CreatedAt                     *time.Time                      `bson:"created_at,omitempty" dynamodbav:"created_at,unixtime,omitempty" json:"created_at,omitempty" diff:"created_at"`
	CreatedBy                     string                          `bson:"created_by,omitempty" dynamodbav:"created_by,omitempty" json:"created_by,omitempty" diff:"created_by"`
	DeploymentConfiguration       *DeploymentConfiguration        `bson:"deployment_configuration,omitempty" dynamodbav:"deployment_configuration,omitempty" json:"deployment_configuration,omitempty" diff:"deployment_configuration"`
	DeploymentController          *DeploymentController           `bson:"deployment_controller,omitempty" dynamodbav:"deployment_controller,omitempty" json:"deployment_controller,omitempty" diff:"deployment_controller"`
	Deployments                   []*Deployment                   `bson:"deployments,omitempty" dynamodbav:"deployments,omitempty" json:"deployments,omitempty" diff:"deployments"`
	DesiredCount                  int32                           `bson:"desired_count,omitempty" dynamodbav:"desired_count,omitempty" json:"desired_count,omitempty" diff:"desired_count"`
	EnableECSManagedTags          bool                            `bson:"enable_ecs_managed_tags,omitempty" dynamodbav:"enable_ecs_managed_tags" json:"enable_ecs_managed_tags,omitempty" diff:"enable_ecs_managed_tags"`
	EnableExecuteCommand          bool                            `bson:"enable_execute_command,omitempty" dynamodbav:"enable_execute_command" json:"enable_execute_command,omitempty" diff:"enable_execute_command"`
	HealthCheckGracePeriodSeconds int32                           `bson:"health_check_grace_period_seconds,omitempty" dynamodbav:"health_check_grace_period_seconds,omitempty" json:"health_check_grace_period_seconds,omitempty" diff:"health_check_grace_period_seconds"`
	LaunchType                    string                          `bson:"launch_type,omitempty" dynamodbav:"launch_type,omitempty" json:"launch_type,omitempty" diff:"launch_type"`
	LoadBalancers                 []*LoadBalancer                 `bson:"load_balancers,omitempty" dynamodbav:"load_balancers,omitempty" json:"load_balancers,omitempty" diff:"load_balancers"`
	NetworkConfiguration          *NetworkConfiguration           `bson:"network_configuration,omitempty" dynamodbav:"network_configuration,omitempty" json:"network_configuration,omitempty" diff:"network_configuration"`
	PendingCount                  int32                           `bson:"pending_count,omitempty" dynamodbav:"pending_count,omitempty" json:"pending_count,omitempty" diff:"pending_count"`
	PlacementConstraints          []*PlacementConstraint          `bson:"placement_constraints,omitempty" dynamodbav:"placement_constraints,omitempty" json:"placement_constraints,omitempty" diff:"placement_constraints"`
	PlacementStrategy             []*PlacementStrategy            `bson:"placement_strategy,omitempty" dynamodbav:"placement_strategy,omitempty" json:"placement_strategy,omitempty" diff:"placement_strategy"`
	PlatformFamily                string                          `bson:"platform_family,omitempty" dynamodbav:"platform_family,omitempty" json:"platform_family,omitempty" diff:"platform_family"`
	PlatformVersion               string                          `bson:"platform_version,omitempty" dynamodbav:"platform_version,omitempty" json:"platform_version,omitempty" diff:"platform_version"`
	PropagateTags                 string                          `bson:"propagate_tags,omitempty" dynamodbav:"propagate_tags,omitempty" json:"propagate_tags,omitempty" diff:"propagate_tags"`
	RoleArn                       string                          `bson:"role_arn,omitempty" dynamodbav:"role_arn,omitempty" json:"role_arn,omitempty" diff:"role_arn"`
	RunningCount                  int32                           `bson:"running_count,omitempty" dynamodbav:"running_count,omitempty" json:"running_count,omitempty" diff:"running_count"`
	SchedulingStrategy            string                          `bson:"scheduling_strategy,omitempty" dynamodbav:"scheduling_strategy,omitempty" json:"scheduling_strategy,omitempty" diff:"scheduling_strategy"`
	ServiceArn                    string                          `bson:"service_arn,omitempty" dynamodbav:"service_arn,omitempty" inventory_primary_key:"true" json:"service_arn,omitempty" diff:"service_arn,identifier"`
	ServiceName                   string                          `bson:"service_name,omitempty" dynamodbav:"service_name,omitempty" json:"service_name,omitempty" diff:"service_name"`
	ServiceRegistries             []*ServiceRegistry              `bson:"service_registries,omitempty" dynamodbav:"service_registries,omitempty" json:"service_registries,omitempty" diff:"service_registries"`
	Status                        string                          `bson:"status,omitempty" dynamodbav:"status,omitempty" json:"status,omitempty" diff:"status"`
	Tags                          map[string]string               `bson:"tags,omitempty" dynamodbav:"tags,omitempty" json:"tags,omitempty" diff:"tags"`
	TaskDefinition                string                          `bson:"task_definition,omitempty" dynamodbav:"task_definition,omitempty" json:"task_definition,omitempty" diff:"task_definition"`
	TaskSets                      []*TaskSet                      `bson:"task_sets,omitempty" dynamodbav:"task_sets,omitempty" json:"task_sets,omitempty" diff:"task_sets"`
	AccountId                     string                          `bson:"account_id,omitempty" dynamodbav:"account_id,omitempty" json:"account_id,omitempty" diff:"account_id"`
	Region                        string                          `bson:"region,omitempty" dynamodbav:"region,omitempty" json:"region,omitempty" diff:"region"`
	ReportTime                    time.Time                       `bson:"report_time,omitempty" dynamodbav:"report_time,unixtime,omitempty" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID                 string                          `bson:"_id,omitempty" dynamodbav:"_id,omitempty" json:"_id,omitempty" diff:"-"`
}
