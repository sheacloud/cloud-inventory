package route53

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/jinzhu/copier"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
)

func PostProcessHostedZone(ctx context.Context, params *localAws.AwsFetchInput, model *HostedZone) error {
	awsClient := params.RegionalClients[params.Region]
	client := awsClient.Route53()

	// get associated VPCs
	result, err := client.GetHostedZone(ctx, &route53.GetHostedZoneInput{
		Id: aws.String(model.Id),
	})
	if err != nil {
		return fmt.Errorf("error calling GetHostedZone: %w", err)
	}
	model.VpcAssociations = make([]*VPC, len(result.VPCs))

	for i, vpc := range result.VPCs {
		vpcModel := new(VPC)
		copier.Copy(vpcModel, &vpc)
		model.VpcAssociations[i] = vpcModel
	}

	hosetdZoneIdParts := strings.Split(model.Id, "/")
	if len(hosetdZoneIdParts) != 3 {
		return fmt.Errorf("invalid hosted zone id: %s", model.Id)
	}

	// get tags
	tagResult, err := client.ListTagsForResource(ctx, &route53.ListTagsForResourceInput{
		ResourceId:   aws.String(hosetdZoneIdParts[2]),
		ResourceType: types.TagResourceTypeHostedzone,
	})
	if err != nil {
		return fmt.Errorf("error calling ListTagsForResource: %w", err)
	}

	if tagResult.ResourceTagSet != nil {
		model.Tags = ConvertTags(tagResult.ResourceTagSet.Tags)
	}

	return nil
}
