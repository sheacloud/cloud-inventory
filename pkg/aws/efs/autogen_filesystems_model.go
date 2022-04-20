//AUTOGENERATED CODE DO NOT EDIT
package efs

import (
	"time"
)

type FileSystem struct {
	CreationTime                 *time.Time        `bson:"creation_time,omitempty" dynamodbav:"creation_time,unixtime,omitempty" json:"creation_time,omitempty" diff:"creation_time"`
	CreationToken                string            `bson:"creation_token,omitempty" dynamodbav:"creation_token,omitempty" json:"creation_token,omitempty" diff:"creation_token"`
	FileSystemId                 string            `bson:"file_system_id,omitempty" dynamodbav:"file_system_id,omitempty" inventory_primary_key:"true" json:"file_system_id,omitempty" diff:"file_system_id,identifier"`
	LifeCycleState               string            `bson:"life_cycle_state,omitempty" dynamodbav:"life_cycle_state,omitempty" json:"life_cycle_state,omitempty" diff:"life_cycle_state"`
	NumberOfMountTargets         int32             `bson:"number_of_mount_targets,omitempty" dynamodbav:"number_of_mount_targets,omitempty" json:"number_of_mount_targets,omitempty" diff:"number_of_mount_targets"`
	OwnerId                      string            `bson:"owner_id,omitempty" dynamodbav:"owner_id,omitempty" json:"owner_id,omitempty" diff:"owner_id"`
	PerformanceMode              string            `bson:"performance_mode,omitempty" dynamodbav:"performance_mode,omitempty" json:"performance_mode,omitempty" diff:"performance_mode"`
	SizeInBytes                  *FileSystemSize   `bson:"size_in_bytes,omitempty" dynamodbav:"size_in_bytes,omitempty" json:"size_in_bytes,omitempty" diff:"size_in_bytes"`
	Tags                         map[string]string `bson:"tags,omitempty" dynamodbav:"tags,omitempty" json:"tags,omitempty" diff:"tags"`
	AvailabilityZoneId           string            `bson:"availability_zone_id,omitempty" dynamodbav:"availability_zone_id,omitempty" json:"availability_zone_id,omitempty" diff:"availability_zone_id"`
	AvailabilityZoneName         string            `bson:"availability_zone_name,omitempty" dynamodbav:"availability_zone_name,omitempty" json:"availability_zone_name,omitempty" diff:"availability_zone_name"`
	Encrypted                    bool              `bson:"encrypted,omitempty" dynamodbav:"encrypted" json:"encrypted,omitempty" diff:"encrypted"`
	FileSystemArn                string            `bson:"file_system_arn,omitempty" dynamodbav:"file_system_arn,omitempty" json:"file_system_arn,omitempty" diff:"file_system_arn"`
	KmsKeyId                     string            `bson:"kms_key_id,omitempty" dynamodbav:"kms_key_id,omitempty" json:"kms_key_id,omitempty" diff:"kms_key_id"`
	Name                         string            `bson:"name,omitempty" dynamodbav:"name,omitempty" json:"name,omitempty" diff:"name"`
	ProvisionedThroughputInMibps float64           `bson:"provisioned_throughput_in_mibps,omitempty" dynamodbav:"provisioned_throughput_in_mibps,omitempty" json:"provisioned_throughput_in_mibps,omitempty" diff:"provisioned_throughput_in_mibps"`
	ThroughputMode               string            `bson:"throughput_mode,omitempty" dynamodbav:"throughput_mode,omitempty" json:"throughput_mode,omitempty" diff:"throughput_mode"`
	AccountId                    string            `bson:"account_id,omitempty" dynamodbav:"account_id,omitempty" json:"account_id,omitempty" diff:"account_id"`
	Region                       string            `bson:"region,omitempty" dynamodbav:"region,omitempty" json:"region,omitempty" diff:"region"`
	ReportTime                   time.Time         `bson:"report_time,omitempty" dynamodbav:"report_time,unixtime,omitempty" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID                string            `bson:"_id,omitempty" dynamodbav:"_id,omitempty" json:"_id,omitempty" diff:"-"`
}
