package ec2

func init() {
	registerCustomVpcEndpointModelPostprocessingFunc(PostProcessVpcEndpointModel)
}

func PostProcessVpcEndpointModel(model *VpcEndpointModel) {
	if model.CreationTimestamp != nil {
		model.CreationTimestampMilli = model.CreationTimestamp.UTC().UnixMilli()
	}
}
