
resource "aws_dynamodb_table" "aws_ec2_addresses" {
  name         = "cloud-inventory-aws-ec2-addresses"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "allocation_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "allocation_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "allocation_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_dhcp_options" {
  name         = "cloud-inventory-aws-ec2-dhcp-options"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "dhcp_options_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "dhcp_options_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "dhcp_options_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_images" {
  name         = "cloud-inventory-aws-ec2-images"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "image_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "image_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "image_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_instances" {
  name         = "cloud-inventory-aws-ec2-instances"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "instance_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "instance_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "instance_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_internet_gateways" {
  name         = "cloud-inventory-aws-ec2-internet-gateways"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "internet_gateway_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "internet_gateway_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "internet_gateway_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_managed_prefix_lists" {
  name         = "cloud-inventory-aws-ec2-managed-prefix-lists"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "prefix_list_arn"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "prefix_list_arn"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "prefix_list_arn"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_nat_gateways" {
  name         = "cloud-inventory-aws-ec2-nat-gateways"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "nat_gateway_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "nat_gateway_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "nat_gateway_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_network_acls" {
  name         = "cloud-inventory-aws-ec2-network-acls"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "network_acl_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "network_acl_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "network_acl_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_network_interfaces" {
  name         = "cloud-inventory-aws-ec2-network-interfaces"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "network_interface_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "network_interface_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "network_interface_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_placement_groups" {
  name         = "cloud-inventory-aws-ec2-placement-groups"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "group_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "group_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "group_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_reserved_instances" {
  name         = "cloud-inventory-aws-ec2-reserved-instances"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "reserved_instances_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "reserved_instances_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "reserved_instances_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_route_tables" {
  name         = "cloud-inventory-aws-ec2-route-tables"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "route_table_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "route_table_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "route_table_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_security_groups" {
  name         = "cloud-inventory-aws-ec2-security-groups"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "group_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "group_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "group_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_subnets" {
  name         = "cloud-inventory-aws-ec2-subnets"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "subnet_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "subnet_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "subnet_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_transit_gateway_peering_attachments" {
  name         = "cloud-inventory-aws-ec2-transit-gateway-peering-attachments"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "transit_gateway_attachment_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "transit_gateway_attachment_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "transit_gateway_attachment_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_transit_gateway_route_tables" {
  name         = "cloud-inventory-aws-ec2-transit-gateway-route-tables"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "transit_gateway_route_table_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "transit_gateway_route_table_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "transit_gateway_route_table_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_transit_gateway_vpc_attachments" {
  name         = "cloud-inventory-aws-ec2-transit-gateway-vpc-attachments"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "transit_gateway_attachment_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "transit_gateway_attachment_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "transit_gateway_attachment_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_transit_gateways" {
  name         = "cloud-inventory-aws-ec2-transit-gateways"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "transit_gateway_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "transit_gateway_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "transit_gateway_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_volumes" {
  name         = "cloud-inventory-aws-ec2-volumes"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "volume_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "volume_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "volume_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_vpc_endpoints" {
  name         = "cloud-inventory-aws-ec2-vpc-endpoints"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "vpc_endpoint_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "vpc_endpoint_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "vpc_endpoint_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_vpc_peering_connections" {
  name         = "cloud-inventory-aws-ec2-vpc-peering-connections"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "vpc_peering_connection_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "vpc_peering_connection_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "vpc_peering_connection_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_vpcs" {
  name         = "cloud-inventory-aws-ec2-vpcs"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "vpc_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "vpc_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "vpc_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ec2_vpn_gateways" {
  name         = "cloud-inventory-aws-ec2-vpn-gateways"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "vpn_gateway_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "vpn_gateway_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "vpn_gateway_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
