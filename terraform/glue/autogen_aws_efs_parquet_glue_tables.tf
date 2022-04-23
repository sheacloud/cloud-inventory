
resource "aws_glue_catalog_table" "aws_efs_filesystems" {
  name          = "aws_efs_filesystems"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/efs/filesystems/"
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
      name       = "creation_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "creation_token"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "file_system_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "life_cycle_state"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "number_of_mount_targets"
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
      name       = "performance_mode"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "size_in_bytes"
      type       = "struct<value:bigint,timestamp:timestamp,value_in_ia:bigint,value_in_standard:bigint>"
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
      name       = "availability_zone_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "availability_zone_name"
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
      name       = "file_system_arn"
      type       = "string"
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
      name       = "name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "provisioned_throughput_in_mibps"
      type       = "double"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "throughput_mode"
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
