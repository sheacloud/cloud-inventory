//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/r3labs/diff/v2"
	"github.com/sheacloud/cloud-inventory/internal/indexedstorage"
	"github.com/sheacloud/cloud-inventory/internal/routes"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud/services/ec2"
)

// ListSecurityGroups godoc
// @Summary      List SecurityGroups
// @Description  get a list of security_groups
// @Tags         aws ec2
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param		 time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after)
// @Param		 time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before' or 'after'." Format(dateTime)
// @Success      200  {array}   ec2.SecurityGroup
// @Failure      400
// @Router       /inventory/aws/ec2/security_groups [get]
func ListSecurityGroups(c *gin.Context, s3Client *awsS3.Client, s3Bucket string) {
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

	s3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "ec2", "security_groups"}, *params.ReportDate, params.GetRequestTimeSelection(), s3Client, new(ec2.SecurityGroup))
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
			obj := result.(*ec2.SecurityGroup)
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

// DiffSecurityGroups godoc
// @Summary      Diff SecurityGroups
// @Description  get a diff of SecurityGroups between two points in time
// @Tags         aws ec2
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
// @Router       /diff/aws/ec2/security_groups [get]
func DiffSecurityGroups(c *gin.Context, s3Client *awsS3.Client, s3Bucket string) {
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

	startS3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "ec2", "security_groups"}, *params.StartReportDate, params.GetRequestStartTimeSelection(), s3Client, new(ec2.SecurityGroup))
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
			obj := result.(*ec2.SecurityGroup)
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

	endS3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "ec2", "security_groups"}, *params.EndReportDate, params.GetRequestEndTimeSelection(), s3Client, new(ec2.SecurityGroup))
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
			obj := result.(*ec2.SecurityGroup)
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

// GetSecurityGroup godoc
// @Summary      Get a specific SecurityGroup
// @Description  Get a specific SecurityGroup by its GroupId
// @Tags         aws ec2
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Param        group_id path string true "The group_id of the SecurityGroup to retrieve"
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param		 time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after)
// @Param		 time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before' or 'after'." Format(dateTime)
// @Success      200  {object}   ec2.SecurityGroup
// @Failure      400
// @Failure 	 404
// @Router       /inventory/aws/ec2/security_groups/{group_id} [get]
func GetSecurityGroup(c *gin.Context, s3Client *awsS3.Client, s3Bucket string) {
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

	id := c.Param("group_id")

	s3DirReader, err := indexedstorage.NewParquetS3DirectoryReader(c.Request.Context(), s3Bucket, []string{"aws", "ec2", "security_groups"}, *params.ReportDate, params.GetRequestTimeSelection(), s3Client, new(ec2.SecurityGroup))
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
		obj := result.(*ec2.SecurityGroup)
		if params.AccountId != nil && obj.AccountId != *params.AccountId {
			continue
		}
		if params.Region != nil && obj.Region != *params.Region {
			continue
		}
		if obj.GroupId == id {
			c.IndentedJSON(200, obj)
			return
		}
	}
	c.AbortWithStatusJSON(404, gin.H{"error": "No SecurityGroup found with group_id " + id})
}
