package cloudwatch

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
)

func PostProcessMetricAlarm(ctx context.Context, params *localAws.AwsFetchInput, model *MetricAlarm) error {
	awsClient := params.RegionalClients[params.Region]
	client := awsClient.CloudWatch()

	// tags
	tagsOutput, err := client.ListTagsForResource(ctx, &cloudwatch.ListTagsForResourceInput{
		ResourceARN: aws.String(model.AlarmArn),
	})
	if err != nil {
		return fmt.Errorf("error calling ListTags: %v", err)
	}
	model.Tags = ConvertTags(tagsOutput.Tags)

	return nil
}
