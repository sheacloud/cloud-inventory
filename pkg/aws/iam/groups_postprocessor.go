package iam

func init() {
	registerCustomGroupModelPostprocessingFunc(PostProcessGroupModel)
}

func PostProcessGroupModel(model *GroupModel) {
	if model.CreateDate != nil {
		model.CreateDateMilli = model.CreateDate.UTC().UnixMilli()
	}
}
