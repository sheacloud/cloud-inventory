service "aws" "ec2" {
    library_path = "github.com/aws/aws-sdk-go-v2/service/ec2"

    datasource "images" {
        primary_object_name = "Image"
        primary_object_field = "ImageId"
        api_function = "DescribeImages"
        primary_object_path = ["Images"]
        models_only = true
        paginate = false

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "instances" {
        primary_object_name = "Instance"
        primary_object_field = "InstanceId"
        api_function = "DescribeInstances"
        primary_object_path = ["Reservations", "Instances"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "network_interfaces" {
        primary_object_name = "NetworkInterface"
        primary_object_field = "NetworkInterfaceId"
        api_function = "DescribeNetworkInterfaces"
        primary_object_path = ["NetworkInterfaces"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "TagSet"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "subnets" {
        primary_object_name = "Subnet"
        primary_object_field = "SubnetId"
        api_function = "DescribeSubnets"
        primary_object_path = ["Subnets"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "vpcs" {
        primary_object_name = "Vpc"
        primary_object_field = "VpcId"
        api_function = "DescribeVpcs"
        primary_object_path = ["Vpcs"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "volumes" {
        primary_object_name = "Volume"
        primary_object_field = "VolumeId"
        api_function = "DescribeVolumes"
        primary_object_path = ["Volumes"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "dhcp_options" {
        primary_object_name = "DhcpOptions"
        primary_object_field = "DhcpOptionsId"
        api_function = "DescribeDhcpOptions"
        primary_object_path = ["DhcpOptions"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "internet_gateways" {
        primary_object_name = "InternetGateway"
        primary_object_field = "InternetGatewayId"
        api_function = "DescribeInternetGateways"
        primary_object_path = ["InternetGateways"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "managed_prefix_lists" {
        primary_object_name = "ManagedPrefixList"
        primary_object_field = "PrefixListArn"
        api_function = "DescribeManagedPrefixLists"
        primary_object_path = ["PrefixLists"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "nat_gateways" {
        primary_object_name = "NatGateway"
        primary_object_field = "NatGatewayId"
        api_function = "DescribeNatGateways"
        primary_object_path = ["NatGateways"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "network_acls" {
        primary_object_name = "NetworkAcl"
        primary_object_field = "NetworkAclId"
        api_function = "DescribeNetworkAcls"
        primary_object_path = ["NetworkAcls"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "route_tables" {
        primary_object_name = "RouteTable"
        primary_object_field = "RouteTableId"
        api_function = "DescribeRouteTables"
        primary_object_path = ["RouteTables"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "security_groups" {
        primary_object_name = "SecurityGroup"
        primary_object_field = "GroupId"
        api_function = "DescribeSecurityGroups"
        primary_object_path = ["SecurityGroups"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "transit_gateways" {
        primary_object_name = "TransitGateway"
        primary_object_field = "TransitGatewayId"
        api_function = "DescribeTransitGateways"
        primary_object_path = ["TransitGateways"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "transit_gateway_route_tables" {
        primary_object_name = "TransitGatewayRouteTable"
        primary_object_field = "TransitGatewayRouteTableId"
        api_function = "DescribeTransitGatewayRouteTables"
        primary_object_path = ["TransitGatewayRouteTables"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "transit_gateway_vpc_attachments" {
        primary_object_name = "TransitGatewayVpcAttachment"
        primary_object_field = "TransitGatewayAttachmentId"
        api_function = "DescribeTransitGatewayVpcAttachments"
        primary_object_path = ["TransitGatewayVpcAttachments"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "transit_gateway_peering_attachments" {
        primary_object_name = "TransitGatewayPeeringAttachment"
        primary_object_field = "TransitGatewayAttachmentId"
        api_function = "DescribeTransitGatewayPeeringAttachments"
        primary_object_path = ["TransitGatewayPeeringAttachments"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "vpc_endpoints" {
        primary_object_name = "VpcEndpoint"
        primary_object_field = "VpcEndpointId"
        api_function = "DescribeVpcEndpoints"
        primary_object_path = ["VpcEndpoints"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "vpc_peering_connections" {
        primary_object_name = "VpcPeeringConnection"
        primary_object_field = "VpcPeeringConnectionId"
        api_function = "DescribeVpcPeeringConnections"
        primary_object_path = ["VpcPeeringConnections"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "vpn_gateways" {
        primary_object_name = "VpnGateway"
        primary_object_field = "VpnGatewayId"
        api_function = "DescribeVpnGateways"
        primary_object_path = ["VpnGateways"]
        models_only = false
        paginate = false

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "reserved_instances" {
        primary_object_name = "ReservedInstances"
        primary_object_field = "ReservedInstancesId"
        api_function = "DescribeReservedInstances"
        primary_object_path = ["ReservedInstances"]
        models_only = false
        paginate = false

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "placement_groups" {
        primary_object_name = "PlacementGroup"
        primary_object_field = "GroupId"
        api_function = "DescribePlacementGroups"
        primary_object_path = ["PlacementGroups"]
        models_only = false
        paginate = false

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "addresses" {
        primary_object_name = "Address"
        primary_object_field = "AllocationId"
        api_function = "DescribeAddresses"
        primary_object_path = ["Addresses"]
        models_only = false
        paginate = false

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }
}