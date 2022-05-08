package kms

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
)

func PostProcessKey(ctx context.Context, params *localAws.AwsFetchInput, model *Key) error {
	awsClient := params.RegionalClients[params.Region]
	client := awsClient.KMS()

	// tags
	tagsResult, err := client.ListResourceTags(ctx, &kms.ListResourceTagsInput{
		KeyId: &model.KeyId,
	})
	if err != nil {
		return fmt.Errorf("error calling ListResourceTags in %s/%s: %w", params.AccountId, params.Region, err)

	}
	model.Tags = ConvertTags(tagsResult.Tags)

	// aliases
	aliasResult, err := client.ListAliases(ctx, &kms.ListAliasesInput{
		KeyId: &model.KeyId,
	})
	if err != nil {
		return fmt.Errorf("error calling ListAliases in %s/%s: %w", params.AccountId, params.Region, err)

	}
	model.Aliases = make([]*AliasListEntry, len(aliasResult.Aliases))
	for i, alias := range aliasResult.Aliases {
		aliasModel := new(AliasListEntry)
		copier.CopyWithOption(&aliasModel, &alias, aws.CopyOption)
		model.Aliases[i] = aliasModel
	}

	return nil
}
