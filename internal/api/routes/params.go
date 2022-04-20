package routes

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sheacloud/cloud-inventory/internal/db"
	"github.com/sheacloud/cloud-inventory/internal/indexedstorage"
)

type AwsQueryParameters struct {
	ReportDate                   *string           `form:"report_date"`
	ReportDateTime               time.Time         `form:"-"`
	AccountId                    *string           `form:"account_id"`
	Region                       *string           `form:"region"`
	TimeSelection                *db.TimeSelection `form:"time_selection"`
	TimeSelectionReferenceString *string           `form:"time_selection_reference"`
	TimeSelectionReference       time.Time         `form:"-"`
	PaginationToken              *string           `form:"pagination_token"`
	PaginationData               *PaginationData   `form:"-"`
	MaxResults                   *int              `form:"max_results"`
}

type PaginationData struct {
	DataFileKeys     []string `json:"data_file_keys"`
	CurrentFileIndex int      `json:"current_file_index"`
	CurrentRowIndex  int      `json:"current_row_index"`
}

func (p *AwsQueryParameters) Process() error {
	var err error
	if p.ReportDate == nil {
		p.ReportDate = aws.String(time.Now().UTC().Format("2006-01-02"))
	}
	fmt.Println(p.ReportDate)
	if p.ReportDateTime, err = time.Parse("2006-01-02", *p.ReportDate); err != nil {
		return err
	}
	if p.TimeSelection == nil {
		selection := db.TimeSelectionLatest
		p.TimeSelection = &selection
		p.TimeSelectionReference = time.Time{}
	}
	if *p.TimeSelection != "latest" && p.TimeSelectionReferenceString == nil {
		return fmt.Errorf("time_selection_reference must be set when time_selection is not 'latest'")
	}
	if p.TimeSelectionReferenceString != nil {
		parsedTime, err := time.Parse(time.RFC3339Nano, *p.TimeSelectionReferenceString)
		if err != nil {
			return fmt.Errorf("time_selection_reference must be a valid RFC3339 timestamp")
		}
		p.TimeSelectionReference = parsedTime
	}
	if p.PaginationToken != nil {
		dataBytes, err := base64.URLEncoding.DecodeString(*p.PaginationToken)
		if err != nil {
			return err
		}
		paginationData := PaginationData{}
		err = json.Unmarshal(dataBytes, &paginationData)
		if err != nil {
			return err
		}
		p.PaginationData = &paginationData
	}
	if p.MaxResults == nil {
		p.MaxResults = aws.Int(100)
	} else {
		if *p.MaxResults < 1 {
			return fmt.Errorf("max_results must be >= 1")
		} else if *p.MaxResults > 100 {
			return fmt.Errorf("max_results must be <= 100")
		}
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
	StartReportDate                   *string           `form:"start_report_date" binding:"required"`
	StartReportDateTime               time.Time         `form:"-"`
	StartTimeSelection                *db.TimeSelection `form:"start_time_selection"`
	StartTimeSelectionReferenceString *string           `form:"start_time_selection_reference"`
	StartTimeSelectionReference       time.Time         `form:"-"`
	EndReportDate                     *string           `form:"end_report_date" binding:"required"`
	EndReportDateTime                 time.Time         `form:"-"`
	EndTimeSelection                  *db.TimeSelection `form:"end_time_selection"`
	EndTimeSelectionReferenceString   *string           `form:"end_time_selection_reference"`
	EndTimeSelectionReference         time.Time         `form:"-"`
	AccountId                         *string           `form:"account_id"`
	Region                            *string           `form:"region"`
}

func (p *AwsDiffParameters) Process() error {
	var err error
	if p.StartReportDateTime, err = time.Parse("2006-01-02", *p.StartReportDate); err != nil {
		return err
	}
	if p.EndReportDateTime, err = time.Parse("2006-01-02", *p.EndReportDate); err != nil {
		return err
	}

	if p.StartTimeSelection == nil {
		tmp := db.TimeSelectionLatest
		p.StartTimeSelection = &tmp
		p.StartTimeSelectionReference = time.Time{}
	}
	if *p.StartTimeSelection != "latest" && p.StartTimeSelectionReferenceString == nil {
		return fmt.Errorf("start_time_selection_reference must be set when start_time_selection is not 'latest'")
	}
	if p.StartTimeSelectionReferenceString != nil {
		parsedTime, err := time.Parse(time.RFC3339Nano, *p.StartTimeSelectionReferenceString)
		if err != nil {
			return fmt.Errorf("start_time_selection_reference must be a valid RFC3339 timestamp")
		}
		p.StartTimeSelectionReference = parsedTime
	}

	if p.EndTimeSelection == nil {
		tmp := db.TimeSelectionLatest
		p.EndTimeSelection = &tmp
		p.EndTimeSelectionReference = time.Time{}
	}
	if *p.EndTimeSelection != "latest" && p.EndTimeSelectionReferenceString == nil {
		return fmt.Errorf("end_time_selection_reference must be set when end_time_selection is not 'latest'")
	}
	if p.EndTimeSelectionReferenceString != nil {
		parsedTime, err := time.Parse(time.RFC3339Nano, *p.EndTimeSelectionReferenceString)
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
