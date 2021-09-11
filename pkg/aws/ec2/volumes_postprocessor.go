package ec2

func init() {
	registerCustomVolumeModelPostprocessingFunc(PostProcessVolumeModel)
}

func PostProcessVolumeModel(model *VolumeModel) {
	if model.CreateTime != nil {
		model.CreateTimeMilli = model.CreateTime.UTC().UnixMilli()
	}

	for _, attachment := range model.Attachments {
		if attachment.AttachTime != nil {
			attachment.AttachTimeMilli = attachment.AttachTime.UTC().UnixMilli()
		}
	}
}
