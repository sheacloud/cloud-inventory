//AUTOGENERATED CODE DO NOT EDIT
package lambda

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

func FetchFunctionConfiguration(ctx context.Context, params *awscloud.AwsFetchInput) *awscloud.AwsFetchOutput {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "lambda",
		Resource:   "filesystems",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime.UTC().UnixMilli(),
	}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.Lambda()

	paginator := lambda.NewListFunctionsPaginator(client, &lambda.ListFunctionsInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling ListFunctions in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.Functions {

			model := new(FunctionConfiguration)
			copier.Copy(&model, &object)

			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime.UTC().UnixMilli()

			err = params.OutputFile.Write(ctx, model)
			if err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error storing FunctionConfiguration model in %s/%s: %w", params.AccountId, params.Region, err))
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
		ResourceName:     "filesystems",
		AccountId:        params.AccountId,
		Region:           params.Region,
	}
}
