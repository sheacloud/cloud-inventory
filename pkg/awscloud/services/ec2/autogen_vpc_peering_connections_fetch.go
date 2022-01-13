//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func FetchVpcPeeringConnection(ctx context.Context, params *awscloud.AwsFetchInput) *awscloud.AwsFetchOutput {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.EC2()

	paginator := ec2.NewDescribeVpcPeeringConnectionsPaginator(client, &ec2.DescribeVpcPeeringConnectionsInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeVpcPeeringConnections in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.VpcPeeringConnections {

			model := new(VpcPeeringConnection)
			copier.Copy(&model, &object)

			model.Tags = ConvertTags(object.Tags)
			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime.UTC().UnixMilli()

			err = params.OutputFile.Write(ctx, model)
			if err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error storing VpcPeeringConnection model in %s/%s: %w", params.AccountId, params.Region, err))
			}
			fetchedResources++
		}

	}

	return &awscloud.AwsFetchOutput{
		FetchingErrors:   fetchingErrors,
		FetchedResources: fetchedResources,
		FailedResources:  failedResources,
		ResourceName:     "vpc_peering_connections",
		AccountId:        params.AccountId,
		Region:           params.Region,
	}
}
