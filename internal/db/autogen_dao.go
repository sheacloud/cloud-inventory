//AUTOGENERATED CODE DO NOT EDIT
package db

import (
	"context"
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

type WriterDAO interface {
	WriteInventoryResults(ctx context.Context, metadata *meta.InventoryResults) error
	WriteIngestionTimestamp(ctx context.Context, metadata *meta.IngestionTimestamp) error
	FinishIndex(ctx context.Context, indices []string, reportDateUnixMilli int64) error
	Finish(ctx context.Context) error
	PutAwsApiGatewayRestApis(ctx context.Context, resources []*apigateway.RestApi) error
	PutAwsApiGatewayV2Apis(ctx context.Context, resources []*apigatewayv2.Api) error
	PutAwsBackupBackupVaults(ctx context.Context, resources []*backup.BackupVault) error
	PutAwsBackupBackupPlans(ctx context.Context, resources []*backup.BackupPlan) error
	PutAwsCloudTrailTrails(ctx context.Context, resources []*cloudtrail.Trail) error
	PutAwsCloudWatchLogsLogGroups(ctx context.Context, resources []*cloudwatchlogs.LogGroup) error
	PutAwsDynamoDBTables(ctx context.Context, resources []*dynamodb.Table) error
	PutAwsEC2Addresses(ctx context.Context, resources []*ec2.Address) error
	PutAwsEC2DhcpOptions(ctx context.Context, resources []*ec2.DhcpOptions) error
	PutAwsEC2Images(ctx context.Context, resources []*ec2.Image) error
	PutAwsEC2Instances(ctx context.Context, resources []*ec2.Instance) error
	PutAwsEC2InternetGateways(ctx context.Context, resources []*ec2.InternetGateway) error
	PutAwsEC2ManagedPrefixLists(ctx context.Context, resources []*ec2.ManagedPrefixList) error
	PutAwsEC2NatGateways(ctx context.Context, resources []*ec2.NatGateway) error
	PutAwsEC2NetworkAcls(ctx context.Context, resources []*ec2.NetworkAcl) error
	PutAwsEC2NetworkInterfaces(ctx context.Context, resources []*ec2.NetworkInterface) error
	PutAwsEC2PlacementGroups(ctx context.Context, resources []*ec2.PlacementGroup) error
	PutAwsEC2ReservedInstances(ctx context.Context, resources []*ec2.ReservedInstances) error
	PutAwsEC2RouteTables(ctx context.Context, resources []*ec2.RouteTable) error
	PutAwsEC2SecurityGroups(ctx context.Context, resources []*ec2.SecurityGroup) error
	PutAwsEC2Subnets(ctx context.Context, resources []*ec2.Subnet) error
	PutAwsEC2TransitGatewayPeeringAttachments(ctx context.Context, resources []*ec2.TransitGatewayPeeringAttachment) error
	PutAwsEC2TransitGatewayRouteTables(ctx context.Context, resources []*ec2.TransitGatewayRouteTable) error
	PutAwsEC2TransitGatewayVpcAttachments(ctx context.Context, resources []*ec2.TransitGatewayVpcAttachment) error
	PutAwsEC2TransitGateways(ctx context.Context, resources []*ec2.TransitGateway) error
	PutAwsEC2Volumes(ctx context.Context, resources []*ec2.Volume) error
	PutAwsEC2VpcEndpoints(ctx context.Context, resources []*ec2.VpcEndpoint) error
	PutAwsEC2VpcPeeringConnections(ctx context.Context, resources []*ec2.VpcPeeringConnection) error
	PutAwsEC2Vpcs(ctx context.Context, resources []*ec2.Vpc) error
	PutAwsEC2VpnGateways(ctx context.Context, resources []*ec2.VpnGateway) error
	PutAwsECSClusters(ctx context.Context, resources []*ecs.Cluster) error
	PutAwsECSServices(ctx context.Context, resources []*ecs.Service) error
	PutAwsECSTasks(ctx context.Context, resources []*ecs.Task) error
	PutAwsEFSFileSystems(ctx context.Context, resources []*efs.FileSystem) error
	PutAwsElastiCacheCacheClusters(ctx context.Context, resources []*elasticache.CacheCluster) error
	PutAwsElasticLoadBalancingLoadBalancers(ctx context.Context, resources []*elasticloadbalancing.LoadBalancer) error
	PutAwsElasticLoadBalancingV2LoadBalancers(ctx context.Context, resources []*elasticloadbalancingv2.LoadBalancer) error
	PutAwsElasticLoadBalancingV2TargetGroups(ctx context.Context, resources []*elasticloadbalancingv2.TargetGroup) error
	PutAwsIAMGroups(ctx context.Context, resources []*iam.Group) error
	PutAwsIAMPolicies(ctx context.Context, resources []*iam.Policy) error
	PutAwsIAMRoles(ctx context.Context, resources []*iam.Role) error
	PutAwsIAMUsers(ctx context.Context, resources []*iam.User) error
	PutAwsLambdaFunctions(ctx context.Context, resources []*lambda.Function) error
	PutAwsRDSDBClusters(ctx context.Context, resources []*rds.DBCluster) error
	PutAwsRDSDBInstances(ctx context.Context, resources []*rds.DBInstance) error
	PutAwsRedshiftClusters(ctx context.Context, resources []*redshift.Cluster) error
	PutAwsRoute53HostedZones(ctx context.Context, resources []*route53.HostedZone) error
	PutAwsS3Buckets(ctx context.Context, resources []*s3.Bucket) error
	PutAwsSNSTopics(ctx context.Context, resources []*sns.Topic) error
	PutAwsSNSSubscriptions(ctx context.Context, resources []*sns.Subscription) error
	PutAwsSQSQueues(ctx context.Context, resources []*sqs.Queue) error
	PutAwsStorageGatewayGateways(ctx context.Context, resources []*storagegateway.Gateway) error
}

type ReaderDAO interface {
	ListAwsApiGatewayRestApis(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*apigateway.RestApi, error)
	GetAwsApiGatewayRestApi(ctx context.Context, reportTimeUnixMilli int64, id string) (*apigateway.RestApi, error)
	GetAwsApiGatewayRestApiReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsApiGatewayRestApiReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsApiGatewayV2Apis(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*apigatewayv2.Api, error)
	GetAwsApiGatewayV2Api(ctx context.Context, reportTimeUnixMilli int64, id string) (*apigatewayv2.Api, error)
	GetAwsApiGatewayV2ApiReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsApiGatewayV2ApiReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsBackupBackupVaults(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*backup.BackupVault, error)
	GetAwsBackupBackupVault(ctx context.Context, reportTimeUnixMilli int64, id string) (*backup.BackupVault, error)
	GetAwsBackupBackupVaultReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsBackupBackupVaultReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsBackupBackupPlans(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*backup.BackupPlan, error)
	GetAwsBackupBackupPlan(ctx context.Context, reportTimeUnixMilli int64, id string) (*backup.BackupPlan, error)
	GetAwsBackupBackupPlanReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsBackupBackupPlanReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsCloudTrailTrails(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*cloudtrail.Trail, error)
	GetAwsCloudTrailTrail(ctx context.Context, reportTimeUnixMilli int64, id string) (*cloudtrail.Trail, error)
	GetAwsCloudTrailTrailReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsCloudTrailTrailReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsCloudWatchLogsLogGroups(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*cloudwatchlogs.LogGroup, error)
	GetAwsCloudWatchLogsLogGroup(ctx context.Context, reportTimeUnixMilli int64, id string) (*cloudwatchlogs.LogGroup, error)
	GetAwsCloudWatchLogsLogGroupReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsCloudWatchLogsLogGroupReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsDynamoDBTables(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*dynamodb.Table, error)
	GetAwsDynamoDBTable(ctx context.Context, reportTimeUnixMilli int64, id string) (*dynamodb.Table, error)
	GetAwsDynamoDBTableReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsDynamoDBTableReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2Addresses(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.Address, error)
	GetAwsEC2Address(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.Address, error)
	GetAwsEC2AddressReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2AddressReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2DhcpOptions(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.DhcpOptions, error)
	GetAwsEC2DhcpOptions(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.DhcpOptions, error)
	GetAwsEC2DhcpOptionsReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2DhcpOptionsReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2Images(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.Image, error)
	GetAwsEC2Image(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.Image, error)
	GetAwsEC2ImageReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2ImageReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2Instances(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.Instance, error)
	GetAwsEC2Instance(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.Instance, error)
	GetAwsEC2InstanceReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2InstanceReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2InternetGateways(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.InternetGateway, error)
	GetAwsEC2InternetGateway(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.InternetGateway, error)
	GetAwsEC2InternetGatewayReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2InternetGatewayReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2ManagedPrefixLists(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.ManagedPrefixList, error)
	GetAwsEC2ManagedPrefixList(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.ManagedPrefixList, error)
	GetAwsEC2ManagedPrefixListReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2ManagedPrefixListReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2NatGateways(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.NatGateway, error)
	GetAwsEC2NatGateway(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.NatGateway, error)
	GetAwsEC2NatGatewayReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2NatGatewayReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2NetworkAcls(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.NetworkAcl, error)
	GetAwsEC2NetworkAcl(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.NetworkAcl, error)
	GetAwsEC2NetworkAclReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2NetworkAclReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2NetworkInterfaces(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.NetworkInterface, error)
	GetAwsEC2NetworkInterface(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.NetworkInterface, error)
	GetAwsEC2NetworkInterfaceReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2NetworkInterfaceReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2PlacementGroups(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.PlacementGroup, error)
	GetAwsEC2PlacementGroup(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.PlacementGroup, error)
	GetAwsEC2PlacementGroupReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2PlacementGroupReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2ReservedInstances(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.ReservedInstances, error)
	GetAwsEC2ReservedInstances(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.ReservedInstances, error)
	GetAwsEC2ReservedInstancesReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2ReservedInstancesReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2RouteTables(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.RouteTable, error)
	GetAwsEC2RouteTable(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.RouteTable, error)
	GetAwsEC2RouteTableReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2RouteTableReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2SecurityGroups(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.SecurityGroup, error)
	GetAwsEC2SecurityGroup(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.SecurityGroup, error)
	GetAwsEC2SecurityGroupReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2SecurityGroupReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2Subnets(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.Subnet, error)
	GetAwsEC2Subnet(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.Subnet, error)
	GetAwsEC2SubnetReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2SubnetReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2TransitGatewayPeeringAttachments(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.TransitGatewayPeeringAttachment, error)
	GetAwsEC2TransitGatewayPeeringAttachment(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.TransitGatewayPeeringAttachment, error)
	GetAwsEC2TransitGatewayPeeringAttachmentReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2TransitGatewayPeeringAttachmentReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2TransitGatewayRouteTables(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.TransitGatewayRouteTable, error)
	GetAwsEC2TransitGatewayRouteTable(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.TransitGatewayRouteTable, error)
	GetAwsEC2TransitGatewayRouteTableReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2TransitGatewayRouteTableReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2TransitGatewayVpcAttachments(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.TransitGatewayVpcAttachment, error)
	GetAwsEC2TransitGatewayVpcAttachment(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.TransitGatewayVpcAttachment, error)
	GetAwsEC2TransitGatewayVpcAttachmentReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2TransitGatewayVpcAttachmentReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2TransitGateways(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.TransitGateway, error)
	GetAwsEC2TransitGateway(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.TransitGateway, error)
	GetAwsEC2TransitGatewayReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2TransitGatewayReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2Volumes(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.Volume, error)
	GetAwsEC2Volume(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.Volume, error)
	GetAwsEC2VolumeReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2VolumeReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2VpcEndpoints(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.VpcEndpoint, error)
	GetAwsEC2VpcEndpoint(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.VpcEndpoint, error)
	GetAwsEC2VpcEndpointReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2VpcEndpointReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2VpcPeeringConnections(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.VpcPeeringConnection, error)
	GetAwsEC2VpcPeeringConnection(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.VpcPeeringConnection, error)
	GetAwsEC2VpcPeeringConnectionReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2VpcPeeringConnectionReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2Vpcs(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.Vpc, error)
	GetAwsEC2Vpc(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.Vpc, error)
	GetAwsEC2VpcReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2VpcReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEC2VpnGateways(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ec2.VpnGateway, error)
	GetAwsEC2VpnGateway(ctx context.Context, reportTimeUnixMilli int64, id string) (*ec2.VpnGateway, error)
	GetAwsEC2VpnGatewayReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEC2VpnGatewayReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsECSClusters(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ecs.Cluster, error)
	GetAwsECSCluster(ctx context.Context, reportTimeUnixMilli int64, id string) (*ecs.Cluster, error)
	GetAwsECSClusterReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsECSClusterReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsECSServices(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ecs.Service, error)
	GetAwsECSService(ctx context.Context, reportTimeUnixMilli int64, id string) (*ecs.Service, error)
	GetAwsECSServiceReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsECSServiceReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsECSTasks(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*ecs.Task, error)
	GetAwsECSTask(ctx context.Context, reportTimeUnixMilli int64, id string) (*ecs.Task, error)
	GetAwsECSTaskReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsECSTaskReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsEFSFileSystems(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*efs.FileSystem, error)
	GetAwsEFSFileSystem(ctx context.Context, reportTimeUnixMilli int64, id string) (*efs.FileSystem, error)
	GetAwsEFSFileSystemReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsEFSFileSystemReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsElastiCacheCacheClusters(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*elasticache.CacheCluster, error)
	GetAwsElastiCacheCacheCluster(ctx context.Context, reportTimeUnixMilli int64, id string) (*elasticache.CacheCluster, error)
	GetAwsElastiCacheCacheClusterReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsElastiCacheCacheClusterReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsElasticLoadBalancingLoadBalancers(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*elasticloadbalancing.LoadBalancer, error)
	GetAwsElasticLoadBalancingLoadBalancer(ctx context.Context, reportTimeUnixMilli int64, id string) (*elasticloadbalancing.LoadBalancer, error)
	GetAwsElasticLoadBalancingLoadBalancerReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsElasticLoadBalancingLoadBalancerReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsElasticLoadBalancingV2LoadBalancers(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*elasticloadbalancingv2.LoadBalancer, error)
	GetAwsElasticLoadBalancingV2LoadBalancer(ctx context.Context, reportTimeUnixMilli int64, id string) (*elasticloadbalancingv2.LoadBalancer, error)
	GetAwsElasticLoadBalancingV2LoadBalancerReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsElasticLoadBalancingV2LoadBalancerReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsElasticLoadBalancingV2TargetGroups(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*elasticloadbalancingv2.TargetGroup, error)
	GetAwsElasticLoadBalancingV2TargetGroup(ctx context.Context, reportTimeUnixMilli int64, id string) (*elasticloadbalancingv2.TargetGroup, error)
	GetAwsElasticLoadBalancingV2TargetGroupReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsElasticLoadBalancingV2TargetGroupReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsIAMGroups(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*iam.Group, error)
	GetAwsIAMGroup(ctx context.Context, reportTimeUnixMilli int64, id string) (*iam.Group, error)
	GetAwsIAMGroupReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsIAMGroupReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsIAMPolicies(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*iam.Policy, error)
	GetAwsIAMPolicy(ctx context.Context, reportTimeUnixMilli int64, id string) (*iam.Policy, error)
	GetAwsIAMPolicyReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsIAMPolicyReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsIAMRoles(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*iam.Role, error)
	GetAwsIAMRole(ctx context.Context, reportTimeUnixMilli int64, id string) (*iam.Role, error)
	GetAwsIAMRoleReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsIAMRoleReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsIAMUsers(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*iam.User, error)
	GetAwsIAMUser(ctx context.Context, reportTimeUnixMilli int64, id string) (*iam.User, error)
	GetAwsIAMUserReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsIAMUserReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsLambdaFunctions(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*lambda.Function, error)
	GetAwsLambdaFunction(ctx context.Context, reportTimeUnixMilli int64, id string) (*lambda.Function, error)
	GetAwsLambdaFunctionReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsLambdaFunctionReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsRDSDBClusters(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*rds.DBCluster, error)
	GetAwsRDSDBCluster(ctx context.Context, reportTimeUnixMilli int64, id string) (*rds.DBCluster, error)
	GetAwsRDSDBClusterReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsRDSDBClusterReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsRDSDBInstances(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*rds.DBInstance, error)
	GetAwsRDSDBInstance(ctx context.Context, reportTimeUnixMilli int64, id string) (*rds.DBInstance, error)
	GetAwsRDSDBInstanceReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsRDSDBInstanceReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsRedshiftClusters(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*redshift.Cluster, error)
	GetAwsRedshiftCluster(ctx context.Context, reportTimeUnixMilli int64, id string) (*redshift.Cluster, error)
	GetAwsRedshiftClusterReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsRedshiftClusterReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsRoute53HostedZones(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*route53.HostedZone, error)
	GetAwsRoute53HostedZone(ctx context.Context, reportTimeUnixMilli int64, id string) (*route53.HostedZone, error)
	GetAwsRoute53HostedZoneReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsRoute53HostedZoneReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsS3Buckets(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*s3.Bucket, error)
	GetAwsS3Bucket(ctx context.Context, reportTimeUnixMilli int64, id string) (*s3.Bucket, error)
	GetAwsS3BucketReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsS3BucketReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsSNSTopics(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*sns.Topic, error)
	GetAwsSNSTopic(ctx context.Context, reportTimeUnixMilli int64, id string) (*sns.Topic, error)
	GetAwsSNSTopicReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsSNSTopicReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsSNSSubscriptions(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*sns.Subscription, error)
	GetAwsSNSSubscription(ctx context.Context, reportTimeUnixMilli int64, id string) (*sns.Subscription, error)
	GetAwsSNSSubscriptionReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsSNSSubscriptionReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsSQSQueues(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*sqs.Queue, error)
	GetAwsSQSQueue(ctx context.Context, reportTimeUnixMilli int64, id string) (*sqs.Queue, error)
	GetAwsSQSQueueReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsSQSQueueReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
	ListAwsStorageGatewayGateways(ctx context.Context, reportTimeUnixMilli int64, accountID, region *string, limit, offset *int64) ([]*storagegateway.Gateway, error)
	GetAwsStorageGatewayGateway(ctx context.Context, reportTimeUnixMilli int64, id string) (*storagegateway.Gateway, error)
	GetAwsStorageGatewayGatewayReportTimes(ctx context.Context, reportDateUnixMilli int64) ([]int64, error)
	GetReferencedAwsStorageGatewayGatewayReportTime(ctx context.Context, reportDateUnixMilli int64, timeSelection TimeSelection, timeReferenceUnixMilli int64) (*int64, error)
}
