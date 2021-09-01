package controller

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type ServiceController interface {
	Process(ctx context.Context, accountId, region string, cfg aws.Config) map[string]error
}
