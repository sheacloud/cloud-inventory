//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_api_routes.tmpl
package ec2

import (
	"github.com/gin-gonic/gin"
	"github.com/r3labs/diff/v2"
	"github.com/sheacloud/cloud-inventory/internal/api/routes"
	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws/ec2"
	"net/url"
	"time"
)

type ListInstancesResponse struct {
	Instances       []*ec2.Instance `json:"instances"`
	PaginationToken *string         `json:"pagination_token,omitempty"`
}

// GetInstancesMetadata godoc
// @Summary      Get Instances Metadata
// @Description  get a list of instances metadata
// @Tags         aws ec2
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.AwsResourceMetadata
// @Failure      400
// @Router       /metadata/aws/ec2/instances [get]
func GetInstancesMetadata(c *gin.Context, dao db.ReaderDAO) {
	reportDateString := c.Query("report_date")
	var reportDate time.Time
	if reportDateString == "" {
		reportDate = time.Now().UTC()
	} else {
		reportDate, _ = time.Parse("2006-01-02", reportDateString)
	}

	reportTimes, err := dao.GetAwsEC2InstanceReportTimes(c, reportDate.UnixMilli())
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	c.IndentedJSON(200, routes.AwsResourceMetadata{
		DateTimes: reportTimes,
		IdField:   "instance_id",
		DisplayFields: []string{
			"instance_id",
		},
	})
}

// ListInstances godoc
// @Summary      List Instances
// @Description  get a list of instances
// @Tags         aws ec2
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param		 time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Param        pagination_token query string false "A pagination token. If this is specified, the next set of results is retrieved. The pagination token is returned in the response."
// @Param        max_results query int false "The maximum number of results to return. Default is 100"
// @Security     ApiKeyAuth
// @Success      200  {object}   ListInstancesResponse
// @Failure      400
// @Router       /inventory/aws/ec2/instances [get]
func ListInstances(c *gin.Context, dao db.ReaderDAO) {
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

	selectedTime, err := dao.GetReferencedAwsEC2InstanceReportTime(c, params.ReportDateUnixMilli, *params.TimeSelection, params.TimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	results, err := dao.ListAwsEC2Instances(c, *selectedTime, params.AccountId, params.Region, nil, nil)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, ListInstancesResponse{
		Instances: results,
	})
}

// GetInstance godoc
// @Summary      Get a specific Instance
// @Description  Get a specific Instance by its InstanceId
// @Tags         aws ec2
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Param        instance_id path string true "The instance_id of the Instance to retrieve"
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param		 time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Security     ApiKeyAuth
// @Success      200  {object}   ec2.Instance
// @Failure      400
// @Failure 	 404
// @Router       /inventory/aws/ec2/instances/{instance_id} [get]
func GetInstance(c *gin.Context, dao db.ReaderDAO) {
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
	id, err := url.QueryUnescape(c.Param("instance_id"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	selectedTime, err := dao.GetReferencedAwsEC2InstanceReportTime(c, params.ReportDateUnixMilli, *params.TimeSelection, params.TimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := dao.GetAwsEC2Instance(c, *selectedTime, id)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, result)
}

// DiffMultiInstances godoc
// @Summary      Diff Instances
// @Description  get a diff of Instances between two points in time
// @Tags         aws ec2
// @Produce      json
// @Param        start_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 start_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 start_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Param        end_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 end_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 end_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.Diff
// @Failure      400
// @Router       /diff/aws/ec2/instances [get]
func DiffMultiInstances(c *gin.Context, dao db.ReaderDAO) {
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

	startSelectedTime, err := dao.GetReferencedAwsEC2InstanceReportTime(c, params.StartReportDateUnixMilli, *params.StartTimeSelection, params.StartTimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	startResults, err := dao.ListAwsEC2Instances(c, *startSelectedTime, params.AccountId, params.Region, nil, nil)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	endSelectedTime, err := dao.GetReferencedAwsEC2InstanceReportTime(c, params.EndReportDateUnixMilli, *params.EndTimeSelection, params.EndTimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	endResults, err := dao.ListAwsEC2Instances(c, *endSelectedTime, params.AccountId, params.Region, nil, nil)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	changelog, err := diff.Diff(startResults, endResults)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, changelog)
}

// DiffSingleInstance godoc
// @Summary      Diff Instance
// @Description  get a diff of Instance between two points in time
// @Tags         aws ec2
// @Produce      json
// @Param        start_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 start_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 start_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Param        end_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 end_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 end_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param        instance_id path string true "The instance_id of the Instance to retrieve"
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.Diff
// @Failure      400
// @Router       /diff/aws/ec2/instances/{instance_id} [get]
func DiffSingleInstance(c *gin.Context, dao db.ReaderDAO) {
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

	id, err := url.QueryUnescape(c.Param("instance_id"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	startSelectedTime, err := dao.GetReferencedAwsEC2InstanceReportTime(c, params.StartReportDateUnixMilli, *params.StartTimeSelection, params.StartTimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	startObject, err := dao.GetAwsEC2Instance(c, *startSelectedTime, id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	endSelectedTime, err := dao.GetReferencedAwsEC2InstanceReportTime(c, params.EndReportDateUnixMilli, *params.EndTimeSelection, params.EndTimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	endObject, err := dao.GetAwsEC2Instance(c, *endSelectedTime, id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	if startObject == nil && endObject == nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "No Instance found with instance_id " + id})
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
