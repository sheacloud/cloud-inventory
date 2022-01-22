//AUTOGENERATED CODE DO NOT EDIT
package elasticache

import (
	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/r3labs/diff/v2"
	"github.com/sheacloud/cloud-inventory/internal/indexedstorage"
	"github.com/sheacloud/cloud-inventory/internal/routes"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud/services/elasticache"
)

// ListCacheClusters godoc
// @Summary      List CacheClusters
// @Description  get a list of cache_clusters
// @Tags         aws elasticache
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param		 time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after)
// @Param		 time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before' or 'after'." Format(dateTime)
// @Success      200  {array}   elasticache.CacheCluster
// @Failure      400
// @Router       /inventory/aws/elasticache/cache_clusters [get]
func ListCacheClusters(c *gin.Context, s3Client *awsS3.Client, s3Bucket string) {
	var params routes.AwsQueryParameters
	if err := c.BindQuery(&params); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	err := params.Process()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	s3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "elasticache", "cache_clusters"}, *params.ReportDate, params.GetRequestTimeSelection(), s3Client, new(elasticache.CacheCluster))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	results := []interface{}{}
	for s3DirReader.HasNextFile() {
		resultInterface, err := s3DirReader.ReadNextFile()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
		results = append(results, resultInterface...)
	}

	// Filter results
	if params.AccountId != nil || params.Region != nil {
		filteredResults := []interface{}{}
		for _, result := range results {
			obj := result.(*elasticache.CacheCluster)
			if params.AccountId != nil && obj.AccountId != *params.AccountId {
				continue
			}
			if params.Region != nil && obj.Region != *params.Region {
				continue
			}
			filteredResults = append(filteredResults, result)
		}
		results = filteredResults
	}

	c.IndentedJSON(200, results)
}

// GetCacheCluster godoc
// @Summary      Get a specific CacheCluster
// @Description  Get a specific CacheCluster by its ARN
// @Tags         aws elasticache
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Param        arn path string true "The arn of the CacheCluster to retrieve"
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param		 time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after)
// @Param		 time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before' or 'after'." Format(dateTime)
// @Success      200  {object}   elasticache.CacheCluster
// @Failure      400
// @Failure 	 404
// @Router       /inventory/aws/elasticache/cache_clusters/{arn} [get]
func GetCacheCluster(c *gin.Context, s3Client *awsS3.Client, s3Bucket string) {
	var params routes.AwsQueryParameters
	if err := c.BindQuery(&params); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	err := params.Process()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("arn")

	s3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "elasticache", "cache_clusters"}, *params.ReportDate, params.GetRequestTimeSelection(), s3Client, new(elasticache.CacheCluster))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	results := []interface{}{}
	for s3DirReader.HasNextFile() {
		resultInterface, err := s3DirReader.ReadNextFile()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
		results = append(results, resultInterface...)
	}

	// Filter results
	for _, result := range results {
		obj := result.(*elasticache.CacheCluster)
		if params.AccountId != nil && obj.AccountId != *params.AccountId {
			continue
		}
		if params.Region != nil && obj.Region != *params.Region {
			continue
		}
		if obj.ARN == id {
			c.IndentedJSON(200, obj)
			return
		}
	}
	c.AbortWithStatusJSON(404, gin.H{"error": "No CacheCluster found with arn " + id})
}

// DiffMultiCacheClusters godoc
// @Summary      Diff CacheClusters
// @Description  get a diff of CacheClusters between two points in time
// @Tags         aws elasticache
// @Produce      json
// @Param        start_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 start_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after)
// @Param		 start_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before' or 'after'." Format(dateTime)
// @Param        end_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 end_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after)
// @Param		 end_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before' or 'after'." Format(dateTime)
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Success      200  {array}   routes.Diff
// @Failure      400
// @Router       /diff/aws/elasticache/cache_clusters [get]
func DiffMultiCacheClusters(c *gin.Context, s3Client *awsS3.Client, s3Bucket string) {
	var params routes.AwsDiffParameters
	if err := c.BindQuery(&params); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	err := params.Process()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	startS3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "elasticache", "cache_clusters"}, *params.StartReportDate, params.GetRequestStartTimeSelection(), s3Client, new(elasticache.CacheCluster))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	startResults := []interface{}{}
	for startS3DirReader.HasNextFile() {
		resultInterface, err := startS3DirReader.ReadNextFile()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
		startResults = append(startResults, resultInterface...)
	}
	// Filter results
	if params.AccountId != nil || params.Region != nil {
		filteredResults := []interface{}{}
		for _, result := range startResults {
			obj := result.(*elasticache.CacheCluster)
			if params.AccountId != nil && obj.AccountId != *params.AccountId {
				continue
			}
			if params.Region != nil && obj.Region != *params.Region {
				continue
			}
			filteredResults = append(filteredResults, result)
		}
		startResults = filteredResults
	}

	endS3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "elasticache", "cache_clusters"}, *params.EndReportDate, params.GetRequestEndTimeSelection(), s3Client, new(elasticache.CacheCluster))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	endResults := []interface{}{}
	for endS3DirReader.HasNextFile() {
		resultInterface, err := endS3DirReader.ReadNextFile()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
		endResults = append(endResults, resultInterface...)
	}
	// Filter results
	if params.AccountId != nil || params.Region != nil {
		filteredResults := []interface{}{}
		for _, result := range endResults {
			obj := result.(*elasticache.CacheCluster)
			if params.AccountId != nil && obj.AccountId != *params.AccountId {
				continue
			}
			if params.Region != nil && obj.Region != *params.Region {
				continue
			}
			filteredResults = append(filteredResults, result)
		}
		endResults = filteredResults
	}

	changelog, err := diff.Diff(startResults, endResults)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, changelog)
}

// DiffSingleCacheCluster godoc
// @Summary      Diff CacheCluster
// @Description  get a diff of CacheCluster between two points in time
// @Tags         aws elasticache
// @Produce      json
// @Param        start_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 start_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after)
// @Param		 start_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before' or 'after'." Format(dateTime)
// @Param        end_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 end_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after)
// @Param		 end_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before' or 'after'." Format(dateTime)
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param        arn path string true "The arn of the CacheCluster to retrieve"
// @Success      200  {array}   routes.Diff
// @Failure      400
// @Router       /diff/aws/elasticache/cache_clusters/{arn} [get]
func DiffSingleCacheCluster(c *gin.Context, s3Client *awsS3.Client, s3Bucket string) {
	var params routes.AwsDiffParameters
	if err := c.BindQuery(&params); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	err := params.Process()
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("arn")

	var startObject *elasticache.CacheCluster
	startS3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "elasticache", "cache_clusters"}, *params.StartReportDate, params.GetRequestStartTimeSelection(), s3Client, new(elasticache.CacheCluster))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	startResults := []interface{}{}
	for startS3DirReader.HasNextFile() {
		resultInterface, err := startS3DirReader.ReadNextFile()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
		startResults = append(startResults, resultInterface...)
	}
	// Filter results
	for _, result := range startResults {
		obj := result.(*elasticache.CacheCluster)
		if params.AccountId != nil && obj.AccountId != *params.AccountId {
			continue
		}
		if params.Region != nil && obj.Region != *params.Region {
			continue
		}
		if obj.ARN == id {
			startObject = obj
			break
		}
	}

	var endObject *elasticache.CacheCluster
	endS3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "elasticache", "cache_clusters"}, *params.EndReportDate, params.GetRequestEndTimeSelection(), s3Client, new(elasticache.CacheCluster))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	endResults := []interface{}{}
	for endS3DirReader.HasNextFile() {
		resultInterface, err := endS3DirReader.ReadNextFile()
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
		endResults = append(endResults, resultInterface...)
	}
	// Filter results
	for _, result := range endResults {
		obj := result.(*elasticache.CacheCluster)
		if params.AccountId != nil && obj.AccountId != *params.AccountId {
			continue
		}
		if params.Region != nil && obj.Region != *params.Region {
			continue
		}
		if obj.ARN == id {
			endObject = obj
			break
		}
	}

	if startObject == nil && endObject == nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "No CacheCluster found with arn " + id})
		return
	} else {
		changelog, err := diff.Diff(startObject, endObject)
		if err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, changelog)
		return
	}
}
