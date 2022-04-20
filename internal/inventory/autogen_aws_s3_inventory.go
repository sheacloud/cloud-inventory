package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/s3"
)

func IngestAwsS3Buckets(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := s3.FetchBuckets(ctx, input)
	if resources != nil {
		err := dao.AWS().S3().PutBuckets(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
