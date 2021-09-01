package ec2

import (
	"context"

	awsec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"gitlab.com/sheacloud-infrastructure/aws/data-warehouse/internal/parquetwriter"
)

func init() {
	Controller.RegisterDataSource("instance", InstanceDataSource)
}

type InstanceModel struct {
	Architecture       string                  `parquet:"name=architecture, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY" copier:"Architecture"`
	IamInstanceProfile IamInstanceProfileModel `parquet:"name=iam_instance_profile"`
	InstanceId         string                  `parquet:"name=instance_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	InstanceType       string                  `parquet:"name=instance_type, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type IamInstanceProfileModel struct {
	Arn string `parquet:"name=arn, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Id  string `parquet:"name=id, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
}

type InstanceDataSourceClient interface {
	DescribeInstances(context.Context, *awsec2.DescribeInstancesInput, ...func(*awsec2.Options)) (*awsec2.DescribeInstancesOutput, error)
}

func InstanceDataSource(ctx context.Context, accountId, region string, client *awsec2.Client, parquetConfig parquetwriter.ParquetConfig) error {
	return instanceDataSource(ctx, accountId, region, client, parquetConfig)
}

// function with client as a specific interface, allowing mocking/testing
func instanceDataSource(ctx context.Context, accountId, region string, client InstanceDataSourceClient, parquetConfig parquetwriter.ParquetConfig) error {
	s3ParquetWriter, err := parquetwriter.NewS3ParquetWriter(new(InstanceModel), accountId, region, "ec2", "instance", parquetConfig)
	if err != nil {
		return err
	}
	defer s3ParquetWriter.Close(ctx)

	paginator := awsec2.NewDescribeInstancesPaginator(client, &awsec2.DescribeInstancesInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":     "ec2",
				"data_source": "instance",
				"account_id":  accountId,
				"region":      region,
				"error":       err,
			}).Error("error calling DescribeInstances")
			return err
		}

		for _, reservation := range output.Reservations {
			for _, instance := range reservation.Instances {
				model := new(InstanceModel)
				copier.Copy(&model, &instance)
				s3ParquetWriter.Write(model)
			}
		}
	}

	return nil
}
