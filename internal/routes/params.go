package routes

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sheacloud/cloud-inventory/internal/indexedstorage"
)

type AwsQueryParameters struct {
	ReportDate                   *string   `form:"report_date"`
	AccountId                    *string   `form:"account_id"`
	Region                       *string   `form:"region"`
	TimeSelection                *string   `form:"time_selection"`
	TimeSelectionReferenceString *string   `form:"time_selection_reference"`
	TimeSelectionReference       time.Time `form:"-"`
}

func (p *AwsQueryParameters) Process() error {
	if p.ReportDate == nil {
		p.ReportDate = aws.String(time.Now().UTC().Format("2006-01-02"))
	}
	if p.TimeSelection == nil {
		p.TimeSelection = aws.String("latest")
		p.TimeSelectionReference = time.Time{}
	}
	if *p.TimeSelection != "latest" && p.TimeSelectionReferenceString == nil {
		return fmt.Errorf("time_selection_reference must be set when time_selection is not 'latest'")
	}
	if p.TimeSelectionReferenceString != nil {
		parsedTime, err := time.Parse(time.RFC3339, *p.TimeSelectionReferenceString)
		if err != nil {
			return fmt.Errorf("time_selection_reference must be a valid RFC3339 timestamp")
		}
		p.TimeSelectionReference = parsedTime
	}

	return nil
}

func (p *AwsQueryParameters) GetRequestTimeSelection() indexedstorage.RequestTimeSelection {
	return indexedstorage.RequestTimeSelection{
		Option:         indexedstorage.RequestTimeSelectionOption(*p.TimeSelection),
		ReferencedTime: p.TimeSelectionReference,
	}
}

type AwsDiffParameters struct {
	StartReportDate                   *string   `form:"start_report_date" binding:"required"`
	StartTimeSelection                *string   `form:"start_time_selection"`
	StartTimeSelectionReferenceString *string   `form:"start_time_selection_reference"`
	StartTimeSelectionReference       time.Time `form:"-"`
	EndReportDate                     *string   `form:"end_report_date" binding:"required"`
	EndTimeSelection                  *string   `form:"end_time_selection"`
	EndTimeSelectionReferenceString   *string   `form:"end_time_selection_reference"`
	EndTimeSelectionReference         time.Time `form:"-"`
	AccountId                         *string   `form:"account_id"`
	Region                            *string   `form:"region"`
}

func (p *AwsDiffParameters) Process() error {
	if p.StartTimeSelection == nil {
		p.StartTimeSelection = aws.String("latest")
		p.StartTimeSelectionReference = time.Time{}
	}
	if *p.StartTimeSelection != "latest" && p.StartTimeSelectionReferenceString == nil {
		return fmt.Errorf("start_time_selection_reference must be set when start_time_selection is not 'latest'")
	}
	if p.StartTimeSelectionReferenceString != nil {
		parsedTime, err := time.Parse(time.RFC3339, *p.StartTimeSelectionReferenceString)
		if err != nil {
			return fmt.Errorf("start_time_selection_reference must be a valid RFC3339 timestamp")
		}
		p.StartTimeSelectionReference = parsedTime
	}

	if p.EndTimeSelection == nil {
		p.EndTimeSelection = aws.String("latest")
		p.EndTimeSelectionReference = time.Time{}
	}
	if *p.EndTimeSelection != "latest" && p.EndTimeSelectionReferenceString == nil {
		return fmt.Errorf("end_time_selection_reference must be set when end_time_selection is not 'latest'")
	}
	if p.EndTimeSelectionReferenceString != nil {
		parsedTime, err := time.Parse(time.RFC3339, *p.EndTimeSelectionReferenceString)
		if err != nil {
			return fmt.Errorf("end_time_selection_reference must be a valid RFC3339 timestamp")
		}
		p.EndTimeSelectionReference = parsedTime
	}

	return nil
}

func (p *AwsDiffParameters) GetRequestStartTimeSelection() indexedstorage.RequestTimeSelection {
	return indexedstorage.RequestTimeSelection{
		Option:         indexedstorage.RequestTimeSelectionOption(*p.StartTimeSelection),
		ReferencedTime: p.StartTimeSelectionReference,
	}
}

func (p *AwsDiffParameters) GetRequestEndTimeSelection() indexedstorage.RequestTimeSelection {
	return indexedstorage.RequestTimeSelection{
		Option:         indexedstorage.RequestTimeSelectionOption(*p.EndTimeSelection),
		ReferencedTime: p.EndTimeSelectionReference,
	}
}
