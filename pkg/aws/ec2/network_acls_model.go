// AUTOGENERATED, DO NOT EDIT
package ec2

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

var customNetworkAclModelPostprocessingFuncs []func(x *NetworkAclModel) = []func(x *NetworkAclModel){}
var customNetworkAclModelFuncsLock sync.Mutex

func registerCustomNetworkAclModelPostprocessingFunc(f func(x *NetworkAclModel)) {
	customNetworkAclModelFuncsLock.Lock()
	defer customNetworkAclModelFuncsLock.Unlock()

	customNetworkAclModelPostprocessingFuncs = append(customNetworkAclModelPostprocessingFuncs, f)
}

func init() {
	Controller.RegisterDataSource("network_acls", NetworkAclDataSource)
}

type NetworkAclModel struct {
	Associations []*NetworkAclAssociationNetworkAclModel `parquet:"name=associations,type=LIST"`
	Entries      []*NetworkAclEntryNetworkAclModel       `parquet:"name=entries,type=LIST"`
	IsDefault    bool                                    `parquet:"name=is_default,type=BOOLEAN"`
	NetworkAclId string                                  `parquet:"name=network_acl_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true"`
	OwnerId      string                                  `parquet:"name=owner_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	Tags         map[string]string                       `parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8"`
	VpcId        string                                  `parquet:"name=vpc_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	AccountId    string                                  `parquet:"name=account_id, type=BYTE_ARRAY, convertedtype=UTF8"`
	Region       string                                  `parquet:"name=region, type=BYTE_ARRAY, convertedtype=UTF8"`
	ReportTime   int64                                   `parquet:"name=report_time, type=INT64, convertedtype=TIMESTAMP_MILLIS"`
}

type NetworkAclAssociationNetworkAclModel struct {
	NetworkAclAssociationId string `parquet:"name=network_acl_association_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	NetworkAclId            string `parquet:"name=network_acl_id,type=BYTE_ARRAY,convertedtype=UTF8"`
	SubnetId                string `parquet:"name=subnet_id,type=BYTE_ARRAY,convertedtype=UTF8"`
}

type NetworkAclEntryNetworkAclModel struct {
	CidrBlock     string                       `parquet:"name=cidr_block,type=BYTE_ARRAY,convertedtype=UTF8"`
	Egress        bool                         `parquet:"name=egress,type=BOOLEAN"`
	IcmpTypeCode  *IcmpTypeCodeNetworkAclModel `parquet:"name=icmp_type_code"`
	Ipv6CidrBlock string                       `parquet:"name=ipv6_cidr_block,type=BYTE_ARRAY,convertedtype=UTF8"`
	PortRange     *PortRangeNetworkAclModel    `parquet:"name=port_range"`
	Protocol      string                       `parquet:"name=protocol,type=BYTE_ARRAY,convertedtype=UTF8"`
	RuleAction    string                       `parquet:"name=rule_action,type=BYTE_ARRAY,convertedtype=UTF8"`
	RuleNumber    int32                        `parquet:"name=rule_number,type=INT32"`
}

type IcmpTypeCodeNetworkAclModel struct {
	Code int32 `parquet:"name=code,type=INT32"`
	Type int32 `parquet:"name=type,type=INT32"`
}

type PortRangeNetworkAclModel struct {
	From int32 `parquet:"name=from,type=INT32"`
	To   int32 `parquet:"name=to,type=INT32"`
}

type TagNetworkAclModel struct {
	Key   string `parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8"`
	Value string `parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8"`
}

func NetworkAclDataSource(ctx context.Context, client *ec2.Client, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new(NetworkAclModel))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := ec2.NewDescribeNetworkAclsPaginator(client, &ec2.DescribeNetworkAclsInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":     storageConfig.Service,
				"data_source": storageConfig.DataSource,
				"account_id":  storageConfig.AccountId,
				"region":      storageConfig.Region,
				"cloud":       storageConfig.Cloud,
				"error":       err,
			}).Error("error calling DescribeNetworkAcls")
			return err
		}

		for _, var0 := range output.NetworkAcls {

			model := new(NetworkAclModel)
			copier.Copy(&model, &var0)

			model.Tags = GetTagMap(var0.Tags)
			model.AccountId = storageConfig.AccountId
			model.Region = storageConfig.Region
			model.ReportTime = reportTime.UTC().UnixMilli()

			for _, f := range customNetworkAclModelPostprocessingFuncs {
				f(model)
			}

			errors := storageContextSet.Store(ctx, model)
			for storageContext, err := range errors {
				storage.LogContextError(storageContext, fmt.Sprintf("Error storing NetworkAclModel: %v", err))
			}
		}

	}

	return nil
}
