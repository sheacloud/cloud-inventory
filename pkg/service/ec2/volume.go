package ec2

import (
	"context"

	awsec2 "github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"gitlab.com/sheacloud-infrastructure/aws/data-warehouse/internal/parquetwriter"
)

func init() {
	Controller.RegisterDataSource("volume", VolumeDataSource)
}

type VolumeModel struct {
	Attachments      []VolumeAttachmentModel `parquet:"name=attachments, type=LIST"`
	AvailabilityZone string                  `parquet:"name=availability_zone, type=BYTE_ARRAY, convertedtype=UTF8"`
	Encrypted        bool                    `parquet:"name=encrypted, type=BOOLEAN"`
	VolumeId         string                  `parquet:"name=volume_id, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type VolumeAttachmentModel struct {
	Device     string `parquet:"name=device, type=BYTE_ARRAY, convertedtype=UTF8"`
	InstanceId string `parquet:"name=instance_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	VolumeId   string `parquet:"name=volume_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	State      string `parquet:"name=state, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type VolumeDataSourceClient interface {
	DescribeVolumes(context.Context, *awsec2.DescribeVolumesInput, ...func(*awsec2.Options)) (*awsec2.DescribeVolumesOutput, error)
}

func VolumeDataSource(ctx context.Context, accountId, region string, client *awsec2.Client, parquetConfig parquetwriter.ParquetConfig) error {
	return volumeDataSource(ctx, accountId, region, client, parquetConfig)
}

// function with client as a specific interface, allowing mocking/testing
func volumeDataSource(ctx context.Context, accountId, region string, client VolumeDataSourceClient, parquetConfig parquetwriter.ParquetConfig) error {
	s3ParquetWriter, err := parquetwriter.NewS3ParquetWriter(new(VolumeModel), accountId, region, "ec2", "volume", parquetConfig)
	if err != nil {
		return err
	}
	defer s3ParquetWriter.Close(ctx)

	paginator := awsec2.NewDescribeVolumesPaginator(client, &awsec2.DescribeVolumesInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":     "ec2",
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
			s3ParquetWriter.Write(model)
		}
	}

	return nil
}
