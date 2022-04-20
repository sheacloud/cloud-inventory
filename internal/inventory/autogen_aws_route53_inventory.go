package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/route53"
)

func IngestAwsRoute53HostedZones(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := route53.FetchHostedZones(ctx, input)
	if resources != nil {
		err := dao.AWS().Route53().PutHostedZones(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
