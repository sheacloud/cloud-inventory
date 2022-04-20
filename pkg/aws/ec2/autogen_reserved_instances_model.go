//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"time"
)

type ReservedInstances struct {
	AvailabilityZone    string             `bson:"availability_zone,omitempty" dynamodbav:"availability_zone,omitempty" json:"availability_zone,omitempty" diff:"availability_zone"`
	CurrencyCode        string             `bson:"currency_code,omitempty" dynamodbav:"currency_code,omitempty" json:"currency_code,omitempty" diff:"currency_code"`
	Duration            int64              `bson:"duration,omitempty" dynamodbav:"duration,omitempty" json:"duration,omitempty" diff:"duration"`
	End                 *time.Time         `bson:"end,omitempty" dynamodbav:"end,unixtime,omitempty" json:"end,omitempty" diff:"end"`
	FixedPrice          float32            `bson:"fixed_price,omitempty" dynamodbav:"fixed_price,omitempty" json:"fixed_price,omitempty" diff:"fixed_price"`
	InstanceCount       int32              `bson:"instance_count,omitempty" dynamodbav:"instance_count,omitempty" json:"instance_count,omitempty" diff:"instance_count"`
	InstanceTenancy     string             `bson:"instance_tenancy,omitempty" dynamodbav:"instance_tenancy,omitempty" json:"instance_tenancy,omitempty" diff:"instance_tenancy"`
	InstanceType        string             `bson:"instance_type,omitempty" dynamodbav:"instance_type,omitempty" json:"instance_type,omitempty" diff:"instance_type"`
	OfferingClass       string             `bson:"offering_class,omitempty" dynamodbav:"offering_class,omitempty" json:"offering_class,omitempty" diff:"offering_class"`
	OfferingType        string             `bson:"offering_type,omitempty" dynamodbav:"offering_type,omitempty" json:"offering_type,omitempty" diff:"offering_type"`
	ProductDescription  string             `bson:"product_description,omitempty" dynamodbav:"product_description,omitempty" json:"product_description,omitempty" diff:"product_description"`
	RecurringCharges    []*RecurringCharge `bson:"recurring_charges,omitempty" dynamodbav:"recurring_charges,omitempty" json:"recurring_charges,omitempty" diff:"recurring_charges"`
	ReservedInstancesId string             `bson:"reserved_instances_id,omitempty" dynamodbav:"reserved_instances_id,omitempty" inventory_primary_key:"true" json:"reserved_instances_id,omitempty" diff:"reserved_instances_id,identifier"`
	Scope               string             `bson:"scope,omitempty" dynamodbav:"scope,omitempty" json:"scope,omitempty" diff:"scope"`
	Start               *time.Time         `bson:"start,omitempty" dynamodbav:"start,unixtime,omitempty" json:"start,omitempty" diff:"start"`
	State               string             `bson:"state,omitempty" dynamodbav:"state,omitempty" json:"state,omitempty" diff:"state"`
	Tags                map[string]string  `bson:"tags,omitempty" dynamodbav:"tags,omitempty" json:"tags,omitempty" diff:"tags"`
	UsagePrice          float32            `bson:"usage_price,omitempty" dynamodbav:"usage_price,omitempty" json:"usage_price,omitempty" diff:"usage_price"`
	AccountId           string             `bson:"account_id,omitempty" dynamodbav:"account_id,omitempty" json:"account_id,omitempty" diff:"account_id"`
	Region              string             `bson:"region,omitempty" dynamodbav:"region,omitempty" json:"region,omitempty" diff:"region"`
	ReportTime          time.Time          `bson:"report_time,omitempty" dynamodbav:"report_time,unixtime,omitempty" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID       string             `bson:"_id,omitempty" dynamodbav:"_id,omitempty" json:"_id,omitempty" diff:"-"`
}