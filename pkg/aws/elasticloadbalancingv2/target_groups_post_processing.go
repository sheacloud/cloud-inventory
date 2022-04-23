package elasticloadbalancingv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/jinzhu/copier"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
)

func PostProcessTargetGroup(ctx context.Context, params *localAws.AwsFetchInput, model *TargetGroup) error {
	awsClient := params.RegionalClients[params.Region]
	client := awsClient.ElasticLoadBalancingV2()

	result, err := client.DescribeTags(ctx, &elasticloadbalancingv2.DescribeTagsInput{
		ResourceArns: []string{model.TargetGroupArn},
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

	targetResult, err := client.DescribeTargetHealth(ctx, &elasticloadbalancingv2.DescribeTargetHealthInput{
		TargetGroupArn: aws.String(model.TargetGroupArn),
	})
	if err != nil {
		return fmt.Errorf("error calling DescribeTargetHealth: %w", err)
	}
	for _, target := range targetResult.TargetHealthDescriptions {
		targetModel := new(TargetHealthDescription)
		copier.CopyWithOption(targetModel, &target, localAws.CopyOption)
		model.Targets = append(model.Targets, targetModel)
	}
	return nil
}
