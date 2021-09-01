resource "aws_glue_catalog_database" "data_warehouse" {
  name = "datawarehouse"
}

resource "aws_glue_catalog_table" "ec2_instances" {
  name          = "ec2_instances"
  database_name = aws_glue_catalog_database.data_warehouse.name

  table_type = "EXTERNAL_TABLE"

  parameters = {
    EXTERNAL              = "TRUE"
    "parquet.compression" = "SNAPPY"
  }

  storage_descriptor {
    location      = "s3://sheacloud-test-parquet/parquet/ec2/instance/"
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
      name = "architecture"
      type = "string"
    }

    columns {
      name = "iam_instance_profile"
      type = "struct<arn:string,id:string>"
    }

    columns {
      name    = "instance_id"
      type    = "string"
      comment = ""
    }

    columns {
      name    = "instance_type"
      type    = "string"
      comment = ""
    }

    columns {
      name    = "subnet_id"
      type    = "string"
      comment = ""
    }

    columns {
        name = "vpc_id"
        type = "string"
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