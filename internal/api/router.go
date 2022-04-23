package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/cloud-inventory/internal/api/routes/aws"
	"github.com/sheacloud/cloud-inventory/internal/db"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Cloud Inventory API
// @version         1.0
// @description     Query Cloud Inventory

// @contact.name   Jon Shea
// @contact.email  cloud-inventory@sheacloud.com

// @license.name  MIT
// @license.url   https://opensource.org/licenses/MIT

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization

// @BasePath  /v1
func GetRouter(dao db.ReaderDAO) *gin.Engine {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*", "https://cloud-inventory.sheacloud.com"}
	corsConfig.AllowHeaders = []string{"*"}
	router.Use(cors.New(corsConfig))

	v1 := router.Group("/v1")

	awsInventoryRouter := v1.Group("/inventory/aws")
	awsDiffRouter := v1.Group("/diff/aws")
	awsMetadataRouter := v1.Group("/metadata/aws")

	aws.AddInventoryRoutes(awsInventoryRouter, dao)
	aws.AddDiffRoutes(awsDiffRouter, dao)
	aws.AddMetadataRoutes(awsMetadataRouter, dao)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
