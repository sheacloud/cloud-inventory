aws_service "autoscaling" {
  service_cap_name         = "AutoScaling"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/autoscaling"
  extra_utilized_functions = ["DescribePolicies"]

  resource "auto_scaling_groups" {
    fetch_function        = "DescribeAutoScalingGroups"
    object_source_name    = "AutoScalingGroup"
    object_singular_name  = "AutoScalingGroup"
    object_plural_name    = "AutoScalingGroups"
    object_unique_id      = "AutoScalingGroupARN"
    object_response_field = "AutoScalingGroups"
    model_only            = false
    pagination            = true
    use_post_processing   = true
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["AutoScalingGroupName"]

    child {
      object_source_name = "ScalingPolicy"
      new_field_name     = "ScalingPolicies"
      field_type         = "list"
    }
  }

  resource "launch_configurations" {
    fetch_function        = "DescribeLaunchConfigurations"
    object_source_name    = "LaunchConfiguration"
    object_singular_name  = "LaunchConfiguration"
    object_plural_name    = "LaunchConfigurations"
    object_unique_id      = "LaunchConfigurationARN"
    object_response_field = "LaunchConfigurations"
    model_only            = false
    pagination            = true
    use_post_processing   = false
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["LaunchConfigurationName"]
  }
}
