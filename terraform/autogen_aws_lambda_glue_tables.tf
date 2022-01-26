
resource "aws_glue_catalog_table" "aws_lambda_functions" {
  name          = "aws_lambda_functions"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/lambda/functions/"
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
      name       = "architectures"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "code_sha256"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "code_size"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "dead_letter_config"
      type       = "struct<target_arn:string>"
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
      name       = "file_system_configs"
      type       = "array<struct<arn:string,local_mount_path:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "function_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "function_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "handler"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "image_config_response"
      type       = "struct<error:struct<error_code:string,message:string>,image_config:struct<command:array<string>,entry_point:array<string>,working_directory:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "kms_key_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "last_modified"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "last_update_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "last_update_status_reason"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "last_update_status_reason_code"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "layers"
      type       = "array<struct<arn:string,code_size:bigint,signing_job_arn:string,signing_profile_version_arn:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "master_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "memory_size"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "package_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "revision_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "role"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "runtime"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "signing_job_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "signing_profile_version_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "state"
      type       = "string"
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
      name       = "state_reason_code"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "timeout"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tracing_config"
      type       = "struct<mode:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "version"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_config"
      type       = "struct<security_group_ids:array<string>,subnet_ids:array<string>,vpc_id:string>"
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
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
