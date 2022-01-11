package awscloud

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud/interfaces"
)

type AwsClient struct {
	ec2Client            *ec2.Client
	cloudwatchlogsClient *cloudwatchlogs.Client
}

func NewAwsClient(cfg aws.Config) *AwsClient {
	return &AwsClient{
		ec2Client:            ec2.NewFromConfig(cfg),
		cloudwatchlogsClient: cloudwatchlogs.NewFromConfig(cfg),
	}
}

func (c *AwsClient) EC2() interfaces.EC2Client {
	return c.ec2Client
}

func (c *AwsClient) CloudWatchLogs() interfaces.CloudWatchLogsClient {
	return c.cloudwatchlogsClient
}
