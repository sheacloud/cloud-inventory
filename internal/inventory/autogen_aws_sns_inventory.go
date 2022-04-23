package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/sns"
)

func IngestAwsSNSTopics(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := sns.FetchTopics(ctx, input)
	if resources != nil {
		err := dao.PutAwsSNSTopics(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsSNSSubscriptions(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := sns.FetchSubscriptions(ctx, input)
	if resources != nil {
		err := dao.PutAwsSNSSubscriptions(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
