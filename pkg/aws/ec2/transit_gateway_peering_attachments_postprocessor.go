package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func init() {
	registerCustomTransitGatewayPeeringAttachmentModelPostprocessingFunc(PostProcessTransitGatewayPeeringAttachmentModel)
}

func PostProcessTransitGatewayPeeringAttachmentModel(ctx context.Context, client *ec2.Client, cfg aws.Config, model *TransitGatewayPeeringAttachmentModel) {
	if model.CreationTime != nil {
		model.CreationTimeMilli = model.CreationTime.UTC().UnixMilli()
	}
}
