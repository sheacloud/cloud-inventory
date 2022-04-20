package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/ecs"
)

func IngestAwsECSClusters(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ecs.FetchClusters(ctx, input)
	if resources != nil {
		err := dao.AWS().ECS().PutClusters(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsECSServices(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ecs.FetchServices(ctx, input)
	if resources != nil {
		err := dao.AWS().ECS().PutServices(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsECSTasks(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ecs.FetchTasks(ctx, input)
	if resources != nil {
		err := dao.AWS().ECS().PutTasks(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
