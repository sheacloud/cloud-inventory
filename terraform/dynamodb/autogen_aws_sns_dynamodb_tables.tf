
resource "aws_dynamodb_table" "aws_sns_topics" {
  name         = "cloud-inventory-aws-sns-topics"
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
    name = "topic_arn"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "topic_arn"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "topic_arn"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_sns_subscriptions" {
  name         = "cloud-inventory-aws-sns-subscriptions"
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
    name = "subscription_arn"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "subscription_arn"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "subscription_arn"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
