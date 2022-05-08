
resource "aws_glue_catalog_table" "aws_secretsmanager_secrets" {
  name          = "aws_secretsmanager_secrets"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/secretsmanager/secrets/"
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
      name       = "arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "created_date"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "deleted_date"
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
      name       = "kms_key_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "last_accessed_date"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "last_changed_date"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "last_rotated_date"
      type       = "timestamp"
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
      name       = "owning_service"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "primary_region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "rotation_enabled"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "rotation_lambda_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "rotation_rules"
      type       = "struct<automatically_after_days:bigint,duration:string,schedule_expression:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "secret_versions_to_stages"
      type       = "map<string,array<string>>"
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
