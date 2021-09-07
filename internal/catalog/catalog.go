package catalog

import (
	"github.com/sheacloud/cloud-inventory/internal/controller"
	"github.com/sheacloud/cloud-inventory/pkg/aws/ec2"
	"github.com/sheacloud/cloud-inventory/pkg/aws/iam"
)

var (
	AwsServiceControllers = []controller.AwsController{
		&ec2.Controller,
		&iam.Controller,
	}

	DatasourceModels = map[string]map[string]map[string]interface{}{
		"aws": {
			"ec2": {
				"instances":          new(ec2.InstanceModel),
				"volumes":            new(ec2.VolumeModel),
				"vpcs":               new(ec2.VpcModel),
				"subnets":            new(ec2.SubnetModel),
				"network_interfaces": new(ec2.NetworkInterfaceModel),
			},
			"iam": {
				"roles": new(iam.RoleModel),
			},
		},
	}
)
