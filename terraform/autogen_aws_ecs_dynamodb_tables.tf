
resource "aws_dynamodb_table" "aws_ecs_clusters" {
  name         = "cloud-inventory-aws-ecs-clusters"
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
    name = "cluster_arn"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "cluster_arn"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "cluster_arn"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ecs_services" {
  name         = "cloud-inventory-aws-ecs-services"
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
    name = "service_arn"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "service_arn"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "service_arn"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_ecs_tasks" {
  name         = "cloud-inventory-aws-ecs-tasks"
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
    name = "task_arn"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "task_arn"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "task_arn"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
