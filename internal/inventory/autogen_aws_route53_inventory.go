package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/route53"
)

func IngestAwsRoute53HostedZones(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := route53.FetchHostedZones(ctx, input)
	if resources != nil {
		err := dao.PutAwsRoute53HostedZones(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
