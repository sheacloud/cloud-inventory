aws_service "cloudformation" {
  service_cap_name         = "CloudFormation"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/cloudformation"
  extra_utilized_functions = ["DescribeStacks"]
  tag_object_name          = "Tag"

  resource "stacks" {
    fetch_function        = "ListStacks"
    object_source_name    = "Stack"
    object_singular_name  = "Stack"
    object_plural_name    = "Stacks"
    object_unique_id      = "StackId"
    object_response_field = "StackSummaries"
    model_only            = true
    pagination            = true
    use_post_processing   = true
    excluded_fields       = []
    convert_tags          = true
    display_fields        = ["StackName"]
  }
}
