package iam

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/jinzhu/copier"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
)

func PostProcessRole(ctx context.Context, params *localAws.AwsFetchInput, model *Role) error {
	awsClient := params.RegionalClients[params.Region]
	client := awsClient.IAM()

	//populate attached role policies
	attachedRolePoliciesPaginator := iam.NewListAttachedRolePoliciesPaginator(client, &iam.ListAttachedRolePoliciesInput{
		RoleName: aws.String(model.RoleName),
	})

	for attachedRolePoliciesPaginator.HasMorePages() {
		output, err := attachedRolePoliciesPaginator.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("error calling ListAttachedRolePolicies: %w", err)
		}

		for _, policy := range output.AttachedPolicies {
			policyModel := new(AttachedPolicy)
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
			return fmt.Errorf("error calling ListRolePolicies: %w", err)
		}

		model.InlinePolicies = append(model.InlinePolicies, output.PolicyNames...)
	}

	// populate tags
	tagsResult, err := client.ListRoleTags(ctx, &iam.ListRoleTagsInput{
		RoleName: aws.String(model.RoleName),
	})
	if err != nil {
		return fmt.Errorf("error calling ListRoleTags: %w", err)
	}
	model.Tags = ConvertTags(tagsResult.Tags)

	return nil
}
