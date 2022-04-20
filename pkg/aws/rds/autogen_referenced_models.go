//AUTOGENERATED CODE DO NOT EDIT
package rds

import (
	"time"
)

type DBClusterRole struct {
	FeatureName string `bson:"feature_name,omitempty" dynamodbav:"feature_name,omitempty" json:"feature_name,omitempty" diff:"feature_name"`
	RoleArn     string `bson:"role_arn,omitempty" dynamodbav:"role_arn,omitempty" json:"role_arn,omitempty" diff:"role_arn"`
	Status      string `bson:"status,omitempty" dynamodbav:"status,omitempty" json:"status,omitempty" diff:"status"`
}

type DBClusterMember struct {
	DBClusterParameterGroupStatus string `bson:"db_cluster_parameter_group_status,omitempty" dynamodbav:"db_cluster_parameter_group_status,omitempty" json:"db_cluster_parameter_group_status,omitempty" diff:"db_cluster_parameter_group_status"`
	DBInstanceIdentifier          string `bson:"db_instance_identifier,omitempty" dynamodbav:"db_instance_identifier,omitempty" json:"db_instance_identifier,omitempty" diff:"db_instance_identifier"`
	IsClusterWriter               bool   `bson:"is_cluster_writer,omitempty" dynamodbav:"is_cluster_writer" json:"is_cluster_writer,omitempty" diff:"is_cluster_writer"`
	PromotionTier                 int32  `bson:"promotion_tier,omitempty" dynamodbav:"promotion_tier,omitempty" json:"promotion_tier,omitempty" diff:"promotion_tier"`
}

type DBClusterOptionGroupStatus struct {
	DBClusterOptionGroupName string `bson:"db_cluster_option_group_name,omitempty" dynamodbav:"db_cluster_option_group_name,omitempty" json:"db_cluster_option_group_name,omitempty" diff:"db_cluster_option_group_name"`
	Status                   string `bson:"status,omitempty" dynamodbav:"status,omitempty" json:"status,omitempty" diff:"status"`
}

type DomainMembership struct {
	Domain      string `bson:"domain,omitempty" dynamodbav:"domain,omitempty" json:"domain,omitempty" diff:"domain"`
	FQDN        string `bson:"fqdn,omitempty" dynamodbav:"fqdn,omitempty" json:"fqdn,omitempty" diff:"fqdn"`
	IAMRoleName string `bson:"iam_role_name,omitempty" dynamodbav:"iam_role_name,omitempty" json:"iam_role_name,omitempty" diff:"iam_role_name"`
	Status      string `bson:"status,omitempty" dynamodbav:"status,omitempty" json:"status,omitempty" diff:"status"`
}

type ClusterPendingModifiedValues struct {
	DBClusterIdentifier              string                        `bson:"db_cluster_identifier,omitempty" dynamodbav:"db_cluster_identifier,omitempty" json:"db_cluster_identifier,omitempty" diff:"db_cluster_identifier"`
	EngineVersion                    string                        `bson:"engine_version,omitempty" dynamodbav:"engine_version,omitempty" json:"engine_version,omitempty" diff:"engine_version"`
	IAMDatabaseAuthenticationEnabled bool                          `bson:"iam_database_authentication_enabled,omitempty" dynamodbav:"iam_database_authentication_enabled" json:"iam_database_authentication_enabled,omitempty" diff:"iam_database_authentication_enabled"`
	MasterUserPassword               string                        `bson:"master_user_password,omitempty" dynamodbav:"master_user_password,omitempty" json:"master_user_password,omitempty" diff:"master_user_password"`
	PendingCloudwatchLogsExports     *PendingCloudwatchLogsExports `bson:"pending_cloudwatch_logs_exports,omitempty" dynamodbav:"pending_cloudwatch_logs_exports,omitempty" json:"pending_cloudwatch_logs_exports,omitempty" diff:"pending_cloudwatch_logs_exports"`
}

type PendingCloudwatchLogsExports struct {
	LogTypesToDisable []string `bson:"log_types_to_disable,omitempty" dynamodbav:"log_types_to_disable,omitempty" json:"log_types_to_disable,omitempty" diff:"log_types_to_disable"`
	LogTypesToEnable  []string `bson:"log_types_to_enable,omitempty" dynamodbav:"log_types_to_enable,omitempty" json:"log_types_to_enable,omitempty" diff:"log_types_to_enable"`
}

type ScalingConfigurationInfo struct {
	AutoPause             bool   `bson:"auto_pause,omitempty" dynamodbav:"auto_pause" json:"auto_pause,omitempty" diff:"auto_pause"`
	MaxCapacity           int32  `bson:"max_capacity,omitempty" dynamodbav:"max_capacity,omitempty" json:"max_capacity,omitempty" diff:"max_capacity"`
	MinCapacity           int32  `bson:"min_capacity,omitempty" dynamodbav:"min_capacity,omitempty" json:"min_capacity,omitempty" diff:"min_capacity"`
	SecondsBeforeTimeout  int32  `bson:"seconds_before_timeout,omitempty" dynamodbav:"seconds_before_timeout,omitempty" json:"seconds_before_timeout,omitempty" diff:"seconds_before_timeout"`
	SecondsUntilAutoPause int32  `bson:"seconds_until_auto_pause,omitempty" dynamodbav:"seconds_until_auto_pause,omitempty" json:"seconds_until_auto_pause,omitempty" diff:"seconds_until_auto_pause"`
	TimeoutAction         string `bson:"timeout_action,omitempty" dynamodbav:"timeout_action,omitempty" json:"timeout_action,omitempty" diff:"timeout_action"`
}

type Tag struct {
	Key   string `bson:"key,omitempty" dynamodbav:"key,omitempty" json:"key,omitempty" diff:"key"`
	Value string `bson:"value,omitempty" dynamodbav:"value,omitempty" json:"value,omitempty" diff:"value"`
}

type VpcSecurityGroupMembership struct {
	Status             string `bson:"status,omitempty" dynamodbav:"status,omitempty" json:"status,omitempty" diff:"status"`
	VpcSecurityGroupId string `bson:"vpc_security_group_id,omitempty" dynamodbav:"vpc_security_group_id,omitempty" json:"vpc_security_group_id,omitempty" diff:"vpc_security_group_id"`
}

type DBInstanceRole struct {
	FeatureName string `bson:"feature_name,omitempty" dynamodbav:"feature_name,omitempty" json:"feature_name,omitempty" diff:"feature_name"`
	RoleArn     string `bson:"role_arn,omitempty" dynamodbav:"role_arn,omitempty" json:"role_arn,omitempty" diff:"role_arn"`
	Status      string `bson:"status,omitempty" dynamodbav:"status,omitempty" json:"status,omitempty" diff:"status"`
}

type DBInstanceAutomatedBackupsReplication struct {
	DBInstanceAutomatedBackupsArn string `bson:"db_instance_automated_backups_arn,omitempty" dynamodbav:"db_instance_automated_backups_arn,omitempty" json:"db_instance_automated_backups_arn,omitempty" diff:"db_instance_automated_backups_arn"`
}

type DBParameterGroupStatus struct {
	DBParameterGroupName string `bson:"db_parameter_group_name,omitempty" dynamodbav:"db_parameter_group_name,omitempty" json:"db_parameter_group_name,omitempty" diff:"db_parameter_group_name"`
	ParameterApplyStatus string `bson:"parameter_apply_status,omitempty" dynamodbav:"parameter_apply_status,omitempty" json:"parameter_apply_status,omitempty" diff:"parameter_apply_status"`
}

type DBSecurityGroupMembership struct {
	DBSecurityGroupName string `bson:"db_security_group_name,omitempty" dynamodbav:"db_security_group_name,omitempty" json:"db_security_group_name,omitempty" diff:"db_security_group_name"`
	Status              string `bson:"status,omitempty" dynamodbav:"status,omitempty" json:"status,omitempty" diff:"status"`
}

type DBSubnetGroup struct {
	DBSubnetGroupArn         string    `bson:"db_subnet_group_arn,omitempty" dynamodbav:"db_subnet_group_arn,omitempty" json:"db_subnet_group_arn,omitempty" diff:"db_subnet_group_arn"`
	DBSubnetGroupDescription string    `bson:"db_subnet_group_description,omitempty" dynamodbav:"db_subnet_group_description,omitempty" json:"db_subnet_group_description,omitempty" diff:"db_subnet_group_description"`
	DBSubnetGroupName        string    `bson:"db_subnet_group_name,omitempty" dynamodbav:"db_subnet_group_name,omitempty" json:"db_subnet_group_name,omitempty" diff:"db_subnet_group_name"`
	SubnetGroupStatus        string    `bson:"subnet_group_status,omitempty" dynamodbav:"subnet_group_status,omitempty" json:"subnet_group_status,omitempty" diff:"subnet_group_status"`
	Subnets                  []*Subnet `bson:"subnets,omitempty" dynamodbav:"subnets,omitempty" json:"subnets,omitempty" diff:"subnets"`
	VpcId                    string    `bson:"vpc_id,omitempty" dynamodbav:"vpc_id,omitempty" json:"vpc_id,omitempty" diff:"vpc_id"`
}

type Subnet struct {
	SubnetAvailabilityZone *AvailabilityZone `bson:"subnet_availability_zone,omitempty" dynamodbav:"subnet_availability_zone,omitempty" json:"subnet_availability_zone,omitempty" diff:"subnet_availability_zone"`
	SubnetIdentifier       string            `bson:"subnet_identifier,omitempty" dynamodbav:"subnet_identifier,omitempty" json:"subnet_identifier,omitempty" diff:"subnet_identifier"`
	SubnetOutpost          *Outpost          `bson:"subnet_outpost,omitempty" dynamodbav:"subnet_outpost,omitempty" json:"subnet_outpost,omitempty" diff:"subnet_outpost"`
	SubnetStatus           string            `bson:"subnet_status,omitempty" dynamodbav:"subnet_status,omitempty" json:"subnet_status,omitempty" diff:"subnet_status"`
}

type AvailabilityZone struct {
	Name string `bson:"name,omitempty" dynamodbav:"name,omitempty" json:"name,omitempty" diff:"name"`
}

type Outpost struct {
	Arn string `bson:"arn,omitempty" dynamodbav:"arn,omitempty" json:"arn,omitempty" diff:"arn"`
}

type Endpoint struct {
	Address      string `bson:"address,omitempty" dynamodbav:"address,omitempty" json:"address,omitempty" diff:"address"`
	HostedZoneId string `bson:"hosted_zone_id,omitempty" dynamodbav:"hosted_zone_id,omitempty" json:"hosted_zone_id,omitempty" diff:"hosted_zone_id"`
	Port         int32  `bson:"port,omitempty" dynamodbav:"port,omitempty" json:"port,omitempty" diff:"port"`
}

type OptionGroupMembership struct {
	OptionGroupName string `bson:"option_group_name,omitempty" dynamodbav:"option_group_name,omitempty" json:"option_group_name,omitempty" diff:"option_group_name"`
	Status          string `bson:"status,omitempty" dynamodbav:"status,omitempty" json:"status,omitempty" diff:"status"`
}

type PendingModifiedValues struct {
	AllocatedStorage                 int32                         `bson:"allocated_storage,omitempty" dynamodbav:"allocated_storage,omitempty" json:"allocated_storage,omitempty" diff:"allocated_storage"`
	AutomationMode                   string                        `bson:"automation_mode,omitempty" dynamodbav:"automation_mode,omitempty" json:"automation_mode,omitempty" diff:"automation_mode"`
	BackupRetentionPeriod            int32                         `bson:"backup_retention_period,omitempty" dynamodbav:"backup_retention_period,omitempty" json:"backup_retention_period,omitempty" diff:"backup_retention_period"`
	CACertificateIdentifier          string                        `bson:"ca_certificate_identifier,omitempty" dynamodbav:"ca_certificate_identifier,omitempty" json:"ca_certificate_identifier,omitempty" diff:"ca_certificate_identifier"`
	DBInstanceClass                  string                        `bson:"db_instance_class,omitempty" dynamodbav:"db_instance_class,omitempty" json:"db_instance_class,omitempty" diff:"db_instance_class"`
	DBInstanceIdentifier             string                        `bson:"db_instance_identifier,omitempty" dynamodbav:"db_instance_identifier,omitempty" json:"db_instance_identifier,omitempty" diff:"db_instance_identifier"`
	DBSubnetGroupName                string                        `bson:"db_subnet_group_name,omitempty" dynamodbav:"db_subnet_group_name,omitempty" json:"db_subnet_group_name,omitempty" diff:"db_subnet_group_name"`
	EngineVersion                    string                        `bson:"engine_version,omitempty" dynamodbav:"engine_version,omitempty" json:"engine_version,omitempty" diff:"engine_version"`
	IAMDatabaseAuthenticationEnabled bool                          `bson:"iam_database_authentication_enabled,omitempty" dynamodbav:"iam_database_authentication_enabled" json:"iam_database_authentication_enabled,omitempty" diff:"iam_database_authentication_enabled"`
	Iops                             int32                         `bson:"iops,omitempty" dynamodbav:"iops,omitempty" json:"iops,omitempty" diff:"iops"`
	LicenseModel                     string                        `bson:"license_model,omitempty" dynamodbav:"license_model,omitempty" json:"license_model,omitempty" diff:"license_model"`
	MasterUserPassword               string                        `bson:"master_user_password,omitempty" dynamodbav:"master_user_password,omitempty" json:"master_user_password,omitempty" diff:"master_user_password"`
	MultiAZ                          bool                          `bson:"multi_az,omitempty" dynamodbav:"multi_az" json:"multi_az,omitempty" diff:"multi_az"`
	PendingCloudwatchLogsExports     *PendingCloudwatchLogsExports `bson:"pending_cloudwatch_logs_exports,omitempty" dynamodbav:"pending_cloudwatch_logs_exports,omitempty" json:"pending_cloudwatch_logs_exports,omitempty" diff:"pending_cloudwatch_logs_exports"`
	Port                             int32                         `bson:"port,omitempty" dynamodbav:"port,omitempty" json:"port,omitempty" diff:"port"`
	ProcessorFeatures                []*ProcessorFeature           `bson:"processor_features,omitempty" dynamodbav:"processor_features,omitempty" json:"processor_features,omitempty" diff:"processor_features"`
	ResumeFullAutomationModeTime     *time.Time                    `bson:"resume_full_automation_mode_time,omitempty" dynamodbav:"resume_full_automation_mode_time,unixtime,omitempty" json:"resume_full_automation_mode_time,omitempty" diff:"resume_full_automation_mode_time"`
	StorageType                      string                        `bson:"storage_type,omitempty" dynamodbav:"storage_type,omitempty" json:"storage_type,omitempty" diff:"storage_type"`
}

type ProcessorFeature struct {
	Name  string `bson:"name,omitempty" dynamodbav:"name,omitempty" json:"name,omitempty" diff:"name"`
	Value string `bson:"value,omitempty" dynamodbav:"value,omitempty" json:"value,omitempty" diff:"value"`
}

type DBInstanceStatusInfo struct {
	Message    string `bson:"message,omitempty" dynamodbav:"message,omitempty" json:"message,omitempty" diff:"message"`
	Normal     bool   `bson:"normal,omitempty" dynamodbav:"normal" json:"normal,omitempty" diff:"normal"`
	Status     string `bson:"status,omitempty" dynamodbav:"status,omitempty" json:"status,omitempty" diff:"status"`
	StatusType string `bson:"status_type,omitempty" dynamodbav:"status_type,omitempty" json:"status_type,omitempty" diff:"status_type"`
}