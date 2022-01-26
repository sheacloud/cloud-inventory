package iam

import (
	"context"

	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessPolicy(ctx context.Context, params *awscloud.AwsFetchInput, model *Policy) error {
	if model.CreateDate != nil {
		model.CreateDateMilli = model.CreateDate.UTC().UnixMilli()
	}

	if model.UpdateDate != nil {
		model.UpdateDateMilli = model.UpdateDate.UTC().UnixMilli()
	}
	return nil
}
