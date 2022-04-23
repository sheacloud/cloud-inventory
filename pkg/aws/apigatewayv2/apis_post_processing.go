package apigatewayv2

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2"
	"github.com/jinzhu/copier"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
)

func PostProcessApi(ctx context.Context, params *localAws.AwsFetchInput, model *Api) error {
	awsClient := params.RegionalClients[params.Region]
	client := awsClient.ApiGatewayV2()

	resourcesResponse, err := client.GetRoutes(ctx, &apigatewayv2.GetRoutesInput{
		ApiId: aws.String(model.ApiId),
	})
	if err != nil {
		return fmt.Errorf("error calling GetRoutes in %s/%s: %w", params.AccountId, params.Region, err)
	}
	model.Routes = []*GetRouteOutput{}
	for _, route := range resourcesResponse.Items {
		resourceModel := new(GetRouteOutput)
		err = copier.CopyWithOption(&resourceModel, &route, localAws.CopyOption)
		if err != nil {
			return fmt.Errorf("error copying route %s in %s/%s: %w", *route.RouteId, params.AccountId, params.Region, err)
		}
		model.Routes = append(model.Routes, resourceModel)
	}

	stagesResposne, err := client.GetStages(ctx, &apigatewayv2.GetStagesInput{
		ApiId: aws.String(model.ApiId),
	})
	if err != nil {
		return fmt.Errorf("error calling GetStages in %s/%s: %w", params.AccountId, params.Region, err)
	}
	model.Stages = []*Stage{}
	for _, stage := range stagesResposne.Items {
		stageModel := new(Stage)
		err = copier.CopyWithOption(&stageModel, &stage, localAws.CopyOption)
		if err != nil {
			return fmt.Errorf("error copying stage %s in %s/%s: %w", *stage.StageName, params.AccountId, params.Region, err)
		}
		model.Stages = append(model.Stages, stageModel)
	}

	integrationsResponse, err := client.GetIntegrations(ctx, &apigatewayv2.GetIntegrationsInput{
		ApiId: aws.String(model.ApiId),
	})
	if err != nil {
		return fmt.Errorf("error calling GetIntegrations in %s/%s: %w", params.AccountId, params.Region, err)
	}
	model.Integrations = []*Integration{}
	for _, integration := range integrationsResponse.Items {
		integrationModel := new(Integration)
		err = copier.CopyWithOption(&integrationModel, &integration, localAws.CopyOption)
		if err != nil {
			return fmt.Errorf("error copying stage %s in %s/%s: %w", *integration.IntegrationId, params.AccountId, params.Region, err)
		}
		model.Integrations = append(model.Integrations, integrationModel)
	}

	authorizersResponse, err := client.GetAuthorizers(ctx, &apigatewayv2.GetAuthorizersInput{
		ApiId: aws.String(model.ApiId),
	})
	if err != nil {
		return fmt.Errorf("error calling GetAuthorizers in %s/%s: %w", params.AccountId, params.Region, err)
	}
	model.Authorizers = []*Authorizer{}
	for _, authorizer := range authorizersResponse.Items {
		authorizerModel := new(Authorizer)
		err = copier.CopyWithOption(&authorizerModel, &authorizer, localAws.CopyOption)
		if err != nil {
			return fmt.Errorf("error copying authorizer %s in %s/%s: %w", *authorizer.AuthorizerId, params.AccountId, params.Region, err)
		}
		model.Authorizers = append(model.Authorizers, authorizerModel)
	}
	return nil
}
