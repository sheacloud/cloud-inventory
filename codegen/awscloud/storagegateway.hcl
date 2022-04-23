aws_service "storagegateway" {
  service_cap_name         = "StorageGateway"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/storagegateway"
  extra_utilized_functions = ["DescribeGatewayInformation"]
  tag_object_name          = "Tag"

  resource "gateways" {
    fetch_function        = "ListGateways"
    object_source_name    = "DescribeGatewayInformationOutput"
    object_singular_name  = "Gateway"
    object_plural_name    = "Gateways"
    object_unique_id      = "GatewayARN"
    object_response_field = "Gateways"
    model_only            = true
    pagination            = true
    use_post_processing   = false
    excluded_fields       = ["ResultMetadata"]
    convert_tags          = true
    display_fields        = ["GatewayName"]
  }
}
