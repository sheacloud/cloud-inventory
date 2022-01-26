package ec2

import (
	"context"

	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessNatGateway(ctx context.Context, params *awscloud.AwsFetchInput, model *NatGateway) error {
	if model.CreateTime != nil {
		model.CreateTimeMilli = model.CreateTime.UTC().UnixMilli()
	}
	if model.DeleteTime != nil {
		model.DeleteTimeMilli = model.DeleteTime.UTC().UnixMilli()
	}

	if model.ProvisionedBandwidth != nil {
		if model.ProvisionedBandwidth.ProvisionTime != nil {
			model.ProvisionedBandwidth.ProvisionTimeMilli = model.ProvisionedBandwidth.ProvisionTime.UTC().UnixMilli()
		}
		if model.ProvisionedBandwidth.RequestTime != nil {
			model.ProvisionedBandwidth.RequestTimeMilli = model.ProvisionedBandwidth.RequestTime.UTC().UnixMilli()
		}
	}

	return nil
}
