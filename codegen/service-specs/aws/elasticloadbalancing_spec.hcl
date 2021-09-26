service "aws" "elasticloadbalancing" {
    library_path = "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"

    datasource "load_balancers" {
        primary_object_name = "LoadBalancerDescription"
        primary_object_field = "LoadBalancerName"
        api_function = "DescribeLoadBalancers"
        primary_object_path = ["LoadBalancerDescriptions"]
        models_only = false
        paginate = true
    }
}