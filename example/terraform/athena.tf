resource "aws_athena_workgroup" "cloud_inventory" {
    name = "cloud-inventory-${random_id.cloud_inventory.hex}"
    force_destroy = true

    configuration {
        enforce_workgroup_configuration = true
        bytes_scanned_cutoff_per_query = 1000000000 // 1 gigabyte
        publish_cloudwatch_metrics_enabled = true
        result_configuration {
            output_location = "s3://${aws_s3_bucket.athena_query_results.id}"
            encryption_configuration {
                encryption_option = "SSE_S3"
            }
        }
    }
}