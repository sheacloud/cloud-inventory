//AUTOGENERATED CODE DO NOT EDIT
package s3

import (
	"time"
)

type Bucket struct {
	CreationDate                      *time.Time                         `bson:"creation_date,omitempty" dynamodbav:"creation_date,unixtime,omitempty" json:"creation_date,omitempty" diff:"creation_date"`
	Name                              string                             `bson:"name,omitempty" dynamodbav:"name,omitempty" inventory_primary_key:"true" json:"name,omitempty" diff:"name,identifier"`
	AccountId                         string                             `bson:"account_id,omitempty" dynamodbav:"account_id,omitempty" json:"account_id,omitempty" diff:"account_id"`
	Region                            string                             `bson:"region,omitempty" dynamodbav:"region,omitempty" json:"region,omitempty" diff:"region"`
	ReportTime                        time.Time                          `bson:"report_time,omitempty" dynamodbav:"report_time,unixtime,omitempty" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID                     string                             `bson:"_id,omitempty" dynamodbav:"_id,omitempty" json:"_id,omitempty" diff:"-"`
	Policy                            string                             `bson:"policy,omitempty" dynamodbav:"policy,omitempty" json:"policy,omitempty" diff:"policy"`
	IsPublic                          bool                               `bson:"is_public,omitempty" dynamodbav:"is_public" json:"is_public,omitempty" diff:"is_public"`
	Tags                              map[string]string                  `bson:"tags,omitempty" dynamodbav:"tags,omitempty" json:"tags,omitempty" diff:"tags"`
	VersioningStatus                  string                             `bson:"versioning_status,omitempty" dynamodbav:"versioning_status,omitempty" json:"versioning_status,omitempty" diff:"versioning_status"`
	MFADeleteStatus                   string                             `bson:"mfa_delete_status,omitempty" dynamodbav:"mfa_delete_status,omitempty" json:"mfa_delete_status,omitempty" diff:"mfa_delete_status"`
	ReplicationConfiguration          *ReplicationConfiguration          `bson:"replication_configuration,omitempty" dynamodbav:"replication_configuration,omitempty" json:"replication_configuration,omitempty" diff:"replication_configuration"`
	AclGrants                         []*Grant                           `bson:"acl_grants,omitempty" dynamodbav:"acl_grants,omitempty" json:"acl_grants,omitempty" diff:"acl_grants"`
	CorsRules                         []*CORSRule                        `bson:"cors_rules,omitempty" dynamodbav:"cors_rules,omitempty" json:"cors_rules,omitempty" diff:"cors_rules"`
	ServerSideEncryptionConfiguration *ServerSideEncryptionConfiguration `bson:"server_side_encryption_configuration,omitempty" dynamodbav:"server_side_encryption_configuration,omitempty" json:"server_side_encryption_configuration,omitempty" diff:"server_side_encryption_configuration"`
	IntelligentTieringConfigurations  []*IntelligentTieringConfiguration `bson:"intelligent_tiering_configurations,omitempty" dynamodbav:"intelligent_tiering_configurations,omitempty" json:"intelligent_tiering_configurations,omitempty" diff:"intelligent_tiering_configurations"`
	InventoryConfigurations           []*InventoryConfiguration          `bson:"inventory_configurations,omitempty" dynamodbav:"inventory_configurations,omitempty" json:"inventory_configurations,omitempty" diff:"inventory_configurations"`
	LifecycleRules                    []*LifecycleRule                   `bson:"lifecycle_rules,omitempty" dynamodbav:"lifecycle_rules,omitempty" json:"lifecycle_rules,omitempty" diff:"lifecycle_rules"`
	Logging                           *LoggingEnabled                    `bson:"logging,omitempty" dynamodbav:"logging,omitempty" json:"logging,omitempty" diff:"logging"`
}
