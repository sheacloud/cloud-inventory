//AUTOGENERATED CODE DO NOT EDIT
package storagegateway

import (
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/cloud-inventory/internal/api/routes"
)

// GetStorageGatewayMetadata godoc
// @Summary      Get StorageGateway Metadata
// @Description  get a list of StorageGateway metadata
// @Tags         aws storagegateway
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.AwsServiceMetadata
// @Failure      400
// @Router       /metadata/aws/storagegateway [get]
func GetStorageGatewayMetadata(c *gin.Context) {
	c.IndentedJSON(200, routes.AwsServiceMetadata{
		Resources: []string{
			"gateways",
		},
	})
}
