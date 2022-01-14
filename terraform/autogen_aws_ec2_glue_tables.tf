
resource "aws_glue_catalog_table" "aws_ec2_images" {
  name          = "aws_ec2_images"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/images/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "architecture"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "block_device_mappings"
      type       = "array<struct<device_name:string,ebs:struct<delete_on_termination:boolean,encrypted:boolean,iops:int,kms_key_id:string,outpost_arn:string,snapshot_id:string,throughput:int,volume_size:int,volume_type:string>,no_device:string,virtual_name:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "boot_mode"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "creation_date"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "deprecation_time"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "description"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ena_support"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "hypervisor"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "image_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "image_location"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "image_owner_alias"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "image_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "kernel_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "owner_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "platform"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "platform_details"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "product_codes"
      type       = "array<struct<product_code_id:string,product_code_type:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "public"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ramdisk_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "root_device_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "root_device_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "sriov_net_support"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state_reason"
      type       = "struct<code:string,message:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "usage_operation"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "virtualization_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_instances" {
  name          = "aws_ec2_instances"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/instances/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "ami_launch_index"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "architecture"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "block_device_mappings"
      type       = "array<struct<device_name:string,ebs:struct<delete_on_termination:boolean,status:string,volume_id:string,attach_time:timestamp>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "boot_mode"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "capacity_reservation_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "capacity_reservation_specification"
      type       = "struct<capacity_reservation_preference:string,capacity_reservation_target:struct<capacity_reservation_id:string,capacity_reservation_resource_group_arn:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "client_token"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cpu_options"
      type       = "struct<core_count:int,threads_per_core:int>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ebs_optimized"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "elastic_gpu_associations"
      type       = "array<struct<elastic_gpu_association_id:string,elastic_gpu_association_state:string,elastic_gpu_association_time:string,elastic_gpu_id:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "elastic_inference_accelerator_associations"
      type       = "array<struct<elastic_inference_accelerator_arn:string,elastic_inference_accelerator_association_id:string,elastic_inference_accelerator_association_state:string,elastic_inference_accelerator_association_time:timestamp>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ena_support"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "enclave_options"
      type       = "struct<enabled:boolean>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "hibernation_options"
      type       = "struct<configured:boolean>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "hypervisor"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "iam_instance_profile"
      type       = "struct<arn:string,id:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "image_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "instance_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "instance_lifecycle"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "instance_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ipv6_address"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "kernel_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "key_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "licenses"
      type       = "array<struct<license_configuration_arn:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "metadata_options"
      type       = "struct<http_endpoint:string,http_protocol_ipv6:string,http_put_response_hop_limit:int,http_tokens:string,state:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "monitoring"
      type       = "struct<state:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "network_interfaces"
      type       = "array<struct<association:struct<carrier_ip:string,customer_owned_ip:string,ip_owner_id:string,public_dns_name:string,public_ip:string>,attachment:struct<attachment_id:string,delete_on_termination:boolean,device_index:int,network_card_index:int,status:string,attach_time:timestamp>,description:string,groups:array<struct<group_id:string,group_name:string>>,interface_type:string,ipv4_prefixes:array<struct<ipv4_prefix:string>>,ipv6_addresses:array<struct<ipv6_address:string>>,ipv6_prefixes:array<struct<ipv6_prefix:string>>,mac_address:string,network_interface_id:string,owner_id:string,private_dns_name:string,private_ip_address:string,private_ip_addresses:array<struct<association:struct<carrier_ip:string,customer_owned_ip:string,ip_owner_id:string,public_dns_name:string,public_ip:string>,primary:boolean,private_dns_name:string,private_ip_address:string>>,source_dest_check:boolean,status:string,subnet_id:string,vpc_id:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "outpost_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "placement"
      type       = "struct<affinity:string,availability_zone:string,group_name:string,host_id:string,host_resource_group_arn:string,partition_number:int,spread_domain:string,tenancy:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "platform"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "platform_details"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "private_dns_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "private_dns_name_options"
      type       = "struct<enable_resource_name_dns_aaaa_record:boolean,enable_resource_name_dns_a_record:boolean,hostname_type:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "private_ip_address"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "product_codes"
      type       = "array<struct<product_code_id:string,product_code_type:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "public_dns_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "public_ip_address"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ramdisk_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "root_device_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "root_device_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "security_groups"
      type       = "array<struct<group_id:string,group_name:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "source_dest_check"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "spot_instance_request_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "sriov_net_support"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state"
      type       = "struct<code:int,name:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state_reason"
      type       = "struct<code:string,message:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state_transition_reason"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "subnet_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "usage_operation"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "virtualization_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "launch_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "usage_operation_update_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_network_interfaces" {
  name          = "aws_ec2_network_interfaces"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/network_interfaces/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "association"
      type       = "struct<allocation_id:string,association_id:string,carrier_ip:string,customer_owned_ip:string,ip_owner_id:string,public_dns_name:string,public_ip:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "attachment"
      type       = "struct<attachment_id:string,delete_on_termination:boolean,device_index:int,instance_id:string,instance_owner_id:string,network_card_index:int,status:string,attach_time:timestamp>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "availability_zone"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "deny_all_igw_traffic"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "description"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "groups"
      type       = "array<struct<group_id:string,group_name:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "interface_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ipv4_prefixes"
      type       = "array<struct<ipv4_prefix:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ipv6_address"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ipv6_addresses"
      type       = "array<struct<ipv6_address:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ipv6_native"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ipv6_prefixes"
      type       = "array<struct<ipv6_prefix:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "mac_address"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "network_interface_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "outpost_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "owner_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "private_dns_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "private_ip_address"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "private_ip_addresses"
      type       = "array<struct<association:struct<allocation_id:string,association_id:string,carrier_ip:string,customer_owned_ip:string,ip_owner_id:string,public_dns_name:string,public_ip:string>,primary:boolean,private_dns_name:string,private_ip_address:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "requester_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "requester_managed"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "source_dest_check"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "subnet_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tag_set_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_subnets" {
  name          = "aws_ec2_subnets"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/subnets/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "assign_ipv6_address_on_creation"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "availability_zone"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "availability_zone_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "available_ip_address_count"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cidr_block"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "customer_owned_ipv4_pool"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "default_for_az"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "enable_dns64"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "enable_lni_at_device_index"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ipv6_cidr_block_association_set"
      type       = "array<struct<association_id:string,ipv6_cidr_block:string,ipv6_cidr_block_state:struct<state:string,status_message:string>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ipv6_native"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "map_customer_owned_ip_on_launch"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "map_public_ip_on_launch"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "outpost_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "owner_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "private_dns_name_options_on_launch"
      type       = "struct<enable_resource_name_dns_aaaa_record:boolean,enable_resource_name_dns_a_record:boolean,hostname_type:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "subnet_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "subnet_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_vpcs" {
  name          = "aws_ec2_vpcs"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/vpcs/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "cidr_block"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cidr_block_association_set"
      type       = "array<struct<association_id:string,cidr_block:string,cidr_block_state:struct<state:string,status_message:string>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "dhcp_options_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "instance_tenancy"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ipv6_cidr_block_association_set"
      type       = "array<struct<association_id:string,ipv6_cidr_block:string,ipv6_cidr_block_state:struct<state:string,status_message:string>,ipv6_pool:string,network_border_group:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "is_default"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "owner_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_volumes" {
  name          = "aws_ec2_volumes"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/volumes/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "attachments"
      type       = "array<struct<delete_on_termination:boolean,device:string,instance_id:string,state:string,volume_id:string,attach_time:timestamp>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "availability_zone"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "encrypted"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "fast_restored"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "iops"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "kms_key_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "multi_attach_enabled"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "outpost_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "size"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "snapshot_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "throughput"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "volume_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "volume_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "create_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_dhcp_options" {
  name          = "aws_ec2_dhcp_options"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/dhcp_options/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "dhcp_configurations"
      type       = "array<struct<key:string,values:array<struct<value:string>>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "dhcp_options_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "owner_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_internet_gateways" {
  name          = "aws_ec2_internet_gateways"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/internet_gateways/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "attachments"
      type       = "array<struct<state:string,vpc_id:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "internet_gateway_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "owner_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_managed_prefix_lists" {
  name          = "aws_ec2_managed_prefix_lists"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/managed_prefix_lists/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "address_family"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "max_entries"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "owner_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "prefix_list_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "prefix_list_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "prefix_list_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state_message"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "version"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_nat_gateways" {
  name          = "aws_ec2_nat_gateways"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/nat_gateways/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "connectivity_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "failure_code"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "failure_message"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "nat_gateway_addresses"
      type       = "array<struct<allocation_id:string,network_interface_id:string,private_ip:string,public_ip:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "nat_gateway_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "provisioned_bandwidth"
      type       = "struct<provisioned:string,requested:string,status:string,provision_time:timestamp,request_time:timestamp>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "subnet_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "create_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "delete_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_network_acls" {
  name          = "aws_ec2_network_acls"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/network_acls/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "associations"
      type       = "array<struct<network_acl_association_id:string,network_acl_id:string,subnet_id:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "entries"
      type       = "array<struct<cidr_block:string,egress:boolean,icmp_type_code:struct<code:int,type:int>,ipv6_cidr_block:string,port_range:struct<from:int,to:int>,protocol:string,rule_action:string,rule_number:int>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "is_default"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "network_acl_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "owner_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_route_tables" {
  name          = "aws_ec2_route_tables"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/route_tables/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "associations"
      type       = "array<struct<association_state:struct<state:string,status_message:string>,gateway_id:string,main:boolean,route_table_association_id:string,route_table_id:string,subnet_id:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "owner_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "propagating_vgws"
      type       = "array<struct<gateway_id:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "route_table_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "routes"
      type       = "array<struct<carrier_gateway_id:string,core_network_arn:string,destination_cidr_block:string,destination_ipv6_cidr_block:string,destination_prefix_list_id:string,egress_only_internet_gateway_id:string,gateway_id:string,instance_id:string,instance_owner_id:string,local_gateway_id:string,nat_gateway_id:string,network_interface_id:string,origin:string,state:string,transit_gateway_id:string,vpc_peering_connection_id:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_security_groups" {
  name          = "aws_ec2_security_groups"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/security_groups/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "description"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "group_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "group_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ip_permissions"
      type       = "array<struct<from_port:int,ip_protocol:string,ip_ranges:array<struct<cidr_ip:string,description:string>>,ipv6_ranges:array<struct<cidr_ipv6:string,description:string>>,prefix_list_ids:array<struct<description:string,prefix_list_id:string>>,to_port:int,user_id_group_pairs:array<struct<description:string,group_id:string,group_name:string,peering_status:string,user_id:string,vpc_id:string,vpc_peering_connection_id:string>>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ip_permissions_egress"
      type       = "array<struct<from_port:int,ip_protocol:string,ip_ranges:array<struct<cidr_ip:string,description:string>>,ipv6_ranges:array<struct<cidr_ipv6:string,description:string>>,prefix_list_ids:array<struct<description:string,prefix_list_id:string>>,to_port:int,user_id_group_pairs:array<struct<description:string,group_id:string,group_name:string,peering_status:string,user_id:string,vpc_id:string,vpc_peering_connection_id:string>>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "owner_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_transit_gateways" {
  name          = "aws_ec2_transit_gateways"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/transit_gateways/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "description"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "options"
      type       = "struct<amazon_side_asn:bigint,association_default_route_table_id:string,auto_accept_shared_attachments:string,default_route_table_association:string,default_route_table_propagation:string,dns_support:string,multicast_support:string,propagation_default_route_table_id:string,transit_gateway_cidr_blocks:array<string>,vpn_ecmp_support:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "owner_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "transit_gateway_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "transit_gateway_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "creation_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_transit_gateway_route_tables" {
  name          = "aws_ec2_transit_gateway_route_tables"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/transit_gateway_route_tables/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "default_association_route_table"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "default_propagation_route_table"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "transit_gateway_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "transit_gateway_route_table_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "creation_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_transit_gateway_vpc_attachments" {
  name          = "aws_ec2_transit_gateway_vpc_attachments"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/transit_gateway_vpc_attachments/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "options"
      type       = "struct<appliance_mode_support:string,dns_support:string,ipv6_support:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "subnet_ids"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "transit_gateway_attachment_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "transit_gateway_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_owner_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "creation_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_transit_gateway_peering_attachments" {
  name          = "aws_ec2_transit_gateway_peering_attachments"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/transit_gateway_peering_attachments/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "accepter_tgw_info"
      type       = "struct<owner_id:string,region:string,transit_gateway_id:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "requester_tgw_info"
      type       = "struct<owner_id:string,region:string,transit_gateway_id:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "status"
      type       = "struct<code:string,message:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "transit_gateway_attachment_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "creation_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_vpc_endpoints" {
  name          = "aws_ec2_vpc_endpoints"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/vpc_endpoints/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "dns_entries"
      type       = "array<struct<dns_name:string,hosted_zone_id:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "groups"
      type       = "array<struct<group_id:string,group_name:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "last_error"
      type       = "struct<code:string,message:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "network_interface_ids"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "owner_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "policy_document"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "private_dns_enabled"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "requester_managed"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "route_table_ids"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "service_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "subnet_ids"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_endpoint_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_endpoint_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "creation_timestamp"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_vpc_peering_connections" {
  name          = "aws_ec2_vpc_peering_connections"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/vpc_peering_connections/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "accepter_vpc_info"
      type       = "struct<cidr_block:string,cidr_block_set:array<struct<cidr_block:string>>,ipv6_cidr_block_set:array<struct<ipv6_cidr_block:string>>,owner_id:string,peering_options:struct<allow_dns_resolution_from_remote_vpc:boolean,allow_egress_from_local_classic_link_to_remote_vpc:boolean,allow_egress_from_local_vpc_to_remote_classic_link:boolean>,region:string,vpc_id:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "requester_vpc_info"
      type       = "struct<cidr_block:string,cidr_block_set:array<struct<cidr_block:string>>,ipv6_cidr_block_set:array<struct<ipv6_cidr_block:string>>,owner_id:string,peering_options:struct<allow_dns_resolution_from_remote_vpc:boolean,allow_egress_from_local_classic_link_to_remote_vpc:boolean,allow_egress_from_local_vpc_to_remote_classic_link:boolean>,region:string,vpc_id:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "status"
      type       = "struct<code:string,message:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_peering_connection_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "expiration_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_vpn_gateways" {
  name          = "aws_ec2_vpn_gateways"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/vpn_gateways/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "amazon_side_asn"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "availability_zone"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_attachments"
      type       = "array<struct<state:string,vpc_id:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpn_gateway_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_reserved_instances" {
  name          = "aws_ec2_reserved_instances"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/reserved_instances/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "availability_zone"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "currency_code"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "duration"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "fixed_price"
      type       = "float"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "instance_count"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "instance_tenancy"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "instance_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "offering_class"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "offering_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "product_description"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "recurring_charges"
      type       = "array<struct<amount:double,frequency:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "reserved_instances_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "scope"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "usage_price"
      type       = "float"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "end"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "start"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_placement_groups" {
  name          = "aws_ec2_placement_groups"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/placement_groups/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "group_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "group_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "partition_count"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "strategy"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_addresses" {
  name          = "aws_ec2_addresses"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/addresses/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "allocation_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "association_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "carrier_ip"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "customer_owned_ip"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "customer_owned_ipv4_pool"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "domain"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "instance_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "network_border_group"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "network_interface_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "network_interface_owner_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "private_ip_address"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "public_ip"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "public_ipv4_pool"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags_old"
      type       = "array<struct<key:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ec2_instance_types" {
  name          = "aws_ec2_instance_types"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ec2/instance_types/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "auto_recovery_supported"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "bare_metal"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "burstable_performance_supported"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "current_generation"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "dedicated_hosts_supported"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ebs_info"
      type       = "struct<ebs_optimized_info:struct<baseline_bandwidth_in_mbps:int,baseline_iops:int,baseline_throughput_in_m_bps:double,maximum_bandwidth_in_mbps:int,maximum_iops:int,maximum_throughput_in_m_bps:double>,ebs_optimized_support:string,encryption_support:string,nvme_support:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "fpga_info"
      type       = "struct<fpgas:array<struct<count:int,manufacturer:string,memory_info:struct<size_in_mi_b:int>,name:string>>,total_fpga_memory_in_mi_b:int>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "free_tier_eligible"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "gpu_info"
      type       = "struct<gpus:array<struct<count:int,manufacturer:string,memory_info:struct<size_in_mi_b:int>,name:string>>,total_gpu_memory_in_mi_b:int>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "hibernation_supported"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "hypervisor"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "inference_accelerator_info"
      type       = "struct<accelerators:array<struct<count:int,manufacturer:string,name:string>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "instance_storage_info"
      type       = "struct<disks:array<struct<count:int,size_in_gb:bigint,type:string>>,encryption_support:string,nvme_support:string,total_size_in_gb:bigint>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "instance_storage_supported"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "instance_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "memory_info"
      type       = "struct<size_in_mi_b:bigint>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "network_info"
      type       = "struct<default_network_card_index:int,efa_info:struct<maximum_efa_interfaces:int>,efa_supported:boolean,ena_support:string,encryption_in_transit_supported:boolean,ipv4_addresses_per_interface:int,ipv6_addresses_per_interface:int,ipv6_supported:boolean,maximum_network_cards:int,maximum_network_interfaces:int,network_cards:array<struct<maximum_network_interfaces:int,network_card_index:int,network_performance:string>>,network_performance:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "placement_group_info"
      type       = "struct<supported_strategies:array<string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "processor_info"
      type       = "struct<supported_architectures:array<string>,sustained_clock_speed_in_ghz:double>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "supported_boot_modes"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "supported_root_device_types"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "supported_usage_classes"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "supported_virtualization_types"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "v_cpu_info"
      type       = "struct<default_cores:int,default_threads_per_core:int,default_v_cpus:int,valid_cores:array<int>,valid_threads_per_core:array<int>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
