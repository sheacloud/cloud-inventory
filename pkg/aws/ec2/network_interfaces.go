package ec2

import (
	"context"
	"fmt"

	awsec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

func init() {
	Controller.RegisterDataSource("network_interfaces", NetworkInterfaceDataSource)
}

type NetworkInterfaceModel struct {
	Association        NetworkInterfaceAssociationModel        `parquet:"name=association"`
	Attachment         NetworkInterfaceAttachmentModel         `parquet:"name=attachment"`
	AvailabilityZone   string                                  `parquet:"name=availability_zone, type=BYTE_ARRAY, convertedtype=UTF8"`
	Description        string                                  `parquet:"name=description, type=BYTE_ARRAY, convertedtype=UTF8"`
	Groups             []GroupIdentifierModel                  `parquet:"name=groups, type=LIST"`
	InterfaceType      string                                  `parquet:"name=interface_type, type=BYTE_ARRAY, convertedtype=UTF8"`
	MacAddress         string                                  `parquet:"name=mac_address, type=BYTE_ARRAY, convertedtype=UTF8"`
	NetworkInterfaceId string                                  `parquet:"name=network_interface_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	OwnerId            string                                  `parquet:"name=owner_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	PrivateDnsName     string                                  `parquet:"name=private_dns_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	PrivateIpAddress   string                                  `parquet:"name=private_ip_address, type=BYTE_ARRAY, convertedtype=UTF8"`
	PrivateIpAddresses []NetworkInterfacePrivateIpAddressModel `parquet:"name=private_ip_addresses, type=LIST"`
	RequesterId        string                                  `parquet:"name=requester_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	RequesterManaged   bool                                    `parquet:"name=requester_managed, type=BOOLEAN"`
	SourceDestCheck    bool                                    `parquet:"name=source_dest_check, type=BOOLEAN"`
	Status             string                                  `parquet:"name=status, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	SubnetId           string                                  `parquet:"name=subnet_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	TagSet             map[string]string                       `parquet:"name=tags, type=MAP, keytype=BYTE_ARRAY, keyconvertedtype=UTF8, valuetype=BYTE_ARRAY, valueconvertedtype=UTF8"`
	VpcId              string                                  `parquet:"name=vpc_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	AccountId          string                                  `parquet:"name=account_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Region             string                                  `parquet:"name=region, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type NetworkInterfacePrivateIpAddressModel struct {
	Association      NetworkInterfaceAssociationModel `parquet:"name=association"`
	Primary          bool                             `parquet:"name=primary, type=BOOLEAN"`
	PrivateDnsName   string                           `parquet:"name=private_dns_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	PrivateIpAddress string                           `parquet:"name=private_ip_address, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type NetworkInterfaceAttachmentModel struct {
	AttachmentId        string `parquet:"name=attachment_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	DeleteOnTermination bool   `parquet:"name=delete_on_termination, type=BOOLEAN"`
	DeviceIndex         int    `parquet:"name=device_index, type=INT32"`
	InstanceId          string `parquet:"name=instance_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	InstanceOwnerId     string `parquet:"name=instance_owner_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	NetworkCardIndex    int    `parquet:"name=network_card_index, type=INT32"`
	Status              string `parquet:"name=status, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type NetworkInterfaceAssociationModel struct {
	AllocationId    string `parquet:"name=allocation_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	AssociationId   string `parquet:"name=association_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	CarrierIp       string `parquet:"name=carrier_ip, type=BYTE_ARRAY, convertedtype=UTF8"`
	CustomerOwnedIp string `parquet:"name=customer_owned_ip, type=BYTE_ARRAY, convertedtype=UTF8"`
	IpOwnerId       string `parquet:"name=ip_owner_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	PublicDnsName   string `parquet:"name=public_dns_name, type=BYTE_ARRAY, convertedtype=UTF8"`
	PublicIp        string `parquet:"name=public_ip, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type NetworkInterfaceDataSourceClient interface {
	DescribeNetworkInterfaces(context.Context, *awsec2.DescribeNetworkInterfacesInput, ...func(*awsec2.Options)) (*awsec2.DescribeNetworkInterfacesOutput, error)
}

func NetworkInterfaceDataSource(ctx context.Context, client *awsec2.Client, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	return networkInterfaceDataSource(ctx, client, storageConfig, storageManager)
}

// function with client as a specific interface, allowing mocking/testing
func networkInterfaceDataSource(ctx context.Context, client NetworkInterfaceDataSourceClient, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(NetworkInterfaceModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := awsec2.NewDescribeNetworkInterfacesPaginator(client, &awsec2.DescribeNetworkInterfacesInput{})

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
			}).Error("error calling DescribeNetworkInterfaces")
			return err
		}

		for _, networkInterface := range output.NetworkInterfaces {
			model := new(NetworkInterfaceModel)
			copier.Copy(&model, &networkInterface)

			model.TagSet = GetTagMap(networkInterface.TagSet)
			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region

			errors := storageContextSet.Store(ctx, model)
			if errors != nil {
				for storageContext, err := range errors {
					storage.LogContextError(storageContext, fmt.Sprintf("Error storing NetworkInterfaceModel: %v", err))
				}
			}
		}
	}

	return nil
}
