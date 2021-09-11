package ec2

func init() {
	registerCustomNetworkInterfaceModelPostprocessingFunc(PostProcessNetworkInterfaceModel)
}

func PostProcessNetworkInterfaceModel(model *NetworkInterfaceModel) {
	if model.Attachment != nil {
		if model.Attachment.AttachTime != nil {
			model.Attachment.AttachTimeMilli = model.Attachment.AttachTime.UTC().UnixMilli()
		}
	}
}
