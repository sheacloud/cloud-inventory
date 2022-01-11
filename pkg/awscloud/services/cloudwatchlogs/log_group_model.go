package cloudwatchlogs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

type LogGroupModel struct {
	Arn               string `parquet:"name=arn,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	CreationTime      int64  `parquet:"name=creation_time,type=INT64"`
	KmsKeyId          string `parquet:"name=kms_key_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	LogGroupName      string `parquet:"name=log_group_name,type=BYTE_ARRAY,convertedtype=UTF8"`
	MetricFilterCount int32  `parquet:"name=metric_filter_count,type=INT32"`
	RetentionInDays   int32  `parquet:"name=retention_in_days,type=INT32"`
	StoredBytes       int64  `parquet:"name=stored_bytes,type=INT64"`
	AccountId         string `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region            string `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime        int64  `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

func FetchLogGroups(ctx context.Context, params *awscloud.AwsFetchInput) *awscloud.AwsFetchOutput {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int

	awsClient := params.RegionalClients[params.Region]
	cloudwatchlogsClient := awsClient.CloudWatchLogs()

	paginator := cloudwatchlogs.NewDescribeLogGroupsPaginator(cloudwatchlogsClient, &cloudwatchlogs.DescribeLogGroupsInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeLogGroups in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, var0 := range output.LogGroups {

			model := new(LogGroupModel)
			copier.Copy(&model, &var0)

			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime.UTC().UnixMilli()

			err = params.OutputFile.Write(ctx, model)
			if err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error storing LogGroup model in %s/%s: %w", params.AccountId, params.Region, err))
			}
			fetchedResources++
		}

	}

	return &awscloud.AwsFetchOutput{
		FetchingErrors:   fetchingErrors,
		FetchedResources: fetchedResources,
		FailedResources:  failedResources,
		ResourceName:     "log_groups",
		AccountId:        params.AccountId,
		Region:           params.Region,
	}
}
