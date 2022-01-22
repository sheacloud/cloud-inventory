# Querying the data via the API

This guide assumes you have already followed the steps outlined in [QUICKSTART_INGESTION.md](QUICKSTART_INGESTION.md)

1. Run the API server locally
```
go run ./cmd/cloud-inventory-api
```

## Query via Curl
```
curl http://localhost:8080/api/v1/inventory/aws/cloudwatchlogs/log_groups
```

## Query via Swagger UI
1. Connect to the Swagger UI endpoint at http://localhost:8080/swagger/index.html
2. Click on a service endpoint and hit "Try it out"
3. Enter any filters you want (none are required for inventory fetches) and hit "Execute"