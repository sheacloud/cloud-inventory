package aws

import (
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

type AwsFetchInput struct {
	AccountId       string
	Region          string
	RegionalClients map[string]AwsClientInterface
	ReportTime      int64
}

type AwsFetchOutputMetadata struct {
	FetchingErrors   []error
	InventoryResults *meta.InventoryResults
}
