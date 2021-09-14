package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func init() {
	registerCustomInstanceProfileModelPostprocessingFunc(PostProcessInstanceProfileModel)
}

func PostProcessInstanceProfileModel(ctx context.Context, client *iam.Client, cfg aws.Config, model *InstanceProfileModel) {
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
