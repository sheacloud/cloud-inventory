
resource "aws_glue_catalog_table" "aws_sns_topics" {
  name          = "aws_sns_topics"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/sns/topics/"
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
      name       = "topic_arn"
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
    columns {
      name       = "delivery_policy"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "display_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "owner"
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
      name       = "subscriptions_confirmed"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "subscriptions_deleted"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "subscriptions_pending"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "effective_delivery_policy"
      type       = "string"
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
      name       = "fifo_topic"
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
resource "aws_glue_catalog_table" "aws_sns_subscriptions" {
  name          = "aws_sns_subscriptions"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/sns/subscriptions/"
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
      name       = "endpoint"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "owner"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "protocol"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "subscription_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "topic_arn"
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
    columns {
      name       = "confirmation_was_authenticated"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "delivery_policy"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "effective_delivery_policy"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "filter_policy"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "pending_confirmation"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "raw_message_delivery"
      type       = "boolean"
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
      name       = "subscription_role_arn"
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
