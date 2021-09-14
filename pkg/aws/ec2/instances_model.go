// AUTOGENERATED, DO NOT EDIT
package ec2

import (
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
	"time"

	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"sync"
)

var customInstanceModelPostprocessingFuncs []func(ctx context.Context, client *ec2.Client, cfg aws.Config, x *InstanceModel) = []func(ctx context.Context, client *ec2.Client, cfg aws.Config, x *InstanceModel){}
var customInstanceModelFuncsLock sync.Mutex

func registerCustomInstanceModelPostprocessingFunc(f func(ctx context.Context, client *ec2.Client, cfg aws.Config, x *InstanceModel)) {
	customInstanceModelFuncsLock.Lock()
	defer customInstanceModelFuncsLock.Unlock()

	customInstanceModelPostprocessingFuncs = append(customInstanceModelPostprocessingFuncs, f)
}

func init() {
	Controller.RegisterDataSource("instances", InstanceDataSource)
}

type InstanceModel struct {
	AmiLaunchIndex                          int32                                                  `parquet:"name=ami_launch_index,type=INT32"`
	Architecture                            string                                                 `parquet:"name=architecture,type=BYTE_ARRAY,convertedtype=UTF8"`
	BlockDeviceMappings                     []*InstanceBlockDeviceMappingInstanceModel             `parquet:"name=block_device_mappings,type=LIST"`
	BootMode                                string                                                 `parquet:"name=boot_mode,type=BYTE_ARRAY,convertedtype=UTF8"`
	CapacityReservationId                   string                                                 `parquet:"name=capacity_reservation_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	CapacityReservationSpecification        *CapacityReservationSpecificationResponseInstanceModel `parquet:"name=capacity_reservation_specification"`
	ClientToken                             string                                                 `parquet:"name=client_token,type=BYTE_ARRAY,convertedtype=UTF8"`
	CpuOptions                              *CpuOptionsInstanceModel                               `parquet:"name=cpu_options"`
	EbsOptimized                            bool                                                   `parquet:"name=ebs_optimized,type=BOOLEAN"`
	ElasticGpuAssociations                  []*ElasticGpuAssociationInstanceModel                  `parquet:"name=elastic_gpu_associations,type=LIST"`
	ElasticInferenceAcceleratorAssociations []*ElasticInferenceAcceleratorAssociationInstanceModel `parquet:"name=elastic_inference_accelerator_associations,type=LIST"`
	EnaSupport                              bool                                                   `parquet:"name=ena_support,type=BOOLEAN"`
	EnclaveOptions                          *EnclaveOptionsInstanceModel                           `parquet:"name=enclave_options"`
	HibernationOptions                      *HibernationOptionsInstanceModel                       `parquet:"name=hibernation_options"`
	Hypervisor                              string                                                 `parquet:"name=hypervisor,type=BYTE_ARRAY,convertedtype=UTF8"`
	IamInstanceProfile                      *IamInstanceProfileInstanceModel                       `parquet:"name=iam_instance_profile"`
	ImageId                                 string                                                 `parquet:"name=image_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	InstanceId                              string                                                 `parquet:"name=instance_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	InstanceLifecycle                       string                                                 `parquet:"name=instance_lifecycle,type=BYTE_ARRAY,convertedtype=UTF8"`
	InstanceType                            string                                                 `parquet:"name=instance_type,type=BYTE_ARRAY,convertedtype=UTF8"`
	KernelId                                string                                                 `parquet:"name=kernel_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	KeyName                                 string                                                 `parquet:"name=key_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	LaunchTime                              *time.Time
	LaunchTimeMilli                         int64                                         `parquet:"name=launch_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	Licenses                                []*LicenseConfigurationInstanceModel          `parquet:"name=licenses,type=LIST"`
	MetadataOptions                         *InstanceMetadataOptionsResponseInstanceModel `parquet:"name=metadata_options"`
	Monitoring                              *MonitoringInstanceModel                      `parquet:"name=monitoring"`
	NetworkInterfaces                       []*InstanceNetworkInterfaceInstanceModel      `parquet:"name=network_interfaces,type=LIST"`
	OutpostArn                              string                                        `parquet:"name=outpost_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	Placement                               *PlacementInstanceModel                       `parquet:"name=placement"`
	Platform                                string                                        `parquet:"name=platform,type=BYTE_ARRAY,convertedtype=UTF8"`
	PrivateDnsName                          string                                        `parquet:"name=private_dns_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	PrivateIpAddress                        string                                        `parquet:"name=private_ip_address,type=BYTE_ARRAY,convertedtype=UTF8"`
	ProductCodes                            []*ProductCodeInstanceModel                   `parquet:"name=product_codes,type=LIST"`
	PublicDnsName                           string                                        `parquet:"name=public_dns_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	PublicIpAddress                         string                                        `parquet:"name=public_ip_address,type=BYTE_ARRAY,convertedtype=UTF8"`
	RamdiskId                               string                                        `parquet:"name=ramdisk_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	RootDeviceName                          string                                        `parquet:"name=root_device_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	RootDeviceType                          string                                        `parquet:"name=root_device_type,type=BYTE_ARRAY,convertedtype=UTF8"`
	SecurityGroups                          []*GroupIdentifierInstanceModel               `parquet:"name=security_groups,type=LIST"`
	SourceDestCheck                         bool                                          `parquet:"name=source_dest_check,type=BOOLEAN"`
	SpotInstanceRequestId                   string                                        `parquet:"name=spot_instance_request_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	SriovNetSupport                         string                                        `parquet:"name=sriov_net_support,type=BYTE_ARRAY,convertedtype=UTF8"`
	State                                   *InstanceStateInstanceModel                   `parquet:"name=state"`
	StateReason                             *StateReasonInstanceModel                     `parquet:"name=state_reason"`
	StateTransitionReason                   string                                        `parquet:"name=state_transition_reason,type=BYTE_ARRAY,convertedtype=UTF8"`
	SubnetId                                string                                        `parquet:"name=subnet_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Tags                                    map[string]string                             `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
	VirtualizationType                      string                                        `parquet:"name=virtualization_type,type=BYTE_ARRAY,convertedtype=UTF8"`
	VpcId                                   string                                        `parquet:"name=vpc_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	AccountId                               string                                        `parquet:"name=account_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Region                                  string                                        `parquet:"name=region, type=BYTE_ARRAY, convertedtype=UTF8"`
	ReportTime                              int64                                         `parquet:"name=report_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
}

type InstanceBlockDeviceMappingInstanceModel struct {
	DeviceName string                               `parquet:"name=device_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	Ebs        *EbsInstanceBlockDeviceInstanceModel `parquet:"name=ebs"`
}

type EbsInstanceBlockDeviceInstanceModel struct {
	AttachTime          *time.Time
	AttachTimeMilli     int64  `parquet:"name=attach_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	DeleteOnTermination bool   `parquet:"name=delete_on_termination,type=BOOLEAN"`
	Status              string `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
	VolumeId            string `parquet:"name=volume_id,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type CapacityReservationSpecificationResponseInstanceModel struct {
	CapacityReservationPreference string                                          `parquet:"name=capacity_reservation_preference,type=BYTE_ARRAY,convertedtype=UTF8"`
	CapacityReservationTarget     *CapacityReservationTargetResponseInstanceModel `parquet:"name=capacity_reservation_target"`
}

type CapacityReservationTargetResponseInstanceModel struct {
	CapacityReservationId               string `parquet:"name=capacity_reservation_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	CapacityReservationResourceGroupArn string `parquet:"name=capacity_reservation_resource_group_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type CpuOptionsInstanceModel struct {
	CoreCount      int32 `parquet:"name=core_count,type=INT32"`
	ThreadsPerCore int32 `parquet:"name=threads_per_core,type=INT32"`
}

type ElasticGpuAssociationInstanceModel struct {
	ElasticGpuAssociationId    string `parquet:"name=elastic_gpu_association_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	ElasticGpuAssociationState string `parquet:"name=elastic_gpu_association_state,type=BYTE_ARRAY,convertedtype=UTF8"`
	ElasticGpuAssociationTime  string `parquet:"name=elastic_gpu_association_time,type=BYTE_ARRAY,convertedtype=UTF8"`
	ElasticGpuId               string `parquet:"name=elastic_gpu_id,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type ElasticInferenceAcceleratorAssociationInstanceModel struct {
	ElasticInferenceAcceleratorArn                  string `parquet:"name=elastic_inference_accelerator_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	ElasticInferenceAcceleratorAssociationId        string `parquet:"name=elastic_inference_accelerator_association_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	ElasticInferenceAcceleratorAssociationState     string `parquet:"name=elastic_inference_accelerator_association_state,type=BYTE_ARRAY,convertedtype=UTF8"`
	ElasticInferenceAcceleratorAssociationTime      *time.Time
	ElasticInferenceAcceleratorAssociationTimeMilli int64 `parquet:"name=elastic_inference_accelerator_association_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
}

type EnclaveOptionsInstanceModel struct {
	Enabled bool `parquet:"name=enabled,type=BOOLEAN"`
}

type HibernationOptionsInstanceModel struct {
	Configured bool `parquet:"name=configured,type=BOOLEAN"`
}

type IamInstanceProfileInstanceModel struct {
	Arn string `parquet:"name=arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	Id  string `parquet:"name=id,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type LicenseConfigurationInstanceModel struct {
	LicenseConfigurationArn string `parquet:"name=license_configuration_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type InstanceMetadataOptionsResponseInstanceModel struct {
	HttpEndpoint            string `parquet:"name=http_endpoint,type=BYTE_ARRAY,convertedtype=UTF8"`
	HttpProtocolIpv6        string `parquet:"name=http_protocol_ipv6,type=BYTE_ARRAY,convertedtype=UTF8"`
	HttpPutResponseHopLimit int32  `parquet:"name=http_put_response_hop_limit,type=INT32"`
	HttpTokens              string `parquet:"name=http_tokens,type=BYTE_ARRAY,convertedtype=UTF8"`
	State                   string `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type MonitoringInstanceModel struct {
	State string `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type InstanceNetworkInterfaceInstanceModel struct {
	Association        *InstanceNetworkInterfaceAssociationInstanceModel `parquet:"name=association"`
	Attachment         *InstanceNetworkInterfaceAttachmentInstanceModel  `parquet:"name=attachment"`
	Description        string                                            `parquet:"name=description,type=BYTE_ARRAY,convertedtype=UTF8"`
	Groups             []*GroupIdentifierInstanceModel                   `parquet:"name=groups,type=LIST"`
	InterfaceType      string                                            `parquet:"name=interface_type,type=BYTE_ARRAY,convertedtype=UTF8"`
	Ipv4Prefixes       []*InstanceIpv4PrefixInstanceModel                `parquet:"name=ipv4_prefixes,type=LIST"`
	Ipv6Addresses      []*InstanceIpv6AddressInstanceModel               `parquet:"name=ipv6_addresses,type=LIST"`
	Ipv6Prefixes       []*InstanceIpv6PrefixInstanceModel                `parquet:"name=ipv6_prefixes,type=LIST"`
	MacAddress         string                                            `parquet:"name=mac_address,type=BYTE_ARRAY,convertedtype=UTF8"`
	NetworkInterfaceId string                                            `parquet:"name=network_interface_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	OwnerId            string                                            `parquet:"name=owner_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	PrivateDnsName     string                                            `parquet:"name=private_dns_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	PrivateIpAddress   string                                            `parquet:"name=private_ip_address,type=BYTE_ARRAY,convertedtype=UTF8"`
	PrivateIpAddresses []*InstancePrivateIpAddressInstanceModel          `parquet:"name=private_ip_addresses,type=LIST"`
	SourceDestCheck    bool                                              `parquet:"name=source_dest_check,type=BOOLEAN"`
	Status             string                                            `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
	SubnetId           string                                            `parquet:"name=subnet_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	VpcId              string                                            `parquet:"name=vpc_id,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type InstanceNetworkInterfaceAssociationInstanceModel struct {
	CarrierIp     string `parquet:"name=carrier_ip,type=BYTE_ARRAY,convertedtype=UTF8"`
	IpOwnerId     string `parquet:"name=ip_owner_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	PublicDnsName string `parquet:"name=public_dns_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	PublicIp      string `parquet:"name=public_ip,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type InstanceNetworkInterfaceAttachmentInstanceModel struct {
	AttachTime          *time.Time
	AttachTimeMilli     int64  `parquet:"name=attach_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	AttachmentId        string `parquet:"name=attachment_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	DeleteOnTermination bool   `parquet:"name=delete_on_termination,type=BOOLEAN"`
	DeviceIndex         int32  `parquet:"name=device_index,type=INT32"`
	NetworkCardIndex    int32  `parquet:"name=network_card_index,type=INT32"`
	Status              string `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type GroupIdentifierInstanceModel struct {
	GroupId   string `parquet:"name=group_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	GroupName string `parquet:"name=group_name,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type InstanceIpv4PrefixInstanceModel struct {
	Ipv4Prefix string `parquet:"name=ipv4_prefix,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type InstanceIpv6AddressInstanceModel struct {
	Ipv6Address string `parquet:"name=ipv6_address,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type InstanceIpv6PrefixInstanceModel struct {
	Ipv6Prefix string `parquet:"name=ipv6_prefix,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type InstancePrivateIpAddressInstanceModel struct {
	Association      *InstanceNetworkInterfaceAssociationInstanceModel `parquet:"name=association"`
	Primary          bool                                              `parquet:"name=primary,type=BOOLEAN"`
	PrivateDnsName   string                                            `parquet:"name=private_dns_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	PrivateIpAddress string                                            `parquet:"name=private_ip_address,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type PlacementInstanceModel struct {
	Affinity             string `parquet:"name=affinity,type=BYTE_ARRAY,convertedtype=UTF8"`
	AvailabilityZone     string `parquet:"name=availability_zone,type=BYTE_ARRAY,convertedtype=UTF8"`
	GroupName            string `parquet:"name=group_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	HostId               string `parquet:"name=host_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	HostResourceGroupArn string `parquet:"name=host_resource_group_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	PartitionNumber      int32  `parquet:"name=partition_number,type=INT32"`
	SpreadDomain         string `parquet:"name=spread_domain,type=BYTE_ARRAY,convertedtype=UTF8"`
	Tenancy              string `parquet:"name=tenancy,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type ProductCodeInstanceModel struct {
	ProductCodeId   string `parquet:"name=product_code_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	ProductCodeType string `parquet:"name=product_code_type,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type InstanceStateInstanceModel struct {
	Code int32  `parquet:"name=code,type=INT32"`
	Name string `parquet:"name=name,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type StateReasonInstanceModel struct {
	Code    string `parquet:"name=code,type=BYTE_ARRAY,convertedtype=UTF8"`
	Message string `parquet:"name=message,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type TagInstanceModel struct {
	Key   string `parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

func InstanceDataSource(ctx context.Context, client *ec2.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(InstanceModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := ec2.NewDescribeInstancesPaginator(client, &ec2.DescribeInstancesInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":     storageConfig.Service,
				"data_source": storageConfig.DataSource,
				"account_id":  storageConfig.AccountId,
				"region":      storageConfig.Region,
				"cloud":       storageConfig.Cloud,
				"error":       err,
			}).Error("error calling DescribeInstances")
			return err
		}

		for _, var0 := range output.Reservations {
			for _, var1 := range var0.Instances {

				model := new(InstanceModel)
				copier.Copy(&model, &var1)

				model.Tags = GetTagMap(var1.Tags)
				model.AccountId = storageConfig.AccountId
				model.Region = storageConfig.Region
				model.ReportTime = reportTime.UTC().UnixMilli()

				for _, f := range customInstanceModelPostprocessingFuncs {
					f(ctx, client, cfg, model)
				}

				errors := storageContextSet.Store(ctx, model)
				for storageContext, err := range errors {
					storage.LogContextError(storageContext, fmt.Sprintf("Error storing InstanceModel: %v", err))
				}
			}
		}

	}

	return nil
}
