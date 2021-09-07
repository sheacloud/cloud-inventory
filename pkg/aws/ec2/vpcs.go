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
	Controller.RegisterDataSource("vpcs", VpcDataSource)
}

type VpcModel struct {
	CidrBlock                   string                             `parquet:"name=cidr_block, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	CidrBlockAssociationSet     []VpcCidrBlockAssociationModel     `parquet:"name=cidr_block_association_set, type=LIST"`
	DhcpOptionsId               string                             `parquet:"name=dhcp_options_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	InstanceTenancy             string                             `parquet:"name=tenancy, type=BYTE_ARRAY, convertedtype=UTF8"`
	Ipv6CidrBlockAssociationSet []VpcIpv6CidrBlockAssociationModel `parquet:"name=ipv6_cidr_block_association_set, type=LIST"`
	IsDefault                   bool                               `parquet:"name=is_default, type=BOOLEAN"`
	OwnerId                     string                             `parquet:"name=owner_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	State                       string                             `parquet:"name=state, type=BYTE_ARRAY, convertedtype=UTF8"`
	Tags                        map[string]string                  `parquet:"name=tags, type=MAP, keytype=BYTE_ARRAY, keyconvertedtype=UTF8, valuetype=BYTE_ARRAY, valueconvertedtype=UTF8"`
	VpcId                       string                             `parquet:"name=vpc_id, type=BYTE_ARRAY, convertedtype=UTF8" inventory_primary_key:"true"`
	AccountId                   string                             `parquet:"name=account_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Region                      string                             `parquet:"name=region, type=BYTE_ARRAY, convertedtype=UTF8"`
	ReportTime                  int64                              `parquet:"name=report_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
}

type VpcIpv6CidrBlockAssociationModel struct {
	AssociationId      string                 `parquet:"name=association_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Ipv6CidrBlock      string                 `parquet:"name=ipv6_cidr_block, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Ipv6CidrBlockState VpcCidrBlockStateModel `parquet:"name=ipv6_cidr_block_state"`
	Ipv6Pool           string                 `parquet:"name=ipv6_pool, type=BYTE_ARRAY, convertedtype=UTF8"`
	NetworkBorderGroup string                 `parquet:"name=network_border_group, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type VpcCidrBlockAssociationModel struct {
	AssociationId  string                 `parquet:"name=association_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	CidrBlock      string                 `parquet:"name=cidr_block, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	CidrBlockState VpcCidrBlockStateModel `parquet:"name=cidr_block_state"`
}

type VpcCidrBlockStateModel struct {
	State         string `parquet:"name=state, type=BYTE_ARRAY, convertedtype=UTF8"`
	StatusMessage string `parquet:"name=association_id, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type VpcDataSourceClient interface {
	DescribeVpcs(context.Context, *awsec2.DescribeVpcsInput, ...func(*awsec2.Options)) (*awsec2.DescribeVpcsOutput, error)
}

func VpcDataSource(ctx context.Context, client *awsec2.Client, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	return vpcDataSource(ctx, client, reportTime, storageConfig, storageManager)
}

// function with client as a specific interface, allowing mocking/testing
func vpcDataSource(ctx context.Context, client VpcDataSourceClient, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(VpcModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := awsec2.NewDescribeVpcsPaginator(client, &awsec2.DescribeVpcsInput{})

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
			}).Error("error calling DescribeVpcs")
			return err
		}

		for _, vpc := range output.Vpcs {
			model := new(VpcModel)
			copier.Copy(&model, &vpc)

			model.Tags = GetTagMap(vpc.Tags)
			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region
			model.ReportTime = reportTime.UnixMilli()

			errors := storageContextSet.Store(ctx, model)
			for storageContext, err := range errors {
				storage.LogContextError(storageContext, fmt.Sprintf("Error storing VpcModel: %v", err))
			}
		}
	}

	return nil
}
