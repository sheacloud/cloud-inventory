package ec2

import (
	"context"

	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

func PostProcessInstance(ctx context.Context, params *awscloud.AwsFetchInput, model *Instance) error {
	//set instance launch time
	if model.LaunchTime != nil {
		model.LaunchTimeMilli = model.LaunchTime.UTC().UnixMilli()
	}

	//set EBS volume attachment time
	for _, b := range model.BlockDeviceMappings {
		if b.Ebs != nil {
			if b.Ebs.AttachTime != nil {
				b.Ebs.AttachTimeMilli = b.Ebs.AttachTime.UTC().UnixMilli()
			}
		}
	}

	// set EIA attachment time
	for _, eia := range model.ElasticInferenceAcceleratorAssociations {
		if eia.ElasticInferenceAcceleratorAssociationTime != nil {
			eia.ElasticInferenceAcceleratorAssociationTimeMilli = eia.ElasticInferenceAcceleratorAssociationTime.UTC().UnixMilli()
		}
	}

	// ENI attachment time
	for _, eni := range model.NetworkInterfaces {
		if eni.Attachment != nil {
			if eni.Attachment.AttachTime != nil {
				eni.Attachment.AttachTimeMilli = eni.Attachment.AttachTime.UTC().UnixMilli()
			}
		}
	}

	return nil
}
