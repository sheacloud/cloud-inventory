service "aws" "elasticloadbalancingv2" {
    library_path = "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"

    datasource "load_balancers" {
        primary_object_name = "LoadBalancer"
        primary_object_field = "LoadBalancerArn"
        api_function = "DescribeLoadBalancers"
        primary_object_path = ["LoadBalancers"]
        models_only = false
        paginate = true

        child {
            resource_name = "Listener"
            resource_field = "Listeners"
            resource_type = "list"
        }

        extra_field {
            name = "Tags"
            type = "map[string]string"
        }
    }

    datasource "target_groups" {
        primary_object_name = "TargetGroup"
        primary_object_field = "TargetGroupArn"
        api_function = "DescribeTargetGroups"
        primary_object_path = ["TargetGroups"]
        models_only = false
        paginate = true

        child {
            resource_name = "TargetHealthDescription"
            resource_field = "Targets"
            resource_type = "list"
        }

        extra_field {
            name = "Tags"
            type = "map[string]string"
        }
    }
}