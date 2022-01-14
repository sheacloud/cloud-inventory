resource "aws_s3_bucket" "cloud_inventory" {
  bucket = "cloud-inventory-${random_id.cloud_inventory.hex}"
  acl    = "private"
  force_destroy = true

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm     = "AES256"
      }
    }
  }
}

resource "aws_s3_bucket" "athena_query_results" {
  bucket = "cloud-inventory-query-results-${random_id.cloud_inventory.hex}"
  acl    = "private"
  force_destroy = true

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm     = "AES256"
      }
    }
  }
}