package awscloud

import (
	"time"

	"github.com/sheacloud/cloud-inventory/internal/indexedstorage"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud/interfaces"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

type AwsFetchInput struct {
	AccountId       string
	Region          string
	RegionalClients map[string]interfaces.AwsClient
	ReportTime      time.Time
	OutputFile      *indexedstorage.ParquetS3File //TODO abstract this away to an interface
}

type AwsFetchOutput struct {
	AccountId        string
	Region           string
	ResourceName     string
	FetchingErrors   []error
	InventoryResults *meta.InventoryResults
}
