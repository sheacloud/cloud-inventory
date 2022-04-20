package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/cloudtrail"
)

func IngestAwsCloudTrailTrails(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := cloudtrail.FetchTrails(ctx, input)
	if resources != nil {
		err := dao.AWS().CloudTrail().PutTrails(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
