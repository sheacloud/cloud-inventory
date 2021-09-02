
resource "aws_glue_catalog_table" "ec2_volumes" {
  name          = "ec2_volumes"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL              = "TRUE"
    "parquet.compression" = "SNAPPY"
  }

  storage_descriptor {
    location      = "s3://sheacloud-test-parquet/parquet/ec2/volumes/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "my-stream"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = 1
      }
    }

    columns {
      name    = "attachments"
      type    = "array<struct<delete_on_termination:boolean,device:string,instance_id:string,volume_id:string,state:string>>"
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
  }

  partition_keys {
    name = "year"
    type = "int"
  }
  partition_keys {
    name = "month"
    type = "int"
  }
  partition_keys {
    name = "day"
    type = "int"
  }
  partition_keys {
    name = "accountid"
    type = "string"
  }
  partition_keys {
    name = "region"
    type = "string"
  }
}
