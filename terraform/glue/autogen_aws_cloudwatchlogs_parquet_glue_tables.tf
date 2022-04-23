
resource "aws_glue_catalog_table" "aws_cloudwatchlogs_log_groups" {
  name          = "aws_cloudwatchlogs_log_groups"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/cloudwatchlogs/log_groups/"
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
      name       = "arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "creation_time"
      type       = "bigint"
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
      name       = "log_group_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "metric_filter_count"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "retention_in_days"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "stored_bytes"
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
