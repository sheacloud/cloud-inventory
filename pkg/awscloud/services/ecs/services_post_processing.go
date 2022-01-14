package ecs

import (
	"context"

	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessService(ctx context.Context, params *awscloud.AwsFetchInput, model *Service) error {
	if model.CreatedAt != nil {
		model.CreatedAtMilli = model.CreatedAt.UTC().UnixMilli()
	}

	for _, deployment := range model.Deployments {
		if deployment.CreatedAt != nil {
			deployment.CreatedAtMilli = deployment.CreatedAt.UTC().UnixMilli()
		}

		if deployment.UpdatedAt != nil {
			deployment.UpdatedAtMilli = deployment.UpdatedAt.UTC().UnixMilli()
		}
	}

	for _, taskSet := range model.TaskSets {
		if taskSet.CreatedAt != nil {
			taskSet.CreatedAtMilli = taskSet.CreatedAt.UTC().UnixMilli()
		}

		if taskSet.StabilityStatusAt != nil {
			taskSet.StabilityStatusAtMilli = taskSet.StabilityStatusAt.UTC().UnixMilli()
		}

		if taskSet.UpdatedAt != nil {
			taskSet.UpdatedAtMilli = taskSet.UpdatedAt.UTC().UnixMilli()
		}
	}

	return nil
}
