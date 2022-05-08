
resource "aws_glue_catalog_table" "aws_cloudfront_distributions" {
  name          = "aws_cloudfront_distributions"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/cloudfront/distributions/"
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
      name       = "arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "distribution_config"
      type       = "struct<caller_reference:string,comment:string,default_cache_behavior:struct<target_origin_id:string,viewer_protocol_policy:string,allowed_methods:struct<items:array<string>,quantity:int,cached_methods:struct<items:array<string>,quantity:int>>,cache_policy_id:string,compress:boolean,default_ttl:bigint,field_level_encryption_id:string,forwarded_values:struct<cookies:struct<forward:string,whitelisted_names:struct<quantity:int,items:array<string>>>,query_string:boolean,headers:struct<quantity:int,items:array<string>>,query_string_cache_keys:struct<quantity:int,items:array<string>>>,function_associations:struct<quantity:int,items:array<struct<event_type:string,function_arn:string>>>,lambda_function_associations:struct<quantity:int,items:array<struct<event_type:string,lambda_function_arn:string,include_body:boolean>>>,max_ttl:bigint,min_ttl:bigint,origin_request_policy_id:string,realtime_log_config_arn:string,response_headers_policy_id:string,smooth_streaming:boolean,trusted_key_groups:struct<enabled:boolean,quantity:int,items:array<string>>,trusted_signers:struct<enabled:boolean,quantity:int,items:array<string>>>,enabled:boolean,origins:struct<items:array<struct<domain_name:string,id:string,connection_attempts:int,connection_timeout:int,custom_headers:struct<quantity:int,items:array<struct<header_name:string,header_value:string>>>,custom_origin_config:struct<http_port:int,https_port:int,origin_protocol_policy:string,origin_keepalive_timeout:int,origin_read_timeout:int,origin_ssl_protocols:struct<items:array<string>,quantity:int>>,origin_path:string,origin_shield:struct<enabled:boolean,origin_shield_region:string>,s3_origin_config:struct<origin_access_identity:string>>>,quantity:int>,aliases:struct<quantity:int,items:array<string>>,cache_behaviors:struct<quantity:int,items:array<struct<path_pattern:string,target_origin_id:string,viewer_protocol_policy:string,allowed_methods:struct<items:array<string>,quantity:int,cached_methods:struct<items:array<string>,quantity:int>>,cache_policy_id:string,compress:boolean,default_ttl:bigint,field_level_encryption_id:string,forwarded_values:struct<cookies:struct<forward:string,whitelisted_names:struct<quantity:int,items:array<string>>>,query_string:boolean,headers:struct<quantity:int,items:array<string>>,query_string_cache_keys:struct<quantity:int,items:array<string>>>,function_associations:struct<quantity:int,items:array<struct<event_type:string,function_arn:string>>>,lambda_function_associations:struct<quantity:int,items:array<struct<event_type:string,lambda_function_arn:string,include_body:boolean>>>,max_ttl:bigint,min_ttl:bigint,origin_request_policy_id:string,realtime_log_config_arn:string,response_headers_policy_id:string,smooth_streaming:boolean,trusted_key_groups:struct<enabled:boolean,quantity:int,items:array<string>>,trusted_signers:struct<enabled:boolean,quantity:int,items:array<string>>>>>,custom_error_responses:struct<quantity:int,items:array<struct<error_code:int,error_caching_min_ttl:bigint,response_code:string,response_page_path:string>>>,default_root_object:string,http_version:string,is_ipv6_enabled:boolean,logging:struct<bucket:string,enabled:boolean,include_cookies:boolean,prefix:string>,origin_groups:struct<quantity:int,items:array<struct<failover_criteria:struct<status_codes:struct<items:array<int>,quantity:int>>,id:string,members:struct<items:array<struct<origin_id:string>>,quantity:int>>>>,price_class:string,restrictions:struct<geo_restriction:struct<quantity:int,restriction_type:string,items:array<string>>>,viewer_certificate:struct<acm_certificate_arn:string,certificate:string,certificate_source:string,cloud_front_default_certificate:boolean,iam_certificate_id:string,minimum_protocol_version:string,ssl_support_method:string>,web_acl_id:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "domain_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "in_progress_invalidation_batches"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "last_modified_time"
      type       = "timestamp"
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
      name       = "active_trusted_key_groups"
      type       = "struct<enabled:boolean,quantity:int,items:array<struct<key_group_id:string,key_pair_ids:struct<quantity:int,items:array<string>>>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "active_trusted_signers"
      type       = "struct<enabled:boolean,quantity:int,items:array<struct<aws_account_number:string,key_pair_ids:struct<quantity:int,items:array<string>>>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "alias_icp_recordals"
      type       = "array<struct<cname:string,icp_recordal_status:string>>"
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
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
