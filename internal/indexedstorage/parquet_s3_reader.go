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
	"github.com/sirupsen/logrus"
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
)

type ParquetS3DirectoryReader struct {
	Bucket            string
	indices           []string
	reportDate        string
	selection         RequestTimeSelection
	Api               *s3.Client
	sampleObject      interface{}
	s3Files           []*reader.ParquetReader
	ctx               context.Context
	currentFileIndex  int
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
		currentFileIndex: 0,
	}

	err := reader.DetermineDataFiles()
	if err != nil {
		return nil, err
	}

	return reader, nil
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
	}

	for _, reportFile := range manifest.ReportFiles {
		if reportFile.ReportTime == desiredReportTime {
			s3FileNames = append(s3FileNames, reportFile.FileName)
		}
	}

	r.desiredReportTime = desiredReportTime

	wg := &sync.WaitGroup{}
	r.s3Files = make([]*reader.ParquetReader, len(s3FileNames))

	for i, fileName := range s3FileNames {
		wg.Add(1)
		go func(i int, fileName string) {
			logrus.Info("creating s3 file reader")
			s3File, err := s3v2.NewS3FileReaderWithClient(r.ctx, r.Api, r.Bucket, fileName)
			if err != nil {
				panic(err)
			}
			logrus.Info("created s3 file reader")

			logrus.Info("creating parquet reader")
			s3Reader, err := reader.NewParquetReader(s3File, r.sampleObject, 16)
			if err != nil {
				panic(err)
			}
			logrus.Info("created parquet reader")
			r.s3Files[i] = s3Reader
			wg.Done()
		}(i, fileName)
	}

	wg.Wait()

	return nil
}

func (r *ParquetS3DirectoryReader) HasNextFile() bool {
	return r.currentFileIndex < len(r.s3Files)
}

func (r *ParquetS3DirectoryReader) ReadNextFile() ([]interface{}, error) {
	if r.currentFileIndex >= len(r.s3Files) {
		return nil, fmt.Errorf("no more files")
	}

	currentFile := r.s3Files[r.currentFileIndex]
	r.currentFileIndex++
	numObjects := currentFile.GetNumRows()

	destSliceValue := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(r.sampleObject)), int(numObjects), int(numObjects))
	destValue := reflect.New(destSliceValue.Type())
	destValue.Elem().Set(destSliceValue)

	err := currentFile.Read(destValue.Interface())
	if err != nil {
		return nil, err
	}

	resultLen := destValue.Elem().Len()
	interfaceList := []interface{}{}
	// convert to list of interfaces, and filter out any incorrect time values
	for i := 0; i < resultLen; i++ {
		reportTimeObj, ok := destValue.Elem().Index(i).Interface().(ReportTimeObject)
		if !ok {
			return nil, fmt.Errorf("cant cast to report time object")
		}
		if reportTimeObj.GetReportTime() == r.desiredReportTime {
			interfaceList = append(interfaceList, destValue.Elem().Index(i).Interface())
		}
	}

	return interfaceList, nil
}
