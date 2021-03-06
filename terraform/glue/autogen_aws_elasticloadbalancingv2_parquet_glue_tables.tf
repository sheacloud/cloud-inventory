
resource "aws_glue_catalog_table" "aws_elasticloadbalancingv2_load_balancers" {
  name          = "aws_elasticloadbalancingv2_load_balancers"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/elasticloadbalancingv2/load_balancers/"
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
      name       = "availability_zones"
      type       = "array<struct<load_balancer_addresses:array<struct<allocation_id:string,i_pv6_address:string,ip_address:string,private_i_pv4_address:string>>,outpost_id:string,subnet_id:string,zone_name:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "canonical_hosted_zone_id"
      type       = "string"
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
      name       = "customer_owned_ipv4_pool"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "dns_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ip_address_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "load_balancer_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "load_balancer_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "scheme"
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
      name       = "state"
      type       = "struct<code:string,reason:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "type"
      type       = "string"
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
    columns {
      name       = "listeners"
      type       = "array<struct<alpn_policy:array<string>,certificates:array<struct<certificate_arn:string,is_default:boolean>>,default_actions:array<struct<type:string,authenticate_cognito_config:struct<user_pool_arn:string,user_pool_client_id:string,user_pool_domain:string,authentication_request_extra_params:map<string,string>,on_unauthenticated_request:string,scope:string,session_cookie_name:string,session_timeout:bigint>,authenticate_oidc_config:struct<authorization_endpoint:string,client_id:string,issuer:string,token_endpoint:string,user_info_endpoint:string,authentication_request_extra_params:map<string,string>,client_secret:string,on_unauthenticated_request:string,scope:string,session_cookie_name:string,session_timeout:bigint,use_existing_client_secret:boolean>,fixed_response_config:struct<status_code:string,content_type:string,message_body:string>,forward_config:struct<target_group_stickiness_config:struct<duration_seconds:int,enabled:boolean>,target_groups:array<struct<target_group_arn:string,weight:int>>>,order:int,redirect_config:struct<status_code:string,host:string,path:string,port:string,protocol:string,query:string>,target_group_arn:string>>,listener_arn:string,load_balancer_arn:string,port:int,protocol:string,ssl_policy:string>>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
resource "aws_glue_catalog_table" "aws_elasticloadbalancingv2_target_groups" {
  name          = "aws_elasticloadbalancingv2_target_groups"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/elasticloadbalancingv2/target_groups/"
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
      name       = "health_check_enabled"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "health_check_interval_seconds"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "health_check_path"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "health_check_port"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "health_check_protocol"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "health_check_timeout_seconds"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "healthy_threshold_count"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "ip_address_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "load_balancer_arns"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "matcher"
      type       = "struct<grpc_code:string,http_code:string>"
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
      name       = "protocol"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "protocol_version"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "target_group_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "target_group_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "target_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "unhealthy_threshold_count"
      type       = "int"
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
    columns {
      name       = "targets"
      type       = "array<struct<health_check_port:string,target:struct<id:string,availability_zone:string,port:int>,target_health:struct<description:string,reason:string,state:string>>>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
