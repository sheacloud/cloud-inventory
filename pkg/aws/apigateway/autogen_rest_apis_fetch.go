//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_fetching_file.tmpl
package apigateway

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

func FetchRestApis(ctx context.Context, params *aws.AwsFetchInput) ([]*RestApi, *aws.AwsFetchOutputMetadata) {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "apigateway",
		Resource:   "rest_apis",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime,
	}
	resources := []*RestApi{}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.ApiGateway()

	paginator := apigateway.NewGetRestApisPaginator(client, &apigateway.GetRestApisInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling GetRestApis in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.Items {

			model := new(RestApi)
			copier.CopyWithOption(&model, &object, aws.CopyOption)

			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime
			model.InventoryUUID = uuid.New().String()

			if err = PostProcessRestApi(ctx, params, model); err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error post-processing RestApi %s %s/%s: %w", model.Id, params.AccountId, params.Region, err))
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
