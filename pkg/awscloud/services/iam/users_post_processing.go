package iam

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/smithy-go"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessUser(ctx context.Context, params *awscloud.AwsFetchInput, model *User) error {
	awsClient := params.RegionalClients[params.Region]
	client := awsClient.IAM()

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
			return fmt.Errorf("error calling ListAttachedUserPolicies: %w", err)
		}

		for _, policy := range output.AttachedPolicies {
			policyModel := new(AttachedPolicy)
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
			return fmt.Errorf("error calling ListUserPolicies: %w", err)
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
			return fmt.Errorf("error calling ListGroupsForUser: %w", err)
		}

		for _, group := range output.Groups {
			model.GroupIds = append(model.GroupIds, *group.GroupId)
		}
	}

	// populate access keys
	model.AccessKeys = make([]*AccessKeyMetadata, 0)
	listAccessKeysPaginator := iam.NewListAccessKeysPaginator(client, &iam.ListAccessKeysInput{
		UserName: aws.String(model.UserName),
	})

	for listAccessKeysPaginator.HasMorePages() {
		output, err := listAccessKeysPaginator.NextPage(ctx)
		if err != nil {
			return fmt.Errorf("error calling ListAccessKeys: %w", err)
		}

		for _, key := range output.AccessKeyMetadata {
			accessKeyModel := new(AccessKeyMetadata)
			copier.Copy(accessKeyModel, &key)

			if accessKeyModel.CreateDate != nil {
				accessKeyModel.CreateDateMilli = accessKeyModel.CreateDate.UTC().UnixMilli()
			}

			model.AccessKeys = append(model.AccessKeys, accessKeyModel)
		}
	}

	// get login profile
	var apiError smithy.APIError

	profileResult, err := client.GetLoginProfile(ctx, &iam.GetLoginProfileInput{
		UserName: aws.String(model.UserName),
	})

	if errors.As(err, &apiError) && apiError.ErrorCode() == "NoSuchEntity" {
		// do nothing
	} else if err != nil {
		return fmt.Errorf("error calling GetLoginProfile: %w", err)
	} else {
		loginProfileModel := new(LoginProfile)
		copier.Copy(loginProfileModel, profileResult.LoginProfile)

		if loginProfileModel.CreateDate != nil {
			loginProfileModel.CreateDateMilli = loginProfileModel.CreateDate.UTC().UnixMilli()
		}

		model.LoginProfile = loginProfileModel
	}

	// populate tags
	tagsResult, err := client.ListUserTags(ctx, &iam.ListUserTagsInput{
		UserName: aws.String(model.UserName),
	})
	if err != nil {
		return fmt.Errorf("error calling ListUserTags: %w", err)
	}
	model.Tags = ConvertTags(tagsResult.Tags)
	return nil
}
