package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/lambda"
)

func IngestAwsLambdaFunctions(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := lambda.FetchFunctions(ctx, input)
	if resources != nil {
		err := dao.AWS().Lambda().PutFunctions(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
