//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_service_metadata_route.tmpl
package cloudtrail

import (
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/cloud-inventory/internal/api/routes"
)

// GetCloudTrailMetadata godoc
// @Summary      Get CloudTrail Metadata
// @Description  get a list of CloudTrail metadata
// @Tags         aws cloudtrail
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.AwsServiceMetadata
// @Failure      400
// @Router       /metadata/aws/cloudtrail [get]
func GetCloudTrailMetadata(c *gin.Context) {
	c.IndentedJSON(200, routes.AwsServiceMetadata{
		Resources: []string{
			"trails",
		},
	})
}
