datasource "aws" "efs" "filesystems" {
    primary_resource_name = "FileSystemDescription"
    primary_resource_field = "FileSystemId"
    api_function = "DescribeFileSystems"
    primary_object_path = ["FileSystems"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/efs"
    models_only = false
    paginate = true
}