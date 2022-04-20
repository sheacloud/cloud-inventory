//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

func FetchManagedPrefixLists(ctx context.Context, params *aws.AwsFetchInput) ([]*ManagedPrefixList, *aws.AwsFetchOutputMetadata) {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "ec2",
		Resource:   "managed_prefix_lists",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime,
	}
	resources := []*ManagedPrefixList{}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.EC2()

	paginator := ec2.NewDescribeManagedPrefixListsPaginator(client, &ec2.DescribeManagedPrefixListsInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeManagedPrefixLists in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.PrefixLists {

			model := new(ManagedPrefixList)
			copier.Copy(&model, &object)

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