
resource "aws_glue_catalog_table" "iam_users" {
  name          = "iam_users"
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
    location      = "s3://${var.bucket_name}/parquet/aws/iam/users/"
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
      name    = "create_date"
      type    = "timestamp"
      comment = ""
    }
    columns {
      name    = "path"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "user_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "user_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "password_last_used"
      type    = "timestamp"
      comment = ""
    }
    columns {
      name    = "permissions_boundary"
      type    = "struct<permissions_boundary_arn:string,permissions_boundary_type:string>"
      comment = ""
    }
    columns {
      name    = "tags"
      type    = "map<string,string>"
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
    columns {
      name    = "access_keys"
      type    = "array<struct<access_key_id:string,create_date:timestamp,status:string>>"
      comment = ""
    }
    columns {
      name    = "login_profile"
      type    = "struct<create_date:timestamp,password_reset_required:boolean>"
      comment = ""
    }
    columns {
      name    = "attached_policies"
      type    = "array<struct<policy_arn:string,policy_name:string>>"
      comment = ""
    }
    columns {
      name    = "inline_policies"
      type    = "array<string>"
      comment = ""
    }
    columns {
      name    = "group_ids"
      type    = "array<string>"
      comment = ""
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
