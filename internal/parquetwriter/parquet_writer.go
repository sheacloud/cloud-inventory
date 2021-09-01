package parquetwriter

import (
	"bytes"
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sirupsen/logrus"
	"github.com/xitongsys/parquet-go-source/buffer"
	"github.com/xitongsys/parquet-go/writer"
)

type PutObjectApi interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

type ParquetConfig struct {
	Bucket          string
	PathPrefix      string
	FileExtension   string
	Api             PutObjectApi
	NumParquetPages int64
	Year            int
	Month           int
	Day             int
}

type S3ParquetWriter struct {
	bucket          string
	pathPrefix      string
	fileExtension   string
	filename        string
	numParquetPages int64
	parquetWriter   *writer.ParquetWriter
	bufferFile      *buffer.BufferFile
	api             PutObjectApi
}

func NewS3ParquetWriter(obj interface{}, accountId, region, service, datasource string, config ParquetConfig) (*S3ParquetWriter, error) {
	buff := bytes.NewBuffer([]byte{})
	bw := &buffer.BufferFile{
		Reader: bytes.NewReader([]byte{}),
		Writer: buff,
	}

	pw, err := writer.NewParquetWriter(bw, obj, config.NumParquetPages)
	if err != nil {
		return nil, err
	}

	w := S3ParquetWriter{
		bucket:          config.Bucket,
		pathPrefix:      fmt.Sprintf("%s%s/%s/year=%v/month=%v/day=%v/accountid=%s/region=%s/", config.PathPrefix, service, datasource, config.Year, config.Month, config.Day, accountId, region),
		fileExtension:   config.FileExtension,
		numParquetPages: config.NumParquetPages,
		bufferFile:      bw,
		parquetWriter:   pw,
		api:             config.Api,
		filename:        "data",
	}

	return &w, nil
}

func (w *S3ParquetWriter) Write(obj interface{}) error {
	return w.parquetWriter.Write(obj)
}

func (w *S3ParquetWriter) Close(ctx context.Context) error {
	w.parquetWriter.WriteStop()

	_, err := w.api.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(w.bucket),
		Key:    aws.String(fmt.Sprintf("%s%s.%s", w.pathPrefix, w.filename, w.fileExtension)),
		Body:   bytes.NewReader(w.bufferFile.Writer.Bytes()),
	})

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("error putting object in S3")
		return err
	}

	return w.bufferFile.Close()
}
