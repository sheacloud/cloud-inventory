
resource "aws_glue_catalog_table" "aws_route53_hosted_zones" {
  name          = "aws_route53_hosted_zones"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/route53/hosted_zones/"
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
      name       = "caller_reference"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "config"
      type       = "struct<comment:string,private_zone:boolean>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "linked_service"
      type       = "struct<description:string,service_principal:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "resource_record_set_count"
      type       = "bigint"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "account_id"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "region"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "report_time"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "vpc_associations"
      type       = "array<struct<vpc_id:string,vpc_region:string>>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
