//AUTOGENERATED CODE DO NOT EDIT
package awscloud

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/aws/aws-sdk-go-v2/service/efs"
	"github.com/aws/aws-sdk-go-v2/service/elasticache"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/redshift"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/sheacloud/cloud-inventory/pkg/awscloud/interfaces"
)

type AwsClient struct {
	cloudwatchlogsClient         *cloudwatchlogs.Client
	dynamodbClient               *dynamodb.Client
	ec2Client                    *ec2.Client
	ecsClient                    *ecs.Client
	efsClient                    *efs.Client
	elasticacheClient            *elasticache.Client
	elasticloadbalancingClient   *elasticloadbalancing.Client
	elasticloadbalancingv2Client *elasticloadbalancingv2.Client
	iamClient                    *iam.Client
	lambdaClient                 *lambda.Client
	rdsClient                    *rds.Client
	redshiftClient               *redshift.Client
	route53Client                *route53.Client
	s3Client                     *s3.Client
}

func NewAwsClient(cfg aws.Config) *AwsClient {
	return &AwsClient{
		cloudwatchlogsClient:         cloudwatchlogs.NewFromConfig(cfg),
		dynamodbClient:               dynamodb.NewFromConfig(cfg),
		ec2Client:                    ec2.NewFromConfig(cfg),
		ecsClient:                    ecs.NewFromConfig(cfg),
		efsClient:                    efs.NewFromConfig(cfg),
		elasticacheClient:            elasticache.NewFromConfig(cfg),
		elasticloadbalancingClient:   elasticloadbalancing.NewFromConfig(cfg),
		elasticloadbalancingv2Client: elasticloadbalancingv2.NewFromConfig(cfg),
		iamClient:                    iam.NewFromConfig(cfg),
		lambdaClient:                 lambda.NewFromConfig(cfg),
		rdsClient:                    rds.NewFromConfig(cfg),
		redshiftClient:               redshift.NewFromConfig(cfg),
		route53Client:                route53.NewFromConfig(cfg),
		s3Client:                     s3.NewFromConfig(cfg),
	}
}

func (c *AwsClient) CloudWatchLogs() interfaces.CloudWatchLogsClient {
	return c.cloudwatchlogsClient
}
func (c *AwsClient) DynamoDB() interfaces.DynamoDBClient {
	return c.dynamodbClient
}
func (c *AwsClient) EC2() interfaces.EC2Client {
	return c.ec2Client
}
func (c *AwsClient) ECS() interfaces.ECSClient {
	return c.ecsClient
}
func (c *AwsClient) EFS() interfaces.EFSClient {
	return c.efsClient
}
func (c *AwsClient) ElastiCache() interfaces.ElastiCacheClient {
	return c.elasticacheClient
}
func (c *AwsClient) ElasticLoadBalancing() interfaces.ElasticLoadBalancingClient {
	return c.elasticloadbalancingClient
}
func (c *AwsClient) ElasticLoadBalancingV2() interfaces.ElasticLoadBalancingV2Client {
	return c.elasticloadbalancingv2Client
}
func (c *AwsClient) IAM() interfaces.IAMClient {
	return c.iamClient
}
func (c *AwsClient) Lambda() interfaces.LambdaClient {
	return c.lambdaClient
}
func (c *AwsClient) RDS() interfaces.RDSClient {
	return c.rdsClient
}
func (c *AwsClient) Redshift() interfaces.RedshiftClient {
	return c.redshiftClient
}
func (c *AwsClient) Route53() interfaces.Route53Client {
	return c.route53Client
}
func (c *AwsClient) S3() interfaces.S3Client {
	return c.s3Client
}
