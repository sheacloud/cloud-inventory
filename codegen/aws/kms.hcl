aws_service "kms" {
  service_cap_name         = "KMS"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/kms"
  extra_utilized_functions = ["DescribeKey", "ListResourceTags", "ListAliases"]

  resource "keys" {
    fetch_function        = "ListKeys"
    object_source_name    = "KeyMetadata"
    object_singular_name  = "Key"
    object_plural_name    = "Keys"
    object_unique_id      = "Arn"
    object_response_field = "Keys"
    model_only            = true
    pagination            = true
    use_post_processing   = true
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["KeyId"]

    child {
      object_source_name = "AliasListEntry"
      new_field_name     = "Aliases"
      field_type         = "list"
    }

    extra_field {
      name = "Tags"
      type = "map[string]string"
    }
  }
}
