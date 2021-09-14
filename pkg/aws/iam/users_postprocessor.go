package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

func init() {
	registerCustomUserModelPostprocessingFunc(PostProcessUserModel)
}

func PostProcessUserModel(ctx context.Context, client *iam.Client, cfg aws.Config, model *UserModel) {
	if model.CreateDate != nil {
		model.CreateDateMilli = model.CreateDate.UTC().UnixMilli()
	}

	if model.PasswordLastUsed != nil {
		model.PasswordLastUsedMilli = model.PasswordLastUsed.UTC().UnixMilli()
	}

	listAttachedUserPoliciesPaginator := iam.NewListAttachedUserPoliciesPaginator(client, &iam.ListAttachedUserPoliciesInput{
		UserName: aws.String(model.UserName),
	})

	for listAttachedUserPoliciesPaginator.HasMorePages() {
		output, err := listAttachedUserPoliciesPaginator.NextPage(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":     "iam",
				"data_source": "users",
				"account_id":  model.AccountId,
				"region":      model.Region,
				"cloud":       "aws",
				"error":       err,
			}).Error("error calling ListAttachedUserPolicies")
			continue
		}

		for _, policy := range output.AttachedPolicies {
			policyModel := new(AttachedPolicyUserModel)
			copier.Copy(policyModel, policy)
			model.AttachedPolicies = append(model.AttachedPolicies, policyModel)
		}
	}

	//populate inline policies
	listUserPoliciesPaginator := iam.NewListUserPoliciesPaginator(client, &iam.ListUserPoliciesInput{
		UserName: aws.String(model.UserName),
	})

	for listUserPoliciesPaginator.HasMorePages() {
		output, err := listUserPoliciesPaginator.NextPage(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":     "iam",
				"data_source": "users",
				"account_id":  model.AccountId,
				"region":      model.Region,
				"cloud":       "aws",
				"error":       err,
			}).Error("error calling ListUserPolicies")
			continue
		}

		model.InlinePolicies = append(model.InlinePolicies, output.PolicyNames...)
	}

	//populate associated groups
	listGroupsForUserPaginator := iam.NewListGroupsForUserPaginator(client, &iam.ListGroupsForUserInput{
		UserName: aws.String(model.UserName),
	})

	for listGroupsForUserPaginator.HasMorePages() {
		output, err := listGroupsForUserPaginator.NextPage(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":     "iam",
				"data_source": "users",
				"account_id":  model.AccountId,
				"region":      model.Region,
				"cloud":       "aws",
				"error":       err,
			}).Error("error calling ListGroupsForUser")
			continue
		}

		for _, group := range output.Groups {
			model.GroupIds = append(model.GroupIds, *group.GroupId)
		}
	}
}
