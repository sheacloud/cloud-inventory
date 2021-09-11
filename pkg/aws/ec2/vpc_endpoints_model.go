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

var customVpcEndpointModelPostprocessingFuncs []func(x *VpcEndpointModel) = []func(x *VpcEndpointModel){}
var customVpcEndpointModelFuncsLock sync.Mutex

func registerCustomVpcEndpointModelPostprocessingFunc(f func(x *VpcEndpointModel)) {
	customVpcEndpointModelFuncsLock.Lock()
	defer customVpcEndpointModelFuncsLock.Unlock()

	customVpcEndpointModelPostprocessingFuncs = append(customVpcEndpointModelPostprocessingFuncs, f)
}

func init() {
	Controller.RegisterDataSource("vpc_endpoints", VpcEndpointDataSource)
}

type VpcEndpointModel struct {
	CreationTimestamp      *time.Time
	CreationTimestampMilli int64                                      `parquet:"name=creation_timestamp, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	DnsEntries             []*DnsEntryVpcEndpointModel                `parquet:"name=dns_entries,type=LIST"`
	Groups                 []*SecurityGroupIdentifierVpcEndpointModel `parquet:"name=groups,type=LIST"`
	LastError              *LastErrorVpcEndpointModel                 `parquet:"name=last_error"`
	NetworkInterfaceIds    []string                                   `parquet:"name=network_interface_ids,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	OwnerId                string                                     `parquet:"name=owner_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	PolicyDocument         string                                     `parquet:"name=policy_document,type=BYTE_ARRAY,convertedtype=UTF8"`
	PrivateDnsEnabled      bool                                       `parquet:"name=private_dns_enabled,type=BOOLEAN"`
	RequesterManaged       bool                                       `parquet:"name=requester_managed,type=BOOLEAN"`
	RouteTableIds          []string                                   `parquet:"name=route_table_ids,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	ServiceName            string                                     `parquet:"name=service_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	State                  string                                     `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
	SubnetIds              []string                                   `parquet:"name=subnet_ids,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	Tags                   map[string]string                          `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
	VpcEndpointId          string                                     `parquet:"name=vpc_endpoint_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	VpcEndpointType        string                                     `parquet:"name=vpc_endpoint_type,type=BYTE_ARRAY,convertedtype=UTF8"`
	VpcId                  string                                     `parquet:"name=vpc_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	AccountId              string                                     `parquet:"name=account_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Region                 string                                     `parquet:"name=region, type=BYTE_ARRAY, convertedtype=UTF8"`
	ReportTime             int64                                      `parquet:"name=report_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
}

type DnsEntryVpcEndpointModel struct {
	DnsName      string `parquet:"name=dns_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	HostedZoneId string `parquet:"name=hosted_zone_id,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type SecurityGroupIdentifierVpcEndpointModel struct {
	GroupId   string `parquet:"name=group_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	GroupName string `parquet:"name=group_name,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type LastErrorVpcEndpointModel struct {
	Code    string `parquet:"name=code,type=BYTE_ARRAY,convertedtype=UTF8"`
	Message string `parquet:"name=message,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type TagVpcEndpointModel struct {
	Key   string `parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

func VpcEndpointDataSource(ctx context.Context, client *ec2.Client, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(VpcEndpointModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := ec2.NewDescribeVpcEndpointsPaginator(client, &ec2.DescribeVpcEndpointsInput{})

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
			}).Error("error calling DescribeVpcEndpoints")
			return err
		}

		for _, var0 := range output.VpcEndpoints {

			model := new(VpcEndpointModel)
			copier.Copy(&model, &var0)

			model.Tags = GetTagMap(var0.Tags)
			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region
			model.ReportTime = reportTime.UTC().UnixMilli()

			for _, f := range customVpcEndpointModelPostprocessingFuncs {
				f(model)
			}

			errors := storageContextSet.Store(ctx, model)
			for storageContext, err := range errors {
				storage.LogContextError(storageContext, fmt.Sprintf("Error storing VpcEndpointModel: %v", err))
			}
		}

	}

	return nil
}
