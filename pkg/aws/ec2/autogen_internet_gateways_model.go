//AUTOGENERATED CODE DO NOT EDIT
package ec2

import (
	"time"
)

type InternetGateway struct {
	Attachments       []*InternetGatewayAttachment `bson:"attachments,omitempty" dynamodbav:"attachments,omitempty" json:"attachments,omitempty" diff:"attachments"`
	InternetGatewayId string                       `bson:"internet_gateway_id,omitempty" dynamodbav:"internet_gateway_id,omitempty" inventory_primary_key:"true" json:"internet_gateway_id,omitempty" diff:"internet_gateway_id,identifier"`
	OwnerId           string                       `bson:"owner_id,omitempty" dynamodbav:"owner_id,omitempty" json:"owner_id,omitempty" diff:"owner_id"`
	Tags              map[string]string            `bson:"tags,omitempty" dynamodbav:"tags,omitempty" json:"tags,omitempty" diff:"tags"`
	AccountId         string                       `bson:"account_id,omitempty" dynamodbav:"account_id,omitempty" json:"account_id,omitempty" diff:"account_id"`
	Region            string                       `bson:"region,omitempty" dynamodbav:"region,omitempty" json:"region,omitempty" diff:"region"`
	ReportTime        time.Time                    `bson:"report_time,omitempty" dynamodbav:"report_time,unixtime,omitempty" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID     string                       `bson:"_id,omitempty" dynamodbav:"_id,omitempty" json:"_id,omitempty" diff:"-"`
}
