resource "aws_dynamodb_table" "meta_ingestion_timestamps" {
  name         = "cloud-inventory-ingestion-timestamps"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "ingestion_key"
  range_key    = "report_time"

  attribute {
    name = "ingestion_key"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
}

resource "aws_dynamodb_table" "meta_inventory_results" {
  name         = "cloud-inventory-inventory-results"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
}
