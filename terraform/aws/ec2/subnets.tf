
resource "aws_glue_catalog_table" "ec2_subnets" {
  name          = "ec2_subnets"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL                        = "TRUE"
    "parquet.compression"           = "SNAPPY"
    "projection.enabled"            = "true"
    "projection.report_date.format" = "yyyy-MM-dd"
    "projection.report_date.range"  = "NOW-3YEARS,NOW"
    "projection.report_date.type"   = "date"
  }

  storage_descriptor {
    location      = "s3://${var.bucket_name}/parquet/aws/ec2/subnets/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name    = "assign_ipv6_address_on_creation"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "availability_zone"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "availability_zone_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "available_ip_address_count"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "cidr_block"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "customer_owned_ipv4_pool"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "default_for_az"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "ipv6_cidr_block_association_set"
      type    = "array<struct<association_id:string,ipv6_cidr_block:string,ipv6_cidr_block_state:struct<state:string,status_message:string>>>"
      comment = ""
    }
    columns {
      name    = "map_customer_owned_ip_on_launch"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "map_public_ip_on_launch"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "outpost_arn"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "owner_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "state"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "subnet_arn"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "subnet_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "vpc_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "tags"
      type    = "map<string,string>"
      comment = ""
    }
    columns {
      name    = "account_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "region"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "report_time"
      type    = "timestamp"
      comment = ""
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
