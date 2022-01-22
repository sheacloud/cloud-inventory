package indexedstorage

type ParquetS3Manifest struct {
	LatestReportTime int64                 `json:"latest_report_time"`
	ReportFiles      []ParquetS3ReportFile `json:"report_files"`
}

type ParquetS3ReportFile struct {
	ReportTime int64  `json:"report_time"`
	FileName   string `json:"file_name"`
}
