package iam

func init() {
	registerCustomInstanceProfileModelPostprocessingFunc(PostProcessInstanceProfileModel)
}

func PostProcessInstanceProfileModel(model *InstanceProfileModel) {
	if model.CreateDate != nil {
		model.CreateDateMilli = model.CreateDate.UTC().UnixMilli()
	}

	for _, role := range model.Roles {
		if role.CreateDate != nil {
			role.CreateDateMilli = role.CreateDate.UTC().UnixMilli()
		}

		if role.RoleLastUsed != nil {
			if role.RoleLastUsed.LastUsedDate != nil {
				role.RoleLastUsed.LastUsedDateMilli = role.RoleLastUsed.LastUsedDate.UTC().UnixMilli()
			}
		}
	}
}
