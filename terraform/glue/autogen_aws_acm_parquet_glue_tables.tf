
resource "aws_glue_catalog_table" "aws_acm_certificates" {
  name          = "aws_acm_certificates"
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
    location      = "s3://${var.s3_bucket_name}/inventory/aws/acm/certificates/"
    input_format  = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetInputFormat"
    output_format = "org.apache.hadoop.hive.ql.io.parquet.MapredParquetOutputFormat"

    ser_de_info {
      name                  = "ion"
      serialization_library = "org.apache.hadoop.hive.ql.io.parquet.serde.ParquetHiveSerDe"
      parameters = {
        "serialization.format" = "1"
      }
    }

    columns {
      name       = "certificate_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "certificate_authority_arn"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "created_at"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "domain_name"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "domain_validation_options"
      type       = "array<struct<domain_name:string,resource_record:struct<name:string,type:string,value:string>,validation_domain:string,validation_emails:array<string>,validation_method:string,validation_status:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "extended_key_usages"
      type       = "array<struct<name:string,oid:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "failure_reason"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "imported_at"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "in_use_by"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "issued_at"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "issuer"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "key_algorithm"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "key_usages"
      type       = "array<struct<name:string>>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "not_after"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "not_before"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "options"
      type       = "struct<certificate_transparency_logging_preference:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "renewal_eligibility"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "renewal_summary"
      type       = "struct<domain_validation_options:array<struct<domain_name:string,resource_record:struct<name:string,type:string,value:string>,validation_domain:string,validation_emails:array<string>,validation_method:string,validation_status:string>>,renewal_status:string,updated_at:timestamp,renewal_status_reason:string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "revocation_reason"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "revoked_at"
      type       = "timestamp"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "serial"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "signature_algorithm"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "status"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "subject"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "subject_alternative_names"
      type       = "array<string>"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "type"
      type       = "string"
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
      name       = "inventory_uuid"
      type       = "string"
      comment    = ""
      parameters = {}
    }
    columns {
      name       = "tags"
      type       = "map<string,string>"
      comment    = ""
      parameters = {}
    }
  }

  partition_keys {
    name = "report_date"
    type = "date"
  }
}
