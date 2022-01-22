
resource "aws_glue_catalog_table" "aws_ecs_clusters" {
  name          = "aws_ecs_clusters"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ecs/clusters/"
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
      name       = "active_services_count"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "attachments"
      type       = "array<struct<details:array<struct<name:string,value:string>>,id:string,status:string,type:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "attachments_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "capacity_providers"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cluster_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cluster_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "configuration"
      type       = "struct<execute_command_configuration:struct<kms_key_id:string,log_configuration:struct<cloud_watch_encryption_enabled:boolean,cloud_watch_log_group_name:string,s3_bucket_name:string,s3_encryption_enabled:boolean,s3_key_prefix:string>,logging:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "default_capacity_provider_strategy"
      type       = "array<struct<capacity_provider:string,base:int,weight:int>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "pending_tasks_count"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "registered_container_instances_count"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "running_tasks_count"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "settings"
      type       = "array<struct<name:string,value:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "statistics"
      type       = "array<struct<name:string,value:string>>"
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
      name       = "tags"
      type       = "map<string,string>"
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
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_ecs_services" {
  name          = "aws_ecs_services"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ecs/services/"
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
      name       = "capacity_provider_strategy"
      type       = "array<struct<capacity_provider:string,base:int,weight:int>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cluster_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "created_by"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "deployment_configuration"
      type       = "struct<deployment_circuit_breaker:struct<enable:boolean,rollback:boolean>,maximum_percent:int,minimum_healthy_percent:int>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "deployment_controller"
      type       = "struct<type:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "deployments"
      type       = "array<struct<capacity_provider_strategy:array<struct<capacity_provider:string,base:int,weight:int>>,desired_count:int,failed_tasks:int,id:string,launch_type:string,network_configuration:struct<awsvpc_configuration:struct<subnets:array<string>,assign_public_ip:string,security_groups:array<string>>>,pending_count:int,platform_family:string,platform_version:string,rollout_state:string,rollout_state_reason:string,running_count:int,status:string,task_definition:string,created_at:timestamp,updated_at:timestamp>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "desired_count"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "enable_ecs_managed_tags"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "enable_execute_command"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "health_check_grace_period_seconds"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "launch_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "load_balancers"
      type       = "array<struct<container_name:string,container_port:int,load_balancer_name:string,target_group_arn:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "network_configuration"
      type       = "struct<awsvpc_configuration:struct<subnets:array<string>,assign_public_ip:string,security_groups:array<string>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "pending_count"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "placement_constraints"
      type       = "array<struct<expression:string,type:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "placement_strategy"
      type       = "array<struct<field:string,type:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "platform_family"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "platform_version"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "propagate_tags"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "role_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "running_count"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "scheduling_strategy"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "service_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "service_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "service_registries"
      type       = "array<struct<container_name:string,container_port:int,port:int,registry_arn:string>>"
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
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "task_definition"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "task_sets"
      type       = "array<struct<capacity_provider_strategy:array<struct<capacity_provider:string,base:int,weight:int>>,cluster_arn:string,computed_desired_count:int,external_id:string,id:string,launch_type:string,load_balancers:array<struct<container_name:string,container_port:int,load_balancer_name:string,target_group_arn:string>>,network_configuration:struct<awsvpc_configuration:struct<subnets:array<string>,assign_public_ip:string,security_groups:array<string>>>,pending_count:int,platform_family:string,platform_version:string,running_count:int,scale:struct<unit:string,value:double>,service_arn:string,service_registries:array<struct<container_name:string,container_port:int,port:int,registry_arn:string>>,stability_status:string,started_by:string,status:string,tags:array<struct<key:string,value:string>>,task_definition:string,task_set_arn:string,created_at:timestamp,stability_status_at:timestamp,updated_at:timestamp>>"
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
      name       = "created_at"
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
resource "aws_glue_catalog_table" "aws_ecs_tasks" {
  name          = "aws_ecs_tasks"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/ecs/tasks/"
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
      name       = "attachments"
      type       = "array<struct<details:array<struct<name:string,value:string>>,id:string,status:string,type:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "attributes"
      type       = "array<struct<name:string,target_id:string,target_type:string,value:string>>"
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
      name       = "capacity_provider_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cluster_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "connectivity"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "container_instance_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "containers"
      type       = "array<struct<container_arn:string,cpu:string,exit_code:int,gpu_ids:array<string>,health_status:string,image:string,image_digest:string,last_status:string,managed_agents:array<struct<last_status:string,name:string,reason:string,last_started_at:timestamp>>,memory:string,memory_reservation:string,name:string,network_bindings:array<struct<bind_ip:string,container_port:int,host_port:int,protocol:string>>,network_interfaces:array<struct<attachment_id:string,ipv6_address:string,private_ipv4_address:string>>,reason:string,runtime_id:string,task_arn:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cpu"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "desired_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "enable_execute_command"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ephemeral_storage"
      type       = "struct<size_in_gi_b:int>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "group"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "health_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "inference_accelerators"
      type       = "array<struct<device_name:string,device_type:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "last_status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "launch_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "memory"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "overrides"
      type       = "struct<container_overrides:array<struct<command:array<string>,cpu:int,environment:array<struct<name:string,value:string>>,environment_files:array<struct<type:string,value:string>>,memory:int,memory_reservation:int,name:string,resource_requirements:array<struct<type:string,value:string>>>>,cpu:string,ephemeral_storage:struct<size_in_gi_b:int>,execution_role_arn:string,inference_accelerator_overrides:array<struct<device_name:string,device_type:string>>,memory:string,task_role_arn:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "platform_family"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "platform_version"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "started_by"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "stop_code"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "stopped_reason"
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
      name       = "task_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "task_definition_arn"
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
      name       = "connectivity_at"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "created_at"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "execution_stopped_at"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "pull_started_at"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "pull_stopped_at"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "started_at"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "stopped_at"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "stopping_at"
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
