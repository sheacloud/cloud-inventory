package ecr

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
)

func PostProcessRepository(ctx context.Context, params *localAws.AwsFetchInput, model *Repository) error {
	awsClient := params.RegionalClients[params.Region]
	client := awsClient.ECR()
	// tags
	tagsOutput, err := client.ListTagsForResource(ctx, &ecr.ListTagsForResourceInput{
		ResourceArn: aws.String(model.RepositoryArn),
	})
	if err != nil {
		return fmt.Errorf("error calling ListTags: %v", err)
	}
	model.Tags = ConvertTags(tagsOutput.Tags)

	return nil
}
