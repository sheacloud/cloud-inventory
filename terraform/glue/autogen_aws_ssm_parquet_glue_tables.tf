
resource "aws_glue_catalog_table" "aws_ssm_parameters" {
  name          = "aws_ssm_parameters"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ssm/parameters/"
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
      name       = "allowed_pattern"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "data_type"
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
      name       = "key_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "last_modified_date"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "last_modified_user"
      type       = "string"
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
      name       = "policies"
      type       = "array<struct<policy_status:string,policy_text:string,policy_type:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tier"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "version"
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
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
