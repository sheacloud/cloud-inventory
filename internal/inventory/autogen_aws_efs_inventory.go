package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/efs"
)

func IngestAwsEFSFileSystems(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := efs.FetchFileSystems(ctx, input)
	if resources != nil {
		err := dao.AWS().EFS().PutFileSystems(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
