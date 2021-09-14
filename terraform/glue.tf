resource "aws_glue_catalog_database" "cloud_inventory" {
  name = "cloud-inventory"
}

module "cloudwatchlogs" {
  source = "./aws/cloudwatchlogs"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}

module "dynamodb" {
  source = "./aws/dynamodb"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}

module "ec2" {
  source = "./aws/ec2"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}

module "ecs" {
  source = "./aws/ecs"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}

module "efs" {
  source = "./aws/efs"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}

module "eks" {
  source = "./aws/eks"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}

module "elasticloadbalancing" {
  source = "./aws/elasticloadbalancing"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}

module "elasticloadbalancingv2" {
  source = "./aws/elasticloadbalancingv2"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}

module "iam" {
  source = "./aws/iam"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}

module "lambda" {
  source = "./aws/lambda"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}

module "rds" {
  source = "./aws/rds"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}

module "redshift" {
  source = "./aws/redshift"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}

module "route53" {
  source = "./aws/route53"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}

module "s3" {
  source = "./aws/s3"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}

module "sns" {
  source = "./aws/sns"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}

module "sqs" {
  source = "./aws/sqs"
  glue_database_name = aws_glue_catalog_database.cloud_inventory.name
  bucket_name = aws_s3_bucket.cloud_inventory.bucket
}