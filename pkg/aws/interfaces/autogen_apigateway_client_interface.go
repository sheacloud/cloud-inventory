//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_service_client_interface_file.tmpl
package interfaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/apigateway"
)

type ApiGatewayClient interface {
	GetTags(ctx context.Context, params *apigateway.GetTagsInput, optFns ...func(*apigateway.Options)) (*apigateway.GetTagsOutput, error)
	GetStages(ctx context.Context, params *apigateway.GetStagesInput, optFns ...func(*apigateway.Options)) (*apigateway.GetStagesOutput, error)
	GetResources(ctx context.Context, params *apigateway.GetResourcesInput, optFns ...func(*apigateway.Options)) (*apigateway.GetResourcesOutput, error)
	GetRestApis(ctx context.Context, params *apigateway.GetRestApisInput, optFns ...func(*apigateway.Options)) (*apigateway.GetRestApisOutput, error)
}
