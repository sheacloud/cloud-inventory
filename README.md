# cloud-inventory

A data warehousing solution for storing periodic cloud infrastructure configuration snapshots in S3, enabling queries through AWS Athena, and custom API, and a web interface.

Differentials of given resources can be calculated between two points in time, allowing you to track changes in your environment.

For a complete list of resources ingested by cloud-inventory, see [IMPLEMENTED_RESOURCES.md](./IMPLEMENTED_RESOURCES.md)

## Testing it out

Read [the quickstart guide](./QUICKSTART_INGESTION.md) to set up a development environment and start ingesting data

Once data is ingested, you can either query the data using [Athena](QUICKSTART_ATHENA.md), the [API](QUICKSTART_API.md), or by viewing the [UI](QUICKSTART_UI.md)

## What it does

cloud-inventory scrapes cloud APIs (just AWS currently, adding Azure support in the future) and stores the response structures in various different backend database options (S3, DynamoDB, MongoDB). These scrapes can occur periodically to allow for historical querying of your cloud resources.

The databases are indexed to allow efficient querying of data for a given date, allowing efficient historical analysis.

cloud-inventory exposes an API for querying data from your desired database, enabling quick and easy access to historic information on your cloud resources across all accounts.

## How it works

cloud-inventory works on various defined data sources such as AWS EC2 instances. Each of these datasources is explicitly defined and has a mapping defined of it's API response to an internal database schema. These database schemas are auto-generated based on the cloud providers own API spec, meaning cloud-inventory ingests all possible information. In some scenarios, cloud-inventory combines multiple cloud provider APIs/resources into a single object. For example, the AWS S3 API exposes buckets and each piece of their configuration as separate API calls. cloud-inventory combines all the API responses for a given bucket into a single object, making it easy to view the full configuration of the resource in a single place.

Each time cloud-inventory runs, it queries each resource against each unique location (i.e. a given region in a given AWS account) and stores the API responses in whichever database(s) you have enabled, indexed by the date that the scrape occured.

You can have multiple databases enabled for writing data. This allows you to write data to a high-performance database like DynamoDB or MongoDB for use with the API, and also to S3 for use with Athena or other big-data tools.