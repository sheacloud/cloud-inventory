aws_service "secretsmanager" {
  service_cap_name         = "SecretsManager"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
  extra_utilized_functions = []
  tag_object_name          = "Tag"


  resource "secrets" {
    fetch_function        = "ListSecrets"
    object_source_name    = "SecretListEntry"
    object_singular_name  = "Secret"
    object_plural_name    = "Secrets"
    object_unique_id      = "ARN"
    object_response_field = "SecretList"
    model_only            = false
    pagination            = true
    use_post_processing   = false
    excluded_fields       = []
    convert_tags          = true
    tag_field_name        = "Tags"
    display_fields        = ["Name"]
  }
}
