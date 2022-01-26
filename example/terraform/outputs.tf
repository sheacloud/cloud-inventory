output "s3_bucket_name" {
    value = aws_s3_bucket.cloud_inventory.bucket
}

output "glue_database_name" {
    value = aws_glue_catalog_database.cloud_inventory.name
}

output "athena_workgroup_name" {
    value = aws_athena_workgroup.cloud_inventory.name
}