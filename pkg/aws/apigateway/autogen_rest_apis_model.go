//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_resource_file.tmpl
package apigateway

type RestApi struct {
	ApiKeySource              string                 `bson:"api_key_source,omitempty" ion:"api_key_source" dynamodbav:"api_key_source,omitempty" parquet:"name=api_key_source,type=BYTE_ARRAY,convertedtype=UTF8" json:"api_key_source,omitempty" diff:"api_key_source"`
	BinaryMediaTypes          []string               `bson:"binary_media_types,omitempty" ion:"binary_media_types" dynamodbav:"binary_media_types,omitempty" parquet:"name=binary_media_types,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8" json:"binary_media_types,omitempty" diff:"binary_media_types"`
	CreatedDate               int64                  `bson:"created_date,omitempty" ion:"created_date" dynamodbav:"created_date,omitempty" parquet:"name=created_date,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"created_date,omitempty" diff:"created_date"`
	Description               string                 `bson:"description,omitempty" ion:"description" dynamodbav:"description,omitempty" parquet:"name=description,type=BYTE_ARRAY,convertedtype=UTF8" json:"description,omitempty" diff:"description"`
	DisableExecuteApiEndpoint bool                   `bson:"disable_execute_api_endpoint,omitempty" ion:"disable_execute_api_endpoint" dynamodbav:"disable_execute_api_endpoint" parquet:"name=disable_execute_api_endpoint,type=BOOLEAN" json:"disable_execute_api_endpoint,omitempty" diff:"disable_execute_api_endpoint"`
	EndpointConfiguration     *EndpointConfiguration `bson:"endpoint_configuration,omitempty" ion:"endpoint_configuration" dynamodbav:"endpoint_configuration,omitempty" parquet:"name=endpoint_configuration" json:"endpoint_configuration,omitempty" diff:"endpoint_configuration"`
	Id                        string                 `bson:"id,omitempty" ion:"id" dynamodbav:"id,omitempty" parquet:"name=id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"id,omitempty" diff:"id,identifier"`
	MinimumCompressionSize    int32                  `bson:"minimum_compression_size,omitempty" ion:"minimum_compression_size" dynamodbav:"minimum_compression_size,omitempty" parquet:"name=minimum_compression_size,type=INT32" json:"minimum_compression_size,omitempty" diff:"minimum_compression_size"`
	Name                      string                 `bson:"name,omitempty" ion:"name" dynamodbav:"name,omitempty" parquet:"name=name,type=BYTE_ARRAY,convertedtype=UTF8" json:"name,omitempty" diff:"name"`
	Policy                    string                 `bson:"policy,omitempty" ion:"policy" dynamodbav:"policy,omitempty" parquet:"name=policy,type=BYTE_ARRAY,convertedtype=UTF8" json:"policy,omitempty" diff:"policy"`
	Tags                      map[string]string      `bson:"tags,omitempty" ion:"tags" dynamodbav:"tags,omitempty" parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"tags,omitempty" diff:"tags"`
	Version                   string                 `bson:"version,omitempty" ion:"version" dynamodbav:"version,omitempty" parquet:"name=version,type=BYTE_ARRAY,convertedtype=UTF8" json:"version,omitempty" diff:"version"`
	Warnings                  []string               `bson:"warnings,omitempty" ion:"warnings" dynamodbav:"warnings,omitempty" parquet:"name=warnings,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8" json:"warnings,omitempty" diff:"warnings"`
	AccountId                 string                 `bson:"account_id,omitempty" ion:"account_id" dynamodbav:"account_id,omitempty" parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"account_id,omitempty" diff:"account_id"`
	Region                    string                 `bson:"region,omitempty" ion:"region" dynamodbav:"region,omitempty" parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8" json:"region,omitempty" diff:"region"`
	ReportTime                int64                  `bson:"report_time,omitempty" ion:"report_time" dynamodbav:"report_time,omitempty" parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID             string                 `bson:"_id,omitempty" ion:"_id" dynamodbav:"_id,omitempty" parquet:"name=inventory_uuid,type=BYTE_ARRAY,convertedtype=UTF8" json:"_id,omitempty" diff:"-"`
	Stages                    []*Stage               `bson:"stages,omitempty" ion:"stages" dynamodbav:"stages,omitempty" parquet:"name=stages,type=MAP,convertedtype=LIST" json:"stages,omitempty" diff:"stages"`
	Resources                 []*Resource            `bson:"resources,omitempty" ion:"resources" dynamodbav:"resources,omitempty" parquet:"name=resources,type=MAP,convertedtype=LIST" json:"resources,omitempty" diff:"resources"`
}
