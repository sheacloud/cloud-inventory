package api

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/cloud-inventory/internal/api/routes/awscloud"
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
// @name                        X-API-Key

// @BasePath  /v1
func GetRouter(s3Client *s3.Client, s3Bucket string) *gin.Engine {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8080", "https://cloud-inventory.sheacloud.com"}
	router.Use(cors.New(corsConfig))

	v1 := router.Group("/v1")

	awsInventoryRouter := v1.Group("/inventory/aws")
	awsDiffRouter := v1.Group("/diff/aws")
	awsMetadataRouter := v1.Group("/metadata/aws")

	awscloud.AddInventoryRoutes(awsInventoryRouter, s3Client, s3Bucket)
	awscloud.AddDiffRoutes(awsDiffRouter, s3Client, s3Bucket)
	awscloud.AddMetadataRoutes(awsMetadataRouter, s3Client, s3Bucket)

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
