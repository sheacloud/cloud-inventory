package redshift

import (
	"context"

	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessCluster(ctx context.Context, params *awscloud.AwsFetchInput, model *Cluster) error {
	if model.ClusterCreateTime != nil {
		model.ClusterCreateTimeMilli = model.ClusterCreateTime.UTC().UnixMilli()
	}

	if model.ExpectedNextSnapshotScheduleTime != nil {
		model.ExpectedNextSnapshotScheduleTimeMilli = model.ExpectedNextSnapshotScheduleTime.UTC().UnixMilli()
	}

	if model.NextMaintenanceWindowStartTime != nil {
		model.NextMaintenanceWindowStartTimeMilli = model.NextMaintenanceWindowStartTime.UTC().UnixMilli()
	}

	for _, deferredMaintenanceWindow := range model.DeferredMaintenanceWindows {
		if deferredMaintenanceWindow.DeferMaintenanceEndTime != nil {
			deferredMaintenanceWindow.DeferMaintenanceEndTimeMilli = deferredMaintenanceWindow.DeferMaintenanceEndTime.UTC().UnixMilli()
		}
		if deferredMaintenanceWindow.DeferMaintenanceStartTime != nil {
			deferredMaintenanceWindow.DeferMaintenanceStartTimeMilli = deferredMaintenanceWindow.DeferMaintenanceStartTime.UTC().UnixMilli()
		}
	}

	return nil
}
