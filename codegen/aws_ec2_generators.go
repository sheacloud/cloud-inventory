package codegen

//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name instances -primary-resource-name Instance -primary-resource-field InstanceId -api-function DescribeInstances -primary-object-path Reservations,Instances -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name network_interfaces -primary-resource-name NetworkInterface -primary-resource-field NetworkInterfaceId -api-function DescribeNetworkInterfaces -primary-object-path NetworkInterfaces -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name subnets -primary-resource-name Subnet -primary-resource-field SubnetId -api-function DescribeSubnets -primary-object-path Subnets -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name vpcs -primary-resource-name Vpc -primary-resource-field VpcId -api-function DescribeVpcs -primary-object-path Vpcs -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name volumes -primary-resource-name Volume -primary-resource-field VolumeId -api-function DescribeVolumes -primary-object-path Volumes -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name dhcp_options -primary-resource-name DhcpOptions -primary-resource-field DhcpOptionsId -api-function DescribeDhcpOptions -primary-object-path DhcpOptions -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name internet_gateways -primary-resource-name InternetGateway -primary-resource-field InternetGatewayId -api-function DescribeInternetGateways -primary-object-path InternetGateways -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name managed_prefix_lists -primary-resource-name ManagedPrefixList -primary-resource-field PrefixListArn -api-function DescribeManagedPrefixLists -primary-object-path PrefixLists -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name nat_gateways -primary-resource-name NatGateway -primary-resource-field NatGatewayId -api-function DescribeNatGateways -primary-object-path NatGateways -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name network_acls -primary-resource-name NetworkAcl -primary-resource-field NetworkAclId -api-function DescribeNetworkAcls -primary-object-path NetworkAcls -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name route_tables -primary-resource-name RouteTable -primary-resource-field RouteTableId -api-function DescribeRouteTables -primary-object-path RouteTables -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name security_groups -primary-resource-name SecurityGroup -primary-resource-field GroupId -api-function DescribeSecurityGroups -primary-object-path SecurityGroups -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg

//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name transit_gateways -primary-resource-name TransitGateway -primary-resource-field TransitGatewayId -api-function DescribeTransitGateways -primary-object-path TransitGateways -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name transit_gateway_route_tables -primary-resource-name TransitGatewayRouteTable -primary-resource-field TransitGatewayRouteTableId -api-function DescribeTransitGatewayRouteTables -primary-object-path TransitGatewayRouteTables -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name transit_gateway_vpc_attachments -primary-resource-name TransitGatewayVpcAttachment -primary-resource-field TransitGatewayAttachmentId -api-function DescribeTransitGatewayVpcAttachments -primary-object-path TransitGatewayVpcAttachments -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name transit_gateway_peering_attachments -primary-resource-name TransitGatewayPeeringAttachment -primary-resource-field TransitGatewayAttachmentId -api-function DescribeTransitGatewayPeeringAttachments -primary-object-path TransitGatewayPeeringAttachments -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg

//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name vpc_endpoints -primary-resource-name VpcEndpoint -primary-resource-field VpcEndpointId -api-function DescribeVpcEndpoints -primary-object-path VpcEndpoints -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg

//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name vpc_peering_connections -primary-resource-name VpcPeeringConnection -primary-resource-field VpcPeeringConnectionId -api-function DescribeVpcPeeringConnections -primary-object-path VpcPeeringConnections -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg

//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name vpn_gateways -primary-resource-name VpnGateway -primary-resource-field VpnGatewayId -api-function DescribeVpnGateways -primary-object-path VpnGateways -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg -paginate=false

//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name reserved_instances -primary-resource-name ReservedInstances -primary-resource-field ReservedInstancesId -api-function DescribeReservedInstances -primary-object-path ReservedInstances -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg -paginate=false

//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name placement_groups -primary-resource-name PlacementGroup -primary-resource-field GroupId -api-function DescribePlacementGroups -primary-object-path PlacementGroups -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg -paginate=false

//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name ec2 -data-source-name addresses -primary-resource-name Address -primary-resource-field AllocationId -api-function DescribeAddresses -primary-object-path Addresses -library-path "github.com/aws/aws-sdk-go-v2/service/ec2" -base-output-path ../pkg -paginate=false

//go:generate go fmt ../pkg/aws/ec2/
