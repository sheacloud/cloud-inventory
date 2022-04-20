//AUTOGENERATED CODE DO NOT EDIT
package sns

import (
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

func ConvertTags(tags []types.Tag) map[string]string {
	tagMap := make(map[string]string)
	for _, tag := range tags {
		tagMap[*tag.Key] = *tag.Value
	}
	return tagMap
}