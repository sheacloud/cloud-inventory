aws_service "apigatewayv2" {
  service_cap_name         = "ApiGatewayV2"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
  extra_utilized_functions = ["GetTags", "GetStages", "GetRoutes", "GetIntegrations", "GetAuthorizers"]

  resource "apis" {
    fetch_function        = "GetApis"
    object_source_name    = "Api"
    object_singular_name  = "Api"
    object_plural_name    = "Apis"
    object_unique_id      = "ApiId"
    object_response_field = "Items"
    model_only            = false
    pagination            = false
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
      object_source_name = "GetRouteOutput"
      new_field_name     = "Routes"
      field_type         = "list"
      excluded_fields    = ["ResultMetadata"]
    }

    child {
      object_source_name = "Integration"
      new_field_name     = "Integrations"
      field_type         = "list"
      excluded_fields    = ["ResponseParameters"]
    }

    child {
      object_source_name = "Authorizer"
      new_field_name     = "Authorizers"
      field_type         = "list"
    }
  }
}
