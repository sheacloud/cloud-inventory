package efs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/efs"
)

func init() {
	registerCustomFileSystemDescriptionModelPostprocessingFunc(PostProcessFileSystemDescriptionModel)
}

func PostProcessFileSystemDescriptionModel(ctx context.Context, client *efs.Client, cfg aws.Config, model *FileSystemDescriptionModel) {
	if model.CreationTime != nil {
		model.CreationTimeMilli = model.CreationTime.UTC().UnixMilli()
	}

	if model.SizeInBytes != nil {
		if model.SizeInBytes.Timestamp != nil {
			model.SizeInBytes.TimestampMilli = model.SizeInBytes.Timestamp.UTC().UnixMilli()
		}
	}
}
