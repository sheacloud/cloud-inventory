service "aws" "efs" {
    library_path = "github.com/aws/aws-sdk-go-v2/service/efs"

    datasource "filesystems" {
        primary_object_name = "FileSystemDescription"
        primary_object_field = "FileSystemId"
        api_function = "DescribeFileSystems"
        primary_object_path = ["FileSystems"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }
}

