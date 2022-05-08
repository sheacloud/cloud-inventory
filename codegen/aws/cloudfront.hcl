aws_service "cloudfront" {
  service_cap_name         = "CloudFront"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/cloudfront"
  extra_utilized_functions = ["GetDistribution", "ListTagsForResource"]
  tag_object_name          = "Tag"
  region_override          = "us-east-1"

  resource "distributions" {
    fetch_function        = "ListDistributions"
    object_source_name    = "Distribution"
    object_singular_name  = "Distribution"
    object_plural_name    = "Distributions"
    object_unique_id      = "ARN"
    object_response_field = "DistributionList.Items"
    model_only            = true
    pagination            = true
    use_post_processing   = false
    excluded_fields       = []
    convert_tags          = true
    display_fields        = ["DomainName", "Id"]

    extra_field {
      name = "Tags"
      type = "map[string]string"
    }
  }
}
