//AUTOGENERATED CODE DO NOT EDIT
package sns

import (
	"github.com/gin-gonic/gin"
	"github.com/r3labs/diff/v2"
	"github.com/sheacloud/cloud-inventory/internal/api/routes"
	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/pkg/aws/sns"
	"net/url"
	"time"
)

type ListSubscriptionsResponse struct {
	Subscriptions   []*sns.Subscription `json:"subscriptions"`
	PaginationToken *string             `json:"pagination_token,omitempty"`
}

// GetSubscriptionsMetadata godoc
// @Summary      Get Subscriptions Metadata
// @Description  get a list of subscriptions metadata
// @Tags         aws sns
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.AwsResourceMetadata
// @Failure      400
// @Router       /metadata/aws/sns/subscriptions [get]
func GetSubscriptionsMetadata(c *gin.Context, dao db.DAO) {
	reportDateString := c.Query("report_date")
	var reportDate time.Time
	if reportDateString == "" {
		reportDate = time.Now().UTC()
	} else {
		reportDate, _ = time.Parse("2006-01-02", reportDateString)
	}

	reportTimes, err := dao.AWS().SNS().GetSubscriptionReportTimes(c, reportDate)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}

	c.IndentedJSON(200, routes.AwsResourceMetadata{
		DateTimes: reportTimes,
		IdField:   "subscription_arn",
		DisplayFields: []string{
			"topic_arn",
			"endpoint",
		},
	})
}

// ListSubscriptions godoc
// @Summary      List Subscriptions
// @Description  get a list of subscriptions
// @Tags         aws sns
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param		 time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Param        pagination_token query string false "A pagination token. If this is specified, the next set of results is retrieved. The pagination token is returned in the response."
// @Param        max_results query int false "The maximum number of results to return. Default is 100"
// @Security     ApiKeyAuth
// @Success      200  {object}   ListSubscriptionsResponse
// @Failure      400
// @Router       /inventory/aws/sns/subscriptions [get]
func ListSubscriptions(c *gin.Context, dao db.DAO) {
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

	selectedTime, err := dao.AWS().SNS().GetReferencedSubscriptionReportTime(c, params.ReportDateTime, *params.TimeSelection, params.TimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	results, err := dao.AWS().SNS().ListSubscriptions(c, *selectedTime, params.AccountId, params.Region, nil, nil)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, ListSubscriptionsResponse{
		Subscriptions: results,
	})
}

// GetSubscription godoc
// @Summary      Get a specific Subscription
// @Description  Get a specific Subscription by its SubscriptionArn
// @Tags         aws sns
// @Produce      json
// @Param        report_date query string false  "Which date to pull data from. Current date by default" Format(date)
// @Param        subscription_arn path string true "The subscription_arn of the Subscription to retrieve"
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param		 time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Security     ApiKeyAuth
// @Success      200  {object}   sns.Subscription
// @Failure      400
// @Failure 	 404
// @Router       /inventory/aws/sns/subscriptions/{subscription_arn} [get]
func GetSubscription(c *gin.Context, dao db.DAO) {
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
	id, err := url.QueryUnescape(c.Param("subscription_arn"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	selectedTime, err := dao.AWS().SNS().GetReferencedSubscriptionReportTime(c, params.ReportDateTime, *params.TimeSelection, params.TimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	result, err := dao.AWS().SNS().GetSubscription(c, *selectedTime, id)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, result)
}

// DiffMultiSubscriptions godoc
// @Summary      Diff Subscriptions
// @Description  get a diff of Subscriptions between two points in time
// @Tags         aws sns
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
// @Router       /diff/aws/sns/subscriptions [get]
func DiffMultiSubscriptions(c *gin.Context, dao db.DAO) {
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

	startSelectedTime, err := dao.AWS().SNS().GetReferencedSubscriptionReportTime(c, params.StartReportDateTime, *params.StartTimeSelection, params.StartTimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	startResults, err := dao.AWS().SNS().ListSubscriptions(c, *startSelectedTime, params.AccountId, params.Region, nil, nil)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	endSelectedTime, err := dao.AWS().SNS().GetReferencedSubscriptionReportTime(c, params.EndReportDateTime, *params.EndTimeSelection, params.EndTimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	endResults, err := dao.AWS().SNS().ListSubscriptions(c, *endSelectedTime, params.AccountId, params.Region, nil, nil)
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

// DiffSingleSubscription godoc
// @Summary      Diff Subscription
// @Description  get a diff of Subscription between two points in time
// @Tags         aws sns
// @Produce      json
// @Param        start_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 start_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 start_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Param        end_report_date query string true  "Which date to pull data from. Current date by default" Format(date)
// @Param		 end_time_selection query string false  "How to select the time range to pull data from. 'latest' by default" Enums(latest, before, after, at)
// @Param		 end_time_selection_reference query string false  "The reference time to use when selecting the time range to pull data from. Only used when time_selection is 'before', 'after', or 'at'." Format(dateTime)
// @Param		 account_id query string false  "A specific account to pull data from. All accounts by default"
// @Param		 region query string false  "A specific region to pull data from. All regions by default"
// @Param        subscription_arn path string true "The subscription_arn of the Subscription to retrieve"
// @Security     ApiKeyAuth
// @Success      200  {array}   routes.Diff
// @Failure      400
// @Router       /diff/aws/sns/subscriptions/{subscription_arn} [get]
func DiffSingleSubscription(c *gin.Context, dao db.DAO) {
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

	id, err := url.QueryUnescape(c.Param("subscription_arn"))
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	startSelectedTime, err := dao.AWS().SNS().GetReferencedSubscriptionReportTime(c, params.StartReportDateTime, *params.StartTimeSelection, params.StartTimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	startObject, err := dao.AWS().SNS().GetSubscription(c, *startSelectedTime, id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	endSelectedTime, err := dao.AWS().SNS().GetReferencedSubscriptionReportTime(c, params.EndReportDateTime, *params.EndTimeSelection, params.EndTimeSelectionReference)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	endObject, err := dao.AWS().SNS().GetSubscription(c, *endSelectedTime, id)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	if startObject == nil && endObject == nil {
		c.AbortWithStatusJSON(404, gin.H{"error": "No Subscription found with subscription_arn " + id})
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
