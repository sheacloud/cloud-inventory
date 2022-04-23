
resource "aws_dynamodb_table" "aws_iam_groups" {
  name         = "cloud-inventory-aws-iam-groups"
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
    name = "group_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "group_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "group_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_iam_policies" {
  name         = "cloud-inventory-aws-iam-policies"
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
    name = "policy_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "policy_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "policy_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_iam_roles" {
  name         = "cloud-inventory-aws-iam-roles"
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
    name = "role_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "role_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "role_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
resource "aws_dynamodb_table" "aws_iam_users" {
  name         = "cloud-inventory-aws-iam-users"
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
    name = "user_id"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "user_id"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "user_id"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
