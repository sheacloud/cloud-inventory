package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

func init() {
	registerCustomGroupModelPostprocessingFunc(PostProcessGroupModel)
}

func PostProcessGroupModel(ctx context.Context, client *iam.Client, cfg aws.Config, model *GroupModel) {
	if model.CreateDate != nil {
		model.CreateDateMilli = model.CreateDate.UTC().UnixMilli()
	}

	listAttachedGroupPoliciesPaginator := iam.NewListAttachedGroupPoliciesPaginator(client, &iam.ListAttachedGroupPoliciesInput{
		GroupName: aws.String(model.GroupName),
	})

	for listAttachedGroupPoliciesPaginator.HasMorePages() {
		output, err := listAttachedGroupPoliciesPaginator.NextPage(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":     "iam",
				"data_source": "groups",
				"account_id":  model.AccountId,
				"region":      model.Region,
				"cloud":       "aws",
				"error":       err,
			}).Error("error calling ListAttachedGroupPolicies")
			continue
		}

		for _, policy := range output.AttachedPolicies {
			policyModel := new(AttachedPolicyGroupModel)
			copier.Copy(policyModel, policy)
			model.AttachedPolicies = append(model.AttachedPolicies, policyModel)
		}
	}

	//populate inline policies
	listGroupPoliciesPaginator := iam.NewListGroupPoliciesPaginator(client, &iam.ListGroupPoliciesInput{
		GroupName: aws.String(model.GroupName),
	})

	for listGroupPoliciesPaginator.HasMorePages() {
		output, err := listGroupPoliciesPaginator.NextPage(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":     "iam",
				"data_source": "groups",
				"account_id":  model.AccountId,
				"region":      model.Region,
				"cloud":       "aws",
				"error":       err,
			}).Error("error calling ListGroupPolicies")
			continue
		}

		model.InlinePolicies = append(model.InlinePolicies, output.PolicyNames...)
	}

	//populate associated groups
	output, err := client.GetGroup(ctx, &iam.GetGroupInput{
		GroupName: aws.String(model.GroupName),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "iam",
			"data_source": "groups",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling GetGroup")
		return
	}

	for _, user := range output.Users {
		model.UserIds = append(model.UserIds, *user.UserId)
	}
}
