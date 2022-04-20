//AUTOGENERATED CODE DO NOT EDIT
package redshift

import (
	"time"
)

type Cluster struct {
	AllowVersionUpgrade                    bool                              `bson:"allow_version_upgrade,omitempty" dynamodbav:"allow_version_upgrade" json:"allow_version_upgrade,omitempty" diff:"allow_version_upgrade"`
	AquaConfiguration                      *AquaConfiguration                `bson:"aqua_configuration,omitempty" dynamodbav:"aqua_configuration,omitempty" json:"aqua_configuration,omitempty" diff:"aqua_configuration"`
	AutomatedSnapshotRetentionPeriod       int32                             `bson:"automated_snapshot_retention_period,omitempty" dynamodbav:"automated_snapshot_retention_period,omitempty" json:"automated_snapshot_retention_period,omitempty" diff:"automated_snapshot_retention_period"`
	AvailabilityZone                       string                            `bson:"availability_zone,omitempty" dynamodbav:"availability_zone,omitempty" json:"availability_zone,omitempty" diff:"availability_zone"`
	AvailabilityZoneRelocationStatus       string                            `bson:"availability_zone_relocation_status,omitempty" dynamodbav:"availability_zone_relocation_status,omitempty" json:"availability_zone_relocation_status,omitempty" diff:"availability_zone_relocation_status"`
	ClusterAvailabilityStatus              string                            `bson:"cluster_availability_status,omitempty" dynamodbav:"cluster_availability_status,omitempty" json:"cluster_availability_status,omitempty" diff:"cluster_availability_status"`
	ClusterCreateTime                      *time.Time                        `bson:"cluster_create_time,omitempty" dynamodbav:"cluster_create_time,unixtime,omitempty" json:"cluster_create_time,omitempty" diff:"cluster_create_time"`
	ClusterIdentifier                      string                            `bson:"cluster_identifier,omitempty" dynamodbav:"cluster_identifier,omitempty" inventory_primary_key:"true" json:"cluster_identifier,omitempty" diff:"cluster_identifier,identifier"`
	ClusterNamespaceArn                    string                            `bson:"cluster_namespace_arn,omitempty" dynamodbav:"cluster_namespace_arn,omitempty" json:"cluster_namespace_arn,omitempty" diff:"cluster_namespace_arn"`
	ClusterNodes                           []*ClusterNode                    `bson:"cluster_nodes,omitempty" dynamodbav:"cluster_nodes,omitempty" json:"cluster_nodes,omitempty" diff:"cluster_nodes"`
	ClusterParameterGroups                 []*ClusterParameterGroupStatus    `bson:"cluster_parameter_groups,omitempty" dynamodbav:"cluster_parameter_groups,omitempty" json:"cluster_parameter_groups,omitempty" diff:"cluster_parameter_groups"`
	ClusterPublicKey                       string                            `bson:"cluster_public_key,omitempty" dynamodbav:"cluster_public_key,omitempty" json:"cluster_public_key,omitempty" diff:"cluster_public_key"`
	ClusterRevisionNumber                  string                            `bson:"cluster_revision_number,omitempty" dynamodbav:"cluster_revision_number,omitempty" json:"cluster_revision_number,omitempty" diff:"cluster_revision_number"`
	ClusterSecurityGroups                  []*ClusterSecurityGroupMembership `bson:"cluster_security_groups,omitempty" dynamodbav:"cluster_security_groups,omitempty" json:"cluster_security_groups,omitempty" diff:"cluster_security_groups"`
	ClusterSnapshotCopyStatus              *ClusterSnapshotCopyStatus        `bson:"cluster_snapshot_copy_status,omitempty" dynamodbav:"cluster_snapshot_copy_status,omitempty" json:"cluster_snapshot_copy_status,omitempty" diff:"cluster_snapshot_copy_status"`
	ClusterStatus                          string                            `bson:"cluster_status,omitempty" dynamodbav:"cluster_status,omitempty" json:"cluster_status,omitempty" diff:"cluster_status"`
	ClusterSubnetGroupName                 string                            `bson:"cluster_subnet_group_name,omitempty" dynamodbav:"cluster_subnet_group_name,omitempty" json:"cluster_subnet_group_name,omitempty" diff:"cluster_subnet_group_name"`
	ClusterVersion                         string                            `bson:"cluster_version,omitempty" dynamodbav:"cluster_version,omitempty" json:"cluster_version,omitempty" diff:"cluster_version"`
	DBName                                 string                            `bson:"db_name,omitempty" dynamodbav:"db_name,omitempty" json:"db_name,omitempty" diff:"db_name"`
	DataTransferProgress                   *DataTransferProgress             `bson:"data_transfer_progress,omitempty" dynamodbav:"data_transfer_progress,omitempty" json:"data_transfer_progress,omitempty" diff:"data_transfer_progress"`
	DefaultIamRoleArn                      string                            `bson:"default_iam_role_arn,omitempty" dynamodbav:"default_iam_role_arn,omitempty" json:"default_iam_role_arn,omitempty" diff:"default_iam_role_arn"`
	DeferredMaintenanceWindows             []*DeferredMaintenanceWindow      `bson:"deferred_maintenance_windows,omitempty" dynamodbav:"deferred_maintenance_windows,omitempty" json:"deferred_maintenance_windows,omitempty" diff:"deferred_maintenance_windows"`
	ElasticIpStatus                        *ElasticIpStatus                  `bson:"elastic_ip_status,omitempty" dynamodbav:"elastic_ip_status,omitempty" json:"elastic_ip_status,omitempty" diff:"elastic_ip_status"`
	ElasticResizeNumberOfNodeOptions       string                            `bson:"elastic_resize_number_of_node_options,omitempty" dynamodbav:"elastic_resize_number_of_node_options,omitempty" json:"elastic_resize_number_of_node_options,omitempty" diff:"elastic_resize_number_of_node_options"`
	Encrypted                              bool                              `bson:"encrypted,omitempty" dynamodbav:"encrypted" json:"encrypted,omitempty" diff:"encrypted"`
	Endpoint                               *Endpoint                         `bson:"endpoint,omitempty" dynamodbav:"endpoint,omitempty" json:"endpoint,omitempty" diff:"endpoint"`
	EnhancedVpcRouting                     bool                              `bson:"enhanced_vpc_routing,omitempty" dynamodbav:"enhanced_vpc_routing" json:"enhanced_vpc_routing,omitempty" diff:"enhanced_vpc_routing"`
	ExpectedNextSnapshotScheduleTime       *time.Time                        `bson:"expected_next_snapshot_schedule_time,omitempty" dynamodbav:"expected_next_snapshot_schedule_time,unixtime,omitempty" json:"expected_next_snapshot_schedule_time,omitempty" diff:"expected_next_snapshot_schedule_time"`
	ExpectedNextSnapshotScheduleTimeStatus string                            `bson:"expected_next_snapshot_schedule_time_status,omitempty" dynamodbav:"expected_next_snapshot_schedule_time_status,omitempty" json:"expected_next_snapshot_schedule_time_status,omitempty" diff:"expected_next_snapshot_schedule_time_status"`
	HsmStatus                              *HsmStatus                        `bson:"hsm_status,omitempty" dynamodbav:"hsm_status,omitempty" json:"hsm_status,omitempty" diff:"hsm_status"`
	IamRoles                               []*ClusterIamRole                 `bson:"iam_roles,omitempty" dynamodbav:"iam_roles,omitempty" json:"iam_roles,omitempty" diff:"iam_roles"`
	KmsKeyId                               string                            `bson:"kms_key_id,omitempty" dynamodbav:"kms_key_id,omitempty" json:"kms_key_id,omitempty" diff:"kms_key_id"`
	MaintenanceTrackName                   string                            `bson:"maintenance_track_name,omitempty" dynamodbav:"maintenance_track_name,omitempty" json:"maintenance_track_name,omitempty" diff:"maintenance_track_name"`
	ManualSnapshotRetentionPeriod          int32                             `bson:"manual_snapshot_retention_period,omitempty" dynamodbav:"manual_snapshot_retention_period,omitempty" json:"manual_snapshot_retention_period,omitempty" diff:"manual_snapshot_retention_period"`
	MasterUsername                         string                            `bson:"master_username,omitempty" dynamodbav:"master_username,omitempty" json:"master_username,omitempty" diff:"master_username"`
	ModifyStatus                           string                            `bson:"modify_status,omitempty" dynamodbav:"modify_status,omitempty" json:"modify_status,omitempty" diff:"modify_status"`
	NextMaintenanceWindowStartTime         *time.Time                        `bson:"next_maintenance_window_start_time,omitempty" dynamodbav:"next_maintenance_window_start_time,unixtime,omitempty" json:"next_maintenance_window_start_time,omitempty" diff:"next_maintenance_window_start_time"`
	NodeType                               string                            `bson:"node_type,omitempty" dynamodbav:"node_type,omitempty" json:"node_type,omitempty" diff:"node_type"`
	NumberOfNodes                          int32                             `bson:"number_of_nodes,omitempty" dynamodbav:"number_of_nodes,omitempty" json:"number_of_nodes,omitempty" diff:"number_of_nodes"`
	PendingActions                         []string                          `bson:"pending_actions,omitempty" dynamodbav:"pending_actions,omitempty" json:"pending_actions,omitempty" diff:"pending_actions"`
	PendingModifiedValues                  *PendingModifiedValues            `bson:"pending_modified_values,omitempty" dynamodbav:"pending_modified_values,omitempty" json:"pending_modified_values,omitempty" diff:"pending_modified_values"`
	PreferredMaintenanceWindow             string                            `bson:"preferred_maintenance_window,omitempty" dynamodbav:"preferred_maintenance_window,omitempty" json:"preferred_maintenance_window,omitempty" diff:"preferred_maintenance_window"`
	PubliclyAccessible                     bool                              `bson:"publicly_accessible,omitempty" dynamodbav:"publicly_accessible" json:"publicly_accessible,omitempty" diff:"publicly_accessible"`
	ReservedNodeExchangeStatus             *ReservedNodeExchangeStatus       `bson:"reserved_node_exchange_status,omitempty" dynamodbav:"reserved_node_exchange_status,omitempty" json:"reserved_node_exchange_status,omitempty" diff:"reserved_node_exchange_status"`
	ResizeInfo                             *ResizeInfo                       `bson:"resize_info,omitempty" dynamodbav:"resize_info,omitempty" json:"resize_info,omitempty" diff:"resize_info"`
	RestoreStatus                          *RestoreStatus                    `bson:"restore_status,omitempty" dynamodbav:"restore_status,omitempty" json:"restore_status,omitempty" diff:"restore_status"`
	SnapshotScheduleIdentifier             string                            `bson:"snapshot_schedule_identifier,omitempty" dynamodbav:"snapshot_schedule_identifier,omitempty" json:"snapshot_schedule_identifier,omitempty" diff:"snapshot_schedule_identifier"`
	SnapshotScheduleState                  string                            `bson:"snapshot_schedule_state,omitempty" dynamodbav:"snapshot_schedule_state,omitempty" json:"snapshot_schedule_state,omitempty" diff:"snapshot_schedule_state"`
	Tags                                   map[string]string                 `bson:"tags,omitempty" dynamodbav:"tags,omitempty" json:"tags,omitempty" diff:"tags"`
	TotalStorageCapacityInMegaBytes        int64                             `bson:"total_storage_capacity_in_mega_bytes,omitempty" dynamodbav:"total_storage_capacity_in_mega_bytes,omitempty" json:"total_storage_capacity_in_mega_bytes,omitempty" diff:"total_storage_capacity_in_mega_bytes"`
	VpcId                                  string                            `bson:"vpc_id,omitempty" dynamodbav:"vpc_id,omitempty" json:"vpc_id,omitempty" diff:"vpc_id"`
	VpcSecurityGroups                      []*VpcSecurityGroupMembership     `bson:"vpc_security_groups,omitempty" dynamodbav:"vpc_security_groups,omitempty" json:"vpc_security_groups,omitempty" diff:"vpc_security_groups"`
	AccountId                              string                            `bson:"account_id,omitempty" dynamodbav:"account_id,omitempty" json:"account_id,omitempty" diff:"account_id"`
	Region                                 string                            `bson:"region,omitempty" dynamodbav:"region,omitempty" json:"region,omitempty" diff:"region"`
	ReportTime                             time.Time                         `bson:"report_time,omitempty" dynamodbav:"report_time,unixtime,omitempty" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID                          string                            `bson:"_id,omitempty" dynamodbav:"_id,omitempty" json:"_id,omitempty" diff:"-"`
}
