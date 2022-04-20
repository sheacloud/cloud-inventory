package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/apigatewayv2"
)

func IngestAwsApiGatewayV2Apis(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := apigatewayv2.FetchApis(ctx, input)
	if resources != nil {
		err := dao.AWS().ApiGatewayV2().PutApis(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
