
resource "aws_glue_catalog_table" "ec2_route_tables" {
  name          = "ec2_route_tables"
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
    location      = "s3://${var.bucket_name}/parquet/aws/ec2/route_tables/"
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
      type    = "array<struct<association_state:struct<state:string,status_message:string>,gateway_id:string,main:boolean,route_table_association_id:string,route_table_id:string,subnet_id:string>>"
      comment = ""
    }
    columns {
      name    = "owner_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "propagating_vgws"
      type    = "array<struct<gateway_id:string>>"
      comment = ""
    }
    columns {
      name    = "route_table_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "routes"
      type    = "array<struct<carrier_gateway_id:string,destination_cidr_block:string,destination_ipv6_cidr_block:string,destination_prefix_list_id:string,egress_only_internet_gateway_id:string,gateway_id:string,instance_id:string,instance_owner_id:string,local_gateway_id:string,nat_gateway_id:string,network_interface_id:string,origin:string,state:string,transit_gateway_id:string,vpc_peering_connection_id:string>>"
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
