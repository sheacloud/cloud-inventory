package ec2

import (
	"context"

	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessNetworkInterface(ctx context.Context, params *awscloud.AwsFetchInput, model *NetworkInterface) error {
	if model.Attachment != nil {
		if model.Attachment.AttachTime != nil {
			model.Attachment.AttachTimeMilli = model.Attachment.AttachTime.UTC().UnixMilli()
		}
	}

	return nil
}
