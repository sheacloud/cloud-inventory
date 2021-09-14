datasource "aws" "rds" "db_clusters" {
    primary_resource_name = "DBCluster"
    primary_resource_field = "DBClusterArn"
    api_function = "DescribeDBClusters"
    primary_object_path = ["DBClusters"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/rds"
    models_only = false
    paginate = true
}

datasource "aws" "rds" "db_instances" {
    primary_resource_name = "DBInstance"
    primary_resource_field = "DBInstanceArn"
    api_function = "DescribeDBInstances"
    primary_object_path = ["DBInstances"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/rds"
    models_only = false
    paginate = true
}