
resource "aws_glue_catalog_table" "aws_sqs_queues" {
  name          = "aws_sqs_queues"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/sqs/queues/"
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
    columns {
      name       = "queue_url"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "policy"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "visibility_timeout"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "maximum_message_size"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "message_retention_period"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "approximate_number_of_messages"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "approximate_number_of_messages_not_visible"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "created_timestamp"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "last_modified_timestamp"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "queue_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "approximate_number_of_messages_delayed"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "delay_seconds"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "receive_message_wait_time_seconds"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "redrive_policy"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "fifo_queue"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "content_based_deduplication"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "kms_master_key_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "kms_data_key_reuse_period_seconds"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "deduplication_scope"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "fifo_throughput_limit"
      type       = "string"
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
