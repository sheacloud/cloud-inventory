//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_service_metadata_route.tmpl
package secretsmanager

import (
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/cloud-inventory/internal/api/routes"
)

// GetSecretsManagerMetadata godoc
// @Summary      Get SecretsManager Metadata
// @Description  get a list of SecretsManager metadata
// @Tags         aws secretsmanager
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.AwsServiceMetadata
// @Failure      400
// @Router       /metadata/aws/secretsmanager [get]
func GetSecretsManagerMetadata(c *gin.Context) {
	c.IndentedJSON(200, routes.AwsServiceMetadata{
		Resources: []string{
			"secrets",
		},
	})
}