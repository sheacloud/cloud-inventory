package s3ion

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/amzn/ion-go/ion"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
)

type S3IonClient struct {
	Bucket     string
	PathPrefix string
	S3Client   *s3.Client
	FilesLock  sync.Mutex
	Files      map[string]*S3IonFile
}

type S3IonFile struct {
	FileKey  string
	Bucket   string
	Buffer   *bytes.Buffer
	Encoder  *ion.Encoder
	Lock     sync.Mutex
	S3Client *s3.Client
}

func NewS3IonClient(client *s3.Client, bucket string) *S3IonClient {
	return &S3IonClient{
		Bucket:     bucket,
		S3Client:   client,
		Files:      make(map[string]*S3IonFile),
		PathPrefix: "inventory",
	}
}

func (f *S3IonFile) Close(ctx context.Context) error {
	f.Lock.Lock()
	defer f.Lock.Unlock()

	err := f.Encoder.Finish()
	if err != nil {
		return err
	}
	_, err = f.S3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: &f.Bucket,
		Key:    &f.FileKey,
		Body:   bytes.NewReader(f.Buffer.Bytes()),
	})

	return err
}

func (c *S3IonClient) GetResourceFile(cloud, service, resource string, reportDateUnixMilli int64) *S3IonFile {
	reportDate := time.UnixMilli(reportDateUnixMilli)
	filePath := fmt.Sprintf("%s/%s/%s/%s/report_date=%s/", c.PathPrefix, cloud, service, resource, reportDate.Format("2006-01-02"))
	c.FilesLock.Lock()
	file, ok := c.Files[filePath]
	if !ok {
		fileUUID := uuid.New().String()
		fileName := filePath + fileUUID + ".ion"
		buffer := bytes.NewBuffer([]byte{})
		file = &S3IonFile{
			FileKey:  fileName,
			Bucket:   c.Bucket,
			Buffer:   buffer,
			Encoder:  ion.NewBinaryEncoder(buffer),
			S3Client: c.S3Client,
		}
		c.Files[filePath] = file
	}
	c.FilesLock.Unlock()

	return file
}

func (c *S3IonClient) CloseResource(ctx context.Context, cloud, service, resource string, reportDate time.Time) error {
	filePath := fmt.Sprintf("%s/%s/%s/%s/report_date=%s/", c.PathPrefix, cloud, service, resource, reportDate.Format("2006-01-02"))
	c.FilesLock.Lock()
	file, ok := c.Files[filePath]
	if !ok {
		c.FilesLock.Unlock()
		return fmt.Errorf("failed to close S3 ion file: file not found")
	}

	delete(c.Files, filePath)
	c.FilesLock.Unlock()

	err := file.Close(ctx)
	if err != nil {
		return err
	}

	return nil
}

func (c *S3IonClient) CloseAll(ctx context.Context) error {
	for _, file := range c.Files {
		err := file.Close(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}
