//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_service_metadata_route.tmpl
package cloudwatch

import (
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/cloud-inventory/internal/api/routes"
)

// GetCloudWatchMetadata godoc
// @Summary      Get CloudWatch Metadata
// @Description  get a list of CloudWatch metadata
// @Tags         aws cloudwatch
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.AwsServiceMetadata
// @Failure      400
// @Router       /metadata/aws/cloudwatch [get]
func GetCloudWatchMetadata(c *gin.Context) {
	c.IndentedJSON(200, routes.AwsServiceMetadata{
		Resources: []string{
			"metric_alarms",
			"composite_alarms",
		},
	})
}