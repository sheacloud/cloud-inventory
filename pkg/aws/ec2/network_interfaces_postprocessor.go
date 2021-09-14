package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func init() {
	registerCustomNetworkInterfaceModelPostprocessingFunc(PostProcessNetworkInterfaceModel)
}

func PostProcessNetworkInterfaceModel(ctx context.Context, client *ec2.Client, cfg aws.Config, model *NetworkInterfaceModel) {
	if model.Attachment != nil {
		if model.Attachment.AttachTime != nil {
			model.Attachment.AttachTimeMilli = model.Attachment.AttachTime.UTC().UnixMilli()
		}
	}
}
