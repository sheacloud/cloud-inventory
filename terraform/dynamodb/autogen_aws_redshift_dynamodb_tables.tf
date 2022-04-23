
resource "aws_dynamodb_table" "aws_redshift_clusters" {
  name         = "cloud-inventory-aws-redshift-clusters"
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
    name = "cluster_identifier"
    type = "S"
  }

  global_secondary_index {
    name            = "report_time_id_index"
    hash_key        = "report_time"
    range_key       = "cluster_identifier"
    projection_type = "ALL"
  }
  global_secondary_index {
    name            = "id_report_time_index"
    hash_key        = "cluster_identifier"
    range_key       = "report_time"
    projection_type = "ALL"
  }
}
