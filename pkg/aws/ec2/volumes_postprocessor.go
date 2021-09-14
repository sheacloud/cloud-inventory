package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func init() {
	registerCustomVolumeModelPostprocessingFunc(PostProcessVolumeModel)
}

func PostProcessVolumeModel(ctx context.Context, client *ec2.Client, cfg aws.Config, model *VolumeModel) {
	if model.CreateTime != nil {
		model.CreateTimeMilli = model.CreateTime.UTC().UnixMilli()
	}

	for _, attachment := range model.Attachments {
		if attachment.AttachTime != nil {
			attachment.AttachTimeMilli = attachment.AttachTime.UTC().UnixMilli()
		}
	}
}
