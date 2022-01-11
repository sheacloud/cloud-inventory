package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud/services/cloudwatchlogs"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud/services/ec2"
)

type AwsCatalogResource struct {
	ResourceName  string
	ResourceModel interface{}
	FetchFunction func(context.Context, *awscloud.AwsFetchInput) *awscloud.AwsFetchOutput
}

type AwsCatalogService struct {
	ServiceName string
	Resources   []AwsCatalogResource
}

var (
	AwsCatalog = []AwsCatalogService{
		{
			ServiceName: "ec2",
			Resources: []AwsCatalogResource{
				{
					ResourceName:  "vpcs",
					ResourceModel: &ec2.VpcModel{},
					FetchFunction: ec2.FetchVpcs,
				},
				{
					ResourceName:  "volumes",
					ResourceModel: &ec2.VolumeModel{},
					FetchFunction: ec2.FetchVolumes,
				},
			},
		},
		{
			ServiceName: "cloudwatchlogs",
			Resources: []AwsCatalogResource{
				{
					ResourceName:  "log_groups",
					ResourceModel: &cloudwatchlogs.LogGroupModel{},
					FetchFunction: cloudwatchlogs.FetchLogGroups,
				},
			},
		},
	}
)
