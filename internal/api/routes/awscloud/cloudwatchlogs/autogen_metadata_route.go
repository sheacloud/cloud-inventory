//AUTOGENERATED CODE DO NOT EDIT
package cloudwatchlogs

import (
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/cloud-inventory/internal/api/routes"
)

// GetCloudWatchLogsMetadata godoc
// @Summary      Get CloudWatchLogs Metadata
// @Description  get a list of CloudWatchLogs metadata
// @Tags         aws cloudwatchlogs
// @Produce      json
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.AwsServiceMetadata
// @Failure      400
// @Router       /metadata/aws/cloudwatchlogs [get]
func GetCloudWatchLogsMetadata(c *gin.Context) {
	c.IndentedJSON(200, routes.AwsServiceMetadata{
		Resources: []string{
			"log_groups",
		},
	})
}
