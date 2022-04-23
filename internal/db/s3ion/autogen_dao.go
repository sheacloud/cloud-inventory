//AUTOGENERATED CODE DO NOT EDIT
package s3ion

import (
	"context"
	"fmt"
	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sheacloud/cloud-inventory/pkg/aws/apigateway"
	"github.com/sheacloud/cloud-inventory/pkg/aws/apigatewayv2"
	"github.com/sheacloud/cloud-inventory/pkg/aws/backup"
	"github.com/sheacloud/cloud-inventory/pkg/aws/cloudtrail"
	"github.com/sheacloud/cloud-inventory/pkg/aws/cloudwatchlogs"
	"github.com/sheacloud/cloud-inventory/pkg/aws/dynamodb"
	"github.com/sheacloud/cloud-inventory/pkg/aws/ec2"
	"github.com/sheacloud/cloud-inventory/pkg/aws/ecs"
	"github.com/sheacloud/cloud-inventory/pkg/aws/efs"
	"github.com/sheacloud/cloud-inventory/pkg/aws/elasticache"
	"github.com/sheacloud/cloud-inventory/pkg/aws/elasticloadbalancing"
	"github.com/sheacloud/cloud-inventory/pkg/aws/elasticloadbalancingv2"
	"github.com/sheacloud/cloud-inventory/pkg/aws/iam"
	"github.com/sheacloud/cloud-inventory/pkg/aws/lambda"
	"github.com/sheacloud/cloud-inventory/pkg/aws/rds"
	"github.com/sheacloud/cloud-inventory/pkg/aws/redshift"
	"github.com/sheacloud/cloud-inventory/pkg/aws/route53"
	"github.com/sheacloud/cloud-inventory/pkg/aws/s3"
	"github.com/sheacloud/cloud-inventory/pkg/aws/sns"
	"github.com/sheacloud/cloud-inventory/pkg/aws/sqs"
	"github.com/sheacloud/cloud-inventory/pkg/aws/storagegateway"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

type S3IonWriterDAO struct {
	ionClient  *S3IonClient
	maxRetries int
}

func NewS3IonWriterDAO(s3Client *awsS3.Client, bucket string, maxRetries int) *S3IonWriterDAO {
	return &S3IonWriterDAO{
		ionClient:  NewS3IonClient(s3Client, bucket),
		maxRetries: maxRetries,
	}
}

func (dao *S3IonWriterDAO) WriteInventoryResults(ctx context.Context, metadata *meta.InventoryResults) error {
	return nil
}

func (dao *S3IonWriterDAO) WriteIngestionTimestamp(ctx context.Context, metadata *meta.IngestionTimestamp) error {
	return nil
}

func (dao *S3IonWriterDAO) Finish(ctx context.Context) error {
	return dao.ionClient.CloseAll(ctx)
}

func (dao *S3IonWriterDAO) PutAwsApiGatewayRestApis(ctx context.Context, resources []*apigateway.RestApi) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "apigateway", "rest_apis", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsApiGatewayV2Apis(ctx context.Context, resources []*apigatewayv2.Api) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "apigatewayv2", "apis", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsBackupBackupVaults(ctx context.Context, resources []*backup.BackupVault) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "backup", "vaults", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsBackupBackupPlans(ctx context.Context, resources []*backup.BackupPlan) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "backup", "plans", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsCloudTrailTrails(ctx context.Context, resources []*cloudtrail.Trail) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "cloudtrail", "trails", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsCloudWatchLogsLogGroups(ctx context.Context, resources []*cloudwatchlogs.LogGroup) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "cloudwatchlogs", "log_groups", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsDynamoDBTables(ctx context.Context, resources []*dynamodb.Table) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "dynamodb", "tables", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2Addresses(ctx context.Context, resources []*ec2.Address) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "addresses", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2DhcpOptions(ctx context.Context, resources []*ec2.DhcpOptions) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "dhcp_options", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2Images(ctx context.Context, resources []*ec2.Image) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "images", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2Instances(ctx context.Context, resources []*ec2.Instance) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "instances", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2InternetGateways(ctx context.Context, resources []*ec2.InternetGateway) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "internet_gateways", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2ManagedPrefixLists(ctx context.Context, resources []*ec2.ManagedPrefixList) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "managed_prefix_lists", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2NatGateways(ctx context.Context, resources []*ec2.NatGateway) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "nat_gateways", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2NetworkAcls(ctx context.Context, resources []*ec2.NetworkAcl) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "network_acls", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2NetworkInterfaces(ctx context.Context, resources []*ec2.NetworkInterface) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "network_interfaces", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2PlacementGroups(ctx context.Context, resources []*ec2.PlacementGroup) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "placement_groups", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2ReservedInstances(ctx context.Context, resources []*ec2.ReservedInstances) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "reserved_instances", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2RouteTables(ctx context.Context, resources []*ec2.RouteTable) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "route_tables", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2SecurityGroups(ctx context.Context, resources []*ec2.SecurityGroup) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "security_groups", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2Subnets(ctx context.Context, resources []*ec2.Subnet) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "subnets", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2TransitGatewayPeeringAttachments(ctx context.Context, resources []*ec2.TransitGatewayPeeringAttachment) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "transit_gateway_peering_attachments", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2TransitGatewayRouteTables(ctx context.Context, resources []*ec2.TransitGatewayRouteTable) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "transit_gateway_route_tables", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2TransitGatewayVpcAttachments(ctx context.Context, resources []*ec2.TransitGatewayVpcAttachment) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "transit_gateway_vpc_attachments", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2TransitGateways(ctx context.Context, resources []*ec2.TransitGateway) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "transit_gateways", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2Volumes(ctx context.Context, resources []*ec2.Volume) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "volumes", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2VpcEndpoints(ctx context.Context, resources []*ec2.VpcEndpoint) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "vpc_endpoints", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2VpcPeeringConnections(ctx context.Context, resources []*ec2.VpcPeeringConnection) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "vpc_peering_connections", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2Vpcs(ctx context.Context, resources []*ec2.Vpc) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "vpcs", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEC2VpnGateways(ctx context.Context, resources []*ec2.VpnGateway) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ec2", "vpn_gateways", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsECSClusters(ctx context.Context, resources []*ecs.Cluster) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ecs", "clusters", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsECSServices(ctx context.Context, resources []*ecs.Service) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ecs", "services", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsECSTasks(ctx context.Context, resources []*ecs.Task) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "ecs", "tasks", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsEFSFileSystems(ctx context.Context, resources []*efs.FileSystem) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "efs", "filesystems", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsElastiCacheCacheClusters(ctx context.Context, resources []*elasticache.CacheCluster) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "elasticache", "cache_clusters", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsElasticLoadBalancingLoadBalancers(ctx context.Context, resources []*elasticloadbalancing.LoadBalancer) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "elasticloadbalancing", "load_balancers", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsElasticLoadBalancingV2LoadBalancers(ctx context.Context, resources []*elasticloadbalancingv2.LoadBalancer) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "elasticloadbalancingv2", "load_balancers", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsElasticLoadBalancingV2TargetGroups(ctx context.Context, resources []*elasticloadbalancingv2.TargetGroup) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "elasticloadbalancingv2", "target_groups", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsIAMGroups(ctx context.Context, resources []*iam.Group) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "iam", "groups", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsIAMPolicies(ctx context.Context, resources []*iam.Policy) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "iam", "policies", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsIAMRoles(ctx context.Context, resources []*iam.Role) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "iam", "roles", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsIAMUsers(ctx context.Context, resources []*iam.User) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "iam", "users", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsLambdaFunctions(ctx context.Context, resources []*lambda.Function) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "lambda", "functions", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsRDSDBClusters(ctx context.Context, resources []*rds.DBCluster) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "rds", "db_clusters", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsRDSDBInstances(ctx context.Context, resources []*rds.DBInstance) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "rds", "db_instances", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsRedshiftClusters(ctx context.Context, resources []*redshift.Cluster) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "redshift", "clusters", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsRoute53HostedZones(ctx context.Context, resources []*route53.HostedZone) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "route53", "hosted_zones", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsS3Buckets(ctx context.Context, resources []*s3.Bucket) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "s3", "buckets", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsSNSTopics(ctx context.Context, resources []*sns.Topic) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "sns", "topics", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsSNSSubscriptions(ctx context.Context, resources []*sns.Subscription) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "sns", "subscriptions", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsSQSQueues(ctx context.Context, resources []*sqs.Queue) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "sqs", "queues", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
func (dao *S3IonWriterDAO) PutAwsStorageGatewayGateways(ctx context.Context, resources []*storagegateway.Gateway) error {
	if len(resources) == 0 {
		return nil
	}
	file := dao.ionClient.GetResourceFile("aws", "storagegateway", "gateways", resources[0].ReportTime)
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Encoder.Encode(resource); err != nil {
			return fmt.Errorf("failed to encode resource: %w", err)
		}
	}

	return nil
}
