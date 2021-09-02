
resource "aws_glue_catalog_table" "ec2_network_interfaces" {
  name          = "ec2_network_interfaces"
  database_name = var.glue_database_name
  table_type    = "EXTERNAL_TABLE"
  parameters = {
    EXTERNAL              = "TRUE"
    "parquet.compression" = "SNAPPY"
  }

  storage_descriptor {
    location      = "s3://sheacloud-test-parquet/parquet/ec2/network_interfaces/"
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
      name    = "association"
      type    = "struct<allocation_id:string,association_id:string,carrier_ip:string,customer_owned_ip:string,ip_owner_id:string,public_dns_name:string,public_ip:string>"
      comment = ""
    }
    columns {
      name    = "attachment"
      type    = "struct<attachment_id:string,delete_on_termination:boolean,device_index:int,instance_id:string,instance_owner_id:string,network_card_index:int,status:string>"
      comment = ""
    }
    columns {
      name    = "availability_zone"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "description"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "groups"
      type    = "array<struct<group_id:string,group_name:string>>"
      comment = ""
    }
    columns {
      name    = "interface_type"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "mac_address"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "network_interface_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "owner_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "private_dns_name"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "private_ip_address"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "private_ip_addresses"
      type    = "array<struct<association:struct<allocation_id:string,association_id:string,carrier_ip:string,customer_owned_ip:string,ip_owner_id:string,public_dns_name:string,public_ip:string>,primary:boolean,private_dns_name:string,private_ip_address:string>>"
      comment = ""
    }
    columns {
      name    = "requester_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "requester_managed"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "source_dest_check"
      type    = "boolean"
      comment = ""
    }
    columns {
      name    = "status"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "subnet_id"
      type    = "string"
      comment = ""
    }
    columns {
      name    = "tags"
      type    = "map<string,string>"
      comment = ""
    }
    columns {
      name    = "vpc_id"
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
