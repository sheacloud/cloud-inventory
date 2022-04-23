package routes

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sheacloud/cloud-inventory/internal/db"
)

type AwsQueryParameters struct {
	ReportDate             *string           `form:"report_date"`
	ReportDateUnixMilli    int64             `form:"-"`
	AccountId              *string           `form:"account_id"`
	Region                 *string           `form:"region"`
	TimeSelection          *db.TimeSelection `form:"time_selection"`
	TimeSelectionReference int64             `form:"time_selection_reference"`
	PaginationToken        *string           `form:"pagination_token"`
	PaginationData         *PaginationData   `form:"-"`
	MaxResults             *int              `form:"max_results"`
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
	reportDate, err := time.Parse("2006-01-02", *p.ReportDate)
	if err != nil {
		return err
	}
	p.ReportDateUnixMilli = reportDate.UnixMilli()

	if p.TimeSelection == nil {
		selection := db.TimeSelectionLatest
		p.TimeSelection = &selection
		p.TimeSelectionReference = 0
	}
	if *p.TimeSelection != "latest" && p.TimeSelectionReference == 0 {
		return fmt.Errorf("time_selection_reference must be set when time_selection is not 'latest'")
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

type AwsDiffParameters struct {
	StartReportDate             *string           `form:"start_report_date" binding:"required"`
	StartReportDateUnixMilli    int64             `form:"-"`
	StartTimeSelection          *db.TimeSelection `form:"start_time_selection"`
	StartTimeSelectionReference int64             `form:"start_time_selection_reference"`
	EndReportDate               *string           `form:"end_report_date" binding:"required"`
	EndReportDateUnixMilli      int64             `form:"-"`
	EndTimeSelection            *db.TimeSelection `form:"end_time_selection"`
	EndTimeSelectionReference   int64             `form:"end_time_selection_reference"`
	AccountId                   *string           `form:"account_id"`
	Region                      *string           `form:"region"`
}

func (p *AwsDiffParameters) Process() error {
	var err error
	startReportDate, err := time.Parse("2006-01-02", *p.StartReportDate)
	if err != nil {
		return err
	}
	p.StartReportDateUnixMilli = startReportDate.UnixMilli()

	endReportDate, err := time.Parse("2006-01-02", *p.EndReportDate)
	if err != nil {
		return err
	}
	p.EndReportDateUnixMilli = endReportDate.UnixMilli()

	if p.StartTimeSelection == nil {
		tmp := db.TimeSelectionLatest
		p.StartTimeSelection = &tmp
		p.StartTimeSelectionReference = 0
	}
	if *p.StartTimeSelection != "latest" && p.StartTimeSelectionReference == 0 {
		return fmt.Errorf("start_time_selection_reference must be set when start_time_selection is not 'latest'")
	}

	if p.EndTimeSelection == nil {
		tmp := db.TimeSelectionLatest
		p.EndTimeSelection = &tmp
		p.EndTimeSelectionReference = 0
	}
	if *p.EndTimeSelection != "latest" && p.EndTimeSelectionReference == 0 {
		return fmt.Errorf("end_time_selection_reference must be set when end_time_selection is not 'latest'")
	}

	return nil
}
