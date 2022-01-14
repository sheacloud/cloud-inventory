package elasticloadbalancing

import (
	"context"

	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessLoadBalancerDescription(ctx context.Context, params *awscloud.AwsFetchInput, model *LoadBalancerDescription) error {
	if model.CreatedTime != nil {
		model.CreatedTimeMilli = model.CreatedTime.UTC().UnixMilli()
	}

	return nil
}
