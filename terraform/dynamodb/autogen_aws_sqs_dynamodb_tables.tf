
resource "aws_dynamodb_table" "aws_sqs_queues" {
  name         = "cloud-inventory-aws-sqs-queues"
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
    name = "queue_arn"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "queue_arn"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "queue_arn"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
