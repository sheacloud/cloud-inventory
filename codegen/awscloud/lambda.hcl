aws_service "lambda" {
    service_cap_name = "Lambda"
    sdk_path = "github.com/aws/aws-sdk-go-v2/service/lambda"
    extra_utilized_functions = []

    resource "filesystems" {
        fetch_function = "ListFunctions"
        object_name = "FunctionConfiguration"
        object_plural_name = "Functions"
        object_unique_id = "FunctionArn"
        object_response_field = "Functions"
        model_only = false
        pagination = true
        use_post_processing = false
        excluded_fields = ["Environment"] // people might put secrets in the environment variables, don't want to ingest those
        convert_tags = false
    }
}