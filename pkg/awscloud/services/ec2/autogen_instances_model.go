//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"time"
)

type Instance struct {
	AmiLaunchIndex                          int32                                     `parquet:"name=ami_launch_index,type=INT32" json:"ami_launch_index" diff:"ami_launch_index"`
	Architecture                            string                                    `parquet:"name=architecture,type=BYTE_ARRAY,convertedtype=UTF8" json:"architecture" diff:"architecture"`
	BlockDeviceMappings                     []*InstanceBlockDeviceMapping             `parquet:"name=block_device_mappings,type=MAP,convertedtype=LIST" json:"block_device_mappings" diff:"block_device_mappings"`
	BootMode                                string                                    `parquet:"name=boot_mode,type=BYTE_ARRAY,convertedtype=UTF8" json:"boot_mode" diff:"boot_mode"`
	CapacityReservationId                   string                                    `parquet:"name=capacity_reservation_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"capacity_reservation_id" diff:"capacity_reservation_id"`
	CapacityReservationSpecification        *CapacityReservationSpecificationResponse `parquet:"name=capacity_reservation_specification" json:"capacity_reservation_specification" diff:"capacity_reservation_specification"`
	ClientToken                             string                                    `parquet:"name=client_token,type=BYTE_ARRAY,convertedtype=UTF8" json:"client_token" diff:"client_token"`
	CpuOptions                              *CpuOptions                               `parquet:"name=cpu_options" json:"cpu_options" diff:"cpu_options"`
	EbsOptimized                            bool                                      `parquet:"name=ebs_optimized,type=BOOLEAN" json:"ebs_optimized" diff:"ebs_optimized"`
	ElasticGpuAssociations                  []*ElasticGpuAssociation                  `parquet:"name=elastic_gpu_associations,type=MAP,convertedtype=LIST" json:"elastic_gpu_associations" diff:"elastic_gpu_associations"`
	ElasticInferenceAcceleratorAssociations []*ElasticInferenceAcceleratorAssociation `parquet:"name=elastic_inference_accelerator_associations,type=MAP,convertedtype=LIST" json:"elastic_inference_accelerator_associations" diff:"elastic_inference_accelerator_associations"`
	EnaSupport                              bool                                      `parquet:"name=ena_support,type=BOOLEAN" json:"ena_support" diff:"ena_support"`
	EnclaveOptions                          *EnclaveOptions                           `parquet:"name=enclave_options" json:"enclave_options" diff:"enclave_options"`
	HibernationOptions                      *HibernationOptions                       `parquet:"name=hibernation_options" json:"hibernation_options" diff:"hibernation_options"`
	Hypervisor                              string                                    `parquet:"name=hypervisor,type=BYTE_ARRAY,convertedtype=UTF8" json:"hypervisor" diff:"hypervisor"`
	IamInstanceProfile                      *IamInstanceProfile                       `parquet:"name=iam_instance_profile" json:"iam_instance_profile" diff:"iam_instance_profile"`
	ImageId                                 string                                    `parquet:"name=image_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"image_id" diff:"image_id"`
	InstanceId                              string                                    `parquet:"name=instance_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"instance_id" diff:"instance_id,identifier"`
	InstanceLifecycle                       string                                    `parquet:"name=instance_lifecycle,type=BYTE_ARRAY,convertedtype=UTF8" json:"instance_lifecycle" diff:"instance_lifecycle"`
	InstanceType                            string                                    `parquet:"name=instance_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"instance_type" diff:"instance_type"`
	Ipv6Address                             string                                    `parquet:"name=ipv6_address,type=BYTE_ARRAY,convertedtype=UTF8" json:"ipv6_address" diff:"ipv6_address"`
	KernelId                                string                                    `parquet:"name=kernel_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"kernel_id" diff:"kernel_id"`
	KeyName                                 string                                    `parquet:"name=key_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"key_name" diff:"key_name"`
	LaunchTime                              *time.Time                                `json:"-"`
	Licenses                                []*LicenseConfiguration                   `parquet:"name=licenses,type=MAP,convertedtype=LIST" json:"licenses" diff:"licenses"`
	MetadataOptions                         *InstanceMetadataOptionsResponse          `parquet:"name=metadata_options" json:"metadata_options" diff:"metadata_options"`
	Monitoring                              *Monitoring                               `parquet:"name=monitoring" json:"monitoring" diff:"monitoring"`
	NetworkInterfaces                       []*InstanceNetworkInterface               `parquet:"name=network_interfaces,type=MAP,convertedtype=LIST" json:"network_interfaces" diff:"network_interfaces"`
	OutpostArn                              string                                    `parquet:"name=outpost_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"outpost_arn" diff:"outpost_arn"`
	Placement                               *Placement                                `parquet:"name=placement" json:"placement" diff:"placement"`
	Platform                                string                                    `parquet:"name=platform,type=BYTE_ARRAY,convertedtype=UTF8" json:"platform" diff:"platform"`
	PlatformDetails                         string                                    `parquet:"name=platform_details,type=BYTE_ARRAY,convertedtype=UTF8" json:"platform_details" diff:"platform_details"`
	PrivateDnsName                          string                                    `parquet:"name=private_dns_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"private_dns_name" diff:"private_dns_name"`
	PrivateDnsNameOptions                   *PrivateDnsNameOptionsResponse            `parquet:"name=private_dns_name_options" json:"private_dns_name_options" diff:"private_dns_name_options"`
	PrivateIpAddress                        string                                    `parquet:"name=private_ip_address,type=BYTE_ARRAY,convertedtype=UTF8" json:"private_ip_address" diff:"private_ip_address"`
	ProductCodes                            []*ProductCode                            `parquet:"name=product_codes,type=MAP,convertedtype=LIST" json:"product_codes" diff:"product_codes"`
	PublicDnsName                           string                                    `parquet:"name=public_dns_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"public_dns_name" diff:"public_dns_name"`
	PublicIpAddress                         string                                    `parquet:"name=public_ip_address,type=BYTE_ARRAY,convertedtype=UTF8" json:"public_ip_address" diff:"public_ip_address"`
	RamdiskId                               string                                    `parquet:"name=ramdisk_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"ramdisk_id" diff:"ramdisk_id"`
	RootDeviceName                          string                                    `parquet:"name=root_device_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"root_device_name" diff:"root_device_name"`
	RootDeviceType                          string                                    `parquet:"name=root_device_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"root_device_type" diff:"root_device_type"`
	SecurityGroups                          []*GroupIdentifier                        `parquet:"name=security_groups,type=MAP,convertedtype=LIST" json:"security_groups" diff:"security_groups"`
	SourceDestCheck                         bool                                      `parquet:"name=source_dest_check,type=BOOLEAN" json:"source_dest_check" diff:"source_dest_check"`
	SpotInstanceRequestId                   string                                    `parquet:"name=spot_instance_request_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"spot_instance_request_id" diff:"spot_instance_request_id"`
	SriovNetSupport                         string                                    `parquet:"name=sriov_net_support,type=BYTE_ARRAY,convertedtype=UTF8" json:"sriov_net_support" diff:"sriov_net_support"`
	State                                   *InstanceState                            `parquet:"name=state" json:"state" diff:"state"`
	StateReason                             *StateReason                              `parquet:"name=state_reason" json:"state_reason" diff:"state_reason"`
	StateTransitionReason                   string                                    `parquet:"name=state_transition_reason,type=BYTE_ARRAY,convertedtype=UTF8" json:"state_transition_reason" diff:"state_transition_reason"`
	SubnetId                                string                                    `parquet:"name=subnet_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"subnet_id" diff:"subnet_id"`
	Tags                                    map[string]string                         `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"tags" diff:"tags"`
	UsageOperation                          string                                    `parquet:"name=usage_operation,type=BYTE_ARRAY,convertedtype=UTF8" json:"usage_operation" diff:"usage_operation"`
	UsageOperationUpdateTime                *time.Time                                `json:"-"`
	VirtualizationType                      string                                    `parquet:"name=virtualization_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"virtualization_type" diff:"virtualization_type"`
	VpcId                                   string                                    `parquet:"name=vpc_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"vpc_id" diff:"vpc_id"`
	AccountId                               string                                    `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"account_id" diff:"account_id"`
	Region                                  string                                    `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8" json:"region" diff:"region"`
	ReportTime                              int64                                     `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"report_time" diff:"report_time,immutable"`
	LaunchTimeMilli                         int64                                     `parquet:"name=launch_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"launch_time" diff:"launch_time"`
	UsageOperationUpdateTimeMilli           int64                                     `parquet:"name=usage_operation_update_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"usage_operation_update_time" diff:"usage_operation_update_time"`
}

func (x *Instance) GetReportTime() int64 {
	return x.ReportTime
}
