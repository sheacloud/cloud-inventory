package storage

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/xitongsys/parquet-go-source/s3v2"
	"github.com/xitongsys/parquet-go/source"
	"github.com/xitongsys/parquet-go/writer"
)

type ParquetS3Writer struct {
	bucket          string
	pathPrefix      string
	fileExtension   string
	numParquetPages int64
	parquetWriter   *writer.ParquetWriter
	s3File          source.ParquetFile
	api             s3v2.S3API
	config          StorageContextConfig
	sampleObj       interface{}
}

func NewParquetS3Writer(bucket, pathPrefix, fileExtension string, numParquetPages int64, api s3v2.S3API, config StorageContextConfig, sampleObj interface{}) (*ParquetS3Writer, error) {
	s3writer := &ParquetS3Writer{
		bucket:          bucket,
		pathPrefix:      pathPrefix,
		fileExtension:   fileExtension,
		numParquetPages: numParquetPages,
		api:             api,
		config:          config,
		sampleObj:       sampleObj,
	}

	// create new writers
	filename := fmt.Sprintf("%s%s.%s", pathPrefix, uuid.New().String(), fileExtension)
	s3File, err := s3v2.NewS3FileWriterWithClient(context.TODO(), api, bucket, filename, nil)
	if err != nil {
		return nil, err
	}

	parquetWriter, err := writer.NewParquetWriter(s3File, sampleObj, numParquetPages)
	if err != nil {
		return nil, err
	}

	s3writer.s3File = s3File
	s3writer.parquetWriter = parquetWriter

	return s3writer, nil
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

func (w *ParquetS3Writer) close() error {
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

type ParquetS3Backend struct {
	Bucket          string
	PathPrefix      string
	FileExtension   string
	Api             s3v2.S3API
	NumParquetPages int64
	writers         map[string]*ParquetS3Writer
	writersLock     sync.Mutex
}

func NewParquetS3Backend(bucket, pathPrefix, fileExtension string, api s3v2.S3API, numParquetPages int64) *ParquetS3Backend {
	writers := make(map[string]*ParquetS3Writer)
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
		err := writer.close()
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
		return serviceWriter, nil
	}
	pathPrefix := fmt.Sprintf("%s%s/%s/%s/report_date=%s/", p.PathPrefix, config.Cloud, config.Service, config.DataSource, config.Date.Format("2006-01-02"))
	w, err := NewParquetS3Writer(p.Bucket, pathPrefix, p.FileExtension, p.NumParquetPages, p.Api, config, sampleObject)
	if err != nil {
		return nil, err
	}
	p.writers[serviceWriterKey] = w

	return w, nil
}
