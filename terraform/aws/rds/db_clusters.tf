
resource "aws_glue_catalog_table" "rds_db_clusters" {
  name          = "rds_db_clusters"
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
    location      = "s3://${var.bucket_name}/parquet/aws/rds/db_clusters/"
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
      name    = "activity_stream_kinesis_stream_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "activity_stream_kms_key_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "activity_stream_mode"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "activity_stream_status"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "allocated_storage"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "associated_roles"
      type    = "array<struct<feature_name:string,role_arn:string,status:string>>"
      comment = ""
    }
    columns {
      name    = "availability_zones"
      type    = "array<string>"
      comment = ""
    }
    columns {
      name    = "backtrack_consumed_change_records"
      type    = "bigint"
      comment = ""
    }
    columns {
      name    = "backtrack_window"
      type    = "bigint"
      comment = ""
    }
    columns {
      name    = "backup_retention_period"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "capacity"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "character_set_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "clone_group_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "copy_tags_to_snapshot"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "cross_account_clone"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "custom_endpoints"
      type    = "array<string>"
      comment = ""
    }
    columns {
      name    = "db_cluster_arn"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "db_cluster_identifier"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "db_cluster_members"
      type    = "array<struct<db_cluster_parameter_group_status:string,db_instance_identifier:string,is_cluster_writer:boolean,promotion_tier:int>>"
      comment = ""
    }
    columns {
      name    = "db_cluster_option_group_memberships"
      type    = "array<struct<db_cluster_option_group_name:string,status:string>>"
      comment = ""
    }
    columns {
      name    = "db_cluster_parameter_group"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "db_subnet_group"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "database_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "db_cluster_resource_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "deletion_protection"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "domain_memberships"
      type    = "array<struct<domain:string,fqdn:string,iam_role_name:string,status:string>>"
      comment = ""
    }
    columns {
      name    = "enabled_cloudwatch_logs_exports"
      type    = "array<string>"
      comment = ""
    }
    columns {
      name    = "endpoint"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "engine"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "engine_mode"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "engine_version"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "global_write_forwarding_requested"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "global_write_forwarding_status"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "hosted_zone_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "http_endpoint_enabled"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "iam_database_authentication_enabled"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "kms_key_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "master_username"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "multi_az"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "pending_modified_values"
      type    = "struct<db_cluster_identifier:string,engine_version:string,iam_database_authentication_enabled:boolean,master_user_password:string,pending_cloudwatch_logs_exports:struct<log_types_to_disable:array<string>,log_types_to_enable:array<string>>>"
      comment = ""
    }
    columns {
      name    = "percent_progress"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "port"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "preferred_backup_window"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "preferred_maintenance_window"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "read_replica_identifiers"
      type    = "array<string>"
      comment = ""
    }
    columns {
      name    = "reader_endpoint"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "replication_source_identifier"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "scaling_configuration_info"
      type    = "struct<auto_pause:boolean,max_capacity:int,min_capacity:int,seconds_until_auto_pause:int,timeout_action:string>"
      comment = ""
    }
    columns {
      name    = "status"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "storage_encrypted"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "vpc_security_groups"
      type    = "array<struct<status:string,vpc_security_group_id:string>>"
      comment = ""
    }
    columns {
      name    = "automatic_restart_time_milli"
      type    = "timestamp"
      comment = ""
    }
    columns {
      name    = "cluster_create_time_milli"
      type    = "timestamp"
      comment = ""
    }
    columns {
      name    = "earliest_backtrack_time_milli"
      type    = "timestamp"
      comment = ""
    }
    columns {
      name    = "earliest_restorable_time_milli"
      type    = "timestamp"
      comment = ""
    }
    columns {
      name    = "latest_restorable_time_milli"
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
