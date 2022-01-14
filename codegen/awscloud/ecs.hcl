aws_service "ecs" {
    service_cap_name = "ECS"
    sdk_path = "github.com/aws/aws-sdk-go-v2/service/ecs"
    extra_utilized_functions = ["DescribeClusters", "DescribeServices", "DescribeTasks"]
    tag_object_name = "Tag"

    resource "clusters" {
        fetch_function = "ListClusters"
        object_name = "Cluster"
        object_unique_id = "ClusterArn"
        object_response_field = "Clusters"
        model_only = true
        pagination = true
        use_post_processing = true
        excluded_fields = []
        convert_tags = false
    }

    resource "services" {
        fetch_function = "ListServices"
        object_name = "Service"
        object_unique_id = "ServiceArn"
        object_response_field = "Services"
        model_only = true
        pagination = true
        use_post_processing = true
        excluded_fields = ["Events"]
        convert_tags = false
    }

    resource "tasks" {
        fetch_function = "ListTasks"
        object_name = "Task"
        object_unique_id = "TaskArn"
        object_response_field = "Tasks"
        model_only = true
        pagination = true
        use_post_processing = true
        excluded_fields = []
        convert_tags = false
    }
}