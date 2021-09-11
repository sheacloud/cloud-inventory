package ec2

func init() {
	registerCustomNatGatewayModelPostprocessingFunc(PostProcessNatGatewayModel)
}

func PostProcessNatGatewayModel(model *NatGatewayModel) {
	if model.CreateTime != nil {
		model.CreateTimeMilli = model.CreateTime.UTC().UnixMilli()
	}
	if model.DeleteTime != nil {
		model.DeleteTimeMilli = model.DeleteTime.UTC().UnixMilli()
	}

	if model.ProvisionedBandwidth != nil {
		if model.ProvisionedBandwidth.ProvisionTime != nil {
			model.ProvisionedBandwidth.ProvisionTimeMilli = model.ProvisionedBandwidth.ProvisionTime.UTC().UnixMilli()
		}
		if model.ProvisionedBandwidth.RequestTime != nil {
			model.ProvisionedBandwidth.RequestTimeMilli = model.ProvisionedBandwidth.RequestTime.UTC().UnixMilli()
		}
	}
}
