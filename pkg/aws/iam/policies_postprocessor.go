package iam

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

func init() {
	registerCustomPolicyModelPostprocessingFunc(PostProcessPolicyModel)
}

func PostProcessPolicyModel(ctx context.Context, client *iam.Client, cfg aws.Config, model *PolicyModel) {
	if model.CreateDate != nil {
		model.CreateDateMilli = model.CreateDate.UTC().UnixMilli()
	}

	if model.UpdateDate != nil {
		model.UpdateDateMilli = model.UpdateDate.UTC().UnixMilli()
	}

	// // TODO determine whether we keep this - this adds a large overhead for calling per-policy (lots of built-in policies in each account)
	// listEntitiesForPolicyPaginator := iam.NewListEntitiesForPolicyPaginator(client, &iam.ListEntitiesForPolicyInput{
	// 	PolicyArn: aws.String(model.Arn),
	// })

	// for listEntitiesForPolicyPaginator.HasMorePages() {
	// 	output, err := listEntitiesForPolicyPaginator.NextPage(ctx)
	// 	if err != nil {
	// 		logrus.WithFields(logrus.Fields{
	// 			"service":     "iam",
	// 			"data_source": "policies",
	// 			"account_id":  model.AccountId,
	// 			"region":      model.Region,
	// 			"cloud":       "aws",
	// 			"error":       err,
	// 		}).Error("error calling ListEntitiesForPolicy")
	// 		continue
	// 	}

	// 	//populate groups
	// 	for _, group := range output.PolicyGroups {
	// 		groupModel := new(PolicyGroupPolicyModel)
	// 		copier.Copy(groupModel, group)
	// 		model.Groups = append(model.Groups, groupModel)
	// 	}

	// 	//populate roles
	// 	for _, role := range output.PolicyRoles {
	// 		roleModel := new(PolicyRolePolicyModel)
	// 		copier.Copy(roleModel, role)
	// 		model.Roles = append(model.Roles, roleModel)
	// 	}

	// 	//populate users
	// 	for _, user := range output.PolicyUsers {
	// 		userModel := new(PolicyUserPolicyModel)
	// 		copier.Copy(userModel, user)
	// 		model.Users = append(model.Users, userModel)
	// 	}
	// }
}
