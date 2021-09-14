package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func init() {
	registerCustomVpcEndpointModelPostprocessingFunc(PostProcessVpcEndpointModel)
}

func PostProcessVpcEndpointModel(ctx context.Context, client *ec2.Client, cfg aws.Config, model *VpcEndpointModel) {
	if model.CreationTimestamp != nil {
		model.CreationTimestampMilli = model.CreationTimestamp.UTC().UnixMilli()
	}
}
