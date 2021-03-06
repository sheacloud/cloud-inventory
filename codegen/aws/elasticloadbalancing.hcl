aws_service "elasticloadbalancing" {
  service_cap_name         = "ElasticLoadBalancing"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
  extra_utilized_functions = []
  tag_object_name          = "Tag"

  resource "load_balancers" {
    fetch_function        = "DescribeLoadBalancers"
    object_source_name    = "LoadBalancerDescription"
    object_singular_name  = "LoadBalancer"
    object_plural_name    = "LoadBalancers"
    object_unique_id      = "LoadBalancerName"
    object_response_field = "LoadBalancerDescriptions"
    model_only            = false
    pagination            = true
    use_post_processing   = false
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["LoadBalancerName"]
  }
}
