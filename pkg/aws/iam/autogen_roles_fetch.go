//AUTOGENERATED CODE DO NOT EDIT
package iam

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

func FetchRoles(ctx context.Context, params *aws.AwsFetchInput) ([]*Role, *aws.AwsFetchOutputMetadata) {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "iam",
		Resource:   "roles",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime,
	}
	resources := []*Role{}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.IAM()

	paginator := iam.NewListRolesPaginator(client, &iam.ListRolesInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling ListRoles in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.Roles {

			model := new(Role)
			copier.CopyWithOption(&model, &object, aws.CopyOption)

			model.Tags = ConvertTags(object.Tags)
			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime
			model.InventoryUUID = uuid.New().String()

			if err = PostProcessRole(ctx, params, model); err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error post-processing Role %s %s/%s: %w", model.RoleId, params.AccountId, params.Region, err))
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
