package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

func init() {
	registerCustomInstanceModelPostprocessingFunc(PostProcessInstanceModel)
}

func PostProcessInstanceModel(ctx context.Context, client *ec2.Client, cfg aws.Config, model *InstanceModel) {
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
}
