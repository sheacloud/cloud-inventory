package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	ec2types "github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/sheacloud/aws-infra-warehouse/internal/parquetwriter"
)

var (
	Controller = Ec2Controller{
		DataSources: map[string]func(ctx context.Context, accountId, region string, client *awsec2.Client, parquetConfig parquetwriter.ParquetConfig) error{},
	}
	serviceName = "ec2"
)

type Ec2Controller struct {
	DataSources map[string]func(ctx context.Context, accountId, region string, client *awsec2.Client, parquetConfig parquetwriter.ParquetConfig) error
}

func (e *Ec2Controller) RegisterDataSource(dataSourceName string, dataSourceFunc func(ctx context.Context, accountId, region string, client *awsec2.Client, parquetConfig parquetwriter.ParquetConfig) error) {
	e.DataSources[dataSourceName] = dataSourceFunc
}

func (e *Ec2Controller) Process(ctx context.Context, accountId, region string, cfg aws.Config, parquetConfig parquetwriter.ParquetConfig) map[string]error {
	ec2Client := awsec2.NewFromConfig(cfg)

	errMap := map[string]error{}

	for dataSourceName, dataSourceFunc := range e.DataSources {
		err := dataSourceFunc(ctx, accountId, region, ec2Client, parquetConfig)
		if err != nil {
			errMap[dataSourceName] = err
		}
	}

	if len(errMap) == 0 {
		return nil
	} else {
		return errMap
	}
}

func GetTagMap(tags []ec2types.Tag) map[string]string {
	tagMap := make(map[string]string)
	for _, tag := range tags {
		tagMap[*tag.Key] = *tag.Value
	}
	return tagMap
}
