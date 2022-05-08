
resource "aws_glue_catalog_table" "aws_autoscaling_auto_scaling_groups" {
  name          = "aws_autoscaling_auto_scaling_groups"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/autoscaling/auto_scaling_groups/"
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
      name       = "auto_scaling_group_name"
      type       = "string"
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
      name       = "created_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "default_cooldown"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "desired_capacity"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "health_check_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "max_size"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "min_size"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "auto_scaling_group_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "capacity_rebalance"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "context"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "default_instance_warmup"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "desired_capacity_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "enabled_metrics"
      type       = "array<struct<granularity:string,metric:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "health_check_grace_period"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "instances"
      type       = "array<struct<availability_zone:string,health_status:string,instance_id:string,lifecycle_state:string,protected_from_scale_in:boolean,instance_type:string,launch_configuration_name:string,launch_template:struct<launch_template_id:string,launch_template_name:string,version:string>,weighted_capacity:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "launch_configuration_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "launch_template"
      type       = "struct<launch_template_id:string,launch_template_name:string,version:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "load_balancer_names"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "max_instance_lifetime"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "mixed_instances_policy"
      type       = "struct<instances_distribution:struct<on_demand_allocation_strategy:string,on_demand_base_capacity:int,on_demand_percentage_above_base_capacity:int,spot_allocation_strategy:string,spot_instance_pools:int,spot_max_price:string>,launch_template:struct<launch_template_specification:struct<launch_template_id:string,launch_template_name:string,version:string>,overrides:array<struct<instance_requirements:struct<memory_mi_b:struct<min:int,max:int>,v_cpu_count:struct<min:int,max:int>,accelerator_count:struct<max:int,min:int>,accelerator_manufacturers:array<string>,accelerator_names:array<string>,accelerator_total_memory_mi_b:struct<max:int,min:int>,accelerator_types:array<string>,bare_metal:string,baseline_ebs_bandwidth_mbps:struct<max:int,min:int>,burstable_performance:string,cpu_manufacturers:array<string>,excluded_instance_types:array<string>,instance_generations:array<string>,local_storage:string,local_storage_types:array<string>,memory_gi_b_per_v_cpu:struct<max:double,min:double>,network_interface_count:struct<max:int,min:int>,on_demand_max_price_percentage_over_lowest_price:int,require_hibernate_support:boolean,spot_max_price_percentage_over_lowest_price:int,total_local_storage_gb:struct<max:double,min:double>>,instance_type:string,launch_template_specification:struct<launch_template_id:string,launch_template_name:string,version:string>,weighted_capacity:string>>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "new_instances_protected_from_scale_in"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "placement_group"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "predicted_capacity"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "service_linked_role_arn"
      type       = "string"
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
      name       = "suspended_processes"
      type       = "array<struct<process_name:string,suspension_reason:string>>"
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
      name       = "target_group_arns"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "termination_policies"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_zone_identifier"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "warm_pool_configuration"
      type       = "struct<instance_reuse_policy:struct<reuse_on_scale_in:boolean>,max_group_prepared_capacity:int,min_size:int,pool_state:string,status:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "warm_pool_size"
      type       = "int"
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
      name       = "scaling_policies"
      type       = "array<struct<adjustment_type:string,alarms:array<struct<alarm_arn:string,alarm_name:string>>,auto_scaling_group_name:string,cooldown:int,enabled:boolean,estimated_instance_warmup:int,metric_aggregation_type:string,min_adjustment_magnitude:int,min_adjustment_step:int,policy_arn:string,policy_name:string,policy_type:string,predictive_scaling_configuration:struct<metric_specifications:array<struct<target_value:double,customized_capacity_metric_specification:struct<metric_data_queries:array<struct<id:string,expression:string,label:string,metric_stat:struct<metric:struct<metric_name:string,namespace:string,dimensions:array<struct<name:string,value:string>>>,stat:string,unit:string>,return_data:boolean>>>,customized_load_metric_specification:struct<metric_data_queries:array<struct<id:string,expression:string,label:string,metric_stat:struct<metric:struct<metric_name:string,namespace:string,dimensions:array<struct<name:string,value:string>>>,stat:string,unit:string>,return_data:boolean>>>,customized_scaling_metric_specification:struct<metric_data_queries:array<struct<id:string,expression:string,label:string,metric_stat:struct<metric:struct<metric_name:string,namespace:string,dimensions:array<struct<name:string,value:string>>>,stat:string,unit:string>,return_data:boolean>>>,predefined_load_metric_specification:struct<predefined_metric_type:string,resource_label:string>,predefined_metric_pair_specification:struct<predefined_metric_type:string,resource_label:string>,predefined_scaling_metric_specification:struct<predefined_metric_type:string,resource_label:string>>>,max_capacity_breach_behavior:string,max_capacity_buffer:int,mode:string,scheduling_buffer_time:int>,scaling_adjustment:int,step_adjustments:array<struct<scaling_adjustment:int,metric_interval_lower_bound:double,metric_interval_upper_bound:double>>,target_tracking_configuration:struct<target_value:double,customized_metric_specification:struct<metric_name:string,namespace:string,statistic:string,dimensions:array<struct<name:string,value:string>>,unit:string>,disable_scale_in:boolean,predefined_metric_specification:struct<predefined_metric_type:string,resource_label:string>>>>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_autoscaling_launch_configurations" {
  name          = "aws_autoscaling_launch_configurations"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/autoscaling/launch_configurations/"
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
      name       = "created_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "image_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "instance_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "launch_configuration_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "associate_public_ip_address"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "block_device_mappings"
      type       = "array<struct<device_name:string,ebs:struct<delete_on_termination:boolean,encrypted:boolean,iops:int,snapshot_id:string,throughput:int,volume_size:int,volume_type:string>,no_device:boolean,virtual_name:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "classic_link_vpc_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "classic_link_vpc_security_groups"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ebs_optimized"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "iam_instance_profile"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "instance_monitoring"
      type       = "struct<enabled:boolean>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "kernel_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "key_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "launch_configuration_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "metadata_options"
      type       = "struct<http_endpoint:string,http_put_response_hop_limit:int,http_tokens:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "placement_tenancy"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ramdisk_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "security_groups"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "spot_price"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "user_data"
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
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
