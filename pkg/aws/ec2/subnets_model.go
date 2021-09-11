// AUTOGENERATED, DO NOT EDIT
package ec2

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

var customSubnetModelPostprocessingFuncs []func(x *SubnetModel) = []func(x *SubnetModel){}
var customSubnetModelFuncsLock sync.Mutex

func registerCustomSubnetModelPostprocessingFunc(f func(x *SubnetModel)) {
	customSubnetModelFuncsLock.Lock()
	defer customSubnetModelFuncsLock.Unlock()

	customSubnetModelPostprocessingFuncs = append(customSubnetModelPostprocessingFuncs, f)
}

func init() {
	Controller.RegisterDataSource("subnets", SubnetDataSource)
}

type SubnetModel struct {
	AssignIpv6AddressOnCreation bool                                         `parquet:"name=assign_ipv6_address_on_creation,type=BOOLEAN"`
	AvailabilityZone            string                                       `parquet:"name=availability_zone,type=BYTE_ARRAY,convertedtype=UTF8"`
	AvailabilityZoneId          string                                       `parquet:"name=availability_zone_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	AvailableIpAddressCount     int32                                        `parquet:"name=available_ip_address_count,type=INT32"`
	CidrBlock                   string                                       `parquet:"name=cidr_block,type=BYTE_ARRAY,convertedtype=UTF8"`
	CustomerOwnedIpv4Pool       string                                       `parquet:"name=customer_owned_ipv4_pool,type=BYTE_ARRAY,convertedtype=UTF8"`
	DefaultForAz                bool                                         `parquet:"name=default_for_az,type=BOOLEAN"`
	Ipv6CidrBlockAssociationSet []*SubnetIpv6CidrBlockAssociationSubnetModel `parquet:"name=ipv6_cidr_block_association_set,type=LIST"`
	MapCustomerOwnedIpOnLaunch  bool                                         `parquet:"name=map_customer_owned_ip_on_launch,type=BOOLEAN"`
	MapPublicIpOnLaunch         bool                                         `parquet:"name=map_public_ip_on_launch,type=BOOLEAN"`
	OutpostArn                  string                                       `parquet:"name=outpost_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	OwnerId                     string                                       `parquet:"name=owner_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	State                       string                                       `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
	SubnetArn                   string                                       `parquet:"name=subnet_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	SubnetId                    string                                       `parquet:"name=subnet_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	Tags                        map[string]string                            `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
	VpcId                       string                                       `parquet:"name=vpc_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	AccountId                   string                                       `parquet:"name=account_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Region                      string                                       `parquet:"name=region, type=BYTE_ARRAY, convertedtype=UTF8"`
	ReportTime                  int64                                        `parquet:"name=report_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
}

type SubnetIpv6CidrBlockAssociationSubnetModel struct {
	AssociationId      string                           `parquet:"name=association_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Ipv6CidrBlock      string                           `parquet:"name=ipv6_cidr_block,type=BYTE_ARRAY,convertedtype=UTF8"`
	Ipv6CidrBlockState *SubnetCidrBlockStateSubnetModel `parquet:"name=ipv6_cidr_block_state"`
}

type SubnetCidrBlockStateSubnetModel struct {
	State         string `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
	StatusMessage string `parquet:"name=status_message,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type TagSubnetModel struct {
	Key   string `parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

func SubnetDataSource(ctx context.Context, client *ec2.Client, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(SubnetModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := ec2.NewDescribeSubnetsPaginator(client, &ec2.DescribeSubnetsInput{})

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

		for _, var0 := range output.Subnets {

			model := new(SubnetModel)
			copier.Copy(&model, &var0)

			model.Tags = GetTagMap(var0.Tags)
			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region
			model.ReportTime = reportTime.UTC().UnixMilli()

			for _, f := range customSubnetModelPostprocessingFuncs {
				f(model)
			}

			errors := storageContextSet.Store(ctx, model)
			for storageContext, err := range errors {
				storage.LogContextError(storageContext, fmt.Sprintf("Error storing SubnetModel: %v", err))
			}
		}

	}

	return nil
}