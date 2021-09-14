package catalog

import (
	"github.com/sheacloud/cloud-inventory/internal/controller"
	"github.com/sheacloud/cloud-inventory/pkg/aws/cloudwatchlogs"
	"github.com/sheacloud/cloud-inventory/pkg/aws/dynamodb"
	"github.com/sheacloud/cloud-inventory/pkg/aws/ec2"
	"github.com/sheacloud/cloud-inventory/pkg/aws/ecs"
	"github.com/sheacloud/cloud-inventory/pkg/aws/efs"
	"github.com/sheacloud/cloud-inventory/pkg/aws/eks"
	"github.com/sheacloud/cloud-inventory/pkg/aws/elasticloadbalancing"
	"github.com/sheacloud/cloud-inventory/pkg/aws/elasticloadbalancingv2"
	"github.com/sheacloud/cloud-inventory/pkg/aws/iam"
	"github.com/sheacloud/cloud-inventory/pkg/aws/lambda"
	"github.com/sheacloud/cloud-inventory/pkg/aws/rds"
	"github.com/sheacloud/cloud-inventory/pkg/aws/redshift"
	"github.com/sheacloud/cloud-inventory/pkg/aws/route53"
	"github.com/sheacloud/cloud-inventory/pkg/aws/s3"
	"github.com/sheacloud/cloud-inventory/pkg/aws/sns"
	"github.com/sheacloud/cloud-inventory/pkg/aws/sqs"
)

var (
	AwsServiceControllers = []controller.AwsController{
		&cloudwatchlogs.Controller,
		&dynamodb.Controller,
		&ec2.Controller,
		&ecs.Controller,
		&efs.Controller,
		&eks.Controller,
		&elasticloadbalancing.Controller,
		&elasticloadbalancingv2.Controller,
		&iam.Controller,
		&lambda.Controller,
		&rds.Controller,
		&redshift.Controller,
		&route53.Controller,
		&s3.Controller,
		&sns.Controller,
		&sqs.Controller,
	}

	DatasourceModels = map[string]map[string]map[string]interface{}{
		"aws": {
			"cloudwatchlogs": {
				"log_groups": new(cloudwatchlogs.LogGroupModel),
			},
			"dynamodb": {},
			"ec2": {
				"instances":                           new(ec2.InstanceModel),
				"volumes":                             new(ec2.VolumeModel),
				"vpcs":                                new(ec2.VpcModel),
				"subnets":                             new(ec2.SubnetModel),
				"network_interfaces":                  new(ec2.NetworkInterfaceModel),
				"dhcp_options":                        new(ec2.DhcpOptionsModel),
				"internet_gateways":                   new(ec2.InternetGatewayModel),
				"managed_prefix_lists":                new(ec2.ManagedPrefixListModel),
				"nat_gateways":                        new(ec2.NatGatewayModel),
				"network_acls":                        new(ec2.NetworkAclModel),
				"route_tables":                        new(ec2.RouteTableModel),
				"security_groups":                     new(ec2.SecurityGroupModel),
				"transit_gateways":                    new(ec2.TransitGatewayModel),
				"transit_gateway_route_tables":        new(ec2.TransitGatewayRouteTableModel),
				"transit_gateway_vpc_attachments":     new(ec2.TransitGatewayVpcAttachmentModel),
				"transit_gateway_peering_attachments": new(ec2.TransitGatewayPeeringAttachmentModel),
				"vpc_endpoints":                       new(ec2.VpcEndpointModel),
				"vpc_peering_connections":             new(ec2.VpcPeeringConnectionModel),
				"vpn_gateways":                        new(ec2.VpnGatewayModel),
				"reserved_instances":                  new(ec2.ReservedInstancesModel),
				"placement_groups":                    new(ec2.PlacementGroupModel),
				"addresses":                           new(ec2.AddressModel),
				"images":                              new(ec2.ImageModel),
			},
			"ecs": {
				"clusters": new(ecs.ClusterModel),
			},
			"efs": {
				"filesystems": new(efs.FileSystemDescriptionModel),
			},
			"eks": {},
			"elasticloadbalancing": {
				"load_balancers": new(elasticloadbalancing.LoadBalancerDescriptionModel),
			},
			"elasticloadbalancingv2": {
				"load_balancers": new(elasticloadbalancingv2.LoadBalancerModel),
				"target_groups":  new(elasticloadbalancingv2.TargetGroupModel),
			},
			"iam": {
				"roles":             new(iam.RoleModel),
				"policies":          new(iam.PolicyModel),
				"users":             new(iam.UserModel),
				"groups":            new(iam.GroupModel),
				"instance_profiles": new(iam.InstanceProfileModel),
			},
			"lambda": {},
			"rds": {
				"db_clusters":  new(rds.DBClusterModel),
				"db_instances": new(rds.DBInstanceModel),
			},
			"redshift": {},
			"route53":  {},
			"s3": {
				"buckets": new(s3.BucketModel),
			},
			"sns": {},
			"sqs": {},
		},
	}
)
