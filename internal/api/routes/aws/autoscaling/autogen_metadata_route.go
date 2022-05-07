//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_service_metadata_route.tmpl
package autoscaling

import (
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/cloud-inventory/internal/api/routes"
)

// GetAutoScalingMetadata godoc
// @Summary      Get AutoScaling Metadata
// @Description  get a list of AutoScaling metadata
// @Tags         aws autoscaling
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.AwsServiceMetadata
// @Failure      400
// @Router       /metadata/aws/autoscaling [get]
func GetAutoScalingMetadata(c *gin.Context) {
	c.IndentedJSON(200, routes.AwsServiceMetadata{
		Resources: []string{
			"auto_scaling_groups",
			"launch_configurations",
		},
	})
}
