datasource "aws" "ecs" "clusters" {
    primary_resource_name = "Cluster"
    primary_resource_field = "ClusterArn"
    api_function = "ListClusters"
    primary_object_path = ["Clusters"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/ecs"
    models_only = true
    paginate = false

    child {
        resource_name = "Service"
        resource_field = "Services"
        resource_type = "list"
    }

    child {
        resource_name = "Task"
        resource_field = "Tasks"
        resource_type = "list"
    }
}