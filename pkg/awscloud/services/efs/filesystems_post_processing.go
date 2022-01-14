package efs

import (
	"context"

	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessFileSystemDescription(ctx context.Context, params *awscloud.AwsFetchInput, model *FileSystemDescription) error {
	if model.CreationTime != nil {
		model.CreationTimeMilli = model.CreationTime.UTC().UnixMilli()
	}

	if model.SizeInBytes != nil {
		if model.SizeInBytes.Timestamp != nil {
			model.SizeInBytes.TimestampMilli = model.SizeInBytes.Timestamp.UTC().UnixMilli()
		}
	}

	return nil
}
