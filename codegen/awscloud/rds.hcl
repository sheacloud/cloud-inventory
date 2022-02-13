aws_service "rds" {
    service_cap_name = "RDS"
    sdk_path = "github.com/aws/aws-sdk-go-v2/service/rds"
    extra_utilized_functions = []
    tag_object_name = "Tag"

    resource "db_clusters" {
        fetch_function = "DescribeDBClusters"
        object_name = "DBCluster"
        object_plural_name = "DBClusters"
        object_unique_id = "DBClusterArn"
        object_response_field = "DBClusters"
        model_only = false
        pagination = true
        use_post_processing = true
        excluded_fields = []
        convert_tags = true
        tag_field_name = "TagList"
        display_fields = ["DBClusterIdentifier"]
    }

    resource "db_instances" {
        fetch_function = "DescribeDBInstances"
        object_name = "DBInstance"
        object_plural_name = "DBInstances"
        object_unique_id = "DBInstanceArn"
        object_response_field = "DBInstances"
        model_only = false
        pagination = true
        use_post_processing = true
        excluded_fields = []
        convert_tags = true
        tag_field_name = "TagList"
        display_fields = ["DBInstanceIdentifier"]
    }
}