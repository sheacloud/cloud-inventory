aws_service "ecr" {
  service_cap_name         = "ECR"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/ecr"
  extra_utilized_functions = ["ListTagsForResource"]
  tag_object_name          = "Tag"

  resource "repositories" {
    fetch_function        = "DescribeRepositories"
    object_source_name    = "Repository"
    object_plural_name    = "Repositories"
    object_unique_id      = "RepositoryArn"
    object_response_field = "Repositories"
    model_only            = false
    pagination            = true
    use_post_processing   = true
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["RepositoryUri"]

    extra_field {
      name = "Tags"
      type = "map[string]string"
    }
  }
}
