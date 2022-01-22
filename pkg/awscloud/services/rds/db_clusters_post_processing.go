package rds

import (
	"context"

	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessDBCluster(ctx context.Context, params *awscloud.AwsFetchInput, model *DBCluster) error {
	if model.AutomaticRestartTime != nil {
		model.AutomaticRestartTimeMilli = model.AutomaticRestartTime.UTC().UnixMilli()
	}

	if model.ClusterCreateTime != nil {
		model.ClusterCreateTimeMilli = model.ClusterCreateTime.UTC().UnixMilli()
	}

	if model.EarliestBacktrackTime != nil {
		model.EarliestBacktrackTimeMilli = model.EarliestBacktrackTime.UTC().UnixMilli()
	}

	if model.EarliestRestorableTime != nil {
		model.EarliestRestorableTimeMilli = model.EarliestRestorableTime.UTC().UnixMilli()
	}

	if model.LatestRestorableTime != nil {
		model.LatestRestorableTimeMilli = model.LatestRestorableTime.UTC().UnixMilli()
	}

	return nil
}
