//AUTOGENERATED CODE DO NOT EDIT
package efs

import (
	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/r3labs/diff/v2"
	"github.com/sheacloud/cloud-inventory/internal/indexedstorage"
	"github.com/sheacloud/cloud-inventory/internal/routes"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud/services/efs"
)

// ListFileSystems godoc
// @Summary      List FileSystems
// @Description  get a list of filesystems
// @Tags         aws efs
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param		 time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after)
// @Param		 time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before' or 'after'." Format(dateTime)
// @Success      200  {array}   efs.FileSystemDescription
// @Failure      400
// @Router       /inventory/aws/efs/filesystems [get]
func ListFileSystems(c *gin.Context, s3Client *awsS3.Client, s3Bucket string) {
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

	s3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "efs", "filesystems"}, *params.ReportDate, params.GetRequestTimeSelection(), s3Client, new(efs.FileSystemDescription))
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
			obj := result.(*efs.FileSystemDescription)
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

// GetFileSystemDescription godoc
// @Summary      Get a specific FileSystemDescription
// @Description  Get a specific FileSystemDescription by its FileSystemId
// @Tags         aws efs
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Param        file_system_id path string true "The file_system_id of the FileSystemDescription to retrieve"
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param		 time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after)
// @Param		 time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before' or 'after'." Format(dateTime)
// @Success      200  {object}   efs.FileSystemDescription
// @Failure      400
// @Failure 	 404
// @Router       /inventory/aws/efs/filesystems/{file_system_id} [get]
func GetFileSystemDescription(c *gin.Context, s3Client *awsS3.Client, s3Bucket string) {
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

	id := c.Param("file_system_id")

	s3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "efs", "filesystems"}, *params.ReportDate, params.GetRequestTimeSelection(), s3Client, new(efs.FileSystemDescription))
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
		obj := result.(*efs.FileSystemDescription)
		if params.AccountId != nil && obj.AccountId != *params.AccountId {
			continue
		}
		if params.Region != nil && obj.Region != *params.Region {
			continue
		}
		if obj.FileSystemId == id {
			c.IndentedJSON(200, obj)
			return
		}
	}
	c.AbortWithStatusJSON(404, gin.H{"error": "No FileSystemDescription found with file_system_id " + id})
}

// DiffMultiFileSystems godoc
// @Summary      Diff FileSystems
// @Description  get a diff of FileSystems between two points in time
// @Tags         aws efs
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
// @Router       /diff/aws/efs/filesystems [get]
func DiffMultiFileSystems(c *gin.Context, s3Client *awsS3.Client, s3Bucket string) {
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

	startS3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "efs", "filesystems"}, *params.StartReportDate, params.GetRequestStartTimeSelection(), s3Client, new(efs.FileSystemDescription))
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
			obj := result.(*efs.FileSystemDescription)
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

	endS3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "efs", "filesystems"}, *params.EndReportDate, params.GetRequestEndTimeSelection(), s3Client, new(efs.FileSystemDescription))
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
			obj := result.(*efs.FileSystemDescription)
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

// DiffSingleFileSystemDescription godoc
// @Summary      Diff FileSystemDescription
// @Description  get a diff of FileSystemDescription between two points in time
// @Tags         aws efs
// @Produce      json
// @Param        start_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 start_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after)
// @Param		 start_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before' or 'after'." Format(dateTime)
// @Param        end_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 end_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after)
// @Param		 end_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before' or 'after'." Format(dateTime)
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param        file_system_id path string true "The file_system_id of the FileSystemDescription to retrieve"
// @Success      200  {array}   routes.Diff
// @Failure      400
// @Router       /diff/aws/efs/filesystems/{file_system_id} [get]
func DiffSingleFileSystemDescription(c *gin.Context, s3Client *awsS3.Client, s3Bucket string) {
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

	id := c.Param("file_system_id")

	var startObject *efs.FileSystemDescription
	startS3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "efs", "filesystems"}, *params.StartReportDate, params.GetRequestStartTimeSelection(), s3Client, new(efs.FileSystemDescription))
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
		obj := result.(*efs.FileSystemDescription)
		if params.AccountId != nil && obj.AccountId != *params.AccountId {
			continue
		}
		if params.Region != nil && obj.Region != *params.Region {
			continue
		}
		if obj.FileSystemId == id {
			startObject = obj
			break
		}
	}

	var endObject *efs.FileSystemDescription
	endS3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "efs", "filesystems"}, *params.EndReportDate, params.GetRequestEndTimeSelection(), s3Client, new(efs.FileSystemDescription))
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
		obj := result.(*efs.FileSystemDescription)
		if params.AccountId != nil && obj.AccountId != *params.AccountId {
			continue
		}
		if params.Region != nil && obj.Region != *params.Region {
			continue
		}
		if obj.FileSystemId == id {
			endObject = obj
			break
		}
	}

	if startObject == nil && endObject == nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "No FileSystemDescription found with file_system_id " + id})
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
