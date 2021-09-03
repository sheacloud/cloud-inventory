package ec2

import (
	"context"
	"fmt"
	"time"

	awsec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

func init() {
	Controller.RegisterDataSource("instances", InstanceDataSource)
}

type InstanceModel struct {
	Architecture          string                             `parquet:"name=architecture, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	EbsOptimized          bool                               `parquet:"name=ebs_optimized, type=BOOLEAN"`
	EnaSupport            bool                               `parquet:"name=ena_support, type=BOOLEAN"`
	BlockDeviceMappings   []*InstanceBlockDeviceMappingModel `parquet:"name=block_device_mappings, type=LIST"`
	Hypervisor            string                             `parquet:"name=hypervisor, type=BYTE_ARRAY, convertedtype=UTF8"`
	IamInstanceProfile    *IamInstanceProfileModel           `parquet:"name=iam_instance_profile"`
	ImageId               string                             `parquet:"name=image_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	InstanceId            string                             `parquet:"name=instance_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	InstanceLifecycle     string                             `parquet:"name=instance_lifecycle, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	InstanceType          string                             `parquet:"name=instance_type, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	KernelId              string                             `parquet:"name=kernel_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	KeyName               string                             `parquet:"name=key_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	LaunchTime            time.Time
	LaunchTimeMillis      int64                            `parquet:"name=launch_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	Monitoring            *MonitoringModel                 `parquet:"name=monitoring"`
	NetworkInterfaces     []*InstanceNetworkInterfaceModel `parquet:"name=network_interfaces, type=LIST"`
	PrivateDnsName        string                           `parquet:"name=private_dns_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	PrivateIpAddress      string                           `parquet:"name=private_ip_address, type=BYTE_ARRAY, convertedtype=UTF8"`
	SecurityGroups        []*GroupIdentifierModel          `parquet:"name=security_groups, type=LIST"`
	SourceDestCheck       bool                             `parquet:"name=source_dest_check, type=BOOLEAN"`
	SpotInstanceRequestId string                           `parquet:"name=spot_instance_request_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	State                 *InstanceStateModel              `parquet:"name=state"`
	StateReason           *StateReasonModel                `parquet:"name=state_reason"`
	StateTransitionReason string                           `parquet:"name=state_transition_reason, type=BYTE_ARRAY, convertedtype=UTF8"`
	SubnetId              string                           `parquet:"name=subnet_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	VirtualizationType    string                           `parquet:"name=virtualization_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	VpcId                 string                           `parquet:"name=vpc_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Tags                  map[string]string                `parquet:"name=tags, type=MAP, keytype=BYTE_ARRAY, keyconvertedtype=UTF8, valuetype=BYTE_ARRAY, valueconvertedtype=UTF8"`
	AccountId             string                           `parquet:"name=account_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Region                string                           `parquet:"name=region, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type StateReasonModel struct {
	Code    string `parquet:"name=code, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Message string `parquet:"name=message, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
}

type InstanceStateModel struct {
	Code int    `parquet:"name=code, type=INT32"`
	Name string `parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
}

type InstanceNetworkInterfaceModel struct {
	Association        InstanceNetworkInterfaceAssociationModel `parquet:"name=association"`
	Attachment         InstanceNetworkInterfaceAttachmentModel  `parquet:"name=attachment"`
	Description        string                                   `parquet:"name=description, type=BYTE_ARRAY, convertedtype=UTF8"`
	Groups             []GroupIdentifierModel                   `parquet:"name=groups, type=LIST"`
	InterfaceType      string                                   `parquet:"name=interface_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	MacAddress         string                                   `parquet:"name=mac_address, type=BYTE_ARRAY, convertedtype=UTF8"`
	NetworkInterfaceId string                                   `parquet:"name=network_interface_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	OwnerId            string                                   `parquet:"name=owner_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	PrivateDnsName     string                                   `parquet:"name=private_dns_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	PrivateIpAddress   string                                   `parquet:"name=private_ip_address, type=BYTE_ARRAY, convertedtype=UTF8"`
	PrivateIpAddresses []InstancePrivateIpAddressModel          `parquet:"name=private_ip_addresses, type=LIST"`
	SourceDestCheck    bool                                     `parquet:"name=source_dest_check, type=BOOLEAN"`
	Status             string                                   `parquet:"name=status, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
}

type InstancePrivateIpAddressModel struct {
	Association      InstanceNetworkInterfaceAssociationModel `parquet:"name=association"`
	Primary          bool                                     `parquet:"name=primary, type=BOOLEAN"`
	PrivateDnsName   string                                   `parquet:"name=private_dns_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	PrivateIpAddress string                                   `parquet:"name=private_ip_address, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type GroupIdentifierModel struct {
	GroupId   string `parquet:"name=group_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	GroupName string `parquet:"name=group_name, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type InstanceNetworkInterfaceAttachmentModel struct {
	AttachmentId        string `parquet:"name=attachment_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	DeleteOnTermination bool   `parquet:"name=delete_on_termination, type=BOOLEAN"`
	DeviceIndex         int    `parquet:"name=device_index, type=INT32"`
	NetworkCardIndex    int    `parquet:"name=network_card_index, type=INT32"`
	Status              string `parquet:"name=status, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type InstanceNetworkInterfaceAssociationModel struct {
	CarrierIp     string `parquet:"name=carrier_ip, type=BYTE_ARRAY, convertedtype=UTF8"`
	IpOwnerId     string `parquet:"name=ip_owner_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	PublicDnsName string `parquet:"name=public_dns_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	PublicIp      string `parquet:"name=public_ip, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type MonitoringModel struct {
	State string `parquet:"name=state, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type InstanceBlockDeviceMappingModel struct {
	DeviceName string                      `parquet:"name=device_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	Ebs        EbsInstanceBlockDeviceModel `parquet:"name=ebs"`
}

type EbsInstanceBlockDeviceModel struct {
	DeviceName  string                                   `parquet:"name=device_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	Ebs         EbsInstanceBlockDeviceSpecificationModel `parquet:"name=ebs"`
	NoDevice    string                                   `parquet:"name=no_device, type=BYTE_ARRAY, convertedtype=UTF8"`
	VirtualName string                                   `parquet:"name=virtual_name, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type EbsInstanceBlockDeviceSpecificationModel struct {
	DeleteOnTermination bool   `parquet:"name=delete_on_termination, type=BOOLEAN"`
	VolumeId            string `parquet:"name=volume_id, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type IamInstanceProfileModel struct {
	Arn string `parquet:"name=arn, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Id  string `parquet:"name=id, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
}

type InstanceDataSourceClient interface {
	DescribeInstances(context.Context, *awsec2.DescribeInstancesInput, ...func(*awsec2.Options)) (*awsec2.DescribeInstancesOutput, error)
}

func InstanceDataSource(ctx context.Context, client *awsec2.Client, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	return instanceDataSource(ctx, client, storageConfig, storageManager)
}

// function with client as a specific interface, allowing mocking/testing
func instanceDataSource(ctx context.Context, client InstanceDataSourceClient, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(InstanceModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := awsec2.NewDescribeInstancesPaginator(client, &awsec2.DescribeInstancesInput{})

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

		for _, reservation := range output.Reservations {
			for _, instance := range reservation.Instances {
				model := new(InstanceModel)
				copier.Copy(&model, &instance)

				model.Tags = GetTagMap(instance.Tags)
				model.LaunchTimeMillis = model.LaunchTime.UTC().UnixMilli()
				model.AccountId = storageConfig.AccountId
				model.Region = storageConfig.Region

				errors := storageContextSet.Store(ctx, model)
				if errors != nil {
					for storageContext, err := range errors {
						storage.LogContextError(storageContext, fmt.Sprintf("Error storing InstanceModel: %v", err))
					}
				}
			}
		}
	}

	return nil
}
