//AUTOGENERATED CODE DO NOT EDIT
package elasticloadbalancing

import (
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/cloud-inventory/internal/api/routes"
)

// GetElasticLoadBalancingMetadata godoc
// @Summary      Get ElasticLoadBalancing Metadata
// @Description  get a list of ElasticLoadBalancing metadata
// @Tags         aws elasticloadbalancing
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.AwsServiceMetadata
// @Failure      400
// @Router       /metadata/aws/elasticloadbalancing [get]
func GetElasticLoadBalancingMetadata(c *gin.Context) {
	c.IndentedJSON(200, routes.AwsServiceMetadata{
		Resources: []string{
			"load_balancers",
		},
	})
}
