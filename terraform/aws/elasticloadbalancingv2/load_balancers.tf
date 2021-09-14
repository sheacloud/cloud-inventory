
resource "aws_glue_catalog_table" "elasticloadbalancingv2_load_balancers" {
  name          = "elasticloadbalancingv2_load_balancers"
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
    location      = "s3://${var.bucket_name}/parquet/aws/elasticloadbalancingv2/load_balancers/"
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
      name    = "availability_zones"
      type    = "array<struct<load_balancer_addresses:array<struct<allocation_id:string,i_pv6_address:string,ip_address:string,private_i_pv4_address:string>>,outpost_id:string,subnet_id:string,zone_name:string>>"
      comment = ""
    }
    columns {
      name    = "canonical_hosted_zone_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "created_time"
      type    = "timestamp"
      comment = ""
    }
    columns {
      name    = "customer_owned_ipv4_pool"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "dns_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "ip_address_type"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "load_balancer_arn"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "load_balancer_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "scheme"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "security_groups"
      type    = "array<string>"
      comment = ""
    }
    columns {
      name    = "state"
      type    = "struct<code:string,reason:string>"
      comment = ""
    }
    columns {
      name    = "type"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "vpc_id"
      type    = "string"
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
      name    = "listeners"
      type    = "array<struct<alpn_policy:array<string>,certificates:array<struct<certificate_arn:string,is_default:boolean>>,default_actions:array<struct<type:string,authenticate_cognito_config:struct<user_pool_arn:string,user_pool_client_id:string,user_pool_domain:string,authentication_request_extra_params:map<string,string>,on_unauthenticated_request:string,scope:string,session_cookie_name:string,session_timeout:bigint>,authenticate_oidc_config:struct<authorization_endpoint:string,client_id:string,issuer:string,token_endpoint:string,user_info_endpoint:string,authentication_request_extra_params:map<string,string>,client_secret:string,on_unauthenticated_request:string,scope:string,session_cookie_name:string,session_timeout:bigint,use_existing_client_secret:boolean>,fixed_response_config:struct<status_code:string,content_type:string,message_body:string>,forward_config:struct<target_group_stickiness_config:struct<duration_seconds:int,enabled:boolean>,target_groups:array<struct<target_group_arn:string,weight:int>>>,order:int,redirect_config:struct<status_code:string,host:string,path:string,port:string,protocol:string,query:string>,target_group_arn:string>>,listener_arn:string,load_balancer_arn:string,port:int,protocol:string,ssl_policy:string>>"
      comment = ""
    }
    columns {
      name    = "tags"
      type    = "map<string,string>"
      comment = ""
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
