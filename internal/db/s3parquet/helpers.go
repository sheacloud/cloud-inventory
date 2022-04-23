package s3parquet

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/xitongsys/parquet-go-source/s3v2"
	"github.com/xitongsys/parquet-go/source"
	"github.com/xitongsys/parquet-go/writer"
)

type S3ParquetClient struct {
	Bucket        string
	PathPrefix    string
	S3Client      s3v2.S3API
	FilesLock     sync.Mutex
	Files         map[string]*S3ParquetFile
	NumProcessors int64
}

type S3ParquetFile struct {
	FileKey  string
	Bucket   string
	s3File   source.ParquetFile
	writer   *writer.ParquetWriter
	Lock     sync.Mutex
	S3Client s3v2.S3API
}

func NewS3ParquetClient(client s3v2.S3API, bucket string, numProcessors int64) *S3ParquetClient {
	return &S3ParquetClient{
		Bucket:        bucket,
		S3Client:      client,
		Files:         make(map[string]*S3ParquetFile),
		PathPrefix:    "inventory",
		NumProcessors: numProcessors,
	}
}

func (f *S3ParquetFile) Write(obj interface{}) error {
	return f.writer.Write(obj)
}

func (f *S3ParquetFile) Close(ctx context.Context) error {
	f.Lock.Lock()
	defer f.Lock.Unlock()

	err := f.writer.WriteStop()
	if err != nil {
		return err
	}

	err = f.s3File.Close()
	if err != nil {
		return err
	}

	return err
}

func (c *S3ParquetClient) GetResourceFile(ctx context.Context, indices []string, reportDateUnixMilli int64, sampleObj interface{}) (*S3ParquetFile, error) {
	reportDate := time.UnixMilli(reportDateUnixMilli)
	indicesPath := strings.Join(indices, "/")
	filePath := fmt.Sprintf("%s/%s/report_date=%s/", c.PathPrefix, indicesPath, reportDate.Format("2006-01-02"))
	c.FilesLock.Lock()
	file, ok := c.Files[filePath]
	if !ok {
		fileUUID := uuid.New().String()
		fileName := filePath + fileUUID + ".parquet"

		s3File, err := s3v2.NewS3FileWriterWithClient(ctx, c.S3Client, c.Bucket, fileName, nil)
		if err != nil {
			return nil, err
		}
		writer, err := writer.NewParquetWriter(s3File, sampleObj, c.NumProcessors)
		if err != nil {
			return nil, err
		}

		file = &S3ParquetFile{
			FileKey:  fileName,
			Bucket:   c.Bucket,
			s3File:   s3File,
			writer:   writer,
			S3Client: c.S3Client,
		}
		c.Files[filePath] = file
	}
	c.FilesLock.Unlock()

	return file, nil
}

func (c *S3ParquetClient) FinishIndex(ctx context.Context, indices []string, reportDateUnixMilli int64) error {
	reportDate := time.UnixMilli(reportDateUnixMilli)
	indicesPath := strings.Join(indices, "/")
	filePath := fmt.Sprintf("%s/%s/report_date=%s/", c.PathPrefix, indicesPath, reportDate.Format("2006-01-02"))
	c.FilesLock.Lock()
	file, ok := c.Files[filePath]
	if !ok {
		c.FilesLock.Unlock()
		return nil
	}

	delete(c.Files, filePath)
	c.FilesLock.Unlock()

	err := file.Close(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *S3ParquetClient) CloseAll(ctx context.Context) error {
	for _, file := range c.Files {
		logrus.WithFields(logrus.Fields{
			"file": file.FileKey,
		}).Info("Closing file")
		err := file.Close(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
