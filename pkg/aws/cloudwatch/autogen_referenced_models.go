//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_referenced_resource_file.tmpl
package cloudwatch

type Dimension struct {
	Name  string `bson:"name,omitempty" ion:"name" dynamodbav:"name,omitempty" parquet:"name=name,type=BYTE_ARRAY,convertedtype=UTF8" json:"name,omitempty" diff:"name"`
	Value string `bson:"value,omitempty" ion:"value" dynamodbav:"value,omitempty" parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8" json:"value,omitempty" diff:"value"`
}

type MetricDataQuery struct {
	Id         string      `bson:"id,omitempty" ion:"id" dynamodbav:"id,omitempty" parquet:"name=id,type=BYTE_ARRAY,convertedtype=UTF8" json:"id,omitempty" diff:"id"`
	AccountId  string      `bson:"account_id,omitempty" ion:"account_id" dynamodbav:"account_id,omitempty" parquet:"name=account_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"account_id,omitempty" diff:"account_id"`
	Expression string      `bson:"expression,omitempty" ion:"expression" dynamodbav:"expression,omitempty" parquet:"name=expression,type=BYTE_ARRAY,convertedtype=UTF8" json:"expression,omitempty" diff:"expression"`
	Label      string      `bson:"label,omitempty" ion:"label" dynamodbav:"label,omitempty" parquet:"name=label,type=BYTE_ARRAY,convertedtype=UTF8" json:"label,omitempty" diff:"label"`
	MetricStat *MetricStat `bson:"metric_stat,omitempty" ion:"metric_stat" dynamodbav:"metric_stat,omitempty" parquet:"name=metric_stat" json:"metric_stat,omitempty" diff:"metric_stat"`
	Period     int32       `bson:"period,omitempty" ion:"period" dynamodbav:"period,omitempty" parquet:"name=period,type=INT32" json:"period,omitempty" diff:"period"`
	ReturnData bool        `bson:"return_data,omitempty" ion:"return_data" dynamodbav:"return_data" parquet:"name=return_data,type=BOOLEAN" json:"return_data,omitempty" diff:"return_data"`
}

type MetricStat struct {
	Metric *Metric `bson:"metric,omitempty" ion:"metric" dynamodbav:"metric,omitempty" parquet:"name=metric" json:"metric,omitempty" diff:"metric"`
	Period int32   `bson:"period,omitempty" ion:"period" dynamodbav:"period,omitempty" parquet:"name=period,type=INT32" json:"period,omitempty" diff:"period"`
	Stat   string  `bson:"stat,omitempty" ion:"stat" dynamodbav:"stat,omitempty" parquet:"name=stat,type=BYTE_ARRAY,convertedtype=UTF8" json:"stat,omitempty" diff:"stat"`
	Unit   string  `bson:"unit,omitempty" ion:"unit" dynamodbav:"unit,omitempty" parquet:"name=unit,type=BYTE_ARRAY,convertedtype=UTF8" json:"unit,omitempty" diff:"unit"`
}

type Metric struct {
	Dimensions []*Dimension `bson:"dimensions,omitempty" ion:"dimensions" dynamodbav:"dimensions,omitempty" parquet:"name=dimensions,type=MAP,convertedtype=LIST" json:"dimensions,omitempty" diff:"dimensions"`
	MetricName string       `bson:"metric_name,omitempty" ion:"metric_name" dynamodbav:"metric_name,omitempty" parquet:"name=metric_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"metric_name,omitempty" diff:"metric_name"`
	Namespace  string       `bson:"namespace,omitempty" ion:"namespace" dynamodbav:"namespace,omitempty" parquet:"name=namespace,type=BYTE_ARRAY,convertedtype=UTF8" json:"namespace,omitempty" diff:"namespace"`
}