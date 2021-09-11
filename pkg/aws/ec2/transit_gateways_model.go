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

var customTransitGatewayModelPostprocessingFuncs []func(x *TransitGatewayModel) = []func(x *TransitGatewayModel){}
var customTransitGatewayModelFuncsLock sync.Mutex

func registerCustomTransitGatewayModelPostprocessingFunc(f func(x *TransitGatewayModel)) {
	customTransitGatewayModelFuncsLock.Lock()
	defer customTransitGatewayModelFuncsLock.Unlock()

	customTransitGatewayModelPostprocessingFuncs = append(customTransitGatewayModelPostprocessingFuncs, f)
}

func init() {
	Controller.RegisterDataSource("transit_gateways", TransitGatewayDataSource)
}

type TransitGatewayModel struct {
	CreationTime      *time.Time
	CreationTimeMilli int64                                     `parquet:"name=creation_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	Description       string                                    `parquet:"name=description,type=BYTE_ARRAY,convertedtype=UTF8"`
	Options           *TransitGatewayOptionsTransitGatewayModel `parquet:"name=options"`
	OwnerId           string                                    `parquet:"name=owner_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	State             string                                    `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
	Tags              map[string]string                         `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
	TransitGatewayArn string                                    `parquet:"name=transit_gateway_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	TransitGatewayId  string                                    `parquet:"name=transit_gateway_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	AccountId         string                                    `parquet:"name=account_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Region            string                                    `parquet:"name=region, type=BYTE_ARRAY, convertedtype=UTF8"`
	ReportTime        int64                                     `parquet:"name=report_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
}

type TransitGatewayOptionsTransitGatewayModel struct {
	AmazonSideAsn                  int64    `parquet:"name=amazon_side_asn,type=INT64"`
	AssociationDefaultRouteTableId string   `parquet:"name=association_default_route_table_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	AutoAcceptSharedAttachments    string   `parquet:"name=auto_accept_shared_attachments,type=BYTE_ARRAY,convertedtype=UTF8"`
	DefaultRouteTableAssociation   string   `parquet:"name=default_route_table_association,type=BYTE_ARRAY,convertedtype=UTF8"`
	DefaultRouteTablePropagation   string   `parquet:"name=default_route_table_propagation,type=BYTE_ARRAY,convertedtype=UTF8"`
	DnsSupport                     string   `parquet:"name=dns_support,type=BYTE_ARRAY,convertedtype=UTF8"`
	MulticastSupport               string   `parquet:"name=multicast_support,type=BYTE_ARRAY,convertedtype=UTF8"`
	PropagationDefaultRouteTableId string   `parquet:"name=propagation_default_route_table_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	TransitGatewayCidrBlocks       []string `parquet:"name=transit_gateway_cidr_blocks,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	VpnEcmpSupport                 string   `parquet:"name=vpn_ecmp_support,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type TagTransitGatewayModel struct {
	Key   string `parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

func TransitGatewayDataSource(ctx context.Context, client *ec2.Client, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(TransitGatewayModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := ec2.NewDescribeTransitGatewaysPaginator(client, &ec2.DescribeTransitGatewaysInput{})

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
			}).Error("error calling DescribeTransitGateways")
			return err
		}

		for _, var0 := range output.TransitGateways {

			model := new(TransitGatewayModel)
			copier.Copy(&model, &var0)

			model.Tags = GetTagMap(var0.Tags)
			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region
			model.ReportTime = reportTime.UTC().UnixMilli()

			for _, f := range customTransitGatewayModelPostprocessingFuncs {
				f(model)
			}

			errors := storageContextSet.Store(ctx, model)
			for storageContext, err := range errors {
				storage.LogContextError(storageContext, fmt.Sprintf("Error storing TransitGatewayModel: %v", err))
			}
		}

	}

	return nil
}
