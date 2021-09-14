package rds

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func init() {
	registerCustomDBClusterModelPostprocessingFunc(PostProcessDBClusterModel)
}

func PostProcessDBClusterModel(ctx context.Context, client *rds.Client, cfg aws.Config, model *DBClusterModel) {
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
}
