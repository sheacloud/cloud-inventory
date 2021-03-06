
resource "aws_glue_catalog_table" "aws_s3_buckets" {
  name          = "aws_s3_buckets"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/s3/buckets/"
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
      name       = "creation_date"
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
      name       = "policy"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "is_public"
      type       = "boolean"
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
      name       = "versioning_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "mfa_delete_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "replication_configuration"
      type       = "struct<role:string,rules:array<struct<destination:struct<bucket:string,access_control_translation:struct<owner:string>,account:string,encryption_configuration:struct<replica_kms_key_id:string>,metrics:struct<status:string,event_threshold:struct<minutes:int>>,replication_time:struct<status:string,time:struct<minutes:int>>,storage_class:string>,status:string,delete_marker_replication:struct<status:string>,existing_object_replication:struct<status:string>,id:string,prefix:string,priority:int,source_selection_criteria:struct<replica_modifications:struct<status:string>,sse_kms_encrypted_objects:struct<status:string>>>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "acl_grants"
      type       = "array<struct<grantee:struct<type:string,display_name:string,email_address:string,id:string,uri:string>,permission:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cors_rules"
      type       = "array<struct<allowed_methods:array<string>,allowed_origins:array<string>,allowed_headers:array<string>,expose_headers:array<string>,id:string,max_age_seconds:int>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "server_side_encryption_configuration"
      type       = "struct<rules:array<struct<apply_server_side_encryption_by_default:struct<sse_algorithm:string,kms_master_key_id:string>,bucket_key_enabled:boolean>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "intelligent_tiering_configurations"
      type       = "array<struct<id:string,status:string,tierings:array<struct<access_tier:string,days:int>>,filter:struct<and:struct<prefix:string,tags:array<struct<key:string,value:string>>>,prefix:string,tag:struct<key:string,value:string>>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "inventory_configurations"
      type       = "array<struct<destination:struct<s3_bucket_destination:struct<bucket:string,format:string,account_id:string,encryption:struct<ssekms:struct<key_id:string>>,prefix:string>>,id:string,included_object_versions:string,is_enabled:boolean,schedule:struct<frequency:string>,filter:struct<prefix:string>,optional_fields:array<string>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "lifecycle_rules"
      type       = "array<struct<status:string,abort_incomplete_multipart_upload:struct<days_after_initiation:int>,expiration:struct<date:timestamp,days:int,expired_object_delete_marker:boolean>,id:string,noncurrent_version_expiration:struct<newer_noncurrent_versions:int,noncurrent_days:int>,noncurrent_version_transitions:array<struct<newer_noncurrent_versions:int,noncurrent_days:int,storage_class:string>>,prefix:string,transitions:array<struct<date:timestamp,days:int,storage_class:string>>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "logging"
      type       = "struct<target_bucket:string,target_prefix:string,target_grants:array<struct<grantee:struct<type:string,display_name:string,email_address:string,id:string,uri:string>,permission:string>>>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
