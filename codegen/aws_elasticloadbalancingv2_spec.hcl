datasource "aws" "elasticloadbalancingv2" "load_balancers" {
    primary_resource_name = "LoadBalancer"
    primary_resource_field = "LoadBalancerArn"
    api_function = "DescribeLoadBalancers"
    primary_object_path = ["LoadBalancers"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
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

datasource "aws" "elasticloadbalancingv2" "target_groups" {
    primary_resource_name = "TargetGroup"
    primary_resource_field = "TargetGroupArn"
    api_function = "DescribeTargetGroups"
    primary_object_path = ["TargetGroups"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
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