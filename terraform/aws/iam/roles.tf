
resource "aws_glue_catalog_table" "iam_roles" {
  name          = "iam_roles"
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
    location      = "s3://${var.bucket_name}/parquet/aws/iam/roles/"
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
      name    = "role_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "role_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "assume_role_policy_document"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "description"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "max_session_duration"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "permissions_boundary"
      type    = "struct<permissions_boundary_arn:string,permissions_boundary_type:string>"
      comment = ""
    }
    columns {
      name    = "role_last_used"
      type    = "struct<last_used_date:timestamp,region:string>"
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
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
