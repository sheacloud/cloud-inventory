package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/rds"
)

func IngestAwsRDSDBClusters(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := rds.FetchDBClusters(ctx, input)
	if resources != nil {
		err := dao.PutAwsRDSDBClusters(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsRDSDBInstances(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := rds.FetchDBInstances(ctx, input)
	if resources != nil {
		err := dao.PutAwsRDSDBInstances(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
