//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_referenced_resource_file.tmpl
package storagegateway

type NetworkInterface struct {
	Ipv4Address string `bson:"ipv4_address,omitempty" ion:"ipv4_address" dynamodbav:"ipv4_address,omitempty" parquet:"name=ipv4_address,type=BYTE_ARRAY,convertedtype=UTF8" json:"ipv4_address,omitempty" diff:"ipv4_address"`
	Ipv6Address string `bson:"ipv6_address,omitempty" ion:"ipv6_address" dynamodbav:"ipv6_address,omitempty" parquet:"name=ipv6_address,type=BYTE_ARRAY,convertedtype=UTF8" json:"ipv6_address,omitempty" diff:"ipv6_address"`
	MacAddress  string `bson:"mac_address,omitempty" ion:"mac_address" dynamodbav:"mac_address,omitempty" parquet:"name=mac_address,type=BYTE_ARRAY,convertedtype=UTF8" json:"mac_address,omitempty" diff:"mac_address"`
}

type Tag struct {
	Key   string `bson:"key,omitempty" ion:"key" dynamodbav:"key,omitempty" parquet:"name=key,type=BYTE_ARRAY,convertedtype=UTF8" json:"key,omitempty" diff:"key"`
	Value string `bson:"value,omitempty" ion:"value" dynamodbav:"value,omitempty" parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8" json:"value,omitempty" diff:"value"`
}
