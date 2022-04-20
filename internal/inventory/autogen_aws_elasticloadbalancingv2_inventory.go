package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/elasticloadbalancingv2"
)

func IngestAwsElasticLoadBalancingV2LoadBalancers(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := elasticloadbalancingv2.FetchLoadBalancers(ctx, input)
	if resources != nil {
		err := dao.AWS().ElasticLoadBalancingV2().PutLoadBalancers(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsElasticLoadBalancingV2TargetGroups(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := elasticloadbalancingv2.FetchTargetGroups(ctx, input)
	if resources != nil {
		err := dao.AWS().ElasticLoadBalancingV2().PutTargetGroups(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
