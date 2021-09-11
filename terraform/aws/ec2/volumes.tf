
resource "aws_glue_catalog_table" "ec2_volumes" {
  name          = "ec2_volumes"
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
    location      = "s3://${var.bucket_name}/parquet/aws/ec2/volumes/"
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
      name    = "attachments"
      type    = "array<struct<attach_time:timestamp,delete_on_termination:boolean,device:string,instance_id:string,state:string,volume_id:string>>"
      comment = ""
    }
    columns {
      name    = "availability_zone"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "create_time"
      type    = "timestamp"
      comment = ""
    }
    columns {
      name    = "encrypted"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "fast_restored"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "iops"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "kms_key_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "multi_attach_enabled"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "outpost_arn"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "size"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "snapshot_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "state"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "tags"
      type    = "map<string,string>"
      comment = ""
    }
    columns {
      name    = "throughput"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "volume_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "volume_type"
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
