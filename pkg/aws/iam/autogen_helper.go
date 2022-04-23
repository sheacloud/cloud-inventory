//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_helpers_file.tmpl
package iam

import (
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
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
