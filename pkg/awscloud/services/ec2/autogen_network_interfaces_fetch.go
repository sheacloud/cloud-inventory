//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func FetchNetworkInterface(ctx context.Context, params *awscloud.AwsFetchInput) *awscloud.AwsFetchOutput {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.EC2()

	paginator := ec2.NewDescribeNetworkInterfacesPaginator(client, &ec2.DescribeNetworkInterfacesInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeNetworkInterfaces in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.NetworkInterfaces {

			model := new(NetworkInterface)
			copier.Copy(&model, &object)

			model.Tags = ConvertTags(object.TagSet)
			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime.UTC().UnixMilli()

			if err = PostProcessNetworkInterface(ctx, params, model); err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error post-processing NetworkInterface %s %s/%s: %w", model.NetworkInterfaceId, params.AccountId, params.Region, err))
				failedResources++
			}

			err = params.OutputFile.Write(ctx, model)
			if err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error storing NetworkInterface model in %s/%s: %w", params.AccountId, params.Region, err))
			}
			fetchedResources++
		}

	}

	return &awscloud.AwsFetchOutput{
		FetchingErrors:   fetchingErrors,
		FetchedResources: fetchedResources,
		FailedResources:  failedResources,
		ResourceName:     "network_interfaces",
		AccountId:        params.AccountId,
		Region:           params.Region,
	}
}
