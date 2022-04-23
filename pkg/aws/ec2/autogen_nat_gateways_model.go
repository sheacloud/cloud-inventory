//AUTOGENERATED CODE DO NOT EDIT
package ec2

type NatGateway struct {
	ConnectivityType     string                `bson:"connectivity_type,omitempty" ion:"connectivity_type" dynamodbav:"connectivity_type,omitempty" parquet:"name=connectivity_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"connectivity_type,omitempty" diff:"connectivity_type"`
	CreateTime           int64                 `bson:"create_time,omitempty" ion:"create_time" dynamodbav:"create_time,omitempty" parquet:"name=create_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"create_time,omitempty" diff:"create_time"`
	DeleteTime           int64                 `bson:"delete_time,omitempty" ion:"delete_time" dynamodbav:"delete_time,omitempty" parquet:"name=delete_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"delete_time,omitempty" diff:"delete_time"`
	FailureCode          string                `bson:"failure_code,omitempty" ion:"failure_code" dynamodbav:"failure_code,omitempty" parquet:"name=failure_code,type=BYTE_ARRAY,convertedtype=UTF8" json:"failure_code,omitempty" diff:"failure_code"`
	FailureMessage       string                `bson:"failure_message,omitempty" ion:"failure_message" dynamodbav:"failure_message,omitempty" parquet:"name=failure_message,type=BYTE_ARRAY,convertedtype=UTF8" json:"failure_message,omitempty" diff:"failure_message"`
	NatGatewayAddresses  []*NatGatewayAddress  `bson:"nat_gateway_addresses,omitempty" ion:"nat_gateway_addresses" dynamodbav:"nat_gateway_addresses,omitempty" parquet:"name=nat_gateway_addresses,type=MAP,convertedtype=LIST" json:"nat_gateway_addresses,omitempty" diff:"nat_gateway_addresses"`
	NatGatewayId         string                `bson:"nat_gateway_id,omitempty" ion:"nat_gateway_id" dynamodbav:"nat_gateway_id,omitempty" parquet:"name=nat_gateway_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"nat_gateway_id,omitempty" diff:"nat_gateway_id,identifier"`
	ProvisionedBandwidth *ProvisionedBandwidth `bson:"provisioned_bandwidth,omitempty" ion:"provisioned_bandwidth" dynamodbav:"provisioned_bandwidth,omitempty" parquet:"name=provisioned_bandwidth" json:"provisioned_bandwidth,omitempty" diff:"provisioned_bandwidth"`
	State                string                `bson:"state,omitempty" ion:"state" dynamodbav:"state,omitempty" parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8" json:"state,omitempty" diff:"state"`
	SubnetId             string                `bson:"subnet_id,omitempty" ion:"subnet_id" dynamodbav:"subnet_id,omitempty" parquet:"name=subnet_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"subnet_id,omitempty" diff:"subnet_id"`
	Tags                 map[string]string     `bson:"tags,omitempty" ion:"tags" dynamodbav:"tags,omitempty" parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"tags,omitempty" diff:"tags"`
	VpcId                string                `bson:"vpc_id,omitempty" ion:"vpc_id" dynamodbav:"vpc_id,omitempty" parquet:"name=vpc_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"vpc_id,omitempty" diff:"vpc_id"`
	AccountId            string                `bson:"account_id,omitempty" ion:"account_id" dynamodbav:"account_id,omitempty" parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"account_id,omitempty" diff:"account_id"`
	Region               string                `bson:"region,omitempty" ion:"region" dynamodbav:"region,omitempty" parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8" json:"region,omitempty" diff:"region"`
	ReportTime           int64                 `bson:"report_time,omitempty" ion:"report_time" dynamodbav:"report_time,omitempty" parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID        string                `bson:"_id,omitempty" ion:"_id" dynamodbav:"_id,omitempty" parquet:"name=inventory_uuid,type=BYTE_ARRAY,convertedtype=UTF8" json:"_id,omitempty" diff:"-"`
}
