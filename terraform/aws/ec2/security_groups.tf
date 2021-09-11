
resource "aws_glue_catalog_table" "ec2_security_groups" {
  name          = "ec2_security_groups"
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
    location      = "s3://${var.bucket_name}/parquet/aws/ec2/security_groups/"
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
      name    = "description"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "group_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "group_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "ip_permissions"
      type    = "array<struct<from_port:int,ip_protocol:string,ip_ranges:array<struct<cidr_ip:string,description:string>>,ipv6_ranges:array<struct<cidr_ipv6:string,description:string>>,prefix_list_ids:array<struct<description:string,prefix_list_id:string>>,to_port:int,user_id_group_pairs:array<struct<description:string,group_id:string,group_name:string,peering_status:string,user_id:string,vpc_id:string,vpc_peering_connection_id:string>>>>"
      comment = ""
    }
    columns {
      name    = "ip_permissions_egress"
      type    = "array<struct<from_port:int,ip_protocol:string,ip_ranges:array<struct<cidr_ip:string,description:string>>,ipv6_ranges:array<struct<cidr_ipv6:string,description:string>>,prefix_list_ids:array<struct<description:string,prefix_list_id:string>>,to_port:int,user_id_group_pairs:array<struct<description:string,group_id:string,group_name:string,peering_status:string,user_id:string,vpc_id:string,vpc_peering_connection_id:string>>>>"
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
