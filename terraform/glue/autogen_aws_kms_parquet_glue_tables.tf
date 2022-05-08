
resource "aws_glue_catalog_table" "aws_kms_keys" {
  name          = "aws_kms_keys"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/kms/keys/"
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
      name       = "key_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "aws_account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cloud_hsm_cluster_id"
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
      name       = "custom_key_store_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "customer_master_key_spec"
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
      name       = "description"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "enabled"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "encryption_algorithms"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "expiration_model"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "key_manager"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "key_spec"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "key_state"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "key_usage"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "mac_algorithms"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "multi_region"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "multi_region_configuration"
      type       = "struct<multi_region_key_type:string,primary_key:struct<arn:string,region:string>,replica_keys:array<struct<arn:string,region:string>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "origin"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "pending_deletion_window_in_days"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "signing_algorithms"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "valid_to"
      type       = "timestamp"
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
      name       = "aliases"
      type       = "array<struct<alias_arn:string,alias_name:string,creation_date:timestamp,last_updated_date:timestamp,target_key_id:string>>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
