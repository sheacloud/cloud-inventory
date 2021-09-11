
resource "aws_glue_catalog_table" "ec2_addresses" {
  name          = "ec2_addresses"
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
    location      = "s3://${var.bucket_name}/parquet/aws/ec2/addresses/"
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
      name    = "allocation_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "association_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "carrier_ip"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "customer_owned_ip"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "customer_owned_ipv4_pool"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "domain"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "instance_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "network_border_group"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "network_interface_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "network_interface_owner_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "private_ip_address"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "public_ip"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "public_ipv4_pool"
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
