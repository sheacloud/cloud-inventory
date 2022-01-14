package elasticloadbalancingv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessLoadBalancer(ctx context.Context, params *awscloud.AwsFetchInput, model *LoadBalancer) error {
	if model.CreatedTime != nil {
		model.CreatedTimeMilli = model.CreatedTime.UTC().UnixMilli()
	}

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
