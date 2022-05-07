aws_service "acm" {
  service_cap_name         = "ACM"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/acm"
  extra_utilized_functions = ["ListTagsForCertificate", "DescribeCertificate"]
  tag_object_name          = "Tag"

  resource "certificates" {
    fetch_function        = "ListCertificates"
    object_source_name    = "CertificateDetail"
    object_singular_name  = "Certificate"
    object_plural_name    = "Certificates"
    object_unique_id      = "CertificateArn"
    object_response_field = "CertificateSummaryList"
    model_only            = true
    pagination            = true
    use_post_processing   = true
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["CertificateArn", "DomainName"]

    extra_field {
      name = "Tags"
      type = "map[string]string"
    }
  }
}
