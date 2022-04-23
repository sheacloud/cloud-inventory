//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_referenced_resource_file.tmpl
package elasticloadbalancingv2

type AvailabilityZone struct {
	LoadBalancerAddresses []*LoadBalancerAddress `bson:"load_balancer_addresses,omitempty" ion:"load_balancer_addresses" dynamodbav:"load_balancer_addresses,omitempty" parquet:"name=load_balancer_addresses,type=MAP,convertedtype=LIST" json:"load_balancer_addresses,omitempty" diff:"load_balancer_addresses"`
	OutpostId             string                 `bson:"outpost_id,omitempty" ion:"outpost_id" dynamodbav:"outpost_id,omitempty" parquet:"name=outpost_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"outpost_id,omitempty" diff:"outpost_id"`
	SubnetId              string                 `bson:"subnet_id,omitempty" ion:"subnet_id" dynamodbav:"subnet_id,omitempty" parquet:"name=subnet_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"subnet_id,omitempty" diff:"subnet_id"`
	ZoneName              string                 `bson:"zone_name,omitempty" ion:"zone_name" dynamodbav:"zone_name,omitempty" parquet:"name=zone_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"zone_name,omitempty" diff:"zone_name"`
}

type LoadBalancerAddress struct {
	AllocationId       string `bson:"allocation_id,omitempty" ion:"allocation_id" dynamodbav:"allocation_id,omitempty" parquet:"name=allocation_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"allocation_id,omitempty" diff:"allocation_id"`
	IPv6Address        string `bson:"i_pv6_address,omitempty" ion:"i_pv6_address" dynamodbav:"i_pv6_address,omitempty" parquet:"name=i_pv6_address,type=BYTE_ARRAY,convertedtype=UTF8" json:"i_pv6_address,omitempty" diff:"i_pv6_address"`
	IpAddress          string `bson:"ip_address,omitempty" ion:"ip_address" dynamodbav:"ip_address,omitempty" parquet:"name=ip_address,type=BYTE_ARRAY,convertedtype=UTF8" json:"ip_address,omitempty" diff:"ip_address"`
	PrivateIPv4Address string `bson:"private_i_pv4_address,omitempty" ion:"private_i_pv4_address" dynamodbav:"private_i_pv4_address,omitempty" parquet:"name=private_i_pv4_address,type=BYTE_ARRAY,convertedtype=UTF8" json:"private_i_pv4_address,omitempty" diff:"private_i_pv4_address"`
}

type LoadBalancerState struct {
	Code   string `bson:"code,omitempty" ion:"code" dynamodbav:"code,omitempty" parquet:"name=code,type=BYTE_ARRAY,convertedtype=UTF8" json:"code,omitempty" diff:"code"`
	Reason string `bson:"reason,omitempty" ion:"reason" dynamodbav:"reason,omitempty" parquet:"name=reason,type=BYTE_ARRAY,convertedtype=UTF8" json:"reason,omitempty" diff:"reason"`
}

type Listener struct {
	AlpnPolicy      []string       `bson:"alpn_policy,omitempty" ion:"alpn_policy" dynamodbav:"alpn_policy,omitempty" parquet:"name=alpn_policy,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8" json:"alpn_policy,omitempty" diff:"alpn_policy"`
	Certificates    []*Certificate `bson:"certificates,omitempty" ion:"certificates" dynamodbav:"certificates,omitempty" parquet:"name=certificates,type=MAP,convertedtype=LIST" json:"certificates,omitempty" diff:"certificates"`
	DefaultActions  []*Action      `bson:"default_actions,omitempty" ion:"default_actions" dynamodbav:"default_actions,omitempty" parquet:"name=default_actions,type=MAP,convertedtype=LIST" json:"default_actions,omitempty" diff:"default_actions"`
	ListenerArn     string         `bson:"listener_arn,omitempty" ion:"listener_arn" dynamodbav:"listener_arn,omitempty" parquet:"name=listener_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"listener_arn,omitempty" diff:"listener_arn"`
	LoadBalancerArn string         `bson:"load_balancer_arn,omitempty" ion:"load_balancer_arn" dynamodbav:"load_balancer_arn,omitempty" parquet:"name=load_balancer_arn,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"load_balancer_arn,omitempty" diff:"load_balancer_arn,identifier"`
	Port            int32          `bson:"port,omitempty" ion:"port" dynamodbav:"port,omitempty" parquet:"name=port,type=INT32" json:"port,omitempty" diff:"port"`
	Protocol        string         `bson:"protocol,omitempty" ion:"protocol" dynamodbav:"protocol,omitempty" parquet:"name=protocol,type=BYTE_ARRAY,convertedtype=UTF8" json:"protocol,omitempty" diff:"protocol"`
	SslPolicy       string         `bson:"ssl_policy,omitempty" ion:"ssl_policy" dynamodbav:"ssl_policy,omitempty" parquet:"name=ssl_policy,type=BYTE_ARRAY,convertedtype=UTF8" json:"ssl_policy,omitempty" diff:"ssl_policy"`
}

type Certificate struct {
	CertificateArn string `bson:"certificate_arn,omitempty" ion:"certificate_arn" dynamodbav:"certificate_arn,omitempty" parquet:"name=certificate_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"certificate_arn,omitempty" diff:"certificate_arn"`
	IsDefault      bool   `bson:"is_default,omitempty" ion:"is_default" dynamodbav:"is_default" parquet:"name=is_default,type=BOOLEAN" json:"is_default,omitempty" diff:"is_default"`
}

type Action struct {
	Type                      string                           `bson:"type,omitempty" ion:"type" dynamodbav:"type,omitempty" parquet:"name=type,type=BYTE_ARRAY,convertedtype=UTF8" json:"type,omitempty" diff:"type"`
	AuthenticateCognitoConfig *AuthenticateCognitoActionConfig `bson:"authenticate_cognito_config,omitempty" ion:"authenticate_cognito_config" dynamodbav:"authenticate_cognito_config,omitempty" parquet:"name=authenticate_cognito_config" json:"authenticate_cognito_config,omitempty" diff:"authenticate_cognito_config"`
	AuthenticateOidcConfig    *AuthenticateOidcActionConfig    `bson:"authenticate_oidc_config,omitempty" ion:"authenticate_oidc_config" dynamodbav:"authenticate_oidc_config,omitempty" parquet:"name=authenticate_oidc_config" json:"authenticate_oidc_config,omitempty" diff:"authenticate_oidc_config"`
	FixedResponseConfig       *FixedResponseActionConfig       `bson:"fixed_response_config,omitempty" ion:"fixed_response_config" dynamodbav:"fixed_response_config,omitempty" parquet:"name=fixed_response_config" json:"fixed_response_config,omitempty" diff:"fixed_response_config"`
	ForwardConfig             *ForwardActionConfig             `bson:"forward_config,omitempty" ion:"forward_config" dynamodbav:"forward_config,omitempty" parquet:"name=forward_config" json:"forward_config,omitempty" diff:"forward_config"`
	Order                     int32                            `bson:"order,omitempty" ion:"order" dynamodbav:"order,omitempty" parquet:"name=order,type=INT32" json:"order,omitempty" diff:"order"`
	RedirectConfig            *RedirectActionConfig            `bson:"redirect_config,omitempty" ion:"redirect_config" dynamodbav:"redirect_config,omitempty" parquet:"name=redirect_config" json:"redirect_config,omitempty" diff:"redirect_config"`
	TargetGroupArn            string                           `bson:"target_group_arn,omitempty" ion:"target_group_arn" dynamodbav:"target_group_arn,omitempty" parquet:"name=target_group_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"target_group_arn,omitempty" diff:"target_group_arn"`
}

type AuthenticateCognitoActionConfig struct {
	UserPoolArn                      string            `bson:"user_pool_arn,omitempty" ion:"user_pool_arn" dynamodbav:"user_pool_arn,omitempty" parquet:"name=user_pool_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"user_pool_arn,omitempty" diff:"user_pool_arn"`
	UserPoolClientId                 string            `bson:"user_pool_client_id,omitempty" ion:"user_pool_client_id" dynamodbav:"user_pool_client_id,omitempty" parquet:"name=user_pool_client_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"user_pool_client_id,omitempty" diff:"user_pool_client_id"`
	UserPoolDomain                   string            `bson:"user_pool_domain,omitempty" ion:"user_pool_domain" dynamodbav:"user_pool_domain,omitempty" parquet:"name=user_pool_domain,type=BYTE_ARRAY,convertedtype=UTF8" json:"user_pool_domain,omitempty" diff:"user_pool_domain"`
	AuthenticationRequestExtraParams map[string]string `bson:"authentication_request_extra_params,omitempty" ion:"authentication_request_extra_params" dynamodbav:"authentication_request_extra_params,omitempty" parquet:"name=authentication_request_extra_params,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"authentication_request_extra_params,omitempty" diff:"authentication_request_extra_params"`
	OnUnauthenticatedRequest         string            `bson:"on_unauthenticated_request,omitempty" ion:"on_unauthenticated_request" dynamodbav:"on_unauthenticated_request,omitempty" parquet:"name=on_unauthenticated_request,type=BYTE_ARRAY,convertedtype=UTF8" json:"on_unauthenticated_request,omitempty" diff:"on_unauthenticated_request"`
	Scope                            string            `bson:"scope,omitempty" ion:"scope" dynamodbav:"scope,omitempty" parquet:"name=scope,type=BYTE_ARRAY,convertedtype=UTF8" json:"scope,omitempty" diff:"scope"`
	SessionCookieName                string            `bson:"session_cookie_name,omitempty" ion:"session_cookie_name" dynamodbav:"session_cookie_name,omitempty" parquet:"name=session_cookie_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"session_cookie_name,omitempty" diff:"session_cookie_name"`
	SessionTimeout                   int64             `bson:"session_timeout,omitempty" ion:"session_timeout" dynamodbav:"session_timeout,omitempty" parquet:"name=session_timeout,type=INT64" json:"session_timeout,omitempty" diff:"session_timeout"`
}

type AuthenticateOidcActionConfig struct {
	AuthorizationEndpoint            string            `bson:"authorization_endpoint,omitempty" ion:"authorization_endpoint" dynamodbav:"authorization_endpoint,omitempty" parquet:"name=authorization_endpoint,type=BYTE_ARRAY,convertedtype=UTF8" json:"authorization_endpoint,omitempty" diff:"authorization_endpoint"`
	ClientId                         string            `bson:"client_id,omitempty" ion:"client_id" dynamodbav:"client_id,omitempty" parquet:"name=client_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"client_id,omitempty" diff:"client_id"`
	Issuer                           string            `bson:"issuer,omitempty" ion:"issuer" dynamodbav:"issuer,omitempty" parquet:"name=issuer,type=BYTE_ARRAY,convertedtype=UTF8" json:"issuer,omitempty" diff:"issuer"`
	TokenEndpoint                    string            `bson:"token_endpoint,omitempty" ion:"token_endpoint" dynamodbav:"token_endpoint,omitempty" parquet:"name=token_endpoint,type=BYTE_ARRAY,convertedtype=UTF8" json:"token_endpoint,omitempty" diff:"token_endpoint"`
	UserInfoEndpoint                 string            `bson:"user_info_endpoint,omitempty" ion:"user_info_endpoint" dynamodbav:"user_info_endpoint,omitempty" parquet:"name=user_info_endpoint,type=BYTE_ARRAY,convertedtype=UTF8" json:"user_info_endpoint,omitempty" diff:"user_info_endpoint"`
	AuthenticationRequestExtraParams map[string]string `bson:"authentication_request_extra_params,omitempty" ion:"authentication_request_extra_params" dynamodbav:"authentication_request_extra_params,omitempty" parquet:"name=authentication_request_extra_params,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"authentication_request_extra_params,omitempty" diff:"authentication_request_extra_params"`
	ClientSecret                     string            `bson:"client_secret,omitempty" ion:"client_secret" dynamodbav:"client_secret,omitempty" parquet:"name=client_secret,type=BYTE_ARRAY,convertedtype=UTF8" json:"client_secret,omitempty" diff:"client_secret"`
	OnUnauthenticatedRequest         string            `bson:"on_unauthenticated_request,omitempty" ion:"on_unauthenticated_request" dynamodbav:"on_unauthenticated_request,omitempty" parquet:"name=on_unauthenticated_request,type=BYTE_ARRAY,convertedtype=UTF8" json:"on_unauthenticated_request,omitempty" diff:"on_unauthenticated_request"`
	Scope                            string            `bson:"scope,omitempty" ion:"scope" dynamodbav:"scope,omitempty" parquet:"name=scope,type=BYTE_ARRAY,convertedtype=UTF8" json:"scope,omitempty" diff:"scope"`
	SessionCookieName                string            `bson:"session_cookie_name,omitempty" ion:"session_cookie_name" dynamodbav:"session_cookie_name,omitempty" parquet:"name=session_cookie_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"session_cookie_name,omitempty" diff:"session_cookie_name"`
	SessionTimeout                   int64             `bson:"session_timeout,omitempty" ion:"session_timeout" dynamodbav:"session_timeout,omitempty" parquet:"name=session_timeout,type=INT64" json:"session_timeout,omitempty" diff:"session_timeout"`
	UseExistingClientSecret          bool              `bson:"use_existing_client_secret,omitempty" ion:"use_existing_client_secret" dynamodbav:"use_existing_client_secret" parquet:"name=use_existing_client_secret,type=BOOLEAN" json:"use_existing_client_secret,omitempty" diff:"use_existing_client_secret"`
}

type FixedResponseActionConfig struct {
	StatusCode  string `bson:"status_code,omitempty" ion:"status_code" dynamodbav:"status_code,omitempty" parquet:"name=status_code,type=BYTE_ARRAY,convertedtype=UTF8" json:"status_code,omitempty" diff:"status_code"`
	ContentType string `bson:"content_type,omitempty" ion:"content_type" dynamodbav:"content_type,omitempty" parquet:"name=content_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"content_type,omitempty" diff:"content_type"`
	MessageBody string `bson:"message_body,omitempty" ion:"message_body" dynamodbav:"message_body,omitempty" parquet:"name=message_body,type=BYTE_ARRAY,convertedtype=UTF8" json:"message_body,omitempty" diff:"message_body"`
}

type ForwardActionConfig struct {
	TargetGroupStickinessConfig *TargetGroupStickinessConfig `bson:"target_group_stickiness_config,omitempty" ion:"target_group_stickiness_config" dynamodbav:"target_group_stickiness_config,omitempty" parquet:"name=target_group_stickiness_config" json:"target_group_stickiness_config,omitempty" diff:"target_group_stickiness_config"`
	TargetGroups                []*TargetGroupTuple          `bson:"target_groups,omitempty" ion:"target_groups" dynamodbav:"target_groups,omitempty" parquet:"name=target_groups,type=MAP,convertedtype=LIST" json:"target_groups,omitempty" diff:"target_groups"`
}

type TargetGroupStickinessConfig struct {
	DurationSeconds int32 `bson:"duration_seconds,omitempty" ion:"duration_seconds" dynamodbav:"duration_seconds,omitempty" parquet:"name=duration_seconds,type=INT32" json:"duration_seconds,omitempty" diff:"duration_seconds"`
	Enabled         bool  `bson:"enabled,omitempty" ion:"enabled" dynamodbav:"enabled" parquet:"name=enabled,type=BOOLEAN" json:"enabled,omitempty" diff:"enabled"`
}

type TargetGroupTuple struct {
	TargetGroupArn string `bson:"target_group_arn,omitempty" ion:"target_group_arn" dynamodbav:"target_group_arn,omitempty" parquet:"name=target_group_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"target_group_arn,omitempty" diff:"target_group_arn"`
	Weight         int32  `bson:"weight,omitempty" ion:"weight" dynamodbav:"weight,omitempty" parquet:"name=weight,type=INT32" json:"weight,omitempty" diff:"weight"`
}

type RedirectActionConfig struct {
	StatusCode string `bson:"status_code,omitempty" ion:"status_code" dynamodbav:"status_code,omitempty" parquet:"name=status_code,type=BYTE_ARRAY,convertedtype=UTF8" json:"status_code,omitempty" diff:"status_code"`
	Host       string `bson:"host,omitempty" ion:"host" dynamodbav:"host,omitempty" parquet:"name=host,type=BYTE_ARRAY,convertedtype=UTF8" json:"host,omitempty" diff:"host"`
	Path       string `bson:"path,omitempty" ion:"path" dynamodbav:"path,omitempty" parquet:"name=path,type=BYTE_ARRAY,convertedtype=UTF8" json:"path,omitempty" diff:"path"`
	Port       string `bson:"port,omitempty" ion:"port" dynamodbav:"port,omitempty" parquet:"name=port,type=BYTE_ARRAY,convertedtype=UTF8" json:"port,omitempty" diff:"port"`
	Protocol   string `bson:"protocol,omitempty" ion:"protocol" dynamodbav:"protocol,omitempty" parquet:"name=protocol,type=BYTE_ARRAY,convertedtype=UTF8" json:"protocol,omitempty" diff:"protocol"`
	Query      string `bson:"query,omitempty" ion:"query" dynamodbav:"query,omitempty" parquet:"name=query,type=BYTE_ARRAY,convertedtype=UTF8" json:"query,omitempty" diff:"query"`
}

type Matcher struct {
	GrpcCode string `bson:"grpc_code,omitempty" ion:"grpc_code" dynamodbav:"grpc_code,omitempty" parquet:"name=grpc_code,type=BYTE_ARRAY,convertedtype=UTF8" json:"grpc_code,omitempty" diff:"grpc_code"`
	HttpCode string `bson:"http_code,omitempty" ion:"http_code" dynamodbav:"http_code,omitempty" parquet:"name=http_code,type=BYTE_ARRAY,convertedtype=UTF8" json:"http_code,omitempty" diff:"http_code"`
}

type TargetHealthDescription struct {
	HealthCheckPort string             `bson:"health_check_port,omitempty" ion:"health_check_port" dynamodbav:"health_check_port,omitempty" parquet:"name=health_check_port,type=BYTE_ARRAY,convertedtype=UTF8" json:"health_check_port,omitempty" diff:"health_check_port"`
	Target          *TargetDescription `bson:"target,omitempty" ion:"target" dynamodbav:"target,omitempty" parquet:"name=target" json:"target,omitempty" diff:"target"`
	TargetHealth    *TargetHealth      `bson:"target_health,omitempty" ion:"target_health" dynamodbav:"target_health,omitempty" parquet:"name=target_health" json:"target_health,omitempty" diff:"target_health"`
}

type TargetDescription struct {
	Id               string `bson:"id,omitempty" ion:"id" dynamodbav:"id,omitempty" parquet:"name=id,type=BYTE_ARRAY,convertedtype=UTF8" json:"id,omitempty" diff:"id"`
	AvailabilityZone string `bson:"availability_zone,omitempty" ion:"availability_zone" dynamodbav:"availability_zone,omitempty" parquet:"name=availability_zone,type=BYTE_ARRAY,convertedtype=UTF8" json:"availability_zone,omitempty" diff:"availability_zone"`
	Port             int32  `bson:"port,omitempty" ion:"port" dynamodbav:"port,omitempty" parquet:"name=port,type=INT32" json:"port,omitempty" diff:"port"`
}

type TargetHealth struct {
	Description string `bson:"description,omitempty" ion:"description" dynamodbav:"description,omitempty" parquet:"name=description,type=BYTE_ARRAY,convertedtype=UTF8" json:"description,omitempty" diff:"description"`
	Reason      string `bson:"reason,omitempty" ion:"reason" dynamodbav:"reason,omitempty" parquet:"name=reason,type=BYTE_ARRAY,convertedtype=UTF8" json:"reason,omitempty" diff:"reason"`
	State       string `bson:"state,omitempty" ion:"state" dynamodbav:"state,omitempty" parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8" json:"state,omitempty" diff:"state"`
}
