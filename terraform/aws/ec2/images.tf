
resource "aws_glue_catalog_table" "ec2_images" {
  name          = "ec2_images"
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
    location      = "s3://${var.bucket_name}/parquet/aws/ec2/images/"
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
      name    = "block_device_mappings"
      type    = "array<struct<device_name:string,ebs:struct<delete_on_termination:boolean,encrypted:boolean,iops:int,kms_key_id:string,outpost_arn:string,snapshot_id:string,throughput:int,volume_size:int,volume_type:string>,no_device:string,virtual_name:string>>"
      comment = ""
    }
    columns {
      name    = "boot_mode"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "creation_date"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "deprecation_time"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "description"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "ena_support"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "hypervisor"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "image_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "image_location"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "image_owner_alias"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "image_type"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "kernel_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "owner_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "platform"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "platform_details"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "product_codes"
      type    = "array<struct<product_code_id:string,product_code_type:string>>"
      comment = ""
    }
    columns {
      name    = "public"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "ramdisk_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "root_device_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "root_device_type"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "sriov_net_support"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "state"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "state_reason"
      type    = "struct<code:string,message:string>"
      comment = ""
    }
    columns {
      name    = "tags"
      type    = "map<string,string>"
      comment = ""
    }
    columns {
      name    = "usage_operation"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "virtualization_type"
      type    = "string"
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
