//AUTOGENERATED CODE DO NOT EDIT
package apigatewayv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

func FetchApis(ctx context.Context, params *aws.AwsFetchInput) ([]*Api, *aws.AwsFetchOutputMetadata) {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "apigatewayv2",
		Resource:   "apis",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime,
	}
	resources := []*Api{}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.ApiGatewayV2()

	result, err := client.GetApis(ctx, &apigatewayv2.GetApisInput{})
	if err != nil {
		fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling GetApis in %s/%s: %w", params.AccountId, params.Region, err))
		inventoryResults.FetchedResources = 0
		inventoryResults.FailedResources = 0
		inventoryResults.HadErrors = true
		return nil, &aws.AwsFetchOutputMetadata{
			FetchingErrors:   fetchingErrors,
			InventoryResults: inventoryResults,
		}
	}

	results := []*apigatewayv2.GetApisOutput{result}
	for _, output := range results {

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling GetApis in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.Items {

			model := new(Api)
			copier.Copy(&model, &object)

			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime
			model.InventoryUUID = uuid.New().String()

			if err = PostProcessApi(ctx, params, model); err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error post-processing Api %s %s/%s: %w", model.ApiId, params.AccountId, params.Region, err))
				failedResources++
			}

			resources = append(resources, model)
			fetchedResources++
		}

	}

	inventoryResults.FetchedResources = fetchedResources
	inventoryResults.FailedResources = failedResources
	inventoryResults.HadErrors = len(fetchingErrors) > 0

	return resources, &aws.AwsFetchOutputMetadata{
		FetchingErrors:   fetchingErrors,
		InventoryResults: inventoryResults,
	}
}
