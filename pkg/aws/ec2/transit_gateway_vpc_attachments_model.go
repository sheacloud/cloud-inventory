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

var customTransitGatewayVpcAttachmentModelPostprocessingFuncs []func(ctx context.Context, client *ec2.Client, cfg aws.Config, x *TransitGatewayVpcAttachmentModel) = []func(ctx context.Context, client *ec2.Client, cfg aws.Config, x *TransitGatewayVpcAttachmentModel){}
var customTransitGatewayVpcAttachmentModelFuncsLock sync.Mutex

func registerCustomTransitGatewayVpcAttachmentModelPostprocessingFunc(f func(ctx context.Context, client *ec2.Client, cfg aws.Config, x *TransitGatewayVpcAttachmentModel)) {
	customTransitGatewayVpcAttachmentModelFuncsLock.Lock()
	defer customTransitGatewayVpcAttachmentModelFuncsLock.Unlock()

	customTransitGatewayVpcAttachmentModelPostprocessingFuncs = append(customTransitGatewayVpcAttachmentModelPostprocessingFuncs, f)
}

func init() {
	Controller.RegisterDataSource("transit_gateway_vpc_attachments", TransitGatewayVpcAttachmentDataSource)
}

type TransitGatewayVpcAttachmentModel struct {
	CreationTime               *time.Time
	Options                    *TransitGatewayVpcAttachmentOptionsTransitGatewayVpcAttachmentModel `parquet:"name=options"`
	State                      string                                                              `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
	SubnetIds                  []string                                                            `parquet:"name=subnet_ids,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8"`
	TagsOld                    []*TagTransitGatewayVpcAttachmentModel
	TransitGatewayAttachmentId string            `parquet:"name=transit_gateway_attachment_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	TransitGatewayId           string            `parquet:"name=transit_gateway_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	VpcId                      string            `parquet:"name=vpc_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	VpcOwnerId                 string            `parquet:"name=vpc_owner_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	CreationTimeMilli          int64             `parquet:"name=creation_time_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	Tags                       map[string]string `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
	AccountId                  string            `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region                     string            `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime                 int64             `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type TransitGatewayVpcAttachmentOptionsTransitGatewayVpcAttachmentModel struct {
	ApplianceModeSupport string `parquet:"name=appliance_mode_support,type=BYTE_ARRAY,convertedtype=UTF8"`
	DnsSupport           string `parquet:"name=dns_support,type=BYTE_ARRAY,convertedtype=UTF8"`
	Ipv6Support          string `parquet:"name=ipv6_support,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type TagTransitGatewayVpcAttachmentModel struct {
	Key   string `parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

func TransitGatewayVpcAttachmentDataSource(ctx context.Context, client *ec2.Client, cfg aws.Config, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(TransitGatewayVpcAttachmentModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := ec2.NewDescribeTransitGatewayVpcAttachmentsPaginator(client, &ec2.DescribeTransitGatewayVpcAttachmentsInput{})

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
			}).Error("error calling DescribeTransitGatewayVpcAttachments")
			return err
		}

		for _, var0 := range output.TransitGatewayVpcAttachments {

			model := new(TransitGatewayVpcAttachmentModel)
			copier.Copy(&model, &var0)

			model.Tags = GetTagMap(var0.Tags)
			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region
			model.ReportTime = reportTime.UTC().UnixMilli()

			for _, f := range customTransitGatewayVpcAttachmentModelPostprocessingFuncs {
				f(ctx, client, cfg, model)
			}

			errors := storageContextSet.Store(ctx, model)
			for storageContext, err := range errors {
				storage.LogContextError(storageContext, fmt.Sprintf("Error storing TransitGatewayVpcAttachmentModel: %v", err))
			}
		}

	}

	return nil
}
