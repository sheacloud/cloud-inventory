aws_service "ecs" {
  service_cap_name         = "ECS"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/ecs"
  extra_utilized_functions = ["DescribeClusters", "DescribeServices", "DescribeTasks"]
  tag_object_name          = "Tag"

  resource "clusters" {
    fetch_function        = "ListClusters"
    object_source_name    = "Cluster"
    object_plural_name    = "Clusters"
    object_unique_id      = "ClusterArn"
    object_response_field = "Clusters"
    model_only            = true
    pagination            = true
    use_post_processing   = true
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["ClusterName"]
  }

  resource "services" {
    fetch_function        = "ListServices"
    object_source_name    = "Service"
    object_plural_name    = "Services"
    object_unique_id      = "ServiceArn"
    object_response_field = "Services"
    model_only            = true
    pagination            = true
    use_post_processing   = true
    excluded_fields       = ["Events"]
    convert_tags          = false
    display_fields        = ["ServiceName"]
  }

  resource "tasks" {
    fetch_function        = "ListTasks"
    object_source_name    = "Task"
    object_plural_name    = "Tasks"
    object_unique_id      = "TaskArn"
    object_response_field = "Tasks"
    model_only            = true
    pagination            = true
    use_post_processing   = true
    excluded_fields       = []
    convert_tags          = false
  }
}
