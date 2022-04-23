
resource "aws_glue_catalog_table" "aws_cloudtrail_trails" {
  name          = "aws_cloudtrail_trails"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/cloudtrail/trails/"
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
      name       = "cloud_watch_logs_log_group_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cloud_watch_logs_role_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "has_custom_event_selectors"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "has_insight_selectors"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "home_region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "include_global_service_events"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "is_multi_region_trail"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "is_organization_trail"
      type       = "boolean"
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
      name       = "log_file_validation_enabled"
      type       = "boolean"
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
      name       = "s3_bucket_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "s3_key_prefix"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "sns_topic_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "sns_topic_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "trail_arn"
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
      name       = "status"
      type       = "struct<is_logging:boolean,latest_cloud_watch_logs_delivery_error:string,latest_cloud_watch_logs_delivery_time:timestamp,latest_delivery_attempt_succeeded:string,latest_delivery_attempt_time:string,latest_delivery_error:string,latest_delivery_time:timestamp,latest_digest_delivery_error:string,latest_digest_delivery_time:timestamp,latest_notification_attempt_succeeded:string,latest_notification_attempt_time:string,latest_notification_error:string,latest_notification_time:timestamp,start_logging_time:timestamp,stop_logging_time:timestamp,time_logging_started:string,time_logging_stopped:string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
