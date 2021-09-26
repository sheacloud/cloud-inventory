
resource "aws_glue_catalog_table" "elasticloadbalancing_load_balancers" {
  name          = "elasticloadbalancing_load_balancers"
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
    location      = "s3://${var.bucket_name}/parquet/aws/elasticloadbalancing/load_balancers/"
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
      name    = "availability_zones"
      type    = "array<string>"
      comment = ""
    }
    columns {
      name    = "backend_server_descriptions"
      type    = "array<struct<instance_port:int,policy_names:array<string>>>"
      comment = ""
    }
    columns {
      name    = "canonical_hosted_zone_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "canonical_hosted_zone_name_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "dns_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "health_check"
      type    = "struct<healthy_threshold:int,interval:int,target:string,timeout:int,unhealthy_threshold:int>"
      comment = ""
    }
    columns {
      name    = "instances"
      type    = "array<struct<instance_id:string>>"
      comment = ""
    }
    columns {
      name    = "listener_descriptions"
      type    = "array<struct<listener:struct<instance_port:int,load_balancer_port:int,protocol:string,instance_protocol:string,ssl_certificate_id:string>,policy_names:array<string>>>"
      comment = ""
    }
    columns {
      name    = "load_balancer_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "policies"
      type    = "struct<app_cookie_stickiness_policies:array<struct<cookie_name:string,policy_name:string>>,lb_cookie_stickiness_policies:array<struct<cookie_expiration_period:bigint,policy_name:string>>,other_policies:array<string>>"
      comment = ""
    }
    columns {
      name    = "scheme"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "security_groups"
      type    = "array<string>"
      comment = ""
    }
    columns {
      name    = "source_security_group"
      type    = "struct<group_name:string,owner_alias:string>"
      comment = ""
    }
    columns {
      name    = "subnets"
      type    = "array<string>"
      comment = ""
    }
    columns {
      name    = "vpc_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "created_time_milli"
      type    = "timestamp"
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
