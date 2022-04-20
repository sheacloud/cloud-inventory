aws_service "cloudtrail" {
  service_cap_name         = "CloudTrail"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/cloudtrail"
  extra_utilized_functions = ["GetTrail", "GetTrailStatus"]

  resource "trails" {
    fetch_function        = "ListTrails"
    object_source_name    = "Trail"
    object_plural_name    = "Trails"
    object_unique_id      = "TrailARN"
    object_response_field = "Trails"
    model_only            = true
    pagination            = false
    use_post_processing   = true
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["Name"]

    child {
      object_source_name = "GetTrailStatusOutput"
      new_field_name     = "Status"
      field_type         = "literal"
      excluded_fields    = ["ResultMetadata"]
    }
  }
}
