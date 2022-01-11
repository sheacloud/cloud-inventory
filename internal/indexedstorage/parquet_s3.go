package indexedstorage

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/xitongsys/parquet-go-source/s3v2"
	"github.com/xitongsys/parquet-go/source"
	"github.com/xitongsys/parquet-go/writer"
)

type ParquetS3File struct {
	bucket          string
	pathPrefix      string
	fileExtension   string
	numParquetPages int64
	parquetWriter   *writer.ParquetWriter
	s3File          source.ParquetFile
	api             s3v2.S3API
	sampleObj       interface{}
	writeLock       sync.Mutex
}

func NewParquetS3File(bucket, pathPrefix, fileExtension string, numParquetPages int64, api s3v2.S3API, sampleObj interface{}) (*ParquetS3File, error) {
	s3writer := &ParquetS3File{
		bucket:          bucket,
		pathPrefix:      pathPrefix,
		fileExtension:   fileExtension,
		numParquetPages: numParquetPages,
		api:             api,
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

type IndexedFileManager struct {
	Bucket          string
	PathPrefix      string
	FileExtension   string
	Api             s3v2.S3API
	NumParquetPages int64
}

func NewIndexedFileManager(bucket, pathPrefix, fileExtension string, api s3v2.S3API, numParquetPages int64) *IndexedFileManager {
	return &IndexedFileManager{
		Bucket:          bucket,
		PathPrefix:      pathPrefix,
		FileExtension:   fileExtension,
		Api:             api,
		NumParquetPages: numParquetPages,
	}
}

func (p *IndexedFileManager) GetIndexedFile(index []string, sampleObject interface{}) (*ParquetS3File, error) {
	pathPrefix := p.PathPrefix + strings.Join(index, "/") + "/"
	newFile, err := NewParquetS3File(p.Bucket, pathPrefix, p.FileExtension, p.NumParquetPages, p.Api, sampleObject)
	if err != nil {
		return nil, err
	}

	return newFile, nil
}
