aws_service "cloudwatchlogs" {
    service_cap_name = "CloudWatchLogs"
    sdk_path = "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
    extra_utilized_functions = []

    resource "log_groups" {
        fetch_function = "DescribeLogGroups"
        object_name = "LogGroup"
        object_unique_id = "Arn"
        object_response_field = "LogGroups"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = false
    }
}