package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/backup"
)

func IngestAwsBackupBackupVaults(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := backup.FetchBackupVaults(ctx, input)
	if resources != nil {
		err := dao.PutAwsBackupBackupVaults(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsBackupBackupPlans(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := backup.FetchBackupPlans(ctx, input)
	if resources != nil {
		err := dao.PutAwsBackupBackupPlans(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
