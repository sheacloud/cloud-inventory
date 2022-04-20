aws_service "apigateway" {
  service_cap_name         = "ApiGateway"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/apigateway"
  extra_utilized_functions = ["GetTags", "GetStages", "GetResources"]

  resource "rest_apis" {
    fetch_function        = "GetRestApis"
    object_source_name    = "RestApi"
    object_singular_name  = "RestApi"
    object_plural_name    = "RestApis"
    object_unique_id      = "Id"
    object_response_field = "Items"
    model_only            = false
    pagination            = true
    use_post_processing   = true
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["Name"]

    child {
      object_source_name = "Stage"
      new_field_name     = "Stages"
      field_type         = "list"
    }

    child {
      object_source_name = "Resource"
      new_field_name     = "Resources"
      field_type         = "list"
    }
  }
}
