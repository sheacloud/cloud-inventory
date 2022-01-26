
resource "aws_glue_catalog_table" "aws_elasticache_cache_clusters" {
  name          = "aws_elasticache_cache_clusters"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/elasticache/cache_clusters/"
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
      name       = "arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "at_rest_encryption_enabled"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "auth_token_enabled"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "auto_minor_version_upgrade"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cache_cluster_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cache_cluster_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cache_node_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cache_nodes"
      type       = "array<struct<cache_node_id:string,cache_node_status:string,customer_availability_zone:string,customer_outpost_arn:string,endpoint:struct<address:string,port:int>,parameter_group_status:string,source_cache_node_id:string,cache_node_create_time:timestamp>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cache_parameter_group"
      type       = "struct<cache_node_ids_to_reboot:array<string>,cache_parameter_group_name:string,parameter_apply_status:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cache_security_groups"
      type       = "array<struct<cache_security_group_name:string,status:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cache_subnet_group_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "client_download_landing_page"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "configuration_endpoint"
      type       = "struct<address:string,port:int>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "engine"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "engine_version"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "log_delivery_configurations"
      type       = "array<struct<destination_details:struct<cloud_watch_logs_details:struct<log_group:string>,kinesis_firehose_details:struct<delivery_stream:string>>,destination_type:string,log_format:string,log_type:string,message:string,status:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "notification_configuration"
      type       = "struct<topic_arn:string,topic_status:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "num_cache_nodes"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "pending_modified_values"
      type       = "struct<auth_token_status:string,cache_node_ids_to_remove:array<string>,cache_node_type:string,engine_version:string,log_delivery_configurations:array<struct<destination_details:struct<cloud_watch_logs_details:struct<log_group:string>,kinesis_firehose_details:struct<delivery_stream:string>>,destination_type:string,log_format:string,log_type:string>>,num_cache_nodes:int>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "preferred_availability_zone"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "preferred_maintenance_window"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "preferred_outpost_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "replication_group_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "replication_group_log_delivery_enabled"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "security_groups"
      type       = "array<struct<security_group_id:string,status:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "snapshot_retention_limit"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "snapshot_window"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "transit_encryption_enabled"
      type       = "boolean"
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
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "auth_token_last_modified_date"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cache_cluster_create_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
