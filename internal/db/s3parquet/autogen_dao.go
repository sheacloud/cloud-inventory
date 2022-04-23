//AUTOGENERATED CODE DO NOT EDIT
package s3parquet

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

type S3ParquetWriterDAO struct {
	parquetClient *S3ParquetClient
}

func NewS3ParquetWriterDAO(s3Client *awsS3.Client, bucket string, numProcessors int64) *S3ParquetWriterDAO {
	return &S3ParquetWriterDAO{
		parquetClient: NewS3ParquetClient(s3Client, bucket, numProcessors),
	}
}

func (dao *S3ParquetWriterDAO) WriteInventoryResults(ctx context.Context, metadata *meta.InventoryResults) error {
	return nil
}

func (dao *S3ParquetWriterDAO) WriteIngestionTimestamp(ctx context.Context, metadata *meta.IngestionTimestamp) error {
	return nil
}

func (dao *S3ParquetWriterDAO) Close(ctx context.Context) error {
	return dao.parquetClient.CloseAll(ctx)
}

func (dao *S3ParquetWriterDAO) PutAwsApiGatewayRestApis(ctx context.Context, resources []*apigateway.RestApi) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "apigateway", "rest_apis", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsApiGatewayV2Apis(ctx context.Context, resources []*apigatewayv2.Api) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "apigatewayv2", "apis", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsBackupBackupVaults(ctx context.Context, resources []*backup.BackupVault) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "backup", "vaults", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsBackupBackupPlans(ctx context.Context, resources []*backup.BackupPlan) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "backup", "plans", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsCloudTrailTrails(ctx context.Context, resources []*cloudtrail.Trail) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "cloudtrail", "trails", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsCloudWatchLogsLogGroups(ctx context.Context, resources []*cloudwatchlogs.LogGroup) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "cloudwatchlogs", "log_groups", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsDynamoDBTables(ctx context.Context, resources []*dynamodb.Table) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "dynamodb", "tables", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2Addresses(ctx context.Context, resources []*ec2.Address) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "addresses", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2DhcpOptions(ctx context.Context, resources []*ec2.DhcpOptions) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "dhcp_options", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2Images(ctx context.Context, resources []*ec2.Image) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "images", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2Instances(ctx context.Context, resources []*ec2.Instance) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "instances", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2InternetGateways(ctx context.Context, resources []*ec2.InternetGateway) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "internet_gateways", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2ManagedPrefixLists(ctx context.Context, resources []*ec2.ManagedPrefixList) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "managed_prefix_lists", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2NatGateways(ctx context.Context, resources []*ec2.NatGateway) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "nat_gateways", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2NetworkAcls(ctx context.Context, resources []*ec2.NetworkAcl) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "network_acls", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2NetworkInterfaces(ctx context.Context, resources []*ec2.NetworkInterface) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "network_interfaces", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2PlacementGroups(ctx context.Context, resources []*ec2.PlacementGroup) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "placement_groups", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2ReservedInstances(ctx context.Context, resources []*ec2.ReservedInstances) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "reserved_instances", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2RouteTables(ctx context.Context, resources []*ec2.RouteTable) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "route_tables", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2SecurityGroups(ctx context.Context, resources []*ec2.SecurityGroup) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "security_groups", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2Subnets(ctx context.Context, resources []*ec2.Subnet) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "subnets", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2TransitGatewayPeeringAttachments(ctx context.Context, resources []*ec2.TransitGatewayPeeringAttachment) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "transit_gateway_peering_attachments", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2TransitGatewayRouteTables(ctx context.Context, resources []*ec2.TransitGatewayRouteTable) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "transit_gateway_route_tables", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2TransitGatewayVpcAttachments(ctx context.Context, resources []*ec2.TransitGatewayVpcAttachment) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "transit_gateway_vpc_attachments", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2TransitGateways(ctx context.Context, resources []*ec2.TransitGateway) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "transit_gateways", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2Volumes(ctx context.Context, resources []*ec2.Volume) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "volumes", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2VpcEndpoints(ctx context.Context, resources []*ec2.VpcEndpoint) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "vpc_endpoints", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2VpcPeeringConnections(ctx context.Context, resources []*ec2.VpcPeeringConnection) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "vpc_peering_connections", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2Vpcs(ctx context.Context, resources []*ec2.Vpc) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "vpcs", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEC2VpnGateways(ctx context.Context, resources []*ec2.VpnGateway) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ec2", "vpn_gateways", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsECSClusters(ctx context.Context, resources []*ecs.Cluster) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ecs", "clusters", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsECSServices(ctx context.Context, resources []*ecs.Service) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ecs", "services", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsECSTasks(ctx context.Context, resources []*ecs.Task) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "ecs", "tasks", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsEFSFileSystems(ctx context.Context, resources []*efs.FileSystem) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "efs", "filesystems", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsElastiCacheCacheClusters(ctx context.Context, resources []*elasticache.CacheCluster) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "elasticache", "cache_clusters", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsElasticLoadBalancingLoadBalancers(ctx context.Context, resources []*elasticloadbalancing.LoadBalancer) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "elasticloadbalancing", "load_balancers", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsElasticLoadBalancingV2LoadBalancers(ctx context.Context, resources []*elasticloadbalancingv2.LoadBalancer) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "elasticloadbalancingv2", "load_balancers", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsElasticLoadBalancingV2TargetGroups(ctx context.Context, resources []*elasticloadbalancingv2.TargetGroup) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "elasticloadbalancingv2", "target_groups", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsIAMGroups(ctx context.Context, resources []*iam.Group) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "iam", "groups", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsIAMPolicies(ctx context.Context, resources []*iam.Policy) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "iam", "policies", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsIAMRoles(ctx context.Context, resources []*iam.Role) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "iam", "roles", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsIAMUsers(ctx context.Context, resources []*iam.User) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "iam", "users", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsLambdaFunctions(ctx context.Context, resources []*lambda.Function) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "lambda", "functions", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsRDSDBClusters(ctx context.Context, resources []*rds.DBCluster) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "rds", "db_clusters", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsRDSDBInstances(ctx context.Context, resources []*rds.DBInstance) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "rds", "db_instances", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsRedshiftClusters(ctx context.Context, resources []*redshift.Cluster) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "redshift", "clusters", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsRoute53HostedZones(ctx context.Context, resources []*route53.HostedZone) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "route53", "hosted_zones", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsS3Buckets(ctx context.Context, resources []*s3.Bucket) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "s3", "buckets", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsSNSTopics(ctx context.Context, resources []*sns.Topic) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "sns", "topics", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsSNSSubscriptions(ctx context.Context, resources []*sns.Subscription) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "sns", "subscriptions", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsSQSQueues(ctx context.Context, resources []*sqs.Queue) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "sqs", "queues", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
func (dao *S3ParquetWriterDAO) PutAwsStorageGatewayGateways(ctx context.Context, resources []*storagegateway.Gateway) error {
	if len(resources) == 0 {
		return nil
	}
	file, err := dao.parquetClient.GetResourceFile(ctx, "aws", "storagegateway", "gateways", resources[0].ReportTime, resources[0])
	if err != nil {
		return err
	}
	file.Lock.Lock()
	defer file.Lock.Unlock()

	for _, resource := range resources {
		if err := file.Write(resource); err != nil {
			return fmt.Errorf("failed to write resource: %w", err)
		}
	}

	return nil
}
