
resource "aws_glue_catalog_table" "rds_db_instances" {
  name          = "rds_db_instances"
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
    location      = "s3://${var.bucket_name}/parquet/aws/rds/db_instances/"
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
      name    = "activity_stream_engine_native_audit_fields_included"
      type    = "boolean"
      comment = ""
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
      name    = "auto_minor_version_upgrade"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "automatic_restart_time"
      type    = "timestamp"
      comment = ""
    }
    columns {
      name    = "availability_zone"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "aws_backup_recovery_point_arn"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "backup_retention_period"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "ca_certificate_identifier"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "character_set_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "copy_tags_to_snapshot"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "customer_owned_ip_enabled"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "db_cluster_identifier"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "db_instance_arn"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "db_instance_automated_backups_replications"
      type    = "array<struct<db_instance_automated_backups_arn:string>>"
      comment = ""
    }
    columns {
      name    = "db_instance_class"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "db_instance_identifier"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "db_instance_status"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "db_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "db_parameter_groups"
      type    = "array<struct<db_parameter_group_name:string,parameter_apply_status:string>>"
      comment = ""
    }
    columns {
      name    = "db_security_groups"
      type    = "array<struct<db_security_group_name:string,status:string>>"
      comment = ""
    }
    columns {
      name    = "db_subnet_group"
      type    = "struct<db_subnet_group_arn:string,db_subnet_group_description:string,db_subnet_group_name:string,subnet_group_status:string,subnets:array<struct<subnet_availability_zone:struct<name:string>,subnet_identifier:string,subnet_outpost:struct<arn:string>,subnet_status:string>>,vpc_id:string>"
      comment = ""
    }
    columns {
      name    = "db_instance_port"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "dbi_resource_id"
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
      type    = "struct<address:string,hosted_zone_id:string,port:int>"
      comment = ""
    }
    columns {
      name    = "engine"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "engine_version"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "enhanced_monitoring_resource_arn"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "iam_database_authentication_enabled"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "instance_create_time"
      type    = "timestamp"
      comment = ""
    }
    columns {
      name    = "iops"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "kms_key_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "latest_restorable_time"
      type    = "timestamp"
      comment = ""
    }
    columns {
      name    = "license_model"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "listener_endpoint"
      type    = "struct<address:string,hosted_zone_id:string,port:int>"
      comment = ""
    }
    columns {
      name    = "master_username"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "max_allocated_storage"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "monitoring_interval"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "monitoring_role_arn"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "multi_az"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "nchar_character_set_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "option_group_memberships"
      type    = "array<struct<option_group_name:string,status:string>>"
      comment = ""
    }
    columns {
      name    = "pending_modified_values"
      type    = "struct<allocated_storage:int,backup_retention_period:int,ca_certificate_identifier:string,db_instance_class:string,db_instance_identifier:string,db_subnet_group_name:string,engine_version:string,iam_database_authentication_enabled:boolean,iops:int,license_model:string,master_user_password:string,multi_az:boolean,pending_cloudwatch_logs_exports:struct<log_types_to_disable:array<string>,log_types_to_enable:array<string>>,port:int,processor_features:array<struct<name:string,value:string>>,storage_type:string>"
      comment = ""
    }
    columns {
      name    = "performance_insights_enabled"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "performance_insights_kms_key_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "performance_insights_retention_period"
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
      name    = "processor_features"
      type    = "array<struct<name:string,value:string>>"
      comment = ""
    }
    columns {
      name    = "promotion_tier"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "publicly_accessible"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "read_replica_db_cluster_identifiers"
      type    = "array<string>"
      comment = ""
    }
    columns {
      name    = "read_replica_db_instance_identifiers"
      type    = "array<string>"
      comment = ""
    }
    columns {
      name    = "read_replica_source_db_instance_identifier"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "replica_mode"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "secondary_availability_zone"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "status_infos"
      type    = "array<struct<message:string,normal:boolean,status:string,status_type:string>>"
      comment = ""
    }
    columns {
      name    = "storage_encrypted"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "storage_type"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "tags"
      type    = "map<string,string>"
      comment = ""
    }
    columns {
      name    = "tde_credential_arn"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "timezone"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "vpc_security_groups"
      type    = "array<struct<status:string,vpc_security_group_id:string>>"
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
