package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

func init() {
	registerCustomRoleModelPostprocessingFunc(PostProcessRoleModel)
}

func PostProcessRoleModel(ctx context.Context, client *iam.Client, cfg aws.Config, model *RoleModel) {
	if model.CreateDate != nil {
		model.CreateDateMilli = model.CreateDate.UTC().UnixMilli()
	}

	if model.RoleLastUsed != nil {
		if model.RoleLastUsed.LastUsedDate != nil {
			model.RoleLastUsed.LastUsedDateMilli = model.RoleLastUsed.LastUsedDate.UTC().UnixMilli()
		}
	}

	//populate attached role policies
	attachedRolePoliciesPaginator := iam.NewListAttachedRolePoliciesPaginator(client, &iam.ListAttachedRolePoliciesInput{
		RoleName: aws.String(model.RoleName),
	})

	for attachedRolePoliciesPaginator.HasMorePages() {
		output, err := attachedRolePoliciesPaginator.NextPage(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":     "iam",
				"data_source": "roles",
				"account_id":  model.AccountId,
				"region":      model.Region,
				"cloud":       "aws",
				"error":       err,
			}).Error("error calling ListAttachedRolePolicies")
			continue
		}

		for _, policy := range output.AttachedPolicies {
			policyModel := new(AttachedPolicyRoleModel)
			copier.Copy(policyModel, policy)
			model.AttachedPolicies = append(model.AttachedPolicies, policyModel)
		}
	}

	//populate inline policies
	rolePoliciesPaginator := iam.NewListRolePoliciesPaginator(client, &iam.ListRolePoliciesInput{
		RoleName: aws.String(model.RoleName),
	})

	for rolePoliciesPaginator.HasMorePages() {
		output, err := rolePoliciesPaginator.NextPage(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":     "iam",
				"data_source": "roles",
				"account_id":  model.AccountId,
				"region":      model.Region,
				"cloud":       "aws",
				"error":       err,
			}).Error("error calling ListRolePolicies")
			continue
		}

		model.InlinePolicies = append(model.InlinePolicies, output.PolicyNames...)
	}
}
