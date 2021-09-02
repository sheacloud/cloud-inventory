package ec2

import (
	"context"
	"time"

	awsec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/aws-infra-warehouse/internal/parquetwriter"
	"github.com/sirupsen/logrus"
)

func init() {
	Controller.RegisterDataSource("volumes", VolumeDataSource)
}

type VolumeModel struct {
	Attachments        []VolumeAttachmentModel `parquet:"name=attachments, type=LIST"`
	AvailabilityZone   string                  `parquet:"name=availability_zone, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	CreateTime         time.Time
	CreateTimeMillis   int64             `parquet:"name=create_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
	Encrypted          bool              `parquet:"name=encrypted, type=BOOLEAN"`
	FastRestored       bool              `parquet:"name=fast_restored, type=BOOLEAN"`
	Iops               int32             `parquet:"name=iops, type=INT32"`
	KmsKeyId           string            `parquet:"name=kms_key_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	MultiAttachEnabled bool              `parquet:"name=multi_attach_enabled, type=BOOLEAN"`
	Size               int32             `parquet:"name=size, type=INT32"`
	SnapshotId         string            `parquet:"name=snapshot_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	State              string            `parquet:"name=state, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
	Tags               map[string]string `parquet:"name=tags, type=MAP, keytype=BYTE_ARRAY, keyconvertedtype=UTF8, valuetype=BYTE_ARRAY, valueconvertedtype=UTF8"`
	Throughput         int32             `parquet:"name=throughput, type=INT32"`
	VolumeId           string            `parquet:"name=volume_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	VolumeType         string            `parquet:"name=volume_type, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
}

type VolumeAttachmentModel struct {
	DeleteOnTermination bool   `parquet:"name=delete_on_termination, type=BOOLEAN"`
	Device              string `parquet:"name=device, type=BYTE_ARRAY, convertedtype=UTF8"`
	InstanceId          string `parquet:"name=instance_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	VolumeId            string `parquet:"name=volume_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	State               string `parquet:"name=state, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN_DICTIONARY"`
}

type VolumeDataSourceClient interface {
	DescribeVolumes(context.Context, *awsec2.DescribeVolumesInput, ...func(*awsec2.Options)) (*awsec2.DescribeVolumesOutput, error)
}

func VolumeDataSource(ctx context.Context, accountId, region string, client *awsec2.Client, parquetConfig parquetwriter.ParquetConfig) error {
	return volumeDataSource(ctx, accountId, region, client, parquetConfig)
}

// function with client as a specific interface, allowing mocking/testing
func volumeDataSource(ctx context.Context, accountId, region string, client VolumeDataSourceClient, parquetConfig parquetwriter.ParquetConfig) error {
	s3ParquetWriter, err := parquetwriter.NewS3ParquetWriter(new(VolumeModel), accountId, region, serviceName, "volumes", parquetConfig)
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
				"data_source": "volume",
				"account_id":  accountId,
				"region":      region,
				"error":       err,
			}).Error("error calling DescribeVolumes")
			return err
		}

		for _, volume := range output.Volumes {
			model := new(VolumeModel)
			copier.Copy(&model, &volume)

			model.Tags = GetTagMap(volume.Tags)
			model.CreateTimeMillis = model.CreateTime.UTC().UnixMilli()

			s3ParquetWriter.Write(model)
		}
	}

	return nil
}
