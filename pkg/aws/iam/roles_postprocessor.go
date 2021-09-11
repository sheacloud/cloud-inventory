package iam

func init() {
	registerCustomRoleModelPostprocessingFunc(PostProcessRoleModel)
}

func PostProcessRoleModel(model *RoleModel) {
	if model.CreateDate != nil {
		model.CreateDateMilli = model.CreateDate.UTC().UnixMilli()
	}

	if model.RoleLastUsed != nil {
		if model.RoleLastUsed.LastUsedDate != nil {
			model.RoleLastUsed.LastUsedDateMilli = model.RoleLastUsed.LastUsedDate.UTC().UnixMilli()
		}
	}
}
