datasource "aws" "cloudwatchlogs" "log_groups" {
    primary_resource_name = "LogGroup"
    primary_resource_field = "Arn"
    api_function = "DescribeLogGroups"
    primary_object_path = ["LogGroups"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
    models_only = false
    paginate = true
}