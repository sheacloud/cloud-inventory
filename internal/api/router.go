package api

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/cloud-inventory/internal/api/routes/awscloud"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func GetRouter(s3Client *s3.Client, s3Bucket string) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	awsInventoryRouter := v1.Group("/inventory/aws")
	awsDiffRouter := v1.Group("/diff/aws")

	awscloud.AddInventoryRoutes(awsInventoryRouter, s3Client, s3Bucket)
	awscloud.AddDiffRoutes(awsDiffRouter, s3Client, s3Bucket)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
