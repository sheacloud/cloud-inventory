
resource "aws_glue_catalog_table" "aws_storagegateway_gateways" {
  name          = "aws_storagegateway_gateways"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/storagegateway/gateways/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "ion"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "cloud_watch_log_group_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "deprecation_date"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ec2_instance_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ec2_instance_region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "endpoint_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "gateway_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "gateway_capacity"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "gateway_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "gateway_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "gateway_network_interfaces"
      type       = "array<struct<ipv4_address:string,ipv6_address:string,mac_address:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "gateway_state"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "gateway_timezone"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "gateway_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "host_environment"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "host_environment_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "last_software_update"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "next_update_availability_date"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "software_updates_end_date"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "supported_gateway_capacities"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_endpoint"
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
      name       = "inventory_uuid"
      type       = "string"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
