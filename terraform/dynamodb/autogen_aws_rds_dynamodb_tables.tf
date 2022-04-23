
resource "aws_dynamodb_table" "aws_rds_db_clusters" {
  name         = "cloud-inventory-aws-rds-db-clusters"
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
    name = "db_cluster_arn"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "db_cluster_arn"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "db_cluster_arn"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_rds_db_instances" {
  name         = "cloud-inventory-aws-rds-db-instances"
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
    name = "db_instance_arn"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "db_instance_arn"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "db_instance_arn"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
