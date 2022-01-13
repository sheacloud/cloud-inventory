package ec2

import (
	"context"

	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessTransitGateway(ctx context.Context, params *awscloud.AwsFetchInput, model *TransitGateway) error {
	if model.CreationTime != nil {
		model.CreationTimeMilli = model.CreationTime.UTC().UnixMilli()
	}

	return nil
}
