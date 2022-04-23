//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_resource_file.tmpl
package ec2

type ReservedInstances struct {
	AvailabilityZone    string             `bson:"availability_zone,omitempty" ion:"availability_zone" dynamodbav:"availability_zone,omitempty" parquet:"name=availability_zone,type=BYTE_ARRAY,convertedtype=UTF8" json:"availability_zone,omitempty" diff:"availability_zone"`
	CurrencyCode        string             `bson:"currency_code,omitempty" ion:"currency_code" dynamodbav:"currency_code,omitempty" parquet:"name=currency_code,type=BYTE_ARRAY,convertedtype=UTF8" json:"currency_code,omitempty" diff:"currency_code"`
	Duration            int64              `bson:"duration,omitempty" ion:"duration" dynamodbav:"duration,omitempty" parquet:"name=duration,type=INT64" json:"duration,omitempty" diff:"duration"`
	End                 int64              `bson:"end,omitempty" ion:"end" dynamodbav:"end,omitempty" parquet:"name=end,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"end,omitempty" diff:"end"`
	FixedPrice          float32            `bson:"fixed_price,omitempty" ion:"fixed_price" dynamodbav:"fixed_price,omitempty" parquet:"name=fixed_price,type=FLOAT" json:"fixed_price,omitempty" diff:"fixed_price"`
	InstanceCount       int32              `bson:"instance_count,omitempty" ion:"instance_count" dynamodbav:"instance_count,omitempty" parquet:"name=instance_count,type=INT32" json:"instance_count,omitempty" diff:"instance_count"`
	InstanceTenancy     string             `bson:"instance_tenancy,omitempty" ion:"instance_tenancy" dynamodbav:"instance_tenancy,omitempty" parquet:"name=instance_tenancy,type=BYTE_ARRAY,convertedtype=UTF8" json:"instance_tenancy,omitempty" diff:"instance_tenancy"`
	InstanceType        string             `bson:"instance_type,omitempty" ion:"instance_type" dynamodbav:"instance_type,omitempty" parquet:"name=instance_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"instance_type,omitempty" diff:"instance_type"`
	OfferingClass       string             `bson:"offering_class,omitempty" ion:"offering_class" dynamodbav:"offering_class,omitempty" parquet:"name=offering_class,type=BYTE_ARRAY,convertedtype=UTF8" json:"offering_class,omitempty" diff:"offering_class"`
	OfferingType        string             `bson:"offering_type,omitempty" ion:"offering_type" dynamodbav:"offering_type,omitempty" parquet:"name=offering_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"offering_type,omitempty" diff:"offering_type"`
	ProductDescription  string             `bson:"product_description,omitempty" ion:"product_description" dynamodbav:"product_description,omitempty" parquet:"name=product_description,type=BYTE_ARRAY,convertedtype=UTF8" json:"product_description,omitempty" diff:"product_description"`
	RecurringCharges    []*RecurringCharge `bson:"recurring_charges,omitempty" ion:"recurring_charges" dynamodbav:"recurring_charges,omitempty" parquet:"name=recurring_charges,type=MAP,convertedtype=LIST" json:"recurring_charges,omitempty" diff:"recurring_charges"`
	ReservedInstancesId string             `bson:"reserved_instances_id,omitempty" ion:"reserved_instances_id" dynamodbav:"reserved_instances_id,omitempty" parquet:"name=reserved_instances_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"reserved_instances_id,omitempty" diff:"reserved_instances_id,identifier"`
	Scope               string             `bson:"scope,omitempty" ion:"scope" dynamodbav:"scope,omitempty" parquet:"name=scope,type=BYTE_ARRAY,convertedtype=UTF8" json:"scope,omitempty" diff:"scope"`
	Start               int64              `bson:"start,omitempty" ion:"start" dynamodbav:"start,omitempty" parquet:"name=start,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"start,omitempty" diff:"start"`
	State               string             `bson:"state,omitempty" ion:"state" dynamodbav:"state,omitempty" parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8" json:"state,omitempty" diff:"state"`
	Tags                map[string]string  `bson:"tags,omitempty" ion:"tags" dynamodbav:"tags,omitempty" parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"tags,omitempty" diff:"tags"`
	UsagePrice          float32            `bson:"usage_price,omitempty" ion:"usage_price" dynamodbav:"usage_price,omitempty" parquet:"name=usage_price,type=FLOAT" json:"usage_price,omitempty" diff:"usage_price"`
	AccountId           string             `bson:"account_id,omitempty" ion:"account_id" dynamodbav:"account_id,omitempty" parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"account_id,omitempty" diff:"account_id"`
	Region              string             `bson:"region,omitempty" ion:"region" dynamodbav:"region,omitempty" parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8" json:"region,omitempty" diff:"region"`
	ReportTime          int64              `bson:"report_time,omitempty" ion:"report_time" dynamodbav:"report_time,omitempty" parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID       string             `bson:"_id,omitempty" ion:"_id" dynamodbav:"_id,omitempty" parquet:"name=inventory_uuid,type=BYTE_ARRAY,convertedtype=UTF8" json:"_id,omitempty" diff:"-"`
}
