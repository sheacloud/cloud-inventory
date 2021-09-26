service "aws" "rds" {
    library_path = "github.com/aws/aws-sdk-go-v2/service/rds"

    datasource "db_clusters" {
        primary_object_name = "DBCluster"
        primary_object_field = "DBClusterArn"
        api_function = "DescribeDBClusters"
        primary_object_path = ["DBClusters"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "TagList"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "db_instances" {
        primary_object_name = "DBInstance"
        primary_object_field = "DBInstanceArn"
        api_function = "DescribeDBInstances"
        primary_object_path = ["DBInstances"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "TagList"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }
}