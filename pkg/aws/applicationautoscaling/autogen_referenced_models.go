//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_referenced_resource_file.tmpl
package applicationautoscaling

type Alarm struct {
	AlarmARN  string `bson:"alarm_arn,omitempty" ion:"alarm_arn" dynamodbav:"alarm_arn,omitempty" parquet:"name=alarm_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"alarm_arn,omitempty" diff:"alarm_arn"`
	AlarmName string `bson:"alarm_name,omitempty" ion:"alarm_name" dynamodbav:"alarm_name,omitempty" parquet:"name=alarm_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"alarm_name,omitempty" diff:"alarm_name"`
}

type StepScalingPolicyConfiguration struct {
	AdjustmentType         string            `bson:"adjustment_type,omitempty" ion:"adjustment_type" dynamodbav:"adjustment_type,omitempty" parquet:"name=adjustment_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"adjustment_type,omitempty" diff:"adjustment_type"`
	Cooldown               int32             `bson:"cooldown,omitempty" ion:"cooldown" dynamodbav:"cooldown,omitempty" parquet:"name=cooldown,type=INT32" json:"cooldown,omitempty" diff:"cooldown"`
	MetricAggregationType  string            `bson:"metric_aggregation_type,omitempty" ion:"metric_aggregation_type" dynamodbav:"metric_aggregation_type,omitempty" parquet:"name=metric_aggregation_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"metric_aggregation_type,omitempty" diff:"metric_aggregation_type"`
	MinAdjustmentMagnitude int32             `bson:"min_adjustment_magnitude,omitempty" ion:"min_adjustment_magnitude" dynamodbav:"min_adjustment_magnitude,omitempty" parquet:"name=min_adjustment_magnitude,type=INT32" json:"min_adjustment_magnitude,omitempty" diff:"min_adjustment_magnitude"`
	StepAdjustments        []*StepAdjustment `bson:"step_adjustments,omitempty" ion:"step_adjustments" dynamodbav:"step_adjustments,omitempty" parquet:"name=step_adjustments,type=MAP,convertedtype=LIST" json:"step_adjustments,omitempty" diff:"step_adjustments"`
}

type StepAdjustment struct {
	ScalingAdjustment        int32   `bson:"scaling_adjustment,omitempty" ion:"scaling_adjustment" dynamodbav:"scaling_adjustment,omitempty" parquet:"name=scaling_adjustment,type=INT32" json:"scaling_adjustment,omitempty" diff:"scaling_adjustment"`
	MetricIntervalLowerBound float64 `bson:"metric_interval_lower_bound,omitempty" ion:"metric_interval_lower_bound" dynamodbav:"metric_interval_lower_bound,omitempty" parquet:"name=metric_interval_lower_bound,type=DOUBLE" json:"metric_interval_lower_bound,omitempty" diff:"metric_interval_lower_bound"`
	MetricIntervalUpperBound float64 `bson:"metric_interval_upper_bound,omitempty" ion:"metric_interval_upper_bound" dynamodbav:"metric_interval_upper_bound,omitempty" parquet:"name=metric_interval_upper_bound,type=DOUBLE" json:"metric_interval_upper_bound,omitempty" diff:"metric_interval_upper_bound"`
}

type TargetTrackingScalingPolicyConfiguration struct {
	TargetValue                   float64                        `bson:"target_value,omitempty" ion:"target_value" dynamodbav:"target_value,omitempty" parquet:"name=target_value,type=DOUBLE" json:"target_value,omitempty" diff:"target_value"`
	CustomizedMetricSpecification *CustomizedMetricSpecification `bson:"customized_metric_specification,omitempty" ion:"customized_metric_specification" dynamodbav:"customized_metric_specification,omitempty" parquet:"name=customized_metric_specification" json:"customized_metric_specification,omitempty" diff:"customized_metric_specification"`
	DisableScaleIn                bool                           `bson:"disable_scale_in,omitempty" ion:"disable_scale_in" dynamodbav:"disable_scale_in" parquet:"name=disable_scale_in,type=BOOLEAN" json:"disable_scale_in,omitempty" diff:"disable_scale_in"`
	PredefinedMetricSpecification *PredefinedMetricSpecification `bson:"predefined_metric_specification,omitempty" ion:"predefined_metric_specification" dynamodbav:"predefined_metric_specification,omitempty" parquet:"name=predefined_metric_specification" json:"predefined_metric_specification,omitempty" diff:"predefined_metric_specification"`
	ScaleInCooldown               int32                          `bson:"scale_in_cooldown,omitempty" ion:"scale_in_cooldown" dynamodbav:"scale_in_cooldown,omitempty" parquet:"name=scale_in_cooldown,type=INT32" json:"scale_in_cooldown,omitempty" diff:"scale_in_cooldown"`
	ScaleOutCooldown              int32                          `bson:"scale_out_cooldown,omitempty" ion:"scale_out_cooldown" dynamodbav:"scale_out_cooldown,omitempty" parquet:"name=scale_out_cooldown,type=INT32" json:"scale_out_cooldown,omitempty" diff:"scale_out_cooldown"`
}

type CustomizedMetricSpecification struct {
	MetricName string             `bson:"metric_name,omitempty" ion:"metric_name" dynamodbav:"metric_name,omitempty" parquet:"name=metric_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"metric_name,omitempty" diff:"metric_name"`
	Namespace  string             `bson:"namespace,omitempty" ion:"namespace" dynamodbav:"namespace,omitempty" parquet:"name=namespace,type=BYTE_ARRAY,convertedtype=UTF8" json:"namespace,omitempty" diff:"namespace"`
	Statistic  string             `bson:"statistic,omitempty" ion:"statistic" dynamodbav:"statistic,omitempty" parquet:"name=statistic,type=BYTE_ARRAY,convertedtype=UTF8" json:"statistic,omitempty" diff:"statistic"`
	Dimensions []*MetricDimension `bson:"dimensions,omitempty" ion:"dimensions" dynamodbav:"dimensions,omitempty" parquet:"name=dimensions,type=MAP,convertedtype=LIST" json:"dimensions,omitempty" diff:"dimensions"`
	Unit       string             `bson:"unit,omitempty" ion:"unit" dynamodbav:"unit,omitempty" parquet:"name=unit,type=BYTE_ARRAY,convertedtype=UTF8" json:"unit,omitempty" diff:"unit"`
}

type MetricDimension struct {
	Name  string `bson:"name,omitempty" ion:"name" dynamodbav:"name,omitempty" parquet:"name=name,type=BYTE_ARRAY,convertedtype=UTF8" json:"name,omitempty" diff:"name"`
	Value string `bson:"value,omitempty" ion:"value" dynamodbav:"value,omitempty" parquet:"name=value,type=BYTE_ARRAY,convertedtype=UTF8" json:"value,omitempty" diff:"value"`
}

type PredefinedMetricSpecification struct {
	PredefinedMetricType string `bson:"predefined_metric_type,omitempty" ion:"predefined_metric_type" dynamodbav:"predefined_metric_type,omitempty" parquet:"name=predefined_metric_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"predefined_metric_type,omitempty" diff:"predefined_metric_type"`
	ResourceLabel        string `bson:"resource_label,omitempty" ion:"resource_label" dynamodbav:"resource_label,omitempty" parquet:"name=resource_label,type=BYTE_ARRAY,convertedtype=UTF8" json:"resource_label,omitempty" diff:"resource_label"`
}
