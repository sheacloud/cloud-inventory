
resource "aws_dynamodb_table" "aws_apigatewayv2_apis" {
  name         = "cloud-inventory-aws-apigatewayv2-apis"
  billing_mode = "PAY_PER_REQUEST"
  hash_key     = "_id"

  attribute {
    name = "_id"
    type = "S"
  }
  attribute {
    name = "report_time"
    type = "N"
  }
  attribute {
    name = "api_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "api_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "api_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
