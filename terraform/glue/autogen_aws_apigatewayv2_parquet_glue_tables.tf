
resource "aws_glue_catalog_table" "aws_apigatewayv2_apis" {
  name          = "aws_apigatewayv2_apis"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/apigatewayv2/apis/"
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
      name       = "name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "protocol_type"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "route_selection_expression"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "api_endpoint"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "api_gateway_managed"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "api_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "api_key_selection_expression"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "cors_configuration"
      type       = "struct<allow_credentials:boolean,allow_headers:array<string>,allow_methods:array<string>,allow_origins:array<string>,expose_headers:array<string>,max_age:int>"
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
      name       = "disable_schema_validation"
      type       = "boolean"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "import_info"
      type       = "array<string>"
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
      type       = "array<struct<stage_name:string,access_log_settings:struct<destination_arn:string,format:string>,api_gateway_managed:boolean,auto_deploy:boolean,client_certificate_id:string,created_date:timestamp,default_route_settings:struct<data_trace_enabled:boolean,detailed_metrics_enabled:boolean,logging_level:string,throttling_burst_limit:int,throttling_rate_limit:double>,deployment_id:string,description:string,last_deployment_status_message:string,last_updated_date:timestamp,route_settings:map<string,struct<data_trace_enabled:boolean,detailed_metrics_enabled:boolean,logging_level:string,throttling_burst_limit:int,throttling_rate_limit:double>>,stage_variables:map<string,string>,tags:map<string,string>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "routes"
      type       = "array<struct<api_gateway_managed:boolean,api_key_required:boolean,authorization_scopes:array<string>,authorization_type:string,authorizer_id:string,model_selection_expression:string,operation_name:string,request_models:map<string,string>,request_parameters:map<string,struct<required:boolean>>,route_id:string,route_key:string,route_response_selection_expression:string,target:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "integrations"
      type       = "array<struct<api_gateway_managed:boolean,connection_id:string,connection_type:string,content_handling_strategy:string,credentials_arn:string,description:string,integration_id:string,integration_method:string,integration_response_selection_expression:string,integration_subtype:string,integration_type:string,integration_uri:string,passthrough_behavior:string,payload_format_version:string,request_parameters:map<string,string>,request_templates:map<string,string>,template_selection_expression:string,timeout_in_millis:int,tls_config:struct<server_name_to_verify:string>>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "authorizers"
      type       = "array<struct<name:string,authorizer_credentials_arn:string,authorizer_id:string,authorizer_payload_format_version:string,authorizer_result_ttl_in_seconds:int,authorizer_type:string,authorizer_uri:string,enable_simple_responses:boolean,identity_source:array<string>,identity_validation_expression:string,jwt_configuration:struct<audience:array<string>,issuer:string>>>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
