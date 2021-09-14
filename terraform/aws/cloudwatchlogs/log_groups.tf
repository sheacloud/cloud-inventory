
resource "aws_glue_catalog_table" "cloudwatchlogs_log_groups" {
  name          = "cloudwatchlogs_log_groups"
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
    location      = "s3://${var.bucket_name}/parquet/aws/cloudwatchlogs/log_groups/"
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
      name    = "arn"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "creation_time"
      type    = "bigint"
      comment = ""
    }
    columns {
      name    = "kms_key_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "log_group_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "metric_filter_count"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "retention_in_days"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "stored_bytes"
      type    = "bigint"
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
