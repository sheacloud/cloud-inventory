package ecs

import (
	"context"

	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessTask(ctx context.Context, params *awscloud.AwsFetchInput, model *Task) error {
	if model.CreatedAt != nil {
		model.CreatedAtMilli = model.CreatedAt.UTC().UnixMilli()
	}

	if model.ConnectivityAt != nil {
		model.ConnectivityAtMilli = model.ConnectivityAt.UTC().UnixMilli()
	}
	if model.CreatedAt != nil {
		model.CreatedAtMilli = model.CreatedAt.UTC().UnixMilli()
	}
	if model.ExecutionStoppedAt != nil {
		model.ExecutionStoppedAtMilli = model.ExecutionStoppedAt.UTC().UnixMilli()
	}
	if model.PullStartedAt != nil {
		model.PullStartedAtMilli = model.PullStartedAt.UTC().UnixMilli()
	}
	if model.PullStoppedAt != nil {
		model.PullStoppedAtMilli = model.PullStoppedAt.UTC().UnixMilli()
	}
	if model.StartedAt != nil {
		model.StartedAtMilli = model.StartedAt.UTC().UnixMilli()
	}
	if model.StoppedAt != nil {
		model.StoppedAtMilli = model.StoppedAt.UTC().UnixMilli()
	}
	if model.StoppingAt != nil {
		model.StoppingAtMilli = model.StoppingAt.UTC().UnixMilli()
	}

	for _, container := range model.Containers {
		for _, managedAgent := range container.ManagedAgents {
			if managedAgent.LastStartedAt != nil {
				managedAgent.LastStartedAtMilli = managedAgent.LastStartedAt.UTC().UnixMilli()
			}
		}
	}

	return nil
}
