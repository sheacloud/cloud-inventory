package ec2

import (
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

func GetTagMap(tags []ec2types.Tag) map[string]string {
	tagMap := make(map[string]string)
	for _, tag := range tags {
		tagMap[*tag.Key] = *tag.Value
	}
	return tagMap
}
