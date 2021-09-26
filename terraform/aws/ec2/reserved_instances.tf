
resource "aws_glue_catalog_table" "ec2_reserved_instances" {
  name          = "ec2_reserved_instances"
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
    location      = "s3://${var.bucket_name}/parquet/aws/ec2/reserved_instances/"
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
      name    = "availability_zone"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "currency_code"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "duration"
      type    = "bigint"
      comment = ""
    }
    columns {
      name    = "fixed_price"
      type    = "float"
      comment = ""
    }
    columns {
      name    = "instance_count"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "instance_tenancy"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "instance_type"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "offering_class"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "offering_type"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "product_description"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "recurring_charges"
      type    = "array<struct<amount:double,frequency:string>>"
      comment = ""
    }
    columns {
      name    = "reserved_instances_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "scope"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "state"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "usage_price"
      type    = "float"
      comment = ""
    }
    columns {
      name    = "end_milli"
      type    = "timestamp"
      comment = ""
    }
    columns {
      name    = "start_milli"
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
