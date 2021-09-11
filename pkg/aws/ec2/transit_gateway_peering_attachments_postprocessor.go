package ec2

func init() {
	registerCustomTransitGatewayPeeringAttachmentModelPostprocessingFunc(PostProcessTransitGatewayPeeringAttachmentModel)
}

func PostProcessTransitGatewayPeeringAttachmentModel(model *TransitGatewayPeeringAttachmentModel) {
	if model.CreationTime != nil {
		model.CreationTimeMilli = model.CreationTime.UTC().UnixMilli()
	}
}
