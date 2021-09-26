
resource "aws_glue_catalog_table" "ec2_transit_gateways" {
  name          = "ec2_transit_gateways"
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
    location      = "s3://${var.bucket_name}/parquet/aws/ec2/transit_gateways/"
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
      name    = "options"
      type    = "struct<amazon_side_asn:bigint,association_default_route_table_id:string,auto_accept_shared_attachments:string,default_route_table_association:string,default_route_table_propagation:string,dns_support:string,multicast_support:string,propagation_default_route_table_id:string,transit_gateway_cidr_blocks:array<string>,vpn_ecmp_support:string>"
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
      name    = "transit_gateway_arn"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "transit_gateway_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "creation_time_milli"
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
