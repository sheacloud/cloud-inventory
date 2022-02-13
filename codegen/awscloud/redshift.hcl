aws_service "redshift" {
    service_cap_name = "Redshift"
    sdk_path = "github.com/aws/aws-sdk-go-v2/service/redshift"
    extra_utilized_functions = []
    tag_object_name = "Tag"

    resource "clusters" {
        fetch_function = "DescribeClusters"
        object_name = "Cluster"
        object_plural_name = "Clusters"
        object_unique_id = "ClusterIdentifier"
        object_response_field = "Clusters"
        model_only = false
        pagination = true
        use_post_processing = true
        excluded_fields = ["MasterUserPassword"]
        convert_tags = true
        tag_field_name = "Tags"
        display_fields = ["ClusterIdentifier"]
    }
}