terraform {
  backend "s3" {
    bucket = "sheacloud-terraform-state"
    key    = "aws/data-warehouse/terraform.tfstate"
    region = "us-east-1"
    dynamodb_table = "sheacloud-terraform-lock"
  }
}
