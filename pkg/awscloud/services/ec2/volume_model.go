package ec2

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
)

type VolumeModel struct {
	Attachments        []*VolumeAttachmentVolumeModel `parquet:"name=attachments,type=MAP,convertedtype=LIST"`
	AvailabilityZone   string                         `parquet:"name=availability_zone,type=BYTE_ARRAY,convertedtype=UTF8"`
	CreateTime         *time.Time
	Encrypted          bool   `parquet:"name=encrypted,type=BOOLEAN"`
	FastRestored       bool   `parquet:"name=fast_restored,type=BOOLEAN"`
	Iops               int32  `parquet:"name=iops,type=INT32"`
	KmsKeyId           string `parquet:"name=kms_key_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	MultiAttachEnabled bool   `parquet:"name=multi_attach_enabled,type=BOOLEAN"`
	OutpostArn         string `parquet:"name=outpost_arn,type=BYTE_ARRAY,convertedtype=UTF8"`
	Size               int32  `parquet:"name=size,type=INT32"`
	SnapshotId         string `parquet:"name=snapshot_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	State              string `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
	TagsOld            []*TagVolumeModel
	Throughput         int32             `parquet:"name=throughput,type=INT32"`
	VolumeId           string            `parquet:"name=volume_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	VolumeType         string            `parquet:"name=volume_type,type=BYTE_ARRAY,convertedtype=UTF8"`
	CreateTimeMilli    int64             `parquet:"name=create_time_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
	Tags               map[string]string `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
	AccountId          string            `parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Region             string            `parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8"`
	ReportTime         int64             `parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type VolumeAttachmentVolumeModel struct {
	AttachTime          *time.Time
	DeleteOnTermination bool   `parquet:"name=delete_on_termination,type=BOOLEAN"`
	Device              string `parquet:"name=device,type=BYTE_ARRAY,convertedtype=UTF8"`
	InstanceId          string `parquet:"name=instance_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	State               string `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8"`
	VolumeId            string `parquet:"name=volume_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	AttachTimeMilli     int64  `parquet:"name=attach_time_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS"`
}

type TagVolumeModel struct {
	Key   string `parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

func FetchVolumes(ctx context.Context, params *awscloud.AwsFetchInput) *awscloud.AwsFetchOutput {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int

	awsClient := params.RegionalClients[params.Region]
	ec2Client := awsClient.EC2()

	paginator := ec2.NewDescribeVolumesPaginator(ec2Client, &ec2.DescribeVolumesInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling DescribeVolumes in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, var0 := range output.Volumes {

			model := new(VolumeModel)
			copier.Copy(&model, &var0)

			model.Tags = GetTagMap(var0.Tags)
			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime.UTC().UnixMilli()

			if err = PostProcessVolume(ctx, params, model); err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error post-processing volume %s %s/%s: %w", model.VolumeId, params.AccountId, params.Region, err))
				failedResources++
			}

			err = params.OutputFile.Write(ctx, model)
			if err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error storing Volume model in %s/%s: %w", params.AccountId, params.Region, err))
			}
			fetchedResources++
		}

	}

	return &awscloud.AwsFetchOutput{
		FetchingErrors:   fetchingErrors,
		FetchedResources: fetchedResources,
		FailedResources:  failedResources,
		ResourceName:     "volumes",
		AccountId:        params.AccountId,
		Region:           params.Region,
	}
}
