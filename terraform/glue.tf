resource "aws_glue_catalog_database" "data_warehouse" {
  name = "datawarehouse"
}

module "ec2" {
  source = "./services/ec2"
  glue_database_name = aws_glue_catalog_database.data_warehouse.name
}