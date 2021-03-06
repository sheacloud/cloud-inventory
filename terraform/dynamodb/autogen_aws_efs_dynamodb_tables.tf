
resource "aws_dynamodb_table" "aws_efs_file_systems" {
  name         = "cloud-inventory-aws-efs-file-systems"
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
    name = "file_system_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "file_system_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "file_system_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
