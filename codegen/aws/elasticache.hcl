aws_service "elasticache" {
  service_cap_name         = "ElastiCache"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/elasticache"
  extra_utilized_functions = ["ListTagsForResource"]
  tag_object_name          = "Tag"

  resource "cache_clusters" {
    fetch_function        = "DescribeCacheClusters"
    object_source_name    = "CacheCluster"
    object_plural_name    = "CacheClusters"
    object_unique_id      = "ARN"
    object_response_field = "CacheClusters"
    model_only            = false
    pagination            = true
    use_post_processing   = true
    excluded_fields       = []
    convert_tags          = false

    extra_field {
      name = "Tags"
      type = "map[string]string"
    }
  }
}
