//AUTOGENERATED CODE DO NOT EDIT
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

type ListNetworkInterfacesResponse struct {
	NetworkInterfaces []*ec2.NetworkInterface `json:"network_interfaces"`
	PaginationToken   *string                 `json:"pagination_token,omitempty"`
}

// GetNetworkInterfacesMetadata godoc
// @Summary      Get NetworkInterfaces Metadata
// @Description  get a list of network_interfaces metadata
// @Tags         aws ec2
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.AwsResourceMetadata
// @Failure      400
// @Router       /metadata/aws/ec2/network_interfaces [get]
func GetNetworkInterfacesMetadata(c *gin.Context, dao db.DAO) {
	reportDateString := c.Query("report_date")
	var reportDate time.Time
	if reportDateString == "" {
		reportDate = time.Now().UTC()
	} else {
		reportDate, _ = time.Parse("2006-01-02", reportDateString)
	}

	reportTimes, err := dao.AWS().EC2().GetNetworkInterfaceReportTimes(c, reportDate)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	c.IndentedJSON(200, routes.AwsResourceMetadata{
		DateTimes: reportTimes,
		IdField:   "network_interface_id",
		DisplayFields: []string{
			"network_interface_id",
		},
	})
}

// ListNetworkInterfaces godoc
// @Summary      List NetworkInterfaces
// @Description  get a list of network_interfaces
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
// @Success      200  {object}   ListNetworkInterfacesResponse
// @Failure      400
// @Router       /inventory/aws/ec2/network_interfaces [get]
func ListNetworkInterfaces(c *gin.Context, dao db.DAO) {
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

	selectedTime, err := dao.AWS().EC2().GetReferencedNetworkInterfaceReportTime(c, params.ReportDateTime, *params.TimeSelection, params.TimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	results, err := dao.AWS().EC2().ListNetworkInterfaces(c, *selectedTime, params.AccountId, params.Region, nil, nil)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, ListNetworkInterfacesResponse{
		NetworkInterfaces: results,
	})
}

// GetNetworkInterface godoc
// @Summary      Get a specific NetworkInterface
// @Description  Get a specific NetworkInterface by its NetworkInterfaceId
// @Tags         aws ec2
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Param        network_interface_id path string true "The network_interface_id of the NetworkInterface to retrieve"
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param		 time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Security     ApiKeyAuth
// @Success      200  {object}   ec2.NetworkInterface
// @Failure      400
// @Failure 	 404
// @Router       /inventory/aws/ec2/network_interfaces/{network_interface_id} [get]
func GetNetworkInterface(c *gin.Context, dao db.DAO) {
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
	id, err := url.QueryUnescape(c.Param("network_interface_id"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	selectedTime, err := dao.AWS().EC2().GetReferencedNetworkInterfaceReportTime(c, params.ReportDateTime, *params.TimeSelection, params.TimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := dao.AWS().EC2().GetNetworkInterface(c, *selectedTime, id)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, result)
}

// DiffMultiNetworkInterfaces godoc
// @Summary      Diff NetworkInterfaces
// @Description  get a diff of NetworkInterfaces between two points in time
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
// @Router       /diff/aws/ec2/network_interfaces [get]
func DiffMultiNetworkInterfaces(c *gin.Context, dao db.DAO) {
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

	startSelectedTime, err := dao.AWS().EC2().GetReferencedNetworkInterfaceReportTime(c, params.StartReportDateTime, *params.StartTimeSelection, params.StartTimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	startResults, err := dao.AWS().EC2().ListNetworkInterfaces(c, *startSelectedTime, params.AccountId, params.Region, nil, nil)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	endSelectedTime, err := dao.AWS().EC2().GetReferencedNetworkInterfaceReportTime(c, params.EndReportDateTime, *params.EndTimeSelection, params.EndTimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	endResults, err := dao.AWS().EC2().ListNetworkInterfaces(c, *endSelectedTime, params.AccountId, params.Region, nil, nil)
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

// DiffSingleNetworkInterface godoc
// @Summary      Diff NetworkInterface
// @Description  get a diff of NetworkInterface between two points in time
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
// @Param        network_interface_id path string true "The network_interface_id of the NetworkInterface to retrieve"
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.Diff
// @Failure      400
// @Router       /diff/aws/ec2/network_interfaces/{network_interface_id} [get]
func DiffSingleNetworkInterface(c *gin.Context, dao db.DAO) {
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

	id, err := url.QueryUnescape(c.Param("network_interface_id"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	startSelectedTime, err := dao.AWS().EC2().GetReferencedNetworkInterfaceReportTime(c, params.StartReportDateTime, *params.StartTimeSelection, params.StartTimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	startObject, err := dao.AWS().EC2().GetNetworkInterface(c, *startSelectedTime, id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	endSelectedTime, err := dao.AWS().EC2().GetReferencedNetworkInterfaceReportTime(c, params.EndReportDateTime, *params.EndTimeSelection, params.EndTimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	endObject, err := dao.AWS().EC2().GetNetworkInterface(c, *endSelectedTime, id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	if startObject == nil && endObject == nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "No NetworkInterface found with network_interface_id " + id})
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