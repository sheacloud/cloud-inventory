//AUTOGENERATED CODE DO NOT EDIT
package elasticache

import (
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
)

func ConvertTags(tags []types.Tag) map[string]string {
	if len(tags) == 0 {
		return nil
	}
	tagMap := make(map[string]string)
	for _, tag := range tags {
		tagMap[*tag.Key] = *tag.Value
	}
	return tagMap
}
