package exampleservice

import (
	"context"

	"github.com/sheacloud/cloud-inventory/pkg/examplecloud"
)

type ExampleResource struct {
	ID         string `parquet:"name=id,type=BYTE_ARRAY,convertedtype=UTF8"`
	AccountID  string `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime int64  `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

func FetchExampleResource(ctx context.Context, params *examplecloud.ExampleCloudFetchInput) *examplecloud.ExampleCloudFetchOutput {
	errors := []error{}
	for _, i := range []string{"1", "2", "3"} {
		resource := ExampleResource{
			ID:         i,
			AccountID:  params.AccountID,
			ReportTime: params.ReportTime.UTC().UnixMilli(),
		}

		err := params.OutputFile.Write(ctx, resource)
		if err != nil {
			errors = append(errors, err)
		}
	}

	return &examplecloud.ExampleCloudFetchOutput{
		FetchingErrors:   errors,
		FetchedResources: 3,
		FailedResources:  0,
	}
}
