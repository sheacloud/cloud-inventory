
resource "aws_glue_catalog_table" "ec2_transit_gateway_peering_attachments" {
  name          = "ec2_transit_gateway_peering_attachments"
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
    location      = "s3://${var.bucket_name}/parquet/aws/ec2/transit_gateway_peering_attachments/"
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
      name    = "accepter_tgw_info"
      type    = "struct<owner_id:string,region:string,transit_gateway_id:string>"
      comment = ""
    }
    columns {
      name    = "requester_tgw_info"
      type    = "struct<owner_id:string,region:string,transit_gateway_id:string>"
      comment = ""
    }
    columns {
      name    = "state"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "status"
      type    = "struct<code:string,message:string>"
      comment = ""
    }
    columns {
      name    = "transit_gateway_attachment_id"
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
