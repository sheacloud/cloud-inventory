//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_fetching_file.tmpl
package efs

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

func FetchFileSystems(ctx context.Context, params *aws.AwsFetchInput) ([]*FileSystem, *aws.AwsFetchOutputMetadata) {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "efs",
		Resource:   "file_systems",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime,
	}
	resources := []*FileSystem{}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.EFS()

	paginator := efs.NewDescribeFileSystemsPaginator(client, &efs.DescribeFileSystemsInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeFileSystems in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.FileSystems {

			model := new(FileSystem)
			copier.CopyWithOption(&model, &object, aws.CopyOption)

			model.Tags = ConvertTags(object.Tags)
			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime
			model.InventoryUUID = uuid.New().String()

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
