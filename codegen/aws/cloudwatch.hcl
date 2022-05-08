aws_service "cloudwatch" {
  service_cap_name         = "CloudWatch"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/cloudwatch"
  extra_utilized_functions = ["ListTagsForResource"]
  tag_object_name          = "Tag"

  resource "metric_alarms" {
    fetch_function        = "DescribeAlarms"
    object_source_name    = "MetricAlarm"
    object_singular_name  = "MetricAlarm"
    object_plural_name    = "MetricAlarms"
    object_unique_id      = "AlarmArn"
    object_response_field = "MetricAlarms"
    model_only            = false
    pagination            = true
    use_post_processing   = true
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["AlarmName"]

    extra_field {
      name = "Tags"
      type = "map[string]string"
    }
  }

  resource "composite_alarms" {
    fetch_function        = "DescribeAlarms"
    object_source_name    = "CompositeAlarm"
    object_singular_name  = "CompositeAlarm"
    object_plural_name    = "CompositeAlarms"
    object_unique_id      = "AlarmArn"
    object_response_field = "CompositeAlarms"
    model_only            = false
    pagination            = true
    use_post_processing   = true
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["AlarmName"]

    extra_field {
      name = "Tags"
      type = "map[string]string"
    }
  }
}
