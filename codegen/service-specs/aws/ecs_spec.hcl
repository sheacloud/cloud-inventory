service "aws" "ecs" {
    library_path = "github.com/aws/aws-sdk-go-v2/service/ecs"

    datasource "clusters" {
        primary_object_name = "Cluster"
        primary_object_field = "ClusterArn"
        api_function = "ListClusters"
        primary_object_path = ["Clusters"]
        
        models_only = true
        paginate = false

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }

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
}

