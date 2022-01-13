//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func FetchInstanceTypeInfo(ctx context.Context, params *awscloud.AwsFetchInput) *awscloud.AwsFetchOutput {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.EC2()

	paginator := ec2.NewDescribeInstanceTypesPaginator(client, &ec2.DescribeInstanceTypesInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeInstanceTypes in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.InstanceTypes {

			model := new(InstanceTypeInfo)
			copier.Copy(&model, &object)

			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime.UTC().UnixMilli()

			err = params.OutputFile.Write(ctx, model)
			if err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error storing InstanceTypeInfo model in %s/%s: %w", params.AccountId, params.Region, err))
			}
			fetchedResources++
		}

	}

	return &awscloud.AwsFetchOutput{
		FetchingErrors:   fetchingErrors,
		FetchedResources: fetchedResources,
		FailedResources:  failedResources,
		ResourceName:     "instance_types",
		AccountId:        params.AccountId,
		Region:           params.Region,
	}
}
