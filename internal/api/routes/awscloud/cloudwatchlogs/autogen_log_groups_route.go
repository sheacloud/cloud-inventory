//AUTOGENERATED CODE DO NOT EDIT
package cloudwatchlogs

import (
	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/r3labs/diff/v2"
	"github.com/sheacloud/cloud-inventory/internal/api/routes"
	"github.com/sheacloud/cloud-inventory/internal/indexedstorage"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud/services/cloudwatchlogs"
	"time"
)

// GetLogGroupsMetadata godoc
// @Summary      Get LogGroups Metadata
// @Description  get a list of log_groups metadata
// @Tags         aws cloudwatchlogs
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Success      200  {array}   routes.AwsResourceMetadata
// @Failure      400
// @Router       /metadata/aws/cloudwatchlogs/log_groups [get]
func GetLogGroupsMetadata(c *gin.Context, s3Client *awsS3.Client, s3Bucket string) {
	reportDate := c.Query("report_date")
	if reportDate == "" {
		reportDate = time.Now().UTC().Format("2006-01-02")
	}

	s3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "cloudwatchlogs", "log_groups"}, reportDate, indexedstorage.RequestTimeSelection{}, s3Client, new(cloudwatchlogs.LogGroup))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	reportTimes, err := s3DirReader.GetAvailableDateTimes()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, routes.AwsResourceMetadata{
		DateTimes: reportTimes,
		IdField:   "arn",
	})
}

// ListLogGroups godoc
// @Summary      List LogGroups
// @Description  get a list of log_groups
// @Tags         aws cloudwatchlogs
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param		 time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Success      200  {array}   cloudwatchlogs.LogGroup
// @Failure      400
// @Router       /inventory/aws/cloudwatchlogs/log_groups [get]
func ListLogGroups(c *gin.Context, s3Client *awsS3.Client, s3Bucket string) {
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

	s3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "cloudwatchlogs", "log_groups"}, *params.ReportDate, params.GetRequestTimeSelection(), s3Client, new(cloudwatchlogs.LogGroup))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = s3DirReader.DetermineDataFiles()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
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
			obj := result.(*cloudwatchlogs.LogGroup)
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

// GetLogGroup godoc
// @Summary      Get a specific LogGroup
// @Description  Get a specific LogGroup by its Arn
// @Tags         aws cloudwatchlogs
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Param        arn path string true "The arn of the LogGroup to retrieve"
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param		 time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Success      200  {object}   cloudwatchlogs.LogGroup
// @Failure      400
// @Failure 	 404
// @Router       /inventory/aws/cloudwatchlogs/log_groups/{arn} [get]
func GetLogGroup(c *gin.Context, s3Client *awsS3.Client, s3Bucket string) {
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

	s3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "cloudwatchlogs", "log_groups"}, *params.ReportDate, params.GetRequestTimeSelection(), s3Client, new(cloudwatchlogs.LogGroup))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = s3DirReader.DetermineDataFiles()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
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
		obj := result.(*cloudwatchlogs.LogGroup)
		if params.AccountId != nil && obj.AccountId != *params.AccountId {
			continue
		}
		if params.Region != nil && obj.Region != *params.Region {
			continue
		}
		if obj.Arn == id {
			c.IndentedJSON(200, obj)
			return
		}
	}
	c.AbortWithStatusJSON(404, gin.H{"error": "No LogGroup found with arn " + id})
}

// DiffMultiLogGroups godoc
// @Summary      Diff LogGroups
// @Description  get a diff of LogGroups between two points in time
// @Tags         aws cloudwatchlogs
// @Produce      json
// @Param        start_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 start_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 start_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Param        end_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 end_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 end_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Success      200  {array}   routes.Diff
// @Failure      400
// @Router       /diff/aws/cloudwatchlogs/log_groups [get]
func DiffMultiLogGroups(c *gin.Context, s3Client *awsS3.Client, s3Bucket string) {
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

	startS3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "cloudwatchlogs", "log_groups"}, *params.StartReportDate, params.GetRequestStartTimeSelection(), s3Client, new(cloudwatchlogs.LogGroup))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = startS3DirReader.DetermineDataFiles()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
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
			obj := result.(*cloudwatchlogs.LogGroup)
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

	endS3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "cloudwatchlogs", "log_groups"}, *params.EndReportDate, params.GetRequestEndTimeSelection(), s3Client, new(cloudwatchlogs.LogGroup))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = endS3DirReader.DetermineDataFiles()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
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
			obj := result.(*cloudwatchlogs.LogGroup)
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

// DiffSingleLogGroup godoc
// @Summary      Diff LogGroup
// @Description  get a diff of LogGroup between two points in time
// @Tags         aws cloudwatchlogs
// @Produce      json
// @Param        start_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 start_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 start_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Param        end_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 end_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 end_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param        arn path string true "The arn of the LogGroup to retrieve"
// @Success      200  {array}   routes.Diff
// @Failure      400
// @Router       /diff/aws/cloudwatchlogs/log_groups/{arn} [get]
func DiffSingleLogGroup(c *gin.Context, s3Client *awsS3.Client, s3Bucket string) {
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

	var startObject *cloudwatchlogs.LogGroup
	startS3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "cloudwatchlogs", "log_groups"}, *params.StartReportDate, params.GetRequestStartTimeSelection(), s3Client, new(cloudwatchlogs.LogGroup))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = startS3DirReader.DetermineDataFiles()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
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
		obj := result.(*cloudwatchlogs.LogGroup)
		if params.AccountId != nil && obj.AccountId != *params.AccountId {
			continue
		}
		if params.Region != nil && obj.Region != *params.Region {
			continue
		}
		if obj.Arn == id {
			startObject = obj
			break
		}
	}

	var endObject *cloudwatchlogs.LogGroup
	endS3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "cloudwatchlogs", "log_groups"}, *params.EndReportDate, params.GetRequestEndTimeSelection(), s3Client, new(cloudwatchlogs.LogGroup))
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	err = endS3DirReader.DetermineDataFiles()
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
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
		obj := result.(*cloudwatchlogs.LogGroup)
		if params.AccountId != nil && obj.AccountId != *params.AccountId {
			continue
		}
		if params.Region != nil && obj.Region != *params.Region {
			continue
		}
		if obj.Arn == id {
			endObject = obj
			break
		}
	}

	if startObject == nil && endObject == nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "No LogGroup found with arn " + id})
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
