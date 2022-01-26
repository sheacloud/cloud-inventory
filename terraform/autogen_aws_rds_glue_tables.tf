
resource "aws_glue_catalog_table" "aws_rds_db_clusters" {
  name          = "aws_rds_db_clusters"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/rds/db_clusters/"
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
      name       = "activity_stream_kinesis_stream_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "activity_stream_kms_key_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "activity_stream_mode"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "activity_stream_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "allocated_storage"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "associated_roles"
      type       = "array<struct<feature_name:string,role_arn:string,status:string>>"
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
      name       = "availability_zones"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "backtrack_consumed_change_records"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "backtrack_window"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "backup_retention_period"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "capacity"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "character_set_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "clone_group_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "copy_tags_to_snapshot"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cross_account_clone"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "custom_endpoints"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_cluster_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_cluster_identifier"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_cluster_instance_class"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_cluster_members"
      type       = "array<struct<db_cluster_parameter_group_status:string,db_instance_identifier:string,is_cluster_writer:boolean,promotion_tier:int>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_cluster_option_group_memberships"
      type       = "array<struct<db_cluster_option_group_name:string,status:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_cluster_parameter_group"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_subnet_group"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "database_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_cluster_resource_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "deletion_protection"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "domain_memberships"
      type       = "array<struct<domain:string,fqdn:string,iam_role_name:string,status:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "enabled_cloudwatch_logs_exports"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "endpoint"
      type       = "string"
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
      name       = "engine_mode"
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
      name       = "global_write_forwarding_requested"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "global_write_forwarding_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "hosted_zone_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "http_endpoint_enabled"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "iam_database_authentication_enabled"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "iops"
      type       = "int"
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
      name       = "master_username"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "monitoring_interval"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "monitoring_role_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "multi_az"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "pending_modified_values"
      type       = "struct<db_cluster_identifier:string,engine_version:string,iam_database_authentication_enabled:boolean,master_user_password:string,pending_cloudwatch_logs_exports:struct<log_types_to_disable:array<string>,log_types_to_enable:array<string>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "percent_progress"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "performance_insights_enabled"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "performance_insights_kms_key_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "performance_insights_retention_period"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "port"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "preferred_backup_window"
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
      name       = "publicly_accessible"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "read_replica_identifiers"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "reader_endpoint"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "replication_source_identifier"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "scaling_configuration_info"
      type       = "struct<auto_pause:boolean,max_capacity:int,min_capacity:int,seconds_before_timeout:int,seconds_until_auto_pause:int,timeout_action:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "storage_encrypted"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "storage_type"
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
      name       = "vpc_security_groups"
      type       = "array<struct<status:string,vpc_security_group_id:string>>"
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
      name       = "automatic_restart_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cluster_create_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "earliest_backtrack_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "earliest_restorable_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "latest_restorable_time"
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
resource "aws_glue_catalog_table" "aws_rds_db_instances" {
  name          = "aws_rds_db_instances"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/rds/db_instances/"
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
      name       = "activity_stream_engine_native_audit_fields_included"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "activity_stream_kinesis_stream_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "activity_stream_kms_key_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "activity_stream_mode"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "activity_stream_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "allocated_storage"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "associated_roles"
      type       = "array<struct<feature_name:string,role_arn:string,status:string>>"
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
      name       = "automation_mode"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "availability_zone"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "aws_backup_recovery_point_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "backup_retention_period"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "backup_target"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ca_certificate_identifier"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "character_set_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "copy_tags_to_snapshot"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "custom_iam_instance_profile"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "customer_owned_ip_enabled"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_cluster_identifier"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_instance_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_instance_automated_backups_replications"
      type       = "array<struct<db_instance_automated_backups_arn:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_instance_class"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_instance_identifier"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_instance_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_parameter_groups"
      type       = "array<struct<db_parameter_group_name:string,parameter_apply_status:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_security_groups"
      type       = "array<struct<db_security_group_name:string,status:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_subnet_group"
      type       = "struct<db_subnet_group_arn:string,db_subnet_group_description:string,db_subnet_group_name:string,subnet_group_status:string,subnets:array<struct<subnet_availability_zone:struct<name:string>,subnet_identifier:string,subnet_outpost:struct<arn:string>,subnet_status:string>>,vpc_id:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "db_instance_port"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "dbi_resource_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "deletion_protection"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "domain_memberships"
      type       = "array<struct<domain:string,fqdn:string,iam_role_name:string,status:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "enabled_cloudwatch_logs_exports"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "endpoint"
      type       = "struct<address:string,hosted_zone_id:string,port:int>"
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
      name       = "enhanced_monitoring_resource_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "iam_database_authentication_enabled"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "iops"
      type       = "int"
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
      name       = "license_model"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "listener_endpoint"
      type       = "struct<address:string,hosted_zone_id:string,port:int>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "master_username"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "max_allocated_storage"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "monitoring_interval"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "monitoring_role_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "multi_az"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "nchar_character_set_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "option_group_memberships"
      type       = "array<struct<option_group_name:string,status:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "pending_modified_values"
      type       = "struct<allocated_storage:int,automation_mode:string,backup_retention_period:int,ca_certificate_identifier:string,db_instance_class:string,db_instance_identifier:string,db_subnet_group_name:string,engine_version:string,iam_database_authentication_enabled:boolean,iops:int,license_model:string,master_user_password:string,multi_az:boolean,pending_cloudwatch_logs_exports:struct<log_types_to_disable:array<string>,log_types_to_enable:array<string>>,port:int,processor_features:array<struct<name:string,value:string>>,storage_type:string,resume_full_automation_mode_time:timestamp>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "performance_insights_enabled"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "performance_insights_kms_key_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "performance_insights_retention_period"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "preferred_backup_window"
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
      name       = "processor_features"
      type       = "array<struct<name:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "promotion_tier"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "publicly_accessible"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "read_replica_db_cluster_identifiers"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "read_replica_db_instance_identifiers"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "read_replica_source_db_instance_identifier"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "replica_mode"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "secondary_availability_zone"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "status_infos"
      type       = "array<struct<message:string,normal:boolean,status:string,status_type:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "storage_encrypted"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "storage_type"
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
      name       = "tde_credential_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "timezone"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_security_groups"
      type       = "array<struct<status:string,vpc_security_group_id:string>>"
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
      name       = "automatic_restart_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "instance_create_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "latest_restorable_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "resume_full_automation_mode_time"
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
