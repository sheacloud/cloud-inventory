//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_service_inventory_file.tmpl
package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/autoscaling"
)

func IngestAwsAutoScalingAutoScalingGroups(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := autoscaling.FetchAutoScalingGroups(ctx, input)
	if resources != nil {
		err := dao.PutAwsAutoScalingAutoScalingGroups(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsAutoScalingLaunchConfigurations(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := autoscaling.FetchLaunchConfigurations(ctx, input)
	if resources != nil {
		err := dao.PutAwsAutoScalingLaunchConfigurations(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}