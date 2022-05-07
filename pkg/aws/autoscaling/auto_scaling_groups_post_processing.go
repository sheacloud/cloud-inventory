package autoscaling

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
	"github.com/jinzhu/copier"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
)

func PostProcessAutoScalingGroup(ctx context.Context, params *localAws.AwsFetchInput, model *AutoScalingGroup) error {
	awsClient := params.RegionalClients[params.Region]
	client := awsClient.AutoScaling()

	paginator := autoscaling.NewDescribePoliciesPaginator(client, &autoscaling.DescribePoliciesInput{
		AutoScalingGroupName: aws.String(model.AutoScalingGroupName),
	})

	model.ScalingPolicies = []*ScalingPolicy{}

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)

		if err != nil {
			return fmt.Errorf("error calling DescribePolicies in %s/%s: %w", params.AccountId, params.Region, err)
		}

		for _, object := range output.ScalingPolicies {
			policyModel := new(ScalingPolicy)
			copier.CopyWithOption(&policyModel, &object, localAws.CopyOption)

			model.ScalingPolicies = append(model.ScalingPolicies, policyModel)
		}
	}

	return nil
}
