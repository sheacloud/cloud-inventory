//AUTOGENERATED CODE DO NOT EDIT
package rds

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

func FetchDBInstance(ctx context.Context, params *awscloud.AwsFetchInput) *awscloud.AwsFetchOutput {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "rds",
		Resource:   "db_instances",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime.UTC().UnixMilli(),
	}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.RDS()

	paginator := rds.NewDescribeDBInstancesPaginator(client, &rds.DescribeDBInstancesInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeDBInstances in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.DBInstances {

			model := new(DBInstance)
			copier.Copy(&model, &object)

			model.Tags = ConvertTags(object.TagList)
			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime.UTC().UnixMilli()

			if err = PostProcessDBInstance(ctx, params, model); err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error post-processing DBInstance %s %s/%s: %w", model.DBInstanceArn, params.AccountId, params.Region, err))
				failedResources++
			}

			err = params.OutputFile.Write(ctx, model)
			if err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error storing DBInstance model in %s/%s: %w", params.AccountId, params.Region, err))
			}
			fetchedResources++
		}

	}

	inventoryResults.FetchedResources = fetchedResources
	inventoryResults.FailedResources = failedResources
	inventoryResults.HadErrors = len(fetchingErrors) > 0

	return &awscloud.AwsFetchOutput{
		FetchingErrors:   fetchingErrors,
		InventoryResults: inventoryResults,
		ResourceName:     "db_instances",
		AccountId:        params.AccountId,
		Region:           params.Region,
	}
}