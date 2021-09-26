
resource "aws_glue_catalog_table" "ecs_clusters" {
  name          = "ecs_clusters"
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
    location      = "s3://${var.bucket_name}/parquet/aws/ecs/clusters/"
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
      name    = "active_services_count"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "attachments"
      type    = "array<struct<details:array<struct<name:string,value:string>>,id:string,status:string,type:string>>"
      comment = ""
    }
    columns {
      name    = "attachments_status"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "capacity_providers"
      type    = "array<string>"
      comment = ""
    }
    columns {
      name    = "cluster_arn"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "cluster_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "configuration"
      type    = "struct<execute_command_configuration:struct<kms_key_id:string,log_configuration:struct<cloud_watch_encryption_enabled:boolean,cloud_watch_log_group_name:string,s3_bucket_name:string,s3_encryption_enabled:boolean,s3_key_prefix:string>,logging:string>>"
      comment = ""
    }
    columns {
      name    = "default_capacity_provider_strategy"
      type    = "array<struct<capacity_provider:string,base:int,weight:int>>"
      comment = ""
    }
    columns {
      name    = "pending_tasks_count"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "registered_container_instances_count"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "running_tasks_count"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "settings"
      type    = "array<struct<name:string,value:string>>"
      comment = ""
    }
    columns {
      name    = "statistics"
      type    = "array<struct<name:string,value:string>>"
      comment = ""
    }
    columns {
      name    = "status"
      type    = "string"
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
    columns {
      name    = "services"
      type    = "array<struct<capacity_provider_strategy:array<struct<capacity_provider:string,base:int,weight:int>>,cluster_arn:string,created_by:string,deployment_configuration:struct<deployment_circuit_breaker:struct<enable:boolean,rollback:boolean>,maximum_percent:int,minimum_healthy_percent:int>,deployment_controller:struct<type:string>,deployments:array<struct<capacity_provider_strategy:array<struct<capacity_provider:string,base:int,weight:int>>,desired_count:int,failed_tasks:int,id:string,launch_type:string,network_configuration:struct<awsvpc_configuration:struct<subnets:array<string>,assign_public_ip:string,security_groups:array<string>>>,pending_count:int,platform_version:string,rollout_state:string,rollout_state_reason:string,running_count:int,status:string,task_definition:string,created_at_milli:timestamp,updated_at_milli:timestamp>>,desired_count:int,enable_ecs_managed_tags:boolean,enable_execute_command:boolean,events:array<struct<id:string,message:string,created_at_milli:timestamp>>,health_check_grace_period_seconds:int,launch_type:string,load_balancers:array<struct<container_name:string,container_port:int,load_balancer_name:string,target_group_arn:string>>,network_configuration:struct<awsvpc_configuration:struct<subnets:array<string>,assign_public_ip:string,security_groups:array<string>>>,pending_count:int,placement_constraints:array<struct<expression:string,type:string>>,placement_strategy:array<struct<field:string,type:string>>,platform_version:string,propagate_tags:string,role_arn:string,running_count:int,scheduling_strategy:string,service_arn:string,service_name:string,service_registries:array<struct<container_name:string,container_port:int,port:int,registry_arn:string>>,status:string,tags:array<struct<key:string,value:string>>,task_definition:string,task_sets:array<struct<capacity_provider_strategy:array<struct<capacity_provider:string,base:int,weight:int>>,cluster_arn:string,computed_desired_count:int,external_id:string,id:string,launch_type:string,load_balancers:array<struct<container_name:string,container_port:int,load_balancer_name:string,target_group_arn:string>>,network_configuration:struct<awsvpc_configuration:struct<subnets:array<string>,assign_public_ip:string,security_groups:array<string>>>,pending_count:int,platform_version:string,running_count:int,scale:struct<unit:string,value:double>,service_arn:string,service_registries:array<struct<container_name:string,container_port:int,port:int,registry_arn:string>>,stability_status:string,started_by:string,status:string,tags:array<struct<key:string,value:string>>,task_definition:string,task_set_arn:string,created_at_milli:timestamp,stability_status_at_milli:timestamp,updated_at_milli:timestamp>>,created_at_milli:timestamp>>"
      comment = ""
    }
    columns {
      name    = "tasks"
      type    = "array<struct<attachments:array<struct<details:array<struct<name:string,value:string>>,id:string,status:string,type:string>>,attributes:array<struct<name:string,target_id:string,target_type:string,value:string>>,availability_zone:string,capacity_provider_name:string,cluster_arn:string,connectivity:string,container_instance_arn:string,containers:array<struct<container_arn:string,cpu:string,exit_code:int,gpu_ids:array<string>,health_status:string,image:string,image_digest:string,last_status:string,managed_agents:array<struct<last_status:string,name:string,reason:string,last_started_at_milli:timestamp>>,memory:string,memory_reservation:string,name:string,network_bindings:array<struct<bind_ip:string,container_port:int,host_port:int,protocol:string>>,network_interfaces:array<struct<attachment_id:string,ipv6_address:string,private_ipv4_address:string>>,reason:string,runtime_id:string,task_arn:string>>,cpu:string,desired_status:string,enable_execute_command:boolean,ephemeral_storage:struct<size_in_gi_b:int>,group:string,health_status:string,inference_accelerators:array<struct<device_name:string,device_type:string>>,last_status:string,launch_type:string,memory:string,overrides:struct<container_overrides:array<struct<command:array<string>,cpu:int,environment:array<struct<name:string,value:string>>,environment_files:array<struct<type:string,value:string>>,memory:int,memory_reservation:int,name:string,resource_requirements:array<struct<type:string,value:string>>>>,cpu:string,ephemeral_storage:struct<size_in_gi_b:int>,execution_role_arn:string,inference_accelerator_overrides:array<struct<device_name:string,device_type:string>>,memory:string,task_role_arn:string>,platform_version:string,started_by:string,stop_code:string,stopped_reason:string,tags:array<struct<key:string,value:string>>,task_arn:string,task_definition_arn:string,version:bigint,connectivity_at_milli:timestamp,created_at_milli:timestamp,execution_stopped_at_milli:timestamp,pull_started_at_milli:timestamp,pull_stopped_at_milli:timestamp,started_at_milli:timestamp,stopped_at_milli:timestamp,stopping_at_milli:timestamp>>"
      comment = ""
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
