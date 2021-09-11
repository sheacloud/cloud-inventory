
resource "aws_glue_catalog_table" "ec2_network_acls" {
  name          = "ec2_network_acls"
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
    location      = "s3://${var.bucket_name}/parquet/aws/ec2/network_acls/"
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
      name    = "associations"
      type    = "array<struct<network_acl_association_id:string,network_acl_id:string,subnet_id:string>>"
      comment = ""
    }
    columns {
      name    = "entries"
      type    = "array<struct<cidr_block:string,egress:boolean,icmp_type_code:struct<code:int,type:int>,ipv6_cidr_block:string,port_range:struct<from:int,to:int>,protocol:string,rule_action:string,rule_number:int>>"
      comment = ""
    }
    columns {
      name    = "is_default"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "network_acl_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "owner_id"
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
