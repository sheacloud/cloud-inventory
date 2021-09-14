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

var customRouteTableModelPostprocessingFuncs []func(ctx context.Context, client *ec2.Client, cfg aws.Config, x *RouteTableModel) = []func(ctx context.Context, client *ec2.Client, cfg aws.Config, x *RouteTableModel){}
var customRouteTableModelFuncsLock sync.Mutex

func registerCustomRouteTableModelPostprocessingFunc(f func(ctx context.Context, client *ec2.Client, cfg aws.Config, x *RouteTableModel)) {
	customRouteTableModelFuncsLock.Lock()
	defer customRouteTableModelFuncsLock.Unlock()

	customRouteTableModelPostprocessingFuncs = append(customRouteTableModelPostprocessingFuncs, f)
}

func init() {
	Controller.RegisterDataSource("route_tables", RouteTableDataSource)
}

type RouteTableModel struct {
	Associations    []*RouteTableAssociationRouteTableModel `parquet:"name=associations,type=LIST"`
	OwnerId         string                                  `parquet:"name=owner_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	PropagatingVgws []*PropagatingVgwRouteTableModel        `parquet:"name=propagating_vgws,type=LIST"`
	RouteTableId    string                                  `parquet:"name=route_table_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	Routes          []*RouteRouteTableModel                 `parquet:"name=routes,type=LIST"`
	Tags            map[string]string                       `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
	VpcId           string                                  `parquet:"name=vpc_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	AccountId       string                                  `parquet:"name=account_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Region          string                                  `parquet:"name=region, type=BYTE_ARRAY, convertedtype=UTF8"`
	ReportTime      int64                                   `parquet:"name=report_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
}

type RouteTableAssociationRouteTableModel struct {
	AssociationState        *RouteTableAssociationStateRouteTableModel `parquet:"name=association_state"`
	GatewayId               string                                     `parquet:"name=gateway_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Main                    bool                                       `parquet:"name=main,type=BOOLEAN"`
	RouteTableAssociationId string                                     `parquet:"name=route_table_association_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	RouteTableId            string                                     `parquet:"name=route_table_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	SubnetId                string                                     `parquet:"name=subnet_id,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type RouteTableAssociationStateRouteTableModel struct {
	State         string `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
	StatusMessage string `parquet:"name=status_message,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type PropagatingVgwRouteTableModel struct {
	GatewayId string `parquet:"name=gateway_id,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type RouteRouteTableModel struct {
	CarrierGatewayId            string `parquet:"name=carrier_gateway_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	DestinationCidrBlock        string `parquet:"name=destination_cidr_block,type=BYTE_ARRAY,convertedtype=UTF8"`
	DestinationIpv6CidrBlock    string `parquet:"name=destination_ipv6_cidr_block,type=BYTE_ARRAY,convertedtype=UTF8"`
	DestinationPrefixListId     string `parquet:"name=destination_prefix_list_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	EgressOnlyInternetGatewayId string `parquet:"name=egress_only_internet_gateway_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	GatewayId                   string `parquet:"name=gateway_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	InstanceId                  string `parquet:"name=instance_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	InstanceOwnerId             string `parquet:"name=instance_owner_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	LocalGatewayId              string `parquet:"name=local_gateway_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	NatGatewayId                string `parquet:"name=nat_gateway_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	NetworkInterfaceId          string `parquet:"name=network_interface_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Origin                      string `parquet:"name=origin,type=BYTE_ARRAY,convertedtype=UTF8"`
	State                       string `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
	TransitGatewayId            string `parquet:"name=transit_gateway_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	VpcPeeringConnectionId      string `parquet:"name=vpc_peering_connection_id,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type TagRouteTableModel struct {
	Key   string `parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

func RouteTableDataSource(ctx context.Context, client *ec2.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(RouteTableModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := ec2.NewDescribeRouteTablesPaginator(client, &ec2.DescribeRouteTablesInput{})

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
			}).Error("error calling DescribeRouteTables")
			return err
		}

		for _, var0 := range output.RouteTables {

			model := new(RouteTableModel)
			copier.Copy(&model, &var0)

			model.Tags = GetTagMap(var0.Tags)
			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region
			model.ReportTime = reportTime.UTC().UnixMilli()

			for _, f := range customRouteTableModelPostprocessingFuncs {
				f(ctx, client, cfg, model)
			}

			errors := storageContextSet.Store(ctx, model)
			for storageContext, err := range errors {
				storage.LogContextError(storageContext, fmt.Sprintf("Error storing RouteTableModel: %v", err))
			}
		}

	}

	return nil
}
