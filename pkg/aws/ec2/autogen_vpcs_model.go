//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_resource_file.tmpl
package ec2

type Vpc struct {
	CidrBlock                   string                         `bson:"cidr_block,omitempty" ion:"cidr_block" dynamodbav:"cidr_block,omitempty" parquet:"name=cidr_block,type=BYTE_ARRAY,convertedtype=UTF8" json:"cidr_block,omitempty" diff:"cidr_block"`
	CidrBlockAssociationSet     []*VpcCidrBlockAssociation     `bson:"cidr_block_association_set,omitempty" ion:"cidr_block_association_set" dynamodbav:"cidr_block_association_set,omitempty" parquet:"name=cidr_block_association_set,type=MAP,convertedtype=LIST" json:"cidr_block_association_set,omitempty" diff:"cidr_block_association_set"`
	DhcpOptionsId               string                         `bson:"dhcp_options_id,omitempty" ion:"dhcp_options_id" dynamodbav:"dhcp_options_id,omitempty" parquet:"name=dhcp_options_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"dhcp_options_id,omitempty" diff:"dhcp_options_id"`
	InstanceTenancy             string                         `bson:"instance_tenancy,omitempty" ion:"instance_tenancy" dynamodbav:"instance_tenancy,omitempty" parquet:"name=instance_tenancy,type=BYTE_ARRAY,convertedtype=UTF8" json:"instance_tenancy,omitempty" diff:"instance_tenancy"`
	Ipv6CidrBlockAssociationSet []*VpcIpv6CidrBlockAssociation `bson:"ipv6_cidr_block_association_set,omitempty" ion:"ipv6_cidr_block_association_set" dynamodbav:"ipv6_cidr_block_association_set,omitempty" parquet:"name=ipv6_cidr_block_association_set,type=MAP,convertedtype=LIST" json:"ipv6_cidr_block_association_set,omitempty" diff:"ipv6_cidr_block_association_set"`
	IsDefault                   bool                           `bson:"is_default,omitempty" ion:"is_default" dynamodbav:"is_default" parquet:"name=is_default,type=BOOLEAN" json:"is_default,omitempty" diff:"is_default"`
	OwnerId                     string                         `bson:"owner_id,omitempty" ion:"owner_id" dynamodbav:"owner_id,omitempty" parquet:"name=owner_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"owner_id,omitempty" diff:"owner_id"`
	State                       string                         `bson:"state,omitempty" ion:"state" dynamodbav:"state,omitempty" parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8" json:"state,omitempty" diff:"state"`
	Tags                        map[string]string              `bson:"tags,omitempty" ion:"tags" dynamodbav:"tags,omitempty" parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"tags,omitempty" diff:"tags"`
	VpcId                       string                         `bson:"vpc_id,omitempty" ion:"vpc_id" dynamodbav:"vpc_id,omitempty" parquet:"name=vpc_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"vpc_id,omitempty" diff:"vpc_id,identifier"`
	AccountId                   string                         `bson:"account_id,omitempty" ion:"account_id" dynamodbav:"account_id,omitempty" parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"account_id,omitempty" diff:"account_id"`
	Region                      string                         `bson:"region,omitempty" ion:"region" dynamodbav:"region,omitempty" parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8" json:"region,omitempty" diff:"region"`
	ReportTime                  int64                          `bson:"report_time,omitempty" ion:"report_time" dynamodbav:"report_time,omitempty" parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID               string                         `bson:"_id,omitempty" ion:"_id" dynamodbav:"_id,omitempty" parquet:"name=inventory_uuid,type=BYTE_ARRAY,convertedtype=UTF8" json:"_id,omitempty" diff:"-"`
}
