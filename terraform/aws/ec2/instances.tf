
resource "aws_glue_catalog_table" "ec2_instances" {
  name          = "ec2_instances"
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
    location      = "s3://${var.bucket_name}/parquet/aws/ec2/instances/"
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
      name    = "architecture"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "ebs_optimized"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "ena_support"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "block_device_mappings"
      type    = "array<struct<device_name:string,ebs:struct<device_name:string,ebs:struct<delete_on_termination:boolean,volume_id:string>,no_device:string,virtual_name:string>>>"
      comment = ""
    }
    columns {
      name    = "hypervisor"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "iam_instance_profile"
      type    = "struct<arn:string,id:string>"
      comment = ""
    }
    columns {
      name    = "image_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "instance_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "instance_lifecycle"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "instance_type"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "kernel_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "key_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "launch_time"
      type    = "timestamp"
      comment = ""
    }
    columns {
      name    = "monitoring"
      type    = "struct<state:string>"
      comment = ""
    }
    columns {
      name    = "network_interfaces"
      type    = "array<struct<association:struct<carrier_ip:string,ip_owner_id:string,public_dns_name:string,public_ip:string>,attachment:struct<attachment_id:string,delete_on_termination:boolean,device_index:int,network_card_index:int,status:string>,description:string,groups:array<struct<group_id:string,group_name:string>>,interface_type:string,mac_address:string,network_interface_id:string,owner_id:string,private_dns_name:string,private_ip_address:string,private_ip_addresses:array<struct<association:struct<carrier_ip:string,ip_owner_id:string,public_dns_name:string,public_ip:string>,primary:boolean,private_dns_name:string,private_ip_address:string>>,source_dest_check:boolean,status:string>>"
      comment = ""
    }
    columns {
      name    = "private_dns_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "private_ip_address"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "security_groups"
      type    = "array<struct<group_id:string,group_name:string>>"
      comment = ""
    }
    columns {
      name    = "source_dest_check"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "spot_instance_request_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "state"
      type    = "struct<code:int,name:string>"
      comment = ""
    }
    columns {
      name    = "state_reason"
      type    = "struct<code:string,message:string>"
      comment = ""
    }
    columns {
      name    = "state_transition_reason"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "subnet_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "virtualization_type"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "vpc_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "tags"
      type    = "map<string,string>"
      comment = ""
    }
    columns {
      name    = "account_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "region"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "report_time"
      type    = "timestamp"
      comment = ""
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
