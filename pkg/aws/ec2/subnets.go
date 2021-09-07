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
	Controller.RegisterDataSource("subnets", SubnetDataSource)
}

type SubnetModel struct {
	AssignIpv6AddressOnCreation bool                                  `parquet:"name=assign_ipv6_address_on_creation, type=BOOLEAN"`
	AvailabilityZone            string                                `parquet:"name=availability_zone, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	AvailabilityZoneId          string                                `parquet:"name=availability_zone_id, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	AvailableIpAddressCount     int32                                 `parquet:"name=available_ip_address_count, type=INT32"`
	CidrBlock                   string                                `parquet:"name=cidr_block, type=BYTE_ARRAY, convertedtype=UTF8"`
	CustomerOwnedIpv4Pool       string                                `parquet:"name=customer_owned_ipv4_pool, type=BYTE_ARRAY, convertedtype=UTF8"`
	DefaultForAz                bool                                  `parquet:"name=default_for_az, type=BOOLEAN"`
	Ipv6CidrBlockAssociationSet []SubnetIpv6CidrBlockAssociationModel `parquet:"name=ipv6_cidr_block_association_set, type=LIST"`
	MapCustomerOwnedIpOnLaunch  bool                                  `parquet:"name=map_customer_owned_ip_on_launch, type=BOOLEAN"`
	MapPublicIpOnLaunch         bool                                  `parquet:"name=map_public_ip_on_launch, type=BOOLEAN"`
	OwnerId                     string                                `parquet:"name=owner_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	State                       string                                `parquet:"name=state, type=BYTE_ARRAY, convertedtype=UTF8"`
	SubnetArn                   string                                `parquet:"name=subnet_arn, type=BYTE_ARRAY, convertedtype=UTF8"`
	SubnetId                    string                                `parquet:"name=subnet_id, type=BYTE_ARRAY, convertedtype=UTF8" inventory_primary_key:"true"`
	Tags                        map[string]string                     `parquet:"name=tags, type=MAP, keytype=BYTE_ARRAY, keyconvertedtype=UTF8, valuetype=BYTE_ARRAY, valueconvertedtype=UTF8"`
	VpcId                       string                                `parquet:"name=vpc_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	AccountId                   string                                `parquet:"name=account_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Region                      string                                `parquet:"name=region, type=BYTE_ARRAY, convertedtype=UTF8"`
	ReportTime                  int64                                 `parquet:"name=report_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
}

type SubnetIpv6CidrBlockAssociationModel struct {
	AssociationId      string                    `parquet:"name=association_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Ipv6CidrBlock      string                    `parquet:"name=ipv6_cidr_block, type=BYTE_ARRAY, convertedtype=UTF8"`
	Ipv6CidrBlockState SubnetCidrBlockStateModel `parquet:"name=ipv6_cidr_block_state"`
}

type SubnetCidrBlockStateModel struct {
	State         string `parquet:"name=state, type=BYTE_ARRAY, convertedtype=UTF8"`
	StatusMessage string `parquet:"name=status_message, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type SubnetDataSourceClient interface {
	DescribeSubnets(context.Context, *awsec2.DescribeSubnetsInput, ...func(*awsec2.Options)) (*awsec2.DescribeSubnetsOutput, error)
}

func SubnetDataSource(ctx context.Context, client *awsec2.Client, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	return subnetDataSource(ctx, client, reportTime, storageConfig, storageManager)
}

// function with client as a specific interface, allowing mocking/testing
func subnetDataSource(ctx context.Context, client SubnetDataSourceClient, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(SubnetModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := awsec2.NewDescribeSubnetsPaginator(client, &awsec2.DescribeSubnetsInput{})

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
			}).Error("error calling DescribeSubnets")
			return err
		}

		for _, subnet := range output.Subnets {
			model := new(SubnetModel)
			copier.Copy(&model, &subnet)

			model.Tags = GetTagMap(subnet.Tags)
			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region
			model.ReportTime = reportTime.UnixMilli()

			errors := storageContextSet.Store(ctx, model)
			for storageContext, err := range errors {
				storage.LogContextError(storageContext, fmt.Sprintf("Error storing SubnetModel: %v", err))
			}
		}
	}

	return nil
}
