package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/iam"
)

func IngestAwsIAMGroups(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := iam.FetchGroups(ctx, input)
	if resources != nil {
		err := dao.AWS().IAM().PutGroups(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsIAMPolicies(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := iam.FetchPolicies(ctx, input)
	if resources != nil {
		err := dao.AWS().IAM().PutPolicies(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsIAMRoles(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := iam.FetchRoles(ctx, input)
	if resources != nil {
		err := dao.AWS().IAM().PutRoles(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsIAMUsers(ctx context.Context, dao db.DAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := iam.FetchUsers(ctx, input)
	if resources != nil {
		err := dao.AWS().IAM().PutUsers(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
