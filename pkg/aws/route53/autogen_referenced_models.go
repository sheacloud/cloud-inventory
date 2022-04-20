//AUTOGENERATED CODE DO NOT EDIT
package route53

type HostedZoneConfig struct {
	Comment     string `bson:"comment,omitempty" dynamodbav:"comment,omitempty" json:"comment,omitempty" diff:"comment"`
	PrivateZone bool   `bson:"private_zone,omitempty" dynamodbav:"private_zone" json:"private_zone,omitempty" diff:"private_zone"`
}

type LinkedService struct {
	Description      string `bson:"description,omitempty" dynamodbav:"description,omitempty" json:"description,omitempty" diff:"description"`
	ServicePrincipal string `bson:"service_principal,omitempty" dynamodbav:"service_principal,omitempty" json:"service_principal,omitempty" diff:"service_principal"`
}

type VPC struct {
	VPCId     string `bson:"vpc_id,omitempty" dynamodbav:"vpc_id,omitempty" json:"vpc_id,omitempty" diff:"vpc_id"`
	VPCRegion string `bson:"vpc_region,omitempty" dynamodbav:"vpc_region,omitempty" json:"vpc_region,omitempty" diff:"vpc_region"`
}