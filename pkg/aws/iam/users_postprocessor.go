package iam

func init() {
	registerCustomUserModelPostprocessingFunc(PostProcessUserModel)
}

func PostProcessUserModel(model *UserModel) {
	if model.CreateDate != nil {
		model.CreateDateMilli = model.CreateDate.UTC().UnixMilli()
	}

	if model.PasswordLastUsed != nil {
		model.PasswordLastUsedMilli = model.PasswordLastUsed.UTC().UnixMilli()
	}
}
