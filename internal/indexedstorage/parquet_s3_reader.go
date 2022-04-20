package indexedstorage

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/xitongsys/parquet-go-source/s3v2"
	"github.com/xitongsys/parquet-go/reader"
)

type ReportTimeObject interface {
	GetReportTime() int64
}

type RequestTimeSelection struct {
	Option         RequestTimeSelectionOption
	ReferencedTime time.Time
}

type RequestTimeSelectionOption string

const (
	RequestTimeSelectionOptionLatest = RequestTimeSelectionOption("latest")
	RequestTimeSelectionOptionBefore = RequestTimeSelectionOption("before")
	RequestTimeSelectionOptionAfter  = RequestTimeSelectionOption("after")
	RequestTimeSelectionOptionAt     = RequestTimeSelectionOption("at")
)

type ParquetS3DirectoryReader struct {
	Bucket            string
	indices           []string
	reportDate        string
	selection         RequestTimeSelection
	Api               *s3.Client
	sampleObject      interface{}
	s3Files           []*reader.ParquetReader
	s3FileNames       []string
	ctx               context.Context
	currentFileIndex  int
	currentRowIndex   int
	desiredReportTime int64
}

func NewParquetS3DirectoryReader(ctx context.Context, bucket string, indices []string, reportDate string, selection RequestTimeSelection, api *s3.Client, sampleObj interface{}) (*ParquetS3DirectoryReader, error) {
	reader := &ParquetS3DirectoryReader{
		Bucket:           bucket,
		indices:          indices,
		reportDate:       reportDate,
		selection:        selection,
		Api:              api,
		sampleObject:     sampleObj,
		ctx:              ctx,
		s3Files:          []*reader.ParquetReader{},
		s3FileNames:      []string{},
		currentFileIndex: 0,
		currentRowIndex:  0,
	}

	return reader, nil
}

func (r *ParquetS3DirectoryReader) GetAvailableDateTimes() ([]string, error) {
	// download the manifest for the given request
	indexPath := strings.Join(r.indices, "/")
	manifestFileName := fmt.Sprintf("manifests/%s/report_date=%s/manifest.json", indexPath, r.reportDate)
	manifestFile, err := r.Api.GetObject(r.ctx, &s3.GetObjectInput{
		Bucket: aws.String(r.Bucket),
		Key:    aws.String(manifestFileName),
	})
	if err != nil {
		return nil, err
	}

	manifest := &ParquetS3Manifest{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(manifestFile.Body)
	err = json.Unmarshal(buf.Bytes(), manifest)
	if err != nil {
		return nil, err
	}

	manifestTimes := []string{}
	for _, reportFile := range manifest.ReportFiles {
		manifestTimes = append(manifestTimes, time.UnixMilli(reportFile.ReportTime).UTC().Format(time.RFC3339Nano))
	}

	return manifestTimes, nil
}

func (r *ParquetS3DirectoryReader) LoadFromPaginationData(dataFileKeys []string, currentFileIndex, currentRowIndex int) error {
	if currentFileIndex >= len(dataFileKeys) {
		return fmt.Errorf("no more files")
	}

	s3FileNames := dataFileKeys[currentFileIndex:]

	wg := &sync.WaitGroup{}
	r.s3Files = make([]*reader.ParquetReader, len(s3FileNames))
	r.s3FileNames = s3FileNames

	// FIXME don't panic on failure to load file, since client could submit bad, handcrafted pagination data with fake file names
	for i, fileName := range s3FileNames {
		wg.Add(1)
		go func(i int, fileName string) {
			s3File, err := s3v2.NewS3FileReaderWithClient(r.ctx, r.Api, r.Bucket, fileName)
			if err != nil {
				panic(err)
			}
			s3Reader, err := reader.NewParquetReader(s3File, r.sampleObject, 16)
			if err != nil {
				panic(err)
			}
			r.s3Files[i] = s3Reader
			wg.Done()
		}(i, fileName)
	}

	wg.Wait()

	r.s3Files[currentFileIndex].SkipRows(int64(currentRowIndex))
	r.currentRowIndex = currentRowIndex
	return nil
}

func (r *ParquetS3DirectoryReader) DetermineDataFiles() error {
	// download the manifest for the given request
	indexPath := strings.Join(r.indices, "/")
	manifestFileName := fmt.Sprintf("manifests/%s/report_date=%s/manifest.json", indexPath, r.reportDate)
	manifestFile, err := r.Api.GetObject(r.ctx, &s3.GetObjectInput{
		Bucket: aws.String(r.Bucket),
		Key:    aws.String(manifestFileName),
	})
	if err != nil {
		return err
	}

	manifest := &ParquetS3Manifest{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(manifestFile.Body)
	err = json.Unmarshal(buf.Bytes(), manifest)
	if err != nil {
		return err
	}

	s3FileNames := []string{}
	// determine which data files to use
	var desiredReportTime int64
	switch r.selection.Option {
	case RequestTimeSelectionOptionLatest:
		desiredReportTime = manifest.LatestReportTime

	case RequestTimeSelectionOptionBefore:
		// determine the timestamp closest to before the requested time
		latestTime := int64(0)
		for _, reportFile := range manifest.ReportFiles {
			if reportFile.ReportTime < r.selection.ReferencedTime.UTC().UnixMilli() && reportFile.ReportTime > latestTime {
				latestTime = reportFile.ReportTime
			}
		}
		if latestTime == 0 {
			return fmt.Errorf("no data from %s before %s", r.reportDate, r.selection.ReferencedTime)
		}
		desiredReportTime = latestTime
	case RequestTimeSelectionOptionAfter:
		// determine the timestamp closest to after the requested time
		var earliestTime int64 = math.MaxInt64
		for _, reportFile := range manifest.ReportFiles {
			if reportFile.ReportTime > r.selection.ReferencedTime.UTC().UnixMilli() && reportFile.ReportTime < earliestTime {
				earliestTime = reportFile.ReportTime
			}
		}
		if earliestTime == math.MaxInt64 {
			return fmt.Errorf("no data from %s after %s", r.reportDate, r.selection.ReferencedTime)
		}
		desiredReportTime = earliestTime
	case RequestTimeSelectionOptionAt:
		desiredReportTime = r.selection.ReferencedTime.UTC().UnixMilli()
		validRequestTime := false
		for _, reportFile := range manifest.ReportFiles {
			if reportFile.ReportTime == desiredReportTime {
				validRequestTime = true
				break
			}
		}
		if !validRequestTime {
			return fmt.Errorf("no data from %s at %s", r.reportDate, r.selection.ReferencedTime)
		}
	}

	for _, reportFile := range manifest.ReportFiles {
		if reportFile.ReportTime == desiredReportTime {
			s3FileNames = append(s3FileNames, reportFile.FileName)
		}
	}

	r.desiredReportTime = desiredReportTime

	wg := &sync.WaitGroup{}
	r.s3FileNames = s3FileNames
	r.s3Files = make([]*reader.ParquetReader, len(s3FileNames))

	for i, fileName := range s3FileNames {
		wg.Add(1)
		go func(i int, fileName string) {
			s3File, err := s3v2.NewS3FileReaderWithClient(r.ctx, r.Api, r.Bucket, fileName)
			if err != nil {
				panic(err)
			}
			s3Reader, err := reader.NewParquetReader(s3File, r.sampleObject, 16)
			if err != nil {
				panic(err)
			}
			r.s3Files[i] = s3Reader
			wg.Done()
		}(i, fileName)
	}

	wg.Wait()

	return nil
}

func (r *ParquetS3DirectoryReader) HasMoreRows() bool {
	// if we read the entirety of the last file, we will have incremented currentFileIndex past the number of files
	return r.currentFileIndex < len(r.s3Files)
}

func (r *ParquetS3DirectoryReader) GetPaginationData() (dataFileKeys []string, currentFileIndex int, currentRowIndex int) {
	dataFileKeys = r.s3FileNames[r.currentFileIndex:]
	currentFileIndex = r.currentFileIndex
	currentRowIndex = r.currentRowIndex
	return dataFileKeys, currentFileIndex, currentRowIndex
}

// ReadRows reads maxResults rows from the current file, and iterates the currentFileIndex and currentRowIndex accordingly so future calls can call next files
func (r *ParquetS3DirectoryReader) ReadRows(maxResults int) ([]interface{}, error) {
	if r.currentFileIndex >= len(r.s3Files) {
		return nil, fmt.Errorf("no more files")
	}

	currentFile := r.s3Files[r.currentFileIndex]

	rowsLeftInFile := currentFile.GetNumRows() - int64(r.currentRowIndex)

	var rowsToRead int
	var newRowIndex int
	var newFileIndex int
	if maxResults == -1 {
		rowsToRead = int(rowsLeftInFile)
		newRowIndex = 0
		newFileIndex = r.currentFileIndex + 1
	} else if rowsLeftInFile <= int64(maxResults) {
		// read everything left in the file
		rowsToRead = int(rowsLeftInFile)
		newRowIndex = 0
		newFileIndex = r.currentFileIndex + 1
	} else {
		rowsToRead = maxResults
		newRowIndex = r.currentRowIndex + maxResults
		newFileIndex = r.currentFileIndex
	}

	destSliceValue := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(r.sampleObject)), rowsToRead, rowsToRead)
	destValue := reflect.New(destSliceValue.Type())
	destValue.Elem().Set(destSliceValue)

	result, err := currentFile.ReadByNumber(rowsToRead)
	if err != nil {
		return nil, err
	}

	r.currentRowIndex = newRowIndex
	r.currentFileIndex = newFileIndex

	return result, nil
}
