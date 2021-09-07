
resource "aws_glue_catalog_table" "ec2_vpcs" {
  name          = "ec2_vpcs"
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
    location      = "s3://${var.bucket_name}/parquet/aws/ec2/vpcs/"
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
      name    = "cidr_block"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "cidr_block_association_set"
      type    = "array<struct<association_id:string,cidr_block:string,cidr_block_state:struct<state:string,association_id:string>>>"
      comment = ""
    }
    columns {
      name    = "dhcp_options_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "tenancy"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "ipv6_cidr_block_association_set"
      type    = "array<struct<association_id:string,ipv6_cidr_block:string,ipv6_cidr_block_state:struct<state:string,association_id:string>,ipv6_pool:string,network_border_group:string>>"
      comment = ""
    }
    columns {
      name    = "is_default"
      type    = "boolean"
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
      name    = "tags"
      type    = "map<string,string>"
      comment = ""
    }
    columns {
      name    = "vpc_id"
      type    = "string"
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
