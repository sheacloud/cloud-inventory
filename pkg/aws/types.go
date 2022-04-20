package aws

import (
	"time"

	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

type AwsFetchInput struct {
	AccountId       string
	Region          string
	RegionalClients map[string]AwsClientInterface
	ReportTime      time.Time
}

type AwsFetchOutputMetadata struct {
	FetchingErrors   []error
	InventoryResults *meta.InventoryResults
}
