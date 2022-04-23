package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/aws/ec2"
)

func IngestAwsEC2Addresses(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchAddresses(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2Addresses(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2DhcpOptions(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchDhcpOptions(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2DhcpOptions(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2Images(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchImages(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2Images(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2Instances(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchInstances(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2Instances(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2InternetGateways(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchInternetGateways(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2InternetGateways(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2ManagedPrefixLists(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchManagedPrefixLists(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2ManagedPrefixLists(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2NatGateways(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchNatGateways(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2NatGateways(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2NetworkAcls(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchNetworkAcls(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2NetworkAcls(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2NetworkInterfaces(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchNetworkInterfaces(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2NetworkInterfaces(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2PlacementGroups(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchPlacementGroups(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2PlacementGroups(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2ReservedInstances(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchReservedInstances(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2ReservedInstances(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2RouteTables(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchRouteTables(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2RouteTables(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2SecurityGroups(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchSecurityGroups(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2SecurityGroups(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2Subnets(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchSubnets(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2Subnets(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2TransitGatewayPeeringAttachments(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchTransitGatewayPeeringAttachments(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2TransitGatewayPeeringAttachments(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2TransitGatewayRouteTables(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchTransitGatewayRouteTables(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2TransitGatewayRouteTables(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2TransitGatewayVpcAttachments(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchTransitGatewayVpcAttachments(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2TransitGatewayVpcAttachments(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2TransitGateways(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchTransitGateways(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2TransitGateways(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2Volumes(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchVolumes(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2Volumes(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2VpcEndpoints(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchVpcEndpoints(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2VpcEndpoints(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2VpcPeeringConnections(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchVpcPeeringConnections(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2VpcPeeringConnections(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2Vpcs(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchVpcs(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2Vpcs(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}

func IngestAwsEC2VpnGateways(ctx context.Context, dao db.WriterDAO, input *aws.AwsFetchInput) (*aws.AwsFetchOutputMetadata, error) {
	resources, metadata := ec2.FetchVpnGateways(ctx, input)
	if resources != nil {
		err := dao.PutAwsEC2VpnGateways(ctx, resources)
		if err != nil {
			return nil, err
		}
	}

	return metadata, nil
}
