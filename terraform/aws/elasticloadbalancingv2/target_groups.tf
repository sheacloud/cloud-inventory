
resource "aws_glue_catalog_table" "elasticloadbalancingv2_target_groups" {
  name          = "elasticloadbalancingv2_target_groups"
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
    location      = "s3://${var.bucket_name}/parquet/aws/elasticloadbalancingv2/target_groups/"
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
      name    = "health_check_enabled"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "health_check_interval_seconds"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "health_check_path"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "health_check_port"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "health_check_protocol"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "health_check_timeout_seconds"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "healthy_threshold_count"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "load_balancer_arns"
      type    = "array<string>"
      comment = ""
    }
    columns {
      name    = "matcher"
      type    = "struct<grpc_code:string,http_code:string>"
      comment = ""
    }
    columns {
      name    = "port"
      type    = "int"
      comment = ""
    }
    columns {
      name    = "protocol"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "protocol_version"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "target_group_arn"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "target_group_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "target_type"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "unhealthy_threshold_count"
      type    = "int"
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
    columns {
      name    = "tags"
      type    = "map<string,string>"
      comment = ""
    }
    columns {
      name    = "targets"
      type    = "array<struct<health_check_port:string,target:struct<id:string,availability_zone:string,port:int>,target_health:struct<description:string,reason:string,state:string>>>"
      comment = ""
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
