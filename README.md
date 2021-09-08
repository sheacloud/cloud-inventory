# cloud-inventory

A data warehousing solution for storing periodic AWS infrastructure configuration exports in S3, leveraging AWS Athena for querying.

## What it does

cloud-inventory scrapes cloud APIs (just AWS currently, intend on adding Azure support soon) and stores the response structures in S3 in a format that is easily queryable by Athena. These scrapes occur periodically to allow for historical querying of your cloud resources.

## How it works

cloud-inventory works on various defined data sources such as AWS EC2 instances. Each of these datasources is explicitly defined and has a mapping defined of it's API resposne to a Parquet format, such as in [instances.go](./pkg/aws/ec2/instances.go).

Each time cloud-inventory runs, it fethces queries each datasource against each unique location (i.e. a given region in a given AWS account) and stores the API responses in parquet files in S3, partitioned by the date that the scrape occured. There are auto-generated views created in Athena to allow for easy querying of this data depending on what you are looking for (all unique resources, just resources found in the most recent scrape, etc).

## Terraform/View auto-generation

Some of the code in this repository is auto-generated by helper programs, [terraform-generator](./cmd/terraform-generator/terraform_generator.go) and [view-generator](./cmd/view-generator/view_generator.go). These programs generate the terraform files which define the AWS Glue tables and the Athena view SQL statements, respectively. This is based on reflection of the Go structures that comprise each datasource.