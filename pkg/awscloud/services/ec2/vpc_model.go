package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

type VpcModel struct {
	CidrBlock                   string                                 `parquet:"name=cidr_block,type=BYTE_ARRAY,convertedtype=UTF8"`
	CidrBlockAssociationSet     []*VpcCidrBlockAssociationVpcModel     `parquet:"name=cidr_block_association_set,type=MAP,convertedtype=LIST"`
	DhcpOptionsId               string                                 `parquet:"name=dhcp_options_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	InstanceTenancy             string                                 `parquet:"name=instance_tenancy,type=BYTE_ARRAY,convertedtype=UTF8"`
	Ipv6CidrBlockAssociationSet []*VpcIpv6CidrBlockAssociationVpcModel `parquet:"name=ipv6_cidr_block_association_set,type=MAP,convertedtype=LIST"`
	IsDefault                   bool                                   `parquet:"name=is_default,type=BOOLEAN"`
	OwnerId                     string                                 `parquet:"name=owner_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	State                       string                                 `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
	TagsOld                     []*TagVpcModel
	VpcId                       string            `parquet:"name=vpc_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	Tags                        map[string]string `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
	AccountId                   string            `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region                      string            `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime                  int64             `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type VpcCidrBlockAssociationVpcModel struct {
	AssociationId  string                     `parquet:"name=association_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	CidrBlock      string                     `parquet:"name=cidr_block,type=BYTE_ARRAY,convertedtype=UTF8"`
	CidrBlockState *VpcCidrBlockStateVpcModel `parquet:"name=cidr_block_state"`
}

type VpcCidrBlockStateVpcModel struct {
	State         string `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
	StatusMessage string `parquet:"name=status_message,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type VpcIpv6CidrBlockAssociationVpcModel struct {
	AssociationId      string                     `parquet:"name=association_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Ipv6CidrBlock      string                     `parquet:"name=ipv6_cidr_block,type=BYTE_ARRAY,convertedtype=UTF8"`
	Ipv6CidrBlockState *VpcCidrBlockStateVpcModel `parquet:"name=ipv6_cidr_block_state"`
	Ipv6Pool           string                     `parquet:"name=ipv6_pool,type=BYTE_ARRAY,convertedtype=UTF8"`
	NetworkBorderGroup string                     `parquet:"name=network_border_group,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type TagVpcModel struct {
	Key   string `parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

func FetchVpcs(ctx context.Context, params *awscloud.AwsFetchInput) *awscloud.AwsFetchOutput {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int

	awsClient := params.RegionalClients[params.Region]
	ec2Client := awsClient.EC2()

	paginator := ec2.NewDescribeVpcsPaginator(ec2Client, &ec2.DescribeVpcsInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeVpcs in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, var0 := range output.Vpcs {

			model := new(VpcModel)
			copier.Copy(&model, &var0)

			model.Tags = GetTagMap(var0.Tags)
			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime.UTC().UnixMilli()

			err = params.OutputFile.Write(ctx, model)
			if err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error storing VPC model in %s/%s: %w", params.AccountId, params.Region, err))
			}
			fetchedResources++
		}

	}

	return &awscloud.AwsFetchOutput{
		FetchingErrors:   fetchingErrors,
		FetchedResources: fetchedResources,
		FailedResources:  failedResources,
		ResourceName:     "vpcs",
		AccountId:        params.AccountId,
		Region:           params.Region,
	}
}
