datasource "aws" "elasticloadbalancing" "load_balancers" {
    primary_resource_name = "LoadBalancerDescription"
    primary_resource_field = "LoadBalancerName"
    api_function = "DescribeLoadBalancers"
    primary_object_path = ["LoadBalancerDescriptions"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
    models_only = false
    paginate = true
}