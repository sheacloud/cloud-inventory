package api

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/sheacloud/cloud-inventory/internal/api/routes/awscloud"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//go:embed web
var staticFS embed.FS

func GetRouter(s3Client *s3.Client, s3Bucket string) *gin.Engine {
	router := gin.Default()

	v1 := router.Group("/api/v1")

	awsInventoryRouter := v1.Group("/inventory/aws")
	awsDiffRouter := v1.Group("/diff/aws")
	awsMetadataRouter := v1.Group("/metadata/aws")

	awscloud.AddInventoryRoutes(awsInventoryRouter, s3Client, s3Bucket)
	awscloud.AddDiffRoutes(awsDiffRouter, s3Client, s3Bucket)
	awscloud.AddMetadataRoutes(awsMetadataRouter, s3Client, s3Bucket)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(static.Serve("/", EmbedFolder(staticFS, "web")))
	return router
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
