//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func FetchPlacementGroup(ctx context.Context, params *awscloud.AwsFetchInput) *awscloud.AwsFetchOutput {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.EC2()

	result, err := client.DescribePlacementGroups(ctx, &ec2.DescribePlacementGroupsInput{})
	if err != nil {
		fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribePlacementGroups in %s/%s: %w", params.AccountId, params.Region, err))
		return &awscloud.AwsFetchOutput{
			FetchingErrors:   fetchingErrors,
			FetchedResources: fetchedResources,
			FailedResources:  failedResources,
			ResourceName:     "placement_groups",
			AccountId:        params.AccountId,
			Region:           params.Region,
		}
	}

	results := []*ec2.DescribePlacementGroupsOutput{result}
	for _, output := range results {

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribePlacementGroups in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.PlacementGroups {

			model := new(PlacementGroup)
			copier.Copy(&model, &object)

			model.Tags = ConvertTags(object.Tags)
			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime.UTC().UnixMilli()

			err = params.OutputFile.Write(ctx, model)
			if err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error storing PlacementGroup model in %s/%s: %w", params.AccountId, params.Region, err))
			}
			fetchedResources++
		}

	}

	return &awscloud.AwsFetchOutput{
		FetchingErrors:   fetchingErrors,
		FetchedResources: fetchedResources,
		FailedResources:  failedResources,
		ResourceName:     "placement_groups",
		AccountId:        params.AccountId,
		Region:           params.Region,
	}
}
