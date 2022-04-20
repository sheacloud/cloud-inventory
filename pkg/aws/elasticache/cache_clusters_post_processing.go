package elasticache

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
)

func PostProcessCacheCluster(ctx context.Context, params *localAws.AwsFetchInput, model *CacheCluster) error {
	awsClient := params.RegionalClients[params.Region]
	client := awsClient.ElastiCache()

	tagResult, err := client.ListTagsForResource(ctx, &elasticache.ListTagsForResourceInput{
		ResourceName: aws.String(model.ARN),
	})
	if err != nil {
		return fmt.Errorf("error calling ListTagsForResource: %w", err)
	}

	model.Tags = ConvertTags(tagResult.TagList)

	return nil
}
