aws_service "dynamodb" {
    service_cap_name = "DynamoDB"
    sdk_path = "github.com/aws/aws-sdk-go-v2/service/dynamodb"
    extra_utilized_functions = ["DescribeTable", "ListTagsOfResource"]
    tag_object_name = "Tag"

    resource "tables" {
        fetch_function = "ListTables"
        object_name = "TableDescription"
        object_unique_id = "TableArn"
        object_response_field = "TableNames"
        model_only = true
        pagination = true
        use_post_processing = false
        excluded_fields = []
        convert_tags = false

        extra_field {
            name = "Tags"
            type = "map[string]string"
        }
    }
}