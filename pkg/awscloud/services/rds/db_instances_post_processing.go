package rds

import (
	"context"

	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessDBInstance(ctx context.Context, params *awscloud.AwsFetchInput, model *DBInstance) error {
	if model.AutomaticRestartTime != nil {
		model.AutomaticRestartTimeMilli = model.AutomaticRestartTime.UTC().UnixMilli()
	}

	if model.InstanceCreateTime != nil {
		model.InstanceCreateTimeMilli = model.InstanceCreateTime.UTC().UnixMilli()
	}

	if model.LatestRestorableTime != nil {
		model.LatestRestorableTimeMilli = model.LatestRestorableTime.UTC().UnixMilli()
	}

	return nil
}
