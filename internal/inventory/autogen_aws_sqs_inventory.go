package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/sqs"
)

func IngestAwsSQSQueues(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := sqs.FetchQueues(ctx, input)
	if resources != nil {
		err := dao.AWS().SQS().PutQueues(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
