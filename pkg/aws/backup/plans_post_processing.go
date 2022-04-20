package backup

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/jinzhu/copier"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
)

func PostProcessBackupPlan(ctx context.Context, params *localAws.AwsFetchInput, model *BackupPlan) error {
	awsClient := params.RegionalClients[params.Region]
	client := awsClient.Backup()

	// tags
	tagsOutput, err := client.ListTags(ctx, &backup.ListTagsInput{
		ResourceArn: aws.String(model.BackupPlanArn),
	})
	if err != nil {
		return fmt.Errorf("error calling ListTags: %v", err)
	}
	model.Tags = tagsOutput.Tags

	selectionsOutput, err := client.ListBackupSelections(ctx, &backup.ListBackupSelectionsInput{
		BackupPlanId: aws.String(model.BackupPlanId),
	})
	if err != nil {
		return fmt.Errorf("error calling ListBackupSelections: %v", err)
	}

	model.Selections = []*BackupSelectionsListMember{}
	for _, selection := range selectionsOutput.BackupSelectionsList {
		selectionModel := new(BackupSelectionsListMember)
		copier.Copy(&selectionModel, selection)

		model.Selections = append(model.Selections, selectionModel)
	}

	return nil
}
