
resource "aws_glue_catalog_table" "aws_cloudformation_stacks" {
  name          = "aws_cloudformation_stacks"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/cloudformation/stacks/"
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
      name       = "stack_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "stack_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "capabilities"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "change_set_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "deletion_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "description"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "disable_rollback"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "drift_information"
      type       = "struct<stack_drift_status:string,last_check_timestamp:timestamp>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "enable_termination_protection"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "last_updated_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "notification_arns"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "outputs"
      type       = "array<struct<description:string,export_name:string,output_key:string,output_value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "parameters"
      type       = "array<struct<parameter_key:string,parameter_value:string,resolved_value:string,use_previous_value:boolean>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "parent_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "role_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "rollback_configuration"
      type       = "struct<monitoring_time_in_minutes:int,rollback_triggers:array<struct<arn:string,type:string>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "root_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "stack_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "stack_status_reason"
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
    columns {
      name       = "timeout_in_minutes"
      type       = "int"
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
