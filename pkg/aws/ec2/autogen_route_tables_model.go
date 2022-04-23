//AUTOGENERATED CODE DO NOT EDIT
package ec2

type RouteTable struct {
	Associations    []*RouteTableAssociation `bson:"associations,omitempty" ion:"associations" dynamodbav:"associations,omitempty" parquet:"name=associations,type=MAP,convertedtype=LIST" json:"associations,omitempty" diff:"associations"`
	OwnerId         string                   `bson:"owner_id,omitempty" ion:"owner_id" dynamodbav:"owner_id,omitempty" parquet:"name=owner_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"owner_id,omitempty" diff:"owner_id"`
	PropagatingVgws []*PropagatingVgw        `bson:"propagating_vgws,omitempty" ion:"propagating_vgws" dynamodbav:"propagating_vgws,omitempty" parquet:"name=propagating_vgws,type=MAP,convertedtype=LIST" json:"propagating_vgws,omitempty" diff:"propagating_vgws"`
	RouteTableId    string                   `bson:"route_table_id,omitempty" ion:"route_table_id" dynamodbav:"route_table_id,omitempty" parquet:"name=route_table_id,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"route_table_id,omitempty" diff:"route_table_id,identifier"`
	Routes          []*Route                 `bson:"routes,omitempty" ion:"routes" dynamodbav:"routes,omitempty" parquet:"name=routes,type=MAP,convertedtype=LIST" json:"routes,omitempty" diff:"routes"`
	Tags            map[string]string        `bson:"tags,omitempty" ion:"tags" dynamodbav:"tags,omitempty" parquet:"name=tags,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"tags,omitempty" diff:"tags"`
	VpcId           string                   `bson:"vpc_id,omitempty" ion:"vpc_id" dynamodbav:"vpc_id,omitempty" parquet:"name=vpc_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"vpc_id,omitempty" diff:"vpc_id"`
	AccountId       string                   `bson:"account_id,omitempty" ion:"account_id" dynamodbav:"account_id,omitempty" parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"account_id,omitempty" diff:"account_id"`
	Region          string                   `bson:"region,omitempty" ion:"region" dynamodbav:"region,omitempty" parquet:"name=region,type=BYTE_ARRAY,convertedtype=UTF8" json:"region,omitempty" diff:"region"`
	ReportTime      int64                    `bson:"report_time,omitempty" ion:"report_time" dynamodbav:"report_time,omitempty" parquet:"name=report_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"report_time,omitempty" diff:"report_time,immutable"`
	InventoryUUID   string                   `bson:"_id,omitempty" ion:"_id" dynamodbav:"_id,omitempty" parquet:"name=inventory_uuid,type=BYTE_ARRAY,convertedtype=UTF8" json:"_id,omitempty" diff:"-"`
}
