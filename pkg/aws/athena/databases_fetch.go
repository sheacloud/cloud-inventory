//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_fetching_file.tmpl
package athena

import (
	"context"
	"fmt"

	awsSdk "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

func FetchDatabases(ctx context.Context, params *aws.AwsFetchInput) ([]*Database, *aws.AwsFetchOutputMetadata) {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "athena",
		Resource:   "databases",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime,
	}
	resources := []*Database{}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.Athena()

	catalogPaginator := athena.NewListDataCatalogsPaginator(client, &athena.ListDataCatalogsInput{})

	for catalogPaginator.HasMorePages() {
		output, err := catalogPaginator.NextPage(ctx)

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling ListDataCatalogs in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, catalog := range output.DataCatalogsSummary {

			paginator := athena.NewListDatabasesPaginator(client, &athena.ListDatabasesInput{
				CatalogName: awsSdk.String(*catalog.CatalogName),
			})

			for paginator.HasMorePages() {
				output, err := paginator.NextPage(ctx)

				if err != nil {
					fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling ListDatabases in %s/%s: %w", params.AccountId, params.Region, err))
					break
				}

				for _, object := range output.DatabaseList {

					model := new(Database)
					copier.CopyWithOption(&model, &object, aws.CopyOption)

					model.DataCatalog = *catalog.CatalogName

					model.AccountId = params.AccountId
					model.Region = params.Region
					model.ReportTime = params.ReportTime
					model.InventoryUUID = uuid.New().String()

					resources = append(resources, model)
					fetchedResources++
				}

			}
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
