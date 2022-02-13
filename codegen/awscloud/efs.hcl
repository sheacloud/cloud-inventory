aws_service "efs" {
    service_cap_name = "EFS"
    sdk_path = "github.com/aws/aws-sdk-go-v2/service/efs"
    extra_utilized_functions = []
    tag_object_name = "Tag"

    resource "filesystems" {
        fetch_function = "DescribeFileSystems"
        object_name = "FileSystemDescription"
        object_plural_name = "FileSystems"
        object_unique_id = "FileSystemId"
        object_response_field = "FileSystems"
        model_only = false
        pagination = true
        use_post_processing = true
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"
        display_fields = ["Name"]
    }
}