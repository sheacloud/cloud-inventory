//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

func FetchReservedInstances(ctx context.Context, params *awscloud.AwsFetchInput) *awscloud.AwsFetchOutput {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud: "aws",
		Service: "ec2",
		Resource: "reserved_instances",
		AccountId: params.AccountId,
		Region: params.Region,
		ReportTime: params.ReportTime.UTC().UnixMilli(),
	}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.EC2()

	

	result, err := client.DescribeReservedInstances(ctx, &ec2.DescribeReservedInstancesInput{})
	if err != nil {
		fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeReservedInstances in %s/%s: %w", params.AccountId, params.Region, err))
		inventoryResults.FetchedResources = 0
		inventoryResults.FailedResources = 0
		inventoryResults.HadErrors = true
		return &awscloud.AwsFetchOutput{
			FetchingErrors:   fetchingErrors,
			InventoryResults: inventoryResults,
			ResourceName:     "reserved_instances",
			AccountId:        params.AccountId,
			Region:           params.Region,
		}
	}

	results := []*ec2.DescribeReservedInstancesOutput{result}
	for _, output := range results {
	
		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeReservedInstances in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.ReservedInstances {

			model := new(ReservedInstances)
			copier.Copy(&model, &object)

			model.Tags = ConvertTags(object.Tags)
			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime.UTC().UnixMilli()

			

			err = params.OutputFile.Write(ctx, model)
			if err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error storing ReservedInstances model in %s/%s: %w", params.AccountId, params.Region, err))
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
		ResourceName:     "reserved_instances",
		AccountId:        params.AccountId,
		Region:           params.Region,
	}
}