aws_service "route53" {
  service_cap_name         = "Route53"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/route53"
  extra_utilized_functions = ["GetHostedZone", "ListTagsForResource"]
  tag_object_name          = "Tag"
  region_override          = "us-east-1"

  resource "hosted_zones" {
    fetch_function        = "ListHostedZones"
    object_source_name    = "HostedZone"
    object_plural_name    = "HostedZones"
    object_unique_id      = "Id"
    object_response_field = "HostedZones"
    model_only            = false
    pagination            = true
    use_post_processing   = true
    excluded_fields       = ["MasterUserPassword"]
    convert_tags          = false
    display_fields        = ["Name"]

    extra_field {
      name = "Tags"
      type = "map[string]string"
    }

    child {
      object_source_name = "VPC"
      new_field_name     = "VpcAssociations"
      field_type         = "list"
    }
  }
}
