package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/elasticloadbalancing"
)

func IngestAwsElasticLoadBalancingLoadBalancers(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := elasticloadbalancing.FetchLoadBalancers(ctx, input)
	if resources != nil {
		err := dao.PutAwsElasticLoadBalancingLoadBalancers(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
