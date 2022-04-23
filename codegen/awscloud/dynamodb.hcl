aws_service "dynamodb" {
  service_cap_name         = "DynamoDB"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/dynamodb"
  extra_utilized_functions = ["DescribeTable", "ListTagsOfResource"]
  tag_object_name          = "Tag"

  resource "tables" {
    fetch_function        = "ListTables"
    object_source_name    = "TableDescription"
    object_singular_name  = "Table"
    object_plural_name    = "Tables"
    object_unique_id      = "TableArn"
    object_response_field = "TableNames"
    model_only            = true
    pagination            = true
    use_post_processing   = false
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["TableName"]

    extra_field {
      name = "Tags"
      type = "map[string]string"
    }
  }
}
