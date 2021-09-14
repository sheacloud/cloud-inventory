package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func init() {
	registerCustomTransitGatewayRouteTableModelPostprocessingFunc(PostProcessTransitGatewayRouteTableModel)
}

func PostProcessTransitGatewayRouteTableModel(ctx context.Context, client *ec2.Client, cfg aws.Config, model *TransitGatewayRouteTableModel) {
	if model.CreationTime != nil {
		model.CreationTimeMilli = model.CreationTime.UTC().UnixMilli()
	}
}
