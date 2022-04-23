package elasticloadbalancingv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
)

func PostProcessLoadBalancer(ctx context.Context, params *localAws.AwsFetchInput, model *LoadBalancer) error {
	awsClient := params.RegionalClients[params.Region]
	client := awsClient.ElasticLoadBalancingV2()

	result, err := client.DescribeTags(ctx, &elasticloadbalancingv2.DescribeTagsInput{
		ResourceArns: []string{model.LoadBalancerArn},
	})
	if err != nil {
		return fmt.Errorf("error calling DescribeTags: %w", err)
	}
	model.Tags = map[string]string{}
	for _, descriptions := range result.TagDescriptions {
		for _, tag := range descriptions.Tags {
			model.Tags[*tag.Key] = *tag.Value
		}
	}

	return nil
}
