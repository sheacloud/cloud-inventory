
resource "aws_glue_catalog_table" "aws_ecr_repositories" {
  name          = "aws_ecr_repositories"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ecr/repositories/"
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
      name       = "created_at"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "encryption_configuration"
      type       = "struct<encryption_type:string,kms_key:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "image_scanning_configuration"
      type       = "struct<scan_on_push:boolean>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "image_tag_mutability"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "registry_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "repository_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "repository_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "repository_uri"
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
