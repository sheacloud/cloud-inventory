package iam

func init() {
	registerCustomPolicyModelPostprocessingFunc(PostProcessPolicyModel)
}

func PostProcessPolicyModel(model *PolicyModel) {
	if model.CreateDate != nil {
		model.CreateDateMilli = model.CreateDate.UTC().UnixMilli()
	}

	if model.UpdateDate != nil {
		model.UpdateDateMilli = model.UpdateDate.UTC().UnixMilli()
	}
}
