resource "aws_s3_bucket" "cloud_inventory" {
  bucket = var.bucket_name
  acl    = "private"

  server_side_encryption_configuration {
    rule {
      apply_server_side_encryption_by_default {
        sse_algorithm     = "AES256"
      }
    }
  }
}

resource "aws_s3_bucket_public_access_block" "cloud_inventory" {
  bucket = aws_s3_bucket.cloud_inventory.id

  block_public_acls   = true
  block_public_policy = true
  ignore_public_acls = true
  restrict_public_buckets = true
}