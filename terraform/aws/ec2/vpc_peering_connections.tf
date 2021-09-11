
resource "aws_glue_catalog_table" "ec2_vpc_peering_connections" {
  name          = "ec2_vpc_peering_connections"
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
    location      = "s3://${var.bucket_name}/parquet/aws/ec2/vpc_peering_connections/"
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
      name    = "accepter_vpc_info"
      type    = "struct<cidr_block:string,cidr_block_set:array<struct<cidr_block:string>>,ipv6_cidr_block_set:array<struct<ipv6_cidr_block:string>>,owner_id:string,peering_options:struct<allow_dns_resolution_from_remote_vpc:boolean,allow_egress_from_local_classic_link_to_remote_vpc:boolean,allow_egress_from_local_vpc_to_remote_classic_link:boolean>,region:string,vpc_id:string>"
      comment = ""
    }
    columns {
      name    = "expiration_time"
      type    = "timestamp"
      comment = ""
    }
    columns {
      name    = "requester_vpc_info"
      type    = "struct<cidr_block:string,cidr_block_set:array<struct<cidr_block:string>>,ipv6_cidr_block_set:array<struct<ipv6_cidr_block:string>>,owner_id:string,peering_options:struct<allow_dns_resolution_from_remote_vpc:boolean,allow_egress_from_local_classic_link_to_remote_vpc:boolean,allow_egress_from_local_vpc_to_remote_classic_link:boolean>,region:string,vpc_id:string>"
      comment = ""
    }
    columns {
      name    = "status"
      type    = "struct<code:string,message:string>"
      comment = ""
    }
    columns {
      name    = "tags"
      type    = "map<string,string>"
      comment = ""
    }
    columns {
      name    = "vpc_peering_connection_id"
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
