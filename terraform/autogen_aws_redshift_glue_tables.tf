
resource "aws_glue_catalog_table" "aws_redshift_clusters" {
  name          = "aws_redshift_clusters"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/redshift/clusters/"
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
      name       = "allow_version_upgrade"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "aqua_configuration"
      type       = "struct<aqua_configuration_status:string,aqua_status:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "automated_snapshot_retention_period"
      type       = "int"
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
      name       = "availability_zone_relocation_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cluster_availability_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cluster_identifier"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cluster_namespace_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cluster_nodes"
      type       = "array<struct<node_role:string,private_ip_address:string,public_ip_address:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cluster_parameter_groups"
      type       = "array<struct<cluster_parameter_status_list:array<struct<parameter_apply_error_description:string,parameter_apply_status:string,parameter_name:string>>,parameter_apply_status:string,parameter_group_name:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cluster_public_key"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cluster_revision_number"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cluster_security_groups"
      type       = "array<struct<cluster_security_group_name:string,status:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cluster_snapshot_copy_status"
      type       = "struct<destination_region:string,manual_snapshot_retention_period:int,retention_period:bigint,snapshot_copy_grant_name:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cluster_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cluster_subnet_group_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cluster_version"
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
      name       = "data_transfer_progress"
      type       = "struct<current_rate_in_mega_bytes_per_second:double,data_transferred_in_mega_bytes:bigint,elapsed_time_in_seconds:bigint,estimated_time_to_completion_in_seconds:bigint,status:string,total_data_in_mega_bytes:bigint>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "default_iam_role_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "deferred_maintenance_windows"
      type       = "array<struct<defer_maintenance_identifier:string,defer_maintenance_end_time:timestamp,defer_maintenance_start_time:timestamp>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "elastic_ip_status"
      type       = "struct<elastic_ip:string,status:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "elastic_resize_number_of_node_options"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "encrypted"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "endpoint"
      type       = "struct<address:string,port:int,vpc_endpoints:array<struct<network_interfaces:array<struct<availability_zone:string,network_interface_id:string,private_ip_address:string,subnet_id:string>>,vpc_endpoint_id:string,vpc_id:string>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "enhanced_vpc_routing"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "expected_next_snapshot_schedule_time_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "hsm_status"
      type       = "struct<hsm_client_certificate_identifier:string,hsm_configuration_identifier:string,status:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "iam_roles"
      type       = "array<struct<apply_status:string,iam_role_arn:string>>"
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
      name       = "maintenance_track_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "manual_snapshot_retention_period"
      type       = "int"
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
      name       = "modify_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "node_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "number_of_nodes"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "pending_actions"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "pending_modified_values"
      type       = "struct<automated_snapshot_retention_period:int,cluster_identifier:string,cluster_type:string,cluster_version:string,encryption_type:string,enhanced_vpc_routing:boolean,maintenance_track_name:string,node_type:string,number_of_nodes:int,publicly_accessible:boolean>"
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
      name       = "reserved_node_exchange_status"
      type       = "struct<reserved_node_exchange_request_id:string,source_reserved_node_count:int,source_reserved_node_id:string,source_reserved_node_type:string,status:string,target_reserved_node_count:int,target_reserved_node_offering_id:string,target_reserved_node_type:string,request_time:timestamp>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "resize_info"
      type       = "struct<allow_cancel_resize:boolean,resize_type:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "restore_status"
      type       = "struct<current_restore_rate_in_mega_bytes_per_second:double,elapsed_time_in_seconds:bigint,estimated_time_to_completion_in_seconds:bigint,progress_in_mega_bytes:bigint,snapshot_size_in_mega_bytes:bigint,status:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "snapshot_schedule_identifier"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "snapshot_schedule_state"
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
      name       = "total_storage_capacity_in_mega_bytes"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_id"
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
      name       = "cluster_create_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "expected_next_snapshot_schedule_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "next_maintenance_window_start_time"
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
