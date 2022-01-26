aws_service "ec2" {
    service_cap_name = "EC2"
    sdk_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    extra_utilized_functions = []
    tag_object_name = "Tag"

    resource "addresses" {
        fetch_function = "DescribeAddresses"
        object_name = "Address"
        object_plural_name = "Addresses"
        object_unique_id = "AllocationId"
        object_response_field = "Addresses"
        model_only = false
        pagination = false
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "dhcp_options" {
        fetch_function = "DescribeDhcpOptions"
        object_name = "DhcpOptions"
        object_plural_name = "DhcpOptions"
        object_unique_id = "DhcpOptionsId"
        object_response_field = "DhcpOptions"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "images" {
        fetch_function = "DescribeImages"
        object_name = "Image"
        object_plural_name = "Images"
        object_unique_id = "ImageId"
        object_response_field = "Images"
        model_only = true
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "instance_types" {
        fetch_function = "DescribeInstanceTypes"
        object_name = "InstanceTypeInfo"
        object_plural_name = "InstanceTypes"
        object_unique_id = "InstanceType"
        object_response_field = "InstanceTypes"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = false
    }

    resource "instances" {
        fetch_function = "DescribeInstances"
        object_name = "Instance"
        object_plural_name = "Instances"
        object_unique_id = "InstanceId"
        object_response_field = "Instances"
        model_only = true
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "internet_gateways" {
        fetch_function = "DescribeInternetGateways"
        object_name = "InternetGateway"
        object_plural_name = "InternetGateways"
        object_unique_id = "InternetGatewayId"
        object_response_field = "InternetGateways"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }    

    resource "managed_prefix_lists" {
        fetch_function = "DescribeManagedPrefixLists"
        object_name = "ManagedPrefixList"
        object_plural_name = "ManagedPrefixLists"
        object_unique_id = "PrefixListArn"
        object_response_field = "PrefixLists"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "nat_gateways" {
        fetch_function = "DescribeNatGateways"
        object_name = "NatGateway"
        object_plural_name = "NatGateways"
        object_unique_id = "NatGatewayId"
        object_response_field = "NatGateways"
        model_only = false
        pagination = true
        use_post_processing = true
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "network_acls" {
        fetch_function = "DescribeNetworkAcls"
        object_name = "NetworkAcl"
        object_plural_name = "NetworkAcls"
        object_unique_id = "NetworkAclId"
        object_response_field = "NetworkAcls"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "network_interfaces" {
        fetch_function = "DescribeNetworkInterfaces"
        object_name = "NetworkInterface"
        object_plural_name = "NetworkInterfaces"
        object_unique_id = "NetworkInterfaceId"
        object_response_field = "NetworkInterfaces"
        model_only = false
        pagination = true
        use_post_processing = true
        excluded_fields = []
        convert_tags = true
        tag_field_name = "TagSet"
    }

    resource "placement_groups" {
        fetch_function = "DescribePlacementGroups"
        object_name = "PlacementGroup"
        object_plural_name = "PlacementGroups"
        object_unique_id = "GroupId"
        object_response_field = "PlacementGroups"
        model_only = false
        pagination = false
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "reserved_instances" {
        fetch_function = "DescribeReservedInstances"
        object_name = "ReservedInstances"
        object_plural_name = "ReservedInstances"
        object_unique_id = "ReservedInstancesId"
        object_response_field = "ReservedInstances"
        model_only = false
        pagination = false
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "route_tables" {
        fetch_function = "DescribeRouteTables"
        object_name = "RouteTable"
        object_plural_name = "RouteTables"
        object_unique_id = "RouteTableId"
        object_response_field = "RouteTables"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "security_groups" {
        fetch_function = "DescribeSecurityGroups"
        object_name = "SecurityGroup"
        object_plural_name = "SecurityGroups"
        object_unique_id = "GroupId"
        object_response_field = "SecurityGroups"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "subnets" {
        fetch_function = "DescribeSubnets"
        object_name = "Subnet"
        object_plural_name = "Subnets"
        object_unique_id = "SubnetId"
        object_response_field = "Subnets"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "transit_gateway_peering_attachments" {
        fetch_function = "DescribeTransitGatewayPeeringAttachments"
        object_name = "TransitGatewayPeeringAttachment"
        object_plural_name = "TransitGatewayPeeringAttachments"
        object_unique_id = "TransitGatewayAttachmentId"
        object_response_field = "TransitGatewayPeeringAttachments"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "transit_gateway_route_tables" {
        fetch_function = "DescribeTransitGatewayRouteTables"
        object_name = "TransitGatewayRouteTable"
        object_plural_name = "TransitGatewayRouteTables"
        object_unique_id = "TransitGatewayRouteTableId"
        object_response_field = "TransitGatewayRouteTables"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "transit_gateway_vpc_attachments" {
        fetch_function = "DescribeTransitGatewayVpcAttachments"
        object_name = "TransitGatewayVpcAttachment"
        object_plural_name = "TransitGatewayVpcAttachments"
        object_unique_id = "TransitGatewayAttachmentId"
        object_response_field = "TransitGatewayVpcAttachments"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "transit_gateways" {
        fetch_function = "DescribeTransitGateways"
        object_name = "TransitGateway"
        object_plural_name = "TransitGateways"
        object_unique_id = "TransitGatewayId"
        object_response_field = "TransitGateways"
        model_only = false
        pagination = true
        use_post_processing = true
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }
    
    resource "volumes" {
        fetch_function = "DescribeVolumes"
        object_name = "Volume"
        object_plural_name = "Volumes"
        object_unique_id = "VolumeId"
        object_response_field = "Volumes"
        model_only = false
        pagination = true
        use_post_processing = true
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "vpc_endpoints" {
        fetch_function = "DescribeVpcEndpoints"
        object_name = "VpcEndpoint"
        object_plural_name = "VpcEndpoints"
        object_unique_id = "VpcEndpointId"
        object_response_field = "VpcEndpoints"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "vpc_peering_connections" {
        fetch_function = "DescribeVpcPeeringConnections"
        object_name = "VpcPeeringConnection"
        object_plural_name = "VpcPeeringConnections"
        object_unique_id = "VpcPeeringConnectionId"
        object_response_field = "VpcPeeringConnections"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "vpcs" {
        fetch_function = "DescribeVpcs"
        object_name = "Vpc"
        object_plural_name = "Vpcs"
        object_unique_id = "VpcId"
        object_response_field = "Vpcs"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "vpn_gateways" {
        fetch_function = "DescribeVpnGateways"
        object_name = "VpnGateway"
        object_plural_name = "VpnGateways"
        object_unique_id = "VpnGatewayId"
        object_response_field = "VpnGateways"
        model_only = false
        pagination = false
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }   
}