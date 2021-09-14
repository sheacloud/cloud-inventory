package elasticloadbalancingv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/sirupsen/logrus"
)

func init() {
	registerCustomLoadBalancerModelPostprocessingFunc(PostProcessLoadBalancerModel)
}

func PostProcessLoadBalancerModel(ctx context.Context, client *elasticloadbalancingv2.Client, cfg aws.Config, model *LoadBalancerModel) {
	if model.CreatedTime != nil {
		model.CreatedTimeMilli = model.CreatedTime.UTC().UnixMilli()
	}

	result, err := client.DescribeTags(ctx, &elasticloadbalancingv2.DescribeTagsInput{
		ResourceArns: []string{model.LoadBalancerArn},
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "elasticloadbalancingv2",
			"data_source": "load_balancers",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling DescribeTags")
		return
	}
	model.Tags = map[string]string{}
	for _, descriptions := range result.TagDescriptions {
		for _, tag := range descriptions.Tags {
			model.Tags[*tag.Key] = *tag.Value
		}
	}
}
