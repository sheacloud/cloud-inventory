package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/storagegateway"
)

func IngestAwsStorageGatewayGateways(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := storagegateway.FetchGateways(ctx, input)
	if resources != nil {
		err := dao.PutAwsStorageGatewayGateways(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
