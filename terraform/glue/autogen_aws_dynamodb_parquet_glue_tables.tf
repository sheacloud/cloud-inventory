
resource "aws_glue_catalog_table" "aws_dynamodb_tables" {
  name          = "aws_dynamodb_tables"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/dynamodb/tables/"
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
      name       = "archival_summary"
      type       = "struct<archival_backup_arn:string,archival_date_time:timestamp,archival_reason:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "attribute_definitions"
      type       = "array<struct<attribute_name:string,attribute_type:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "billing_mode_summary"
      type       = "struct<billing_mode:string,last_update_to_pay_per_request_date_time:timestamp>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "creation_date_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "global_secondary_indexes"
      type       = "array<struct<backfilling:boolean,index_arn:string,index_name:string,index_size_bytes:bigint,index_status:string,item_count:bigint,key_schema:array<struct<attribute_name:string,key_type:string>>,projection:struct<non_key_attributes:array<string>,projection_type:string>,provisioned_throughput:struct<last_decrease_date_time:timestamp,last_increase_date_time:timestamp,number_of_decreases_today:bigint,read_capacity_units:bigint,write_capacity_units:bigint>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "global_table_version"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "item_count"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "key_schema"
      type       = "array<struct<attribute_name:string,key_type:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "latest_stream_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "latest_stream_label"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "local_secondary_indexes"
      type       = "array<struct<index_arn:string,index_name:string,index_size_bytes:bigint,item_count:bigint,key_schema:array<struct<attribute_name:string,key_type:string>>,projection:struct<non_key_attributes:array<string>,projection_type:string>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "provisioned_throughput"
      type       = "struct<last_decrease_date_time:timestamp,last_increase_date_time:timestamp,number_of_decreases_today:bigint,read_capacity_units:bigint,write_capacity_units:bigint>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "replicas"
      type       = "array<struct<global_secondary_indexes:array<struct<index_name:string,provisioned_throughput_override:struct<read_capacity_units:bigint>>>,kms_master_key_id:string,provisioned_throughput_override:struct<read_capacity_units:bigint>,region_name:string,replica_inaccessible_date_time:timestamp,replica_status:string,replica_status_description:string,replica_status_percent_progress:string,replica_table_class_summary:struct<last_update_date_time:timestamp,table_class:string>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "restore_summary"
      type       = "struct<restore_date_time:timestamp,restore_in_progress:boolean,source_backup_arn:string,source_table_arn:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "sse_description"
      type       = "struct<inaccessible_encryption_date_time:timestamp,kms_master_key_arn:string,sse_type:string,status:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "stream_specification"
      type       = "struct<stream_enabled:boolean,stream_view_type:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "table_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "table_class_summary"
      type       = "struct<last_update_date_time:timestamp,table_class:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "table_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "table_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "table_size_bytes"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "table_status"
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
