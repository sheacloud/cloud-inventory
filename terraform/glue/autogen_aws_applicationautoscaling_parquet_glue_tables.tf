
resource "aws_glue_catalog_table" "aws_applicationautoscaling_scaling_policies" {
  name          = "aws_applicationautoscaling_scaling_policies"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/applicationautoscaling/scaling_policies/"
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
      name       = "policy_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "policy_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "policy_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "resource_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "scalable_dimension"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "service_namespace"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "alarms"
      type       = "array<struct<alarm_arn:string,alarm_name:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "step_scaling_policy_configuration"
      type       = "struct<adjustment_type:string,cooldown:int,metric_aggregation_type:string,min_adjustment_magnitude:int,step_adjustments:array<struct<scaling_adjustment:int,metric_interval_lower_bound:double,metric_interval_upper_bound:double>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "target_tracking_scaling_policy_configuration"
      type       = "struct<target_value:double,customized_metric_specification:struct<metric_name:string,namespace:string,statistic:string,dimensions:array<struct<name:string,value:string>>,unit:string>,disable_scale_in:boolean,predefined_metric_specification:struct<predefined_metric_type:string,resource_label:string>,scale_in_cooldown:int,scale_out_cooldown:int>"
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
