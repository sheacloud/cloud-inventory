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
			},
			"iam": {
				"roles":             new(iam.RoleModel),
				"policies":          new(iam.PolicyModel),
				"users":             new(iam.UserModel),
				"groups":            new(iam.GroupModel),
				"instance_profiles": new(iam.InstanceProfileModel),
			},
		},
	}
)
