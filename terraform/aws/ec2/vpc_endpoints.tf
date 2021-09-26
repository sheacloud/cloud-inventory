
resource "aws_glue_catalog_table" "ec2_vpc_endpoints" {
  name          = "ec2_vpc_endpoints"
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
    location      = "s3://${var.bucket_name}/parquet/aws/ec2/vpc_endpoints/"
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
      name    = "dns_entries"
      type    = "array<struct<dns_name:string,hosted_zone_id:string>>"
      comment = ""
    }
    columns {
      name    = "groups"
      type    = "array<struct<group_id:string,group_name:string>>"
      comment = ""
    }
    columns {
      name    = "last_error"
      type    = "struct<code:string,message:string>"
      comment = ""
    }
    columns {
      name    = "network_interface_ids"
      type    = "array<string>"
      comment = ""
    }
    columns {
      name    = "owner_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "policy_document"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "private_dns_enabled"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "requester_managed"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "route_table_ids"
      type    = "array<string>"
      comment = ""
    }
    columns {
      name    = "service_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "state"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "subnet_ids"
      type    = "array<string>"
      comment = ""
    }
    columns {
      name    = "vpc_endpoint_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "vpc_endpoint_type"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "vpc_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "creation_timestamp_milli"
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
