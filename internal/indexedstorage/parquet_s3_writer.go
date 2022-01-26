package indexedstorage

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
	"github.com/google/uuid"
	"github.com/xitongsys/parquet-go-source/s3v2"
	"github.com/xitongsys/parquet-go/source"
	"github.com/xitongsys/parquet-go/writer"
)

type ParquetS3File struct {
	bucket        string
	indices       []string
	reportDate    string
	parquetWriter *writer.ParquetWriter
	s3File        source.ParquetFile
	s3FileKey     string
	api           s3v2.S3API
	sampleObj     interface{}
	writeLock     sync.Mutex
	reportTime    int64
}

func NewParquetS3File(bucket string, indices []string, reportDate string, reportTime int64, api s3v2.S3API, sampleObj interface{}) (*ParquetS3File, error) {
	s3writer := &ParquetS3File{
		bucket:     bucket,
		indices:    indices,
		reportDate: reportDate,
		reportTime: reportTime,
		api:        api,
		sampleObj:  sampleObj,
	}

	// create new writers
	indexPath := strings.Join(indices, "/")
	filename := fmt.Sprintf("inventory/%s/report_date=%s/%s.parquet", indexPath, reportDate, uuid.New().String())
	s3writer.s3FileKey = filename
	s3File, err := s3v2.NewS3FileWriterWithClient(context.TODO(), api, bucket, filename, nil)
	if err != nil {
		return nil, err
	}

	parquetWriter, err := writer.NewParquetWriter(s3File, sampleObj, 32)
	if err != nil {
		return nil, err
	}

	s3writer.s3File = s3File
	s3writer.parquetWriter = parquetWriter

	return s3writer, nil
}

func (w *ParquetS3File) Write(ctx context.Context, obj interface{}) error {
	w.writeLock.Lock()
	defer w.writeLock.Unlock()

	return w.parquetWriter.Write(obj)
}

func (w *ParquetS3File) Close() error {
	if w.parquetWriter != nil {
		err := w.parquetWriter.WriteStop()
		if err != nil {
			return err
		}
	}
	if w.s3File != nil {
		err := w.s3File.Close()
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *ParquetS3File) UpdateManifest(ctx context.Context) error {
	indexPath := strings.Join(w.indices, "/")
	manifestFileName := fmt.Sprintf("manifests/%s/report_date=%s/manifest.json", indexPath, w.reportDate)
	// try and get current manifest

	manifestFile, err := w.api.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(w.bucket),
		Key:    aws.String(manifestFileName),
	})

	manifest := &ParquetS3Manifest{}

	var apiError smithy.APIError
	if errors.As(err, &apiError) && apiError.ErrorCode() == "NoSuchKey" {
		// create a new manifest
		manifest.LatestReportTime = w.reportTime
		manifest.ReportFiles = []ParquetS3ReportFile{
			{
				ReportTime: w.reportTime,
				FileName:   w.s3FileKey,
			},
		}
	} else if err != nil {
		return fmt.Errorf("error calling GetObject: %w", err)
	} else {
		// update the existing manifest
		buf := new(bytes.Buffer)
		buf.ReadFrom(manifestFile.Body)
		err = json.Unmarshal(buf.Bytes(), manifest)
		if err != nil {
			return fmt.Errorf("error unmarshalling manifest: %w", err)
		}

		// update manifest
		manifest.LatestReportTime = w.reportTime
		manifest.ReportFiles = append(manifest.ReportFiles, ParquetS3ReportFile{
			ReportTime: w.reportTime,
			FileName:   w.s3FileKey,
		})
	}

	// write the manifest to S3
	manifestBytes, err := json.Marshal(manifest)
	if err != nil {
		return fmt.Errorf("error marshalling manifest: %w", err)
	}
	_, err = w.api.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(w.bucket),
		Key:    aws.String(manifestFileName),
		Body:   bytes.NewReader(manifestBytes),
	})
	if err != nil {
		return fmt.Errorf("error calling PutObject: %w", err)
	}

	return nil
}

type IndexedFileManager struct {
	Bucket        string
	PathPrefix    string
	FileExtension string
	Api           s3v2.S3API
}

func NewIndexedFileManager(bucket, pathPrefix, fileExtension string, api s3v2.S3API) *IndexedFileManager {
	return &IndexedFileManager{
		Bucket:        bucket,
		PathPrefix:    pathPrefix,
		FileExtension: fileExtension,
		Api:           api,
	}
}

func (p *IndexedFileManager) GetIndexedFile(indices []string, reportDate string, reportTime int64, sampleObject interface{}) (*ParquetS3File, error) {
	newFile, err := NewParquetS3File(p.Bucket, indices, reportDate, reportTime, p.Api, sampleObject)
	if err != nil {
		return nil, err
	}

	return newFile, nil
}
