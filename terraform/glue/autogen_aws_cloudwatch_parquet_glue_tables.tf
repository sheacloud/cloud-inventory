
resource "aws_glue_catalog_table" "aws_cloudwatch_metric_alarms" {
  name          = "aws_cloudwatch_metric_alarms"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/cloudwatch/metric_alarms/"
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
      name       = "actions_enabled"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "alarm_actions"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "alarm_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "alarm_configuration_updated_timestamp"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "alarm_description"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "alarm_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "comparison_operator"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "datapoints_to_alarm"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "dimensions"
      type       = "array<struct<name:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "evaluate_low_sample_count_percentile"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "evaluation_periods"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "extended_statistic"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "insufficient_data_actions"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "metric_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "metrics"
      type       = "array<struct<id:string,account_id:string,expression:string,label:string,metric_stat:struct<metric:struct<dimensions:array<struct<name:string,value:string>>,metric_name:string,namespace:string>,period:int,stat:string,unit:string>,period:int,return_data:boolean>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "namespace"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ok_actions"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "period"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state_reason"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state_reason_data"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state_updated_timestamp"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state_value"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "statistic"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "threshold"
      type       = "double"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "threshold_metric_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "treat_missing_data"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "unit"
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
resource "aws_glue_catalog_table" "aws_cloudwatch_composite_alarms" {
  name          = "aws_cloudwatch_composite_alarms"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/cloudwatch/composite_alarms/"
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
      name       = "actions_enabled"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "alarm_actions"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "alarm_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "alarm_configuration_updated_timestamp"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "alarm_description"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "alarm_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "alarm_rule"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "insufficient_data_actions"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ok_actions"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state_reason"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state_reason_data"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state_updated_timestamp"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state_value"
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
