aws_service "ssm" {
  service_cap_name         = "SSM"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/ssm"
  extra_utilized_functions = []
  tag_object_name          = "Tag"


  resource "parameters" {
    fetch_function        = "DescribeParameters"
    object_source_name    = "ParameterMetadata"
    object_singular_name  = "Parameter"
    object_plural_name    = "Parameters"
    object_unique_id      = "Name"
    object_response_field = "Parameters"
    model_only            = false
    pagination            = true
    use_post_processing   = false
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["Name"]
  }
}
