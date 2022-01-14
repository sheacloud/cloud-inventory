aws_service "elasticloadbalancingv2" {
    service_cap_name = "ElasticLoadBalancingV2"
    sdk_path = "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
    extra_utilized_functions = ["DescribeTags", "DescribeTargetHealth"]
    tag_object_name = "Tag"

    resource "load_balancers" {
        fetch_function = "DescribeLoadBalancers"
        object_name = "LoadBalancer"
        object_unique_id = "LoadBalancerArn"
        object_response_field = "LoadBalancers"
        model_only = false
        pagination = true
        use_post_processing = true
        excluded_fields = []
        convert_tags = false

        child {
            object_name = "Listener"
            new_field_name = "Listeners"
            field_type = "list"
        }

        extra_field {
            name = "Tags"
            type = "map[string]string"
        }
    }

    resource "target_groups" {
        fetch_function = "DescribeTargetGroups"
        object_name = "TargetGroup"
        object_unique_id = "TargetGroupArn"
        object_response_field = "TargetGroups"
        model_only = false
        pagination = true
        use_post_processing = true
        excluded_fields = []
        convert_tags = false

        child {
            object_name = "TargetHealthDescription"
            new_field_name = "Targets"
            field_type = "list"
        }

        extra_field {
            name = "Tags"
            type = "map[string]string"
        }
    }
}