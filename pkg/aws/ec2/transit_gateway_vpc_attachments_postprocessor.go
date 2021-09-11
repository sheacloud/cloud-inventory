package ec2

func init() {
	registerCustomTransitGatewayVpcAttachmentModelPostprocessingFunc(PostProcessTransitGatewayVpcAttachmentModel)
}

func PostProcessTransitGatewayVpcAttachmentModel(model *TransitGatewayVpcAttachmentModel) {
	if model.CreationTime != nil {
		model.CreationTimeMilli = model.CreationTime.UTC().UnixMilli()
	}
}
