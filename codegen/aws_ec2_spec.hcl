datasource "aws" "ec2" "images" {
    primary_resource_name = "Image"
    primary_resource_field = "ImageId"
    api_function = "DescribeImages"
    primary_object_path = ["Images"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = true
    paginate = false
}

datasource "aws" "ec2" "instances" {
    primary_resource_name = "Instance"
    primary_resource_field = "InstanceId"
    api_function = "DescribeInstances"
    primary_object_path = ["Reservations", "Instances"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "network_interfaces" {
    primary_resource_name = "NetworkInterface"
    primary_resource_field = "NetworkInterfaceId"
    api_function = "DescribeNetworkInterfaces"
    primary_object_path = ["NetworkInterfaces"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "subnets" {
    primary_resource_name = "Subnet"
    primary_resource_field = "SubnetId"
    api_function = "DescribeSubnets"
    primary_object_path = ["Subnets"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "vpcs" {
    primary_resource_name = "Vpc"
    primary_resource_field = "VpcId"
    api_function = "DescribeVpcs"
    primary_object_path = ["Vpcs"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "volumes" {
    primary_resource_name = "Volume"
    primary_resource_field = "VolumeId"
    api_function = "DescribeVolumes"
    primary_object_path = ["Volumes"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "dhcp_options" {
    primary_resource_name = "DhcpOptions"
    primary_resource_field = "DhcpOptionsId"
    api_function = "DescribeDhcpOptions"
    primary_object_path = ["DhcpOptions"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "internet_gateways" {
    primary_resource_name = "InternetGateway"
    primary_resource_field = "InternetGatewayId"
    api_function = "DescribeInternetGateways"
    primary_object_path = ["InternetGateways"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "managed_prefix_lists" {
    primary_resource_name = "ManagedPrefixList"
    primary_resource_field = "PrefixListArn"
    api_function = "DescribeManagedPrefixLists"
    primary_object_path = ["PrefixLists"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "nat_gateways" {
    primary_resource_name = "NatGateway"
    primary_resource_field = "NatGatewayId"
    api_function = "DescribeNatGateways"
    primary_object_path = ["NatGateways"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "network_acls" {
    primary_resource_name = "NetworkAcl"
    primary_resource_field = "NetworkAclId"
    api_function = "DescribeNetworkAcls"
    primary_object_path = ["NetworkAcls"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "route_tables" {
    primary_resource_name = "RouteTable"
    primary_resource_field = "RouteTableId"
    api_function = "DescribeRouteTables"
    primary_object_path = ["RouteTables"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "security_groups" {
    primary_resource_name = "SecurityGroup"
    primary_resource_field = "GroupId"
    api_function = "DescribeSecurityGroups"
    primary_object_path = ["SecurityGroups"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "transit_gateways" {
    primary_resource_name = "TransitGateway"
    primary_resource_field = "TransitGatewayId"
    api_function = "DescribeTransitGateways"
    primary_object_path = ["TransitGateways"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "transit_gateway_route_tables" {
    primary_resource_name = "TransitGatewayRouteTable"
    primary_resource_field = "TransitGatewayRouteTableId"
    api_function = "DescribeTransitGatewayRouteTables"
    primary_object_path = ["TransitGatewayRouteTables"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "transit_gateway_vpc_attachments" {
    primary_resource_name = "TransitGatewayVpcAttachment"
    primary_resource_field = "TransitGatewayAttachmentId"
    api_function = "DescribeTransitGatewayVpcAttachments"
    primary_object_path = ["TransitGatewayVpcAttachments"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "transit_gateway_peering_attachments" {
    primary_resource_name = "TransitGatewayPeeringAttachment"
    primary_resource_field = "TransitGatewayAttachmentId"
    api_function = "DescribeTransitGatewayPeeringAttachments"
    primary_object_path = ["TransitGatewayPeeringAttachments"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "vpc_endpoints" {
    primary_resource_name = "VpcEndpoint"
    primary_resource_field = "VpcEndpointId"
    api_function = "DescribeVpcEndpoints"
    primary_object_path = ["VpcEndpoints"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "vpc_peering_connections" {
    primary_resource_name = "VpcPeeringConnection"
    primary_resource_field = "VpcPeeringConnectionId"
    api_function = "DescribeVpcPeeringConnections"
    primary_object_path = ["VpcPeeringConnections"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = true
}

datasource "aws" "ec2" "vpn_gateways" {
    primary_resource_name = "VpnGateway"
    primary_resource_field = "VpnGatewayId"
    api_function = "DescribeVpnGateways"
    primary_object_path = ["VpnGateways"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = false
}

datasource "aws" "ec2" "reserved_instances" {
    primary_resource_name = "ReservedInstances"
    primary_resource_field = "ReservedInstancesId"
    api_function = "DescribeReservedInstances"
    primary_object_path = ["ReservedInstances"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = false
}

datasource "aws" "ec2" "placement_groups" {
    primary_resource_name = "PlacementGroup"
    primary_resource_field = "GroupId"
    api_function = "DescribePlacementGroups"
    primary_object_path = ["PlacementGroups"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = false
}

datasource "aws" "ec2" "addresses" {
    primary_resource_name = "Address"
    primary_resource_field = "AllocationId"
    api_function = "DescribeAddresses"
    primary_object_path = ["Addresses"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    models_only = false
    paginate = false
}