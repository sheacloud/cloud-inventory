//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_fetching_file.tmpl
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

func FetchVpnGateways(ctx context.Context, params *aws.AwsFetchInput) ([]*VpnGateway, *aws.AwsFetchOutputMetadata) {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "ec2",
		Resource:   "vpn_gateways",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime,
	}
	resources := []*VpnGateway{}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.EC2()

	result, err := client.DescribeVpnGateways(ctx, &ec2.DescribeVpnGatewaysInput{})
	if err != nil {
		fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeVpnGateways in %s/%s: %w", params.AccountId, params.Region, err))
		inventoryResults.FetchedResources = 0
		inventoryResults.FailedResources = 0
		inventoryResults.HadErrors = true
		return nil, &aws.AwsFetchOutputMetadata{
			FetchingErrors:   fetchingErrors,
			InventoryResults: inventoryResults,
		}
	}

	results := []*ec2.DescribeVpnGatewaysOutput{result}
	for _, output := range results {

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeVpnGateways in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.VpnGateways {

			model := new(VpnGateway)
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
