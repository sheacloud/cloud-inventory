service "aws" "cloudwatchlogs" {
    library_path = "github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"

    datasource "log_groups" {
        primary_object_name = "LogGroup"
        primary_object_field = "Arn"
        api_function = "DescribeLogGroups"
        primary_object_path = ["LogGroups"]
        
        models_only = false
        paginate = true
    }
}