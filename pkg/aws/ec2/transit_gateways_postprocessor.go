package ec2

func init() {
	registerCustomTransitGatewayModelPostprocessingFunc(PostProcessTransitGatewayModel)
}

func PostProcessTransitGatewayModel(model *TransitGatewayModel) {
	if model.CreationTime != nil {
		model.CreationTimeMilli = model.CreationTime.UTC().UnixMilli()
	}
}
