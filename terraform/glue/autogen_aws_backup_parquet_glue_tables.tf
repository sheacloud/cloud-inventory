
resource "aws_glue_catalog_table" "aws_backup_backup_vaults" {
  name          = "aws_backup_backup_vaults"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/backup/backup_vaults/"
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
      name       = "backup_vault_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "backup_vault_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "creation_date"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "creator_request_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "encryption_key_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "lock_date"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "locked"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "max_retention_days"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "min_retention_days"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "number_of_recovery_points"
      type       = "bigint"
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
resource "aws_glue_catalog_table" "aws_backup_backup_plans" {
  name          = "aws_backup_backup_plans"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/backup/backup_plans/"
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
      name       = "advanced_backup_settings"
      type       = "array<struct<backup_options:map<string,string>,resource_type:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "backup_plan_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "backup_plan_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "backup_plan_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "creation_date"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "creator_request_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "deletion_date"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "last_execution_date"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "version_id"
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
    columns {
      name       = "selections"
      type       = "array<struct<backup_plan_id:string,creation_date:timestamp,creator_request_id:string,iam_role_arn:string,selection_id:string,selection_name:string>>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
