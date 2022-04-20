//AUTOGENERATED CODE DO NOT EDIT
package efs

import (
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/cloud-inventory/internal/api/routes"
)

// GetEFSMetadata godoc
// @Summary      Get EFS Metadata
// @Description  get a list of EFS metadata
// @Tags         aws efs
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.AwsServiceMetadata
// @Failure      400
// @Router       /metadata/aws/efs [get]
func GetEFSMetadata(c *gin.Context) {
	c.IndentedJSON(200, routes.AwsServiceMetadata{
		Resources: []string{
			"filesystems",
		},
	})
}