
resource "aws_glue_catalog_table" "aws_iam_groups" {
  name          = "aws_iam_groups"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/iam/groups/"
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
      name       = "arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "group_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "group_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "path"
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
      name       = "inline_policies"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "user_ids"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "attached_policies"
      type       = "array<struct<policy_arn:string,policy_name:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "create_date"
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
resource "aws_glue_catalog_table" "aws_iam_policies" {
  name          = "aws_iam_policies"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/iam/policies/"
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
      name       = "arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "attachment_count"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "default_version_id"
      type       = "string"
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
      name       = "is_attachable"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "path"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "permissions_boundary_usage_count"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "policy_id"
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
      name       = "tags"
      type       = "map<string,string>"
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
      name       = "create_date"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "update_date"
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
resource "aws_glue_catalog_table" "aws_iam_roles" {
  name          = "aws_iam_roles"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/iam/roles/"
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
      name       = "arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "path"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "role_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "role_name"
      type       = "string"
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
      name       = "max_session_duration"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "permissions_boundary"
      type       = "struct<permissions_boundary_arn:string,permissions_boundary_type:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "role_last_used"
      type       = "struct<region:string,last_used_date:timestamp>"
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
      name       = "inline_policies"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "attached_policies"
      type       = "array<struct<policy_arn:string,policy_name:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "create_date"
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
resource "aws_glue_catalog_table" "aws_iam_users" {
  name          = "aws_iam_users"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/iam/users/"
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
      name       = "arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "path"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "user_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "user_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "permissions_boundary"
      type       = "struct<permissions_boundary_arn:string,permissions_boundary_type:string>"
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
      name       = "inline_policies"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "group_ids"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "access_keys"
      type       = "array<struct<access_key_id:string,status:string,create_date:timestamp>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "login_profile"
      type       = "struct<password_reset_required:boolean,create_date:timestamp>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "attached_policies"
      type       = "array<struct<policy_arn:string,policy_name:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "create_date"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "password_last_used"
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
