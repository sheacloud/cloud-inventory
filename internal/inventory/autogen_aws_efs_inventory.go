//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_service_inventory_file.tmpl
package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/efs"
)

func IngestAwsEFSFileSystems(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := efs.FetchFileSystems(ctx, input)
	if resources != nil {
		err := dao.PutAwsEFSFileSystems(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
