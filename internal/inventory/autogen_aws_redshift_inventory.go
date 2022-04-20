package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/redshift"
)

func IngestAwsRedshiftClusters(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := redshift.FetchClusters(ctx, input)
	if resources != nil {
		err := dao.AWS().Redshift().PutClusters(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}