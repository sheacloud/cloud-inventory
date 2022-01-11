package ec2

import (
	"context"
	"fmt"

	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessVolume(ctx context.Context, params *awscloud.AwsFetchInput, model *VolumeModel) error {
	if model.CreateTime != nil {
		model.CreateTimeMilli = model.CreateTime.UTC().UnixMilli()
	}

	for _, attachment := range model.Attachments {
		if attachment.AttachTime != nil {
			attachment.AttachTimeMilli = attachment.AttachTime.UTC().UnixMilli()
		}
	}

	return fmt.Errorf("test")
}
