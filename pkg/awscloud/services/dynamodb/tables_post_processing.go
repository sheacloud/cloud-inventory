package dynamodb

import (
	"context"

	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessTableDescription(ctx context.Context, params *awscloud.AwsFetchInput, model *TableDescription) error {
	if model.CreationDateTime != nil {
		model.CreationDateTimeMilli = model.CreationDateTime.UTC().UnixMilli()
	}

	if model.ArchivalSummary != nil {
		if model.ArchivalSummary.ArchivalDateTime != nil {
			model.ArchivalSummary.ArchivalDateTimeMilli = model.ArchivalSummary.ArchivalDateTime.UTC().UnixMilli()
		}
	}

	if model.BillingModeSummary != nil {
		if model.BillingModeSummary.LastUpdateToPayPerRequestDateTime != nil {
			model.BillingModeSummary.LastUpdateToPayPerRequestDateTimeMilli = model.BillingModeSummary.LastUpdateToPayPerRequestDateTime.UTC().UnixMilli()
		}
	}

	if model.ProvisionedThroughput != nil {
		if model.ProvisionedThroughput.LastDecreaseDateTime != nil {
			model.ProvisionedThroughput.LastDecreaseDateTimeMilli = model.ProvisionedThroughput.LastDecreaseDateTime.UTC().UnixMilli()
		}
		if model.ProvisionedThroughput.LastIncreaseDateTime != nil {
			model.ProvisionedThroughput.LastIncreaseDateTimeMilli = model.ProvisionedThroughput.LastIncreaseDateTime.UTC().UnixMilli()
		}
	}

	for _, replica := range model.Replicas {
		if replica.ReplicaInaccessibleDateTime != nil {
			replica.ReplicaInaccessibleDateTimeMilli = replica.ReplicaInaccessibleDateTime.UTC().UnixMilli()
		}
	}

	if model.RestoreSummary != nil {
		if model.RestoreSummary.RestoreDateTime != nil {
			model.RestoreSummary.RestoreDateTimeMilli = model.RestoreSummary.RestoreDateTime.UTC().UnixMilli()
		}
	}

	if model.SSEDescription != nil {
		if model.SSEDescription.InaccessibleEncryptionDateTime != nil {
			model.SSEDescription.InaccessibleEncryptionDateTimeMilli = model.SSEDescription.InaccessibleEncryptionDateTime.UTC().UnixMilli()
		}
	}

	return nil
}
