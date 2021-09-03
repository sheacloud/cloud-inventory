package controller

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sheacloud/cloud-inventory/internal/storage"
)

type AwsController interface {
	Process(ctx context.Context, accountId, region string, date time.Time, cfg aws.Config, storageManager *storage.StorageManager) map[string]error
	GetName() string
}
