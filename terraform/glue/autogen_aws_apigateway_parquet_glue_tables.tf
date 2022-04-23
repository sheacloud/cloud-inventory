
resource "aws_glue_catalog_table" "aws_apigateway_rest_apis" {
  name          = "aws_apigateway_rest_apis"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/apigateway/rest_apis/"
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
      name       = "api_key_source"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "binary_media_types"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "created_date"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "description"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "disable_execute_api_endpoint"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "endpoint_configuration"
      type       = "struct<types:array<string>,vpc_endpoint_ids:array<string>>"
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
      name       = "minimum_compression_size"
      type       = "int"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "policy"
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
      name       = "version"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "warnings"
      type       = "array<string>"
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
      name       = "stages"
      type       = "array<struct<access_log_settings:struct<destination_arn:string,format:string>,cache_cluster_enabled:boolean,cache_cluster_size:string,cache_cluster_status:string,canary_settings:struct<deployment_id:string,percent_traffic:double,stage_variable_overrides:map<string,string>,use_stage_cache:boolean>,client_certificate_id:string,created_date:timestamp,deployment_id:string,description:string,documentation_version:string,last_updated_date:timestamp,method_settings:map<string,struct<cache_data_encrypted:boolean,cache_ttl_in_seconds:int,caching_enabled:boolean,data_trace_enabled:boolean,logging_level:string,metrics_enabled:boolean,require_authorization_for_cache_control:boolean,throttling_burst_limit:int,throttling_rate_limit:double,unauthorized_cache_control_header_strategy:string>>,stage_name:string,tags:map<string,string>,tracing_enabled:boolean,variables:map<string,string>,web_acl_arn:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "resources"
      type       = "array<struct<id:string,parent_id:string,path:string,path_part:string,resource_methods:map<string,struct<api_key_required:boolean,authorization_scopes:array<string>,authorization_type:string,authorizer_id:string,http_method:string,method_integration:struct<cache_key_parameters:array<string>,cache_namespace:string,connection_id:string,connection_type:string,content_handling:string,credentials:string,http_method:string,integration_responses:map<string,struct<content_handling:string,response_parameters:map<string,string>,response_templates:map<string,string>,selection_pattern:string,status_code:string>>,passthrough_behavior:string,request_parameters:map<string,string>,request_templates:map<string,string>,timeout_in_millis:int,tls_config:struct<insecure_skip_verification:boolean>,type:string,uri:string>,method_responses:map<string,struct<response_models:map<string,string>,response_parameters:map<string,boolean>,status_code:string>>,operation_name:string,request_models:map<string,string>,request_parameters:map<string,boolean>,request_validator_id:string>>>>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
