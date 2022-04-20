package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/dynamodb"
)

func IngestAwsDynamoDBTables(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := dynamodb.FetchTables(ctx, input)
	if resources != nil {
		err := dao.AWS().DynamoDB().PutTables(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
