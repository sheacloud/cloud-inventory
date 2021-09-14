package elasticloadbalancing

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
)

func init() {
	registerCustomLoadBalancerDescriptionModelPostprocessingFunc(PostProcessLoadBalancerDescriptionModel)
}

func PostProcessLoadBalancerDescriptionModel(ctx context.Context, client *elasticloadbalancing.Client, cfg aws.Config, model *LoadBalancerDescriptionModel) {
	if model.CreatedTime != nil {
		model.CreatedTimeMilli = model.CreatedTime.UTC().UnixMilli()
	}
}
