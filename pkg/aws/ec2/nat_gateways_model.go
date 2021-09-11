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

var customNatGatewayModelPostprocessingFuncs []func(x *NatGatewayModel) = []func(x *NatGatewayModel){}
var customNatGatewayModelFuncsLock sync.Mutex

func registerCustomNatGatewayModelPostprocessingFunc(f func(x *NatGatewayModel)) {
	customNatGatewayModelFuncsLock.Lock()
	defer customNatGatewayModelFuncsLock.Unlock()

	customNatGatewayModelPostprocessingFuncs = append(customNatGatewayModelPostprocessingFuncs, f)
}

func init() {
	Controller.RegisterDataSource("nat_gateways", NatGatewayDataSource)
}

type NatGatewayModel struct {
	ConnectivityType     string `parquet:"name=connectivity_type,type=BYTE_ARRAY,convertedtype=UTF8"`
	CreateTime           *time.Time
	CreateTimeMilli      int64 `parquet:"name=create_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	DeleteTime           *time.Time
	DeleteTimeMilli      int64                                `parquet:"name=delete_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	FailureCode          string                               `parquet:"name=failure_code,type=BYTE_ARRAY,convertedtype=UTF8"`
	FailureMessage       string                               `parquet:"name=failure_message,type=BYTE_ARRAY,convertedtype=UTF8"`
	NatGatewayAddresses  []*NatGatewayAddressNatGatewayModel  `parquet:"name=nat_gateway_addresses,type=LIST"`
	NatGatewayId         string                               `parquet:"name=nat_gateway_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	ProvisionedBandwidth *ProvisionedBandwidthNatGatewayModel `parquet:"name=provisioned_bandwidth"`
	State                string                               `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
	SubnetId             string                               `parquet:"name=subnet_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Tags                 map[string]string                    `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
	VpcId                string                               `parquet:"name=vpc_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	AccountId            string                               `parquet:"name=account_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Region               string                               `parquet:"name=region, type=BYTE_ARRAY, convertedtype=UTF8"`
	ReportTime           int64                                `parquet:"name=report_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
}

type NatGatewayAddressNatGatewayModel struct {
	AllocationId       string `parquet:"name=allocation_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	NetworkInterfaceId string `parquet:"name=network_interface_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	PrivateIp          string `parquet:"name=private_ip,type=BYTE_ARRAY,convertedtype=UTF8"`
	PublicIp           string `parquet:"name=public_ip,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type ProvisionedBandwidthNatGatewayModel struct {
	ProvisionTime      *time.Time
	ProvisionTimeMilli int64  `parquet:"name=provision_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	Provisioned        string `parquet:"name=provisioned,type=BYTE_ARRAY,convertedtype=UTF8"`
	RequestTime        *time.Time
	RequestTimeMilli   int64  `parquet:"name=request_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	Requested          string `parquet:"name=requested,type=BYTE_ARRAY,convertedtype=UTF8"`
	Status             string `parquet:"name=status,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type TagNatGatewayModel struct {
	Key   string `parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

func NatGatewayDataSource(ctx context.Context, client *ec2.Client, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(NatGatewayModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := ec2.NewDescribeNatGatewaysPaginator(client, &ec2.DescribeNatGatewaysInput{})

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
			}).Error("error calling DescribeNatGateways")
			return err
		}

		for _, var0 := range output.NatGateways {

			model := new(NatGatewayModel)
			copier.Copy(&model, &var0)

			model.Tags = GetTagMap(var0.Tags)
			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region
			model.ReportTime = reportTime.UTC().UnixMilli()

			for _, f := range customNatGatewayModelPostprocessingFuncs {
				f(model)
			}

			errors := storageContextSet.Store(ctx, model)
			for storageContext, err := range errors {
				storage.LogContextError(storageContext, fmt.Sprintf("Error storing NatGatewayModel: %v", err))
			}
		}

	}

	return nil
}