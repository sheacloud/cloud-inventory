aws_service "ec2" {
    service_cap_name = "EC2"
    sdk_path = "github.com/aws/aws-sdk-go-v2/service/ec2"
    extra_utilized_functions = []
    tag_object_name = "Tag"

    resource "images" {
        fetch_function = "DescribeImages"
        object_name = "Image"
        object_unique_id = "ImageId"
        object_response_field = "Images"
        model_only = true
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "instances" {
        fetch_function = "DescribeInstances"
        object_name = "Instance"
        object_unique_id = "InstanceId"
        object_response_field = "Instances"
        model_only = true
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "network_interfaces" {
        fetch_function = "DescribeNetworkInterfaces"
        object_name = "NetworkInterface"
        object_unique_id = "NetworkInterfaceId"
        object_response_field = "NetworkInterfaces"
        model_only = false
        pagination = true
        use_post_processing = true
        excluded_fields = []
        convert_tags = true
        tag_field_name = "TagSet"
    }

    resource "subnets" {
        fetch_function = "DescribeSubnets"
        object_name = "Subnet"
        object_unique_id = "SubnetId"
        object_response_field = "Subnets"
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
        object_unique_id = "VpcId"
        object_response_field = "Vpcs"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "volumes" {
        fetch_function = "DescribeVolumes"
        object_name = "Volume"
        object_unique_id = "VolumeId"
        object_response_field = "Volumes"
        model_only = false
        pagination = true
        use_post_processing = true
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "dhcp_options" {
        fetch_function = "DescribeDhcpOptions"
        object_name = "DhcpOptions"
        object_unique_id = "DhcpOptionsId"
        object_response_field = "DhcpOptions"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "internet_gateways" {
        fetch_function = "DescribeInternetGateways"
        object_name = "InternetGateway"
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
        object_unique_id = "NetworkAclId"
        object_response_field = "NetworkAcls"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "route_tables" {
        fetch_function = "DescribeRouteTables"
        object_name = "RouteTable"
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
        object_unique_id = "GroupId"
        object_response_field = "SecurityGroups"
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
        object_unique_id = "TransitGatewayId"
        object_response_field = "TransitGateways"
        model_only = false
        pagination = true
        use_post_processing = true
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "transit_gateway_route_tables" {
        fetch_function = "DescribeTransitGatewayRouteTables"
        object_name = "TransitGatewayRouteTable"
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
        object_unique_id = "TransitGatewayVpcAttachmentId"
        object_response_field = "TransitGatewayVpcAttachments"
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
        object_unique_id = "TransitGatewayPeeringAttachmentId"
        object_response_field = "TransitGatewayPeeringAttachments"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "vpc_endpoints" {
        fetch_function = "DescribeVpcEndpoints"
        object_name = "VpcEndpoint"
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
        object_unique_id = "VpcPeeringConnectionId"
        object_response_field = "VpcPeeringConnections"
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
        object_unique_id = "VpnGatewayId"
        object_response_field = "VpnGateways"
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
        object_unique_id = "ReservedInstancesId"
        object_response_field = "ReservedInstances"
        model_only = false
        pagination = false
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "placement_groups" {
        fetch_function = "DescribePlacementGroups"
        object_name = "PlacementGroup"
        object_unique_id = "GroupId"
        object_response_field = "PlacementGroups"
        model_only = false
        pagination = false
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "addresses" {
        fetch_function = "DescribeAddresses"
        object_name = "Address"
        object_unique_id = "AllocationId"
        object_response_field = "Addresses"
        model_only = false
        pagination = false
        use_post_processing = false
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
    }

    resource "instance_types" {
        fetch_function = "DescribeInstanceTypes"
        object_name = "InstanceTypeInfo"
        object_unique_id = "InstanceType"
        object_response_field = "InstanceTypes"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = false
    }
}