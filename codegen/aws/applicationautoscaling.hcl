aws_service "applicationautoscaling" {
  service_cap_name         = "ApplicationAutoScaling"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/applicationautoscaling"
  extra_utilized_functions = []

  resource "scaling_policies" {
    fetch_function        = "DescribeScalingPolicies"
    object_source_name    = "ScalingPolicy"
    object_singular_name  = "ScalingPolicy"
    object_plural_name    = "ScalingPolicies"
    object_unique_id      = "PolicyARN"
    object_response_field = "ScalingPolicies"
    model_only            = true
    pagination            = true
    use_post_processing   = false
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["ServiceNamespace", "PolicyType", "PolicyName"]
  }
}
