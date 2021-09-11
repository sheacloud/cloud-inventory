package ec2

func init() {
	registerCustomTransitGatewayRouteTableModelPostprocessingFunc(PostProcessTransitGatewayRouteTableModel)
}

func PostProcessTransitGatewayRouteTableModel(model *TransitGatewayRouteTableModel) {
	if model.CreationTime != nil {
		model.CreationTimeMilli = model.CreationTime.UTC().UnixMilli()
	}
}
