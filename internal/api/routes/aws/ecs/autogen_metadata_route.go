//AUTOGENERATED CODE DO NOT EDIT
package ecs

import (
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/cloud-inventory/internal/api/routes"
)

// GetECSMetadata godoc
// @Summary      Get ECS Metadata
// @Description  get a list of ECS metadata
// @Tags         aws ecs
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.AwsServiceMetadata
// @Failure      400
// @Router       /metadata/aws/ecs [get]
func GetECSMetadata(c *gin.Context) {
	c.IndentedJSON(200, routes.AwsServiceMetadata{
		Resources: []string{
			"clusters",
			"services",
			"tasks",
		},
	})
}