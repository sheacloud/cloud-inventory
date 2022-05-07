//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_resource_file.tmpl
package athena

type DataCatalog struct {
	Name          string            `bson:"name,omitempty" ion:"name" dynamodbav:"name,omitempty" parquet:"name=name,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"name,omitempty" diff:"name,identifier"`
	Type          string            `bson:"type,omitempty" ion:"type" dynamodbav:"type,omitempty" parquet:"name=type,type=BYTE_ARRAY,convertedtype=UTF8" json:"type,omitempty" diff:"type"`
	Description   string            `bson:"description,omitempty" ion:"description" dynamodbav:"description,omitempty" parquet:"name=description,type=BYTE_ARRAY,convertedtype=UTF8" json:"description,omitempty" diff:"description"`
	Parameters    map[string]string `bson:"parameters,omitempty" ion:"parameters" dynamodbav:"parameters,omitempty" parquet:"name=parameters,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"parameters,omitempty" diff:"parameters"`
	AccountId     string            `bson:"account_id,omitempty" ion:"account_id" dynamodbav:"account_id,omitempty" parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"account_id,omitempty" diff:"account_id"`
	Region        string            `bson:"region,omitempty" ion:"region" dynamodbav:"region,omitempty" parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8" json:"region,omitempty" diff:"region"`
	ReportTime    int64             `bson:"report_time,omitempty" ion:"report_time" dynamodbav:"report_time,omitempty" parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID string            `bson:"_id,omitempty" ion:"_id" dynamodbav:"_id,omitempty" parquet:"name=inventory_uuid,type=BYTE_ARRAY,convertedtype=UTF8" json:"_id,omitempty" diff:"-"`
}
