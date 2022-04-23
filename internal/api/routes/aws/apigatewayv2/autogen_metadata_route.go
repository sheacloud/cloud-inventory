//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_service_metadata_route.tmpl
package apigatewayv2

import (
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/cloud-inventory/internal/api/routes"
)

// GetApiGatewayV2Metadata godoc
// @Summary      Get ApiGatewayV2 Metadata
// @Description  get a list of ApiGatewayV2 metadata
// @Tags         aws apigatewayv2
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.AwsServiceMetadata
// @Failure      400
// @Router       /metadata/aws/apigatewayv2 [get]
func GetApiGatewayV2Metadata(c *gin.Context) {
	c.IndentedJSON(200, routes.AwsServiceMetadata{
		Resources: []string{
			"apis",
		},
	})
}
