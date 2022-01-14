package iam

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessGroup(ctx context.Context, params *awscloud.AwsFetchInput, model *Group) error {
	awsClient := params.RegionalClients[params.Region]
	client := awsClient.IAM()

	if model.CreateDate != nil {
		model.CreateDateMilli = model.CreateDate.UTC().UnixMilli()
	}

	listAttachedGroupPoliciesPaginator := iam.NewListAttachedGroupPoliciesPaginator(client, &iam.ListAttachedGroupPoliciesInput{
		GroupName: aws.String(model.GroupName),
	})

	for listAttachedGroupPoliciesPaginator.HasMorePages() {
		output, err := listAttachedGroupPoliciesPaginator.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("failed to list attached group policies: %w", err)
		}

		for _, policy := range output.AttachedPolicies {
			policyModel := new(AttachedPolicy)
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
			return fmt.Errorf("failed to list group policies: %w", err)
		}

		model.InlinePolicies = append(model.InlinePolicies, output.PolicyNames...)
	}

	//populate associated groups
	output, err := client.GetGroup(ctx, &iam.GetGroupInput{
		GroupName: aws.String(model.GroupName),
	})
	if err != nil {
		return fmt.Errorf("error calling GetGroup: %w", err)
	}

	for _, user := range output.Users {
		model.UserIds = append(model.UserIds, *user.UserId)
	}

	return nil
}
