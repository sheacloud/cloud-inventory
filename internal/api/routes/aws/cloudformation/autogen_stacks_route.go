//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_api_routes.tmpl
package cloudformation

import (
	"github.com/gin-gonic/gin"
	"github.com/r3labs/diff/v2"
	"github.com/sheacloud/cloud-inventory/internal/api/routes"
	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws/cloudformation"
	"net/url"
	"time"
)

type ListStacksResponse struct {
	Stacks          []*cloudformation.Stack `json:"stacks"`
	PaginationToken *string                 `json:"pagination_token,omitempty"`
}

// GetStacksMetadata godoc
// @Summary      Get Stacks Metadata
// @Description  get a list of stacks metadata
// @Tags         aws cloudformation
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.AwsResourceMetadata
// @Failure      400
// @Router       /metadata/aws/cloudformation/stacks [get]
func GetStacksMetadata(c *gin.Context, dao db.ReaderDAO) {
	reportDateString := c.Query("report_date")
	var reportDate time.Time
	if reportDateString == "" {
		reportDate = time.Now().UTC()
	} else {
		reportDate, _ = time.Parse("2006-01-02", reportDateString)
	}

	reportTimes, err := dao.GetAwsCloudFormationStackReportTimes(c, reportDate.UnixMilli())
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	c.IndentedJSON(200, routes.AwsResourceMetadata{
		DateTimes: reportTimes,
		IdField:   "stack_id",
		DisplayFields: []string{
			"stack_name",
		},
	})
}

// ListStacks godoc
// @Summary      List Stacks
// @Description  get a list of stacks
// @Tags         aws cloudformation
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param		 time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Param        pagination_token query string false "A pagination token. If this is specified, the next set of results is retrieved. The pagination token is returned in the response."
// @Param        max_results query int false "The maximum number of results to return. Default is 100"
// @Security     ApiKeyAuth
// @Success      200  {object}   ListStacksResponse
// @Failure      400
// @Router       /inventory/aws/cloudformation/stacks [get]
func ListStacks(c *gin.Context, dao db.ReaderDAO) {
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

	selectedTime, err := dao.GetReferencedAwsCloudFormationStackReportTime(c, params.ReportDateUnixMilli, *params.TimeSelection, params.TimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	results, err := dao.ListAwsCloudFormationStacks(c, *selectedTime, params.AccountId, params.Region, nil, nil)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, ListStacksResponse{
		Stacks: results,
	})
}

// GetStack godoc
// @Summary      Get a specific Stack
// @Description  Get a specific Stack by its StackId
// @Tags         aws cloudformation
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Param        stack_id path string true "The stack_id of the Stack to retrieve"
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param		 time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Security     ApiKeyAuth
// @Success      200  {object}   cloudformation.Stack
// @Failure      400
// @Failure 	 404
// @Router       /inventory/aws/cloudformation/stacks/{stack_id} [get]
func GetStack(c *gin.Context, dao db.ReaderDAO) {
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
	id, err := url.QueryUnescape(c.Param("stack_id"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	selectedTime, err := dao.GetReferencedAwsCloudFormationStackReportTime(c, params.ReportDateUnixMilli, *params.TimeSelection, params.TimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := dao.GetAwsCloudFormationStack(c, *selectedTime, id)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, result)
}

// DiffMultiStacks godoc
// @Summary      Diff Stacks
// @Description  get a diff of Stacks between two points in time
// @Tags         aws cloudformation
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
// @Router       /diff/aws/cloudformation/stacks [get]
func DiffMultiStacks(c *gin.Context, dao db.ReaderDAO) {
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

	startSelectedTime, err := dao.GetReferencedAwsCloudFormationStackReportTime(c, params.StartReportDateUnixMilli, *params.StartTimeSelection, params.StartTimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	startResults, err := dao.ListAwsCloudFormationStacks(c, *startSelectedTime, params.AccountId, params.Region, nil, nil)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	endSelectedTime, err := dao.GetReferencedAwsCloudFormationStackReportTime(c, params.EndReportDateUnixMilli, *params.EndTimeSelection, params.EndTimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	endResults, err := dao.ListAwsCloudFormationStacks(c, *endSelectedTime, params.AccountId, params.Region, nil, nil)
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

// DiffSingleStack godoc
// @Summary      Diff Stack
// @Description  get a diff of Stack between two points in time
// @Tags         aws cloudformation
// @Produce      json
// @Param        start_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 start_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 start_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Param        end_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 end_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 end_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param        stack_id path string true "The stack_id of the Stack to retrieve"
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.Diff
// @Failure      400
// @Router       /diff/aws/cloudformation/stacks/{stack_id} [get]
func DiffSingleStack(c *gin.Context, dao db.ReaderDAO) {
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

	id, err := url.QueryUnescape(c.Param("stack_id"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	startSelectedTime, err := dao.GetReferencedAwsCloudFormationStackReportTime(c, params.StartReportDateUnixMilli, *params.StartTimeSelection, params.StartTimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	startObject, err := dao.GetAwsCloudFormationStack(c, *startSelectedTime, id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	endSelectedTime, err := dao.GetReferencedAwsCloudFormationStackReportTime(c, params.EndReportDateUnixMilli, *params.EndTimeSelection, params.EndTimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	endObject, err := dao.GetAwsCloudFormationStack(c, *endSelectedTime, id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	if startObject == nil && endObject == nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "No Stack found with stack_id " + id})
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
