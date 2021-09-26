
resource "aws_glue_catalog_table" "ec2_nat_gateways" {
  name          = "ec2_nat_gateways"
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
    location      = "s3://${var.bucket_name}/parquet/aws/ec2/nat_gateways/"
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
      name    = "connectivity_type"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "failure_code"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "failure_message"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "nat_gateway_addresses"
      type    = "array<struct<allocation_id:string,network_interface_id:string,private_ip:string,public_ip:string>>"
      comment = ""
    }
    columns {
      name    = "nat_gateway_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "provisioned_bandwidth"
      type    = "struct<provisioned:string,requested:string,status:string,provision_time_milli:timestamp,request_time_milli:timestamp>"
      comment = ""
    }
    columns {
      name    = "state"
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
      name    = "create_time_milli"
      type    = "timestamp"
      comment = ""
    }
    columns {
      name    = "delete_time_milli"
      type    = "timestamp"
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
