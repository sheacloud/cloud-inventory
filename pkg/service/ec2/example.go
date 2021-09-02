package ec2

import (
	"context"

	awsec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/aws-infra-warehouse/internal/parquetwriter"
	"github.com/sirupsen/logrus"
)

func init() {
	Controller.RegisterDataSource("example", ExampleDataSource)
}

type ExampleModel struct {
	Field string            `parquet:"name=field, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Tags  map[string]string `parquet:"name=tags, type=MAP, keytype=BYTE_ARRAY, keyconvertedtype=UTF8, valuetype=BYTE_ARRAY, valueconvertedtype=UTF8"`
}

type ExampleDataSourceClient interface {
	DescribeVolumes(context.Context, *awsec2.DescribeVolumesInput, ...func(*awsec2.Options)) (*awsec2.DescribeVolumesOutput, error)
}

func ExampleDataSource(ctx context.Context, accountId, region string, client *awsec2.Client, parquetConfig parquetwriter.ParquetConfig) error {
	return exampleDataSource(ctx, accountId, region, client, parquetConfig)
}

// function with client as a specific interface, allowing mocking/testing
func exampleDataSource(ctx context.Context, accountId, region string, client ExampleDataSourceClient, parquetConfig parquetwriter.ParquetConfig) error {
	s3ParquetWriter, err := parquetwriter.NewS3ParquetWriter(new(ExampleModel), accountId, region, serviceName, "example", parquetConfig)
	if err != nil {
		return err
	}
	defer s3ParquetWriter.Close(ctx)

	paginator := awsec2.NewDescribeVolumesPaginator(client, &awsec2.DescribeVolumesInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":     serviceName,
				"data_source": "example",
				"account_id":  accountId,
				"region":      region,
				"error":       err,
			}).Error("error calling DescribeVolumes")
			return err
		}

		for _, volume := range output.Volumes {
			model := new(ExampleModel)
			copier.Copy(&model, &volume)

			model.Tags = GetTagMap(volume.Tags)

			s3ParquetWriter.Write(model)
		}
	}

	return nil
}
