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

var customAddressModelPostprocessingFuncs []func(ctx context.Context, client *ec2.Client, cfg aws.Config, x *AddressModel) = []func(ctx context.Context, client *ec2.Client, cfg aws.Config, x *AddressModel){}
var customAddressModelFuncsLock sync.Mutex

func registerCustomAddressModelPostprocessingFunc(f func(ctx context.Context, client *ec2.Client, cfg aws.Config, x *AddressModel)) {
	customAddressModelFuncsLock.Lock()
	defer customAddressModelFuncsLock.Unlock()

	customAddressModelPostprocessingFuncs = append(customAddressModelPostprocessingFuncs, f)
}

func init() {
	Controller.RegisterDataSource("addresses", AddressDataSource)
}

type AddressModel struct {
	AllocationId            string `parquet:"name=allocation_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	AssociationId           string `parquet:"name=association_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	CarrierIp               string `parquet:"name=carrier_ip,type=BYTE_ARRAY,convertedtype=UTF8"`
	CustomerOwnedIp         string `parquet:"name=customer_owned_ip,type=BYTE_ARRAY,convertedtype=UTF8"`
	CustomerOwnedIpv4Pool   string `parquet:"name=customer_owned_ipv4_pool,type=BYTE_ARRAY,convertedtype=UTF8"`
	Domain                  string `parquet:"name=domain,type=BYTE_ARRAY,convertedtype=UTF8"`
	InstanceId              string `parquet:"name=instance_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	NetworkBorderGroup      string `parquet:"name=network_border_group,type=BYTE_ARRAY,convertedtype=UTF8"`
	NetworkInterfaceId      string `parquet:"name=network_interface_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	NetworkInterfaceOwnerId string `parquet:"name=network_interface_owner_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	PrivateIpAddress        string `parquet:"name=private_ip_address,type=BYTE_ARRAY,convertedtype=UTF8"`
	PublicIp                string `parquet:"name=public_ip,type=BYTE_ARRAY,convertedtype=UTF8"`
	PublicIpv4Pool          string `parquet:"name=public_ipv4_pool,type=BYTE_ARRAY,convertedtype=UTF8"`
	TagsOld                 []*TagAddressModel
	Tags                    map[string]string `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
	AccountId               string            `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region                  string            `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime              int64             `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type TagAddressModel struct {
	Key   string `parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

func AddressDataSource(ctx context.Context, client *ec2.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(AddressModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	params := &ec2.DescribeAddressesInput{}

	result, err := client.DescribeAddresses(ctx, params)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     storageConfig.Service,
			"data_source": storageConfig.DataSource,
			"account_id":  storageConfig.AccountId,
			"region":      storageConfig.Region,
			"cloud":       storageConfig.Cloud,
			"error":       err,
		}).Error("error calling DescribeAddresses")
		return err
	}

	results := []*ec2.DescribeAddressesOutput{result}
	for _, output := range results {

		for _, var0 := range output.Addresses {

			model := new(AddressModel)
			copier.Copy(&model, &var0)

			model.Tags = GetTagMap(var0.Tags)
			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region
			model.ReportTime = reportTime.UTC().UnixMilli()

			for _, f := range customAddressModelPostprocessingFuncs {
				f(ctx, client, cfg, model)
			}

			errors := storageContextSet.Store(ctx, model)
			for storageContext, err := range errors {
				storage.LogContextError(storageContext, fmt.Sprintf("Error storing AddressModel: %v", err))
			}
		}

	}

	return nil
}
