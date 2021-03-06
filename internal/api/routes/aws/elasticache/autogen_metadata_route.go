//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_service_metadata_route.tmpl
package elasticache

import (
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/cloud-inventory/internal/api/routes"
)

// GetElastiCacheMetadata godoc
// @Summary      Get ElastiCache Metadata
// @Description  get a list of ElastiCache metadata
// @Tags         aws elasticache
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.AwsServiceMetadata
// @Failure      400
// @Router       /metadata/aws/elasticache [get]
func GetElastiCacheMetadata(c *gin.Context) {
	c.IndentedJSON(200, routes.AwsServiceMetadata{
		Resources: []string{
			"cache_clusters",
		},
	})
}
