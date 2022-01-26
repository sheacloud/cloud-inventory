# cloud-inventory

A data warehousing solution for storing periodic cloud infrastructure configuration snapshots in S3, enabling queries through AWS Athena or a custom API.

Differentials of given resources can be calculated between two points in time, allowing you to track changes in your environment.

For a complete list of resources ingested by cloud-inventory, see [IMPLEMENTED_RESOURCES.md](./IMPLEMENTED_RESOURCES.md)

## Testing it out

Read [the quickstart guide](./QUICKSTART_INGESTION.md) to set up a development environment and start ingesting data

Once data is ingested into S3, you can either query the data using [Athena](QUICKSTART_ATHENA.md) or using the [API](QUICKSTART_API.md)

## What it does

cloud-inventory scrapes cloud APIs (just AWS currently, intend on adding Azure support soon) and stores the response structures in S3 in Apache Parquet format, indexed by date. These scrapes can occur periodically to allow for historical querying of your cloud resources.

## How it works

cloud-inventory works on various defined data sources such as AWS EC2 instances. Each of these datasources is explicitly defined and has a mapping defined of it's API resposne to a Parquet format.

Each time cloud-inventory runs, it queries each datasource against each unique location (i.e. a given region in a given AWS account) and stores the API responses in parquet files in S3, partitioned by the date that the scrape occured. There are auto-generated views created in Athena to allow for easy querying of this data depending on what you are looking for (all unique resources, just resources found in the most recent scrape, etc).

In addition to querying with Athena, there is an API which allows retrieving the data for a given time period, or even performing diffs of resources between two points in time.