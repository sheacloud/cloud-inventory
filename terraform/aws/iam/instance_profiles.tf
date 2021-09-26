
resource "aws_glue_catalog_table" "iam_instance_profiles" {
  name          = "iam_instance_profiles"
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
    location      = "s3://${var.bucket_name}/parquet/aws/iam/instance_profiles/"
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
      name    = "instance_profile_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "instance_profile_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "path"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "roles"
      type    = "array<struct<arn:string,path:string,role_id:string,role_name:string,assume_role_policy_document:string,description:string,max_session_duration:int,permissions_boundary:struct<permissions_boundary_arn:string,permissions_boundary_type:string>,role_last_used:struct<region:string,last_used_date_milli:timestamp>,create_date_milli:timestamp,tags:map<string,string>>>"
      comment = ""
    }
    columns {
      name    = "create_date_milli"
      type    = "timestamp"
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
