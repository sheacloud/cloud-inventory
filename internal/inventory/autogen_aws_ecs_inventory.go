package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/ecs"
)

func IngestAwsECSClusters(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ecs.FetchClusters(ctx, input)
	if resources != nil {
		err := dao.PutAwsECSClusters(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsECSServices(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ecs.FetchServices(ctx, input)
	if resources != nil {
		err := dao.PutAwsECSServices(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsECSTasks(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ecs.FetchTasks(ctx, input)
	if resources != nil {
		err := dao.PutAwsECSTasks(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
