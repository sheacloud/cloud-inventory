package apigateway

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/jinzhu/copier"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
)

func PostProcessRestApi(ctx context.Context, params *localAws.AwsFetchInput, model *RestApi) error {
	awsClient := params.RegionalClients[params.Region]
	client := awsClient.ApiGateway()

	resourcesResponse, err := client.GetResources(ctx, &apigateway.GetResourcesInput{
		RestApiId: aws.String(model.Id),
	})
	if err != nil {
		return fmt.Errorf("error calling GetResources in %s/%s: %w", params.AccountId, params.Region, err)
	}
	model.Resources = []*Resource{}
	for _, resource := range resourcesResponse.Items {
		resourceModel := new(Resource)
		err = copier.Copy(&resourceModel, &resource)
		if err != nil {
			return fmt.Errorf("error copying resource %s in %s/%s: %w", *resource.Id, params.AccountId, params.Region, err)
		}
		model.Resources = append(model.Resources, resourceModel)
	}

	stagesResposne, err := client.GetStages(ctx, &apigateway.GetStagesInput{
		RestApiId: aws.String(model.Id),
	})
	if err != nil {
		return fmt.Errorf("error calling GetStages in %s/%s: %w", params.AccountId, params.Region, err)
	}
	model.Stages = []*Stage{}
	for _, stage := range stagesResposne.Item {
		stageModel := new(Stage)
		err = copier.Copy(&stageModel, &stage)
		if err != nil {
			return fmt.Errorf("error copying stage %s in %s/%s: %w", *stage.StageName, params.AccountId, params.Region, err)
		}
		model.Stages = append(model.Stages, stageModel)
	}

	return nil
}
