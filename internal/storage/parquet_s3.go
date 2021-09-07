package storage

import (
	"bytes"
	"context"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xitongsys/parquet-go-source/buffer"
	"github.com/xitongsys/parquet-go/writer"
)

type S3Api interface {
	PutObject(ctx context.Context, params *s3.PutObjectInput, optFns ...func(*s3.Options)) (*s3.PutObjectOutput, error)
}

type ParquetS3Writer struct {
	bucket          string
	pathPrefix      string
	fileExtension   string
	filename        string
	numParquetPages int64
	parquetWriter   *writer.ParquetWriter
	bufferFile      *buffer.BufferFile
	api             S3Api
	config          StorageContextConfig
}

func (w *ParquetS3Writer) GetConfig() StorageContextConfig {
	return w.config
}

func (w *ParquetS3Writer) Store(ctx context.Context, obj interface{}) error {
	return w.parquetWriter.Write(obj)
}

//Close is a no-op since multiple datasources might use it, close() is the real closure method
func (w *ParquetS3Writer) Close(ctx context.Context) error {
	return nil
}

func (w *ParquetS3Writer) close(ctx context.Context) error {
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

type ParquetS3Backend struct {
	Bucket          string
	PathPrefix      string
	FileExtension   string
	Api             S3Api
	NumParquetPages int64
	writers         map[string]ParquetS3Writer
	writersLock     sync.Mutex
}

func NewParquetS3Backend(bucket, pathPrefix, fileExtension string, api S3Api, numParquetPages int64) *ParquetS3Backend {
	writers := make(map[string]ParquetS3Writer)
	return &ParquetS3Backend{
		Bucket:          bucket,
		PathPrefix:      pathPrefix,
		FileExtension:   fileExtension,
		Api:             api,
		NumParquetPages: numParquetPages,
		writers:         writers,
	}
}

func (p *ParquetS3Backend) CloseStorageContexts(ctx context.Context) []error {
	errors := []error{}
	for _, writer := range p.writers {
		err := writer.close(ctx)
		if err != nil {
			errors = append(errors, err)
		}
	}

	if len(errors) == 0 {
		return nil
	} else {
		return errors
	}
}

func (p *ParquetS3Backend) GetStorageContext(config StorageContextConfig, sampleObject interface{}) (StorageBackendContext, error) {
	p.writersLock.Lock()
	defer p.writersLock.Unlock()

	serviceWriterKey := fmt.Sprintf("%s_%s_%s_%s", config.Date.Format("2006-01-02"), config.Cloud, config.Service, config.DataSource)
	serviceWriter, ok := p.writers[serviceWriterKey]
	if ok {
		return &serviceWriter, nil
	}

	buff := bytes.NewBuffer([]byte{})
	bw := &buffer.BufferFile{
		Reader: bytes.NewReader([]byte{}),
		Writer: buff,
	}

	pw, err := writer.NewParquetWriter(bw, sampleObject, p.NumParquetPages)
	if err != nil {
		return nil, err
	}

	w := ParquetS3Writer{
		bucket:          p.Bucket,
		pathPrefix:      fmt.Sprintf("%s%s/%s/%s/report_date=%s/", p.PathPrefix, config.Cloud, config.Service, config.DataSource, config.Date.Format("2006-01-02")),
		fileExtension:   p.FileExtension,
		numParquetPages: p.NumParquetPages,
		bufferFile:      bw,
		parquetWriter:   pw,
		api:             p.Api,
		filename:        uuid.New().String(),
		config:          config,
	}
	p.writers[serviceWriterKey] = w

	return &w, nil
}
