//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_referenced_resource_file.tmpl
package elasticloadbalancing

type BackendServerDescription struct {
	InstancePort int32    `bson:"instance_port,omitempty" ion:"instance_port" dynamodbav:"instance_port,omitempty" parquet:"name=instance_port,type=INT32" json:"instance_port,omitempty" diff:"instance_port"`
	PolicyNames  []string `bson:"policy_names,omitempty" ion:"policy_names" dynamodbav:"policy_names,omitempty" parquet:"name=policy_names,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8" json:"policy_names,omitempty" diff:"policy_names"`
}

type HealthCheck struct {
	HealthyThreshold   int32  `bson:"healthy_threshold,omitempty" ion:"healthy_threshold" dynamodbav:"healthy_threshold,omitempty" parquet:"name=healthy_threshold,type=INT32" json:"healthy_threshold,omitempty" diff:"healthy_threshold"`
	Interval           int32  `bson:"interval,omitempty" ion:"interval" dynamodbav:"interval,omitempty" parquet:"name=interval,type=INT32" json:"interval,omitempty" diff:"interval"`
	Target             string `bson:"target,omitempty" ion:"target" dynamodbav:"target,omitempty" parquet:"name=target,type=BYTE_ARRAY,convertedtype=UTF8" json:"target,omitempty" diff:"target"`
	Timeout            int32  `bson:"timeout,omitempty" ion:"timeout" dynamodbav:"timeout,omitempty" parquet:"name=timeout,type=INT32" json:"timeout,omitempty" diff:"timeout"`
	UnhealthyThreshold int32  `bson:"unhealthy_threshold,omitempty" ion:"unhealthy_threshold" dynamodbav:"unhealthy_threshold,omitempty" parquet:"name=unhealthy_threshold,type=INT32" json:"unhealthy_threshold,omitempty" diff:"unhealthy_threshold"`
}

type Instance struct {
	InstanceId string `bson:"instance_id,omitempty" ion:"instance_id" dynamodbav:"instance_id,omitempty" parquet:"name=instance_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"instance_id,omitempty" diff:"instance_id"`
}

type ListenerDescription struct {
	Listener    *Listener `bson:"listener,omitempty" ion:"listener" dynamodbav:"listener,omitempty" parquet:"name=listener" json:"listener,omitempty" diff:"listener"`
	PolicyNames []string  `bson:"policy_names,omitempty" ion:"policy_names" dynamodbav:"policy_names,omitempty" parquet:"name=policy_names,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8" json:"policy_names,omitempty" diff:"policy_names"`
}

type Listener struct {
	InstancePort     int32  `bson:"instance_port,omitempty" ion:"instance_port" dynamodbav:"instance_port,omitempty" parquet:"name=instance_port,type=INT32" json:"instance_port,omitempty" diff:"instance_port"`
	LoadBalancerPort int32  `bson:"load_balancer_port,omitempty" ion:"load_balancer_port" dynamodbav:"load_balancer_port,omitempty" parquet:"name=load_balancer_port,type=INT32" json:"load_balancer_port,omitempty" diff:"load_balancer_port"`
	Protocol         string `bson:"protocol,omitempty" ion:"protocol" dynamodbav:"protocol,omitempty" parquet:"name=protocol,type=BYTE_ARRAY,convertedtype=UTF8" json:"protocol,omitempty" diff:"protocol"`
	InstanceProtocol string `bson:"instance_protocol,omitempty" ion:"instance_protocol" dynamodbav:"instance_protocol,omitempty" parquet:"name=instance_protocol,type=BYTE_ARRAY,convertedtype=UTF8" json:"instance_protocol,omitempty" diff:"instance_protocol"`
	SSLCertificateId string `bson:"ssl_certificate_id,omitempty" ion:"ssl_certificate_id" dynamodbav:"ssl_certificate_id,omitempty" parquet:"name=ssl_certificate_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"ssl_certificate_id,omitempty" diff:"ssl_certificate_id"`
}

type Policies struct {
	AppCookieStickinessPolicies []*AppCookieStickinessPolicy `bson:"app_cookie_stickiness_policies,omitempty" ion:"app_cookie_stickiness_policies" dynamodbav:"app_cookie_stickiness_policies,omitempty" parquet:"name=app_cookie_stickiness_policies,type=MAP,convertedtype=LIST" json:"app_cookie_stickiness_policies,omitempty" diff:"app_cookie_stickiness_policies"`
	LBCookieStickinessPolicies  []*LBCookieStickinessPolicy  `bson:"lb_cookie_stickiness_policies,omitempty" ion:"lb_cookie_stickiness_policies" dynamodbav:"lb_cookie_stickiness_policies,omitempty" parquet:"name=lb_cookie_stickiness_policies,type=MAP,convertedtype=LIST" json:"lb_cookie_stickiness_policies,omitempty" diff:"lb_cookie_stickiness_policies"`
	OtherPolicies               []string                     `bson:"other_policies,omitempty" ion:"other_policies" dynamodbav:"other_policies,omitempty" parquet:"name=other_policies,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8" json:"other_policies,omitempty" diff:"other_policies"`
}

type AppCookieStickinessPolicy struct {
	CookieName string `bson:"cookie_name,omitempty" ion:"cookie_name" dynamodbav:"cookie_name,omitempty" parquet:"name=cookie_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"cookie_name,omitempty" diff:"cookie_name"`
	PolicyName string `bson:"policy_name,omitempty" ion:"policy_name" dynamodbav:"policy_name,omitempty" parquet:"name=policy_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"policy_name,omitempty" diff:"policy_name"`
}

type LBCookieStickinessPolicy struct {
	CookieExpirationPeriod int64  `bson:"cookie_expiration_period,omitempty" ion:"cookie_expiration_period" dynamodbav:"cookie_expiration_period,omitempty" parquet:"name=cookie_expiration_period,type=INT64" json:"cookie_expiration_period,omitempty" diff:"cookie_expiration_period"`
	PolicyName             string `bson:"policy_name,omitempty" ion:"policy_name" dynamodbav:"policy_name,omitempty" parquet:"name=policy_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"policy_name,omitempty" diff:"policy_name"`
}

type SourceSecurityGroup struct {
	GroupName  string `bson:"group_name,omitempty" ion:"group_name" dynamodbav:"group_name,omitempty" parquet:"name=group_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"group_name,omitempty" diff:"group_name"`
	OwnerAlias string `bson:"owner_alias,omitempty" ion:"owner_alias" dynamodbav:"owner_alias,omitempty" parquet:"name=owner_alias,type=BYTE_ARRAY,convertedtype=UTF8" json:"owner_alias,omitempty" diff:"owner_alias"`
}
