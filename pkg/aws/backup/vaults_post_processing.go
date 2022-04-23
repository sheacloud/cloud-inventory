package backup

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
)

func PostProcessBackupVault(ctx context.Context, params *localAws.AwsFetchInput, model *BackupVault) error {
	awsClient := params.RegionalClients[params.Region]
	client := awsClient.Backup()

	// tags
	tagsOutput, err := client.ListTags(ctx, &backup.ListTagsInput{
		ResourceArn: aws.String(model.BackupVaultArn),
	})
	if err != nil {
		return fmt.Errorf("error calling ListTags: %v", err)
	}
	model.Tags = tagsOutput.Tags
	return nil
}
