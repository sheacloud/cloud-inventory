package examplecloud

import (
	"time"

	"github.com/sheacloud/cloud-inventory/internal/indexedstorage"
)

type ExampleCloudFetchInput struct {
	AccountID  string
	Client     ExampleClient
	ReportTime time.Time
	OutputFile *indexedstorage.ParquetS3File
}

type ExampleCloudFetchOutput struct {
	FetchingErrors   []error
	FetchedResources int
	FailedResources  int
}
