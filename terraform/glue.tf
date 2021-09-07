resource "aws_glue_catalog_database" "cloud_inventory" {
  name = "cloud-inventory"
}

module "ec2" {
  source = "./aws/ec2"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}

module "iam" {
  source = "./aws/iam"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}