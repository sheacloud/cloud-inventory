package elasticache

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessCacheCluster(ctx context.Context, params *awscloud.AwsFetchInput, model *CacheCluster) error {
	if model.AuthTokenLastModifiedDate != nil {
		model.AuthTokenLastModifiedDateMilli = model.AuthTokenLastModifiedDate.UTC().UnixMilli()
	}
	if model.CacheClusterCreateTime != nil {
		model.CacheClusterCreateTimeMilli = model.CacheClusterCreateTime.UTC().UnixMilli()
	}

	for _, node := range model.CacheNodes {
		if node.CacheNodeCreateTime != nil {
			node.CacheNodeCreateTimeMilli = node.CacheNodeCreateTime.UTC().UnixMilli()
		}
	}

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
