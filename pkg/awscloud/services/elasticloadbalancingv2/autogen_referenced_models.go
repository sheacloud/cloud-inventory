//AUTOGENERATED CODE DO NOT EDIT
package elasticloadbalancingv2

type AvailabilityZone struct {
	LoadBalancerAddresses []*LoadBalancerAddress `parquet:"name=load_balancer_addresses,type=MAP,convertedtype=LIST" json:"load_balancer_addresses" diff:"load_balancer_addresses"`
	OutpostId             string                 `parquet:"name=outpost_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"outpost_id" diff:"outpost_id"`
	SubnetId              string                 `parquet:"name=subnet_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"subnet_id" diff:"subnet_id"`
	ZoneName              string                 `parquet:"name=zone_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"zone_name" diff:"zone_name"`
}

type LoadBalancerAddress struct {
	AllocationId       string `parquet:"name=allocation_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"allocation_id" diff:"allocation_id"`
	IPv6Address        string `parquet:"name=i_pv6_address,type=BYTE_ARRAY,convertedtype=UTF8" json:"i_pv6_address" diff:"i_pv6_address"`
	IpAddress          string `parquet:"name=ip_address,type=BYTE_ARRAY,convertedtype=UTF8" json:"ip_address" diff:"ip_address"`
	PrivateIPv4Address string `parquet:"name=private_i_pv4_address,type=BYTE_ARRAY,convertedtype=UTF8" json:"private_i_pv4_address" diff:"private_i_pv4_address"`
}

type LoadBalancerState struct {
	Code   string `parquet:"name=code,type=BYTE_ARRAY,convertedtype=UTF8" json:"code" diff:"code"`
	Reason string `parquet:"name=reason,type=BYTE_ARRAY,convertedtype=UTF8" json:"reason" diff:"reason"`
}

type Listener struct {
	AlpnPolicy      []string       `parquet:"name=alpn_policy,type=MAP,convertedtype=LIST,valuetype=BYTE_ARRAY,valueconvertedtype=UTF8" json:"alpn_policy" diff:"alpn_policy"`
	Certificates    []*Certificate `parquet:"name=certificates,type=MAP,convertedtype=LIST" json:"certificates" diff:"certificates"`
	DefaultActions  []*Action      `parquet:"name=default_actions,type=MAP,convertedtype=LIST" json:"default_actions" diff:"default_actions"`
	ListenerArn     string         `parquet:"name=listener_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"listener_arn" diff:"listener_arn"`
	LoadBalancerArn string         `parquet:"name=load_balancer_arn,type=BYTE_ARRAY,convertedtype=UTF8" inventory_primary_key:"true" json:"load_balancer_arn" diff:"load_balancer_arn,identifier"`
	Port            int32          `parquet:"name=port,type=INT32" json:"port" diff:"port"`
	Protocol        string         `parquet:"name=protocol,type=BYTE_ARRAY,convertedtype=UTF8" json:"protocol" diff:"protocol"`
	SslPolicy       string         `parquet:"name=ssl_policy,type=BYTE_ARRAY,convertedtype=UTF8" json:"ssl_policy" diff:"ssl_policy"`
}

type Certificate struct {
	CertificateArn string `parquet:"name=certificate_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"certificate_arn" diff:"certificate_arn"`
	IsDefault      bool   `parquet:"name=is_default,type=BOOLEAN" json:"is_default" diff:"is_default"`
}

type Action struct {
	Type                      string                           `parquet:"name=type,type=BYTE_ARRAY,convertedtype=UTF8" json:"type" diff:"type"`
	AuthenticateCognitoConfig *AuthenticateCognitoActionConfig `parquet:"name=authenticate_cognito_config" json:"authenticate_cognito_config" diff:"authenticate_cognito_config"`
	AuthenticateOidcConfig    *AuthenticateOidcActionConfig    `parquet:"name=authenticate_oidc_config" json:"authenticate_oidc_config" diff:"authenticate_oidc_config"`
	FixedResponseConfig       *FixedResponseActionConfig       `parquet:"name=fixed_response_config" json:"fixed_response_config" diff:"fixed_response_config"`
	ForwardConfig             *ForwardActionConfig             `parquet:"name=forward_config" json:"forward_config" diff:"forward_config"`
	Order                     int32                            `parquet:"name=order,type=INT32" json:"order" diff:"order"`
	RedirectConfig            *RedirectActionConfig            `parquet:"name=redirect_config" json:"redirect_config" diff:"redirect_config"`
	TargetGroupArn            string                           `parquet:"name=target_group_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"target_group_arn" diff:"target_group_arn"`
}

type AuthenticateCognitoActionConfig struct {
	UserPoolArn                      string            `parquet:"name=user_pool_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"user_pool_arn" diff:"user_pool_arn"`
	UserPoolClientId                 string            `parquet:"name=user_pool_client_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"user_pool_client_id" diff:"user_pool_client_id"`
	UserPoolDomain                   string            `parquet:"name=user_pool_domain,type=BYTE_ARRAY,convertedtype=UTF8" json:"user_pool_domain" diff:"user_pool_domain"`
	AuthenticationRequestExtraParams map[string]string `parquet:"name=authentication_request_extra_params,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"authentication_request_extra_params" diff:"authentication_request_extra_params"`
	OnUnauthenticatedRequest         string            `parquet:"name=on_unauthenticated_request,type=BYTE_ARRAY,convertedtype=UTF8" json:"on_unauthenticated_request" diff:"on_unauthenticated_request"`
	Scope                            string            `parquet:"name=scope,type=BYTE_ARRAY,convertedtype=UTF8" json:"scope" diff:"scope"`
	SessionCookieName                string            `parquet:"name=session_cookie_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"session_cookie_name" diff:"session_cookie_name"`
	SessionTimeout                   int64             `parquet:"name=session_timeout,type=INT64" json:"session_timeout" diff:"session_timeout"`
}

type AuthenticateOidcActionConfig struct {
	AuthorizationEndpoint            string            `parquet:"name=authorization_endpoint,type=BYTE_ARRAY,convertedtype=UTF8" json:"authorization_endpoint" diff:"authorization_endpoint"`
	ClientId                         string            `parquet:"name=client_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"client_id" diff:"client_id"`
	Issuer                           string            `parquet:"name=issuer,type=BYTE_ARRAY,convertedtype=UTF8" json:"issuer" diff:"issuer"`
	TokenEndpoint                    string            `parquet:"name=token_endpoint,type=BYTE_ARRAY,convertedtype=UTF8" json:"token_endpoint" diff:"token_endpoint"`
	UserInfoEndpoint                 string            `parquet:"name=user_info_endpoint,type=BYTE_ARRAY,convertedtype=UTF8" json:"user_info_endpoint" diff:"user_info_endpoint"`
	AuthenticationRequestExtraParams map[string]string `parquet:"name=authentication_request_extra_params,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"authentication_request_extra_params" diff:"authentication_request_extra_params"`
	ClientSecret                     string            `parquet:"name=client_secret,type=BYTE_ARRAY,convertedtype=UTF8" json:"client_secret" diff:"client_secret"`
	OnUnauthenticatedRequest         string            `parquet:"name=on_unauthenticated_request,type=BYTE_ARRAY,convertedtype=UTF8" json:"on_unauthenticated_request" diff:"on_unauthenticated_request"`
	Scope                            string            `parquet:"name=scope,type=BYTE_ARRAY,convertedtype=UTF8" json:"scope" diff:"scope"`
	SessionCookieName                string            `parquet:"name=session_cookie_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"session_cookie_name" diff:"session_cookie_name"`
	SessionTimeout                   int64             `parquet:"name=session_timeout,type=INT64" json:"session_timeout" diff:"session_timeout"`
	UseExistingClientSecret          bool              `parquet:"name=use_existing_client_secret,type=BOOLEAN" json:"use_existing_client_secret" diff:"use_existing_client_secret"`
}

type FixedResponseActionConfig struct {
	StatusCode  string `parquet:"name=status_code,type=BYTE_ARRAY,convertedtype=UTF8" json:"status_code" diff:"status_code"`
	ContentType string `parquet:"name=content_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"content_type" diff:"content_type"`
	MessageBody string `parquet:"name=message_body,type=BYTE_ARRAY,convertedtype=UTF8" json:"message_body" diff:"message_body"`
}

type ForwardActionConfig struct {
	TargetGroupStickinessConfig *TargetGroupStickinessConfig `parquet:"name=target_group_stickiness_config" json:"target_group_stickiness_config" diff:"target_group_stickiness_config"`
	TargetGroups                []*TargetGroupTuple          `parquet:"name=target_groups,type=MAP,convertedtype=LIST" json:"target_groups" diff:"target_groups"`
}

type TargetGroupStickinessConfig struct {
	DurationSeconds int32 `parquet:"name=duration_seconds,type=INT32" json:"duration_seconds" diff:"duration_seconds"`
	Enabled         bool  `parquet:"name=enabled,type=BOOLEAN" json:"enabled" diff:"enabled"`
}

type TargetGroupTuple struct {
	TargetGroupArn string `parquet:"name=target_group_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"target_group_arn" diff:"target_group_arn"`
	Weight         int32  `parquet:"name=weight,type=INT32" json:"weight" diff:"weight"`
}

type RedirectActionConfig struct {
	StatusCode string `parquet:"name=status_code,type=BYTE_ARRAY,convertedtype=UTF8" json:"status_code" diff:"status_code"`
	Host       string `parquet:"name=host,type=BYTE_ARRAY,convertedtype=UTF8" json:"host" diff:"host"`
	Path       string `parquet:"name=path,type=BYTE_ARRAY,convertedtype=UTF8" json:"path" diff:"path"`
	Port       string `parquet:"name=port,type=BYTE_ARRAY,convertedtype=UTF8" json:"port" diff:"port"`
	Protocol   string `parquet:"name=protocol,type=BYTE_ARRAY,convertedtype=UTF8" json:"protocol" diff:"protocol"`
	Query      string `parquet:"name=query,type=BYTE_ARRAY,convertedtype=UTF8" json:"query" diff:"query"`
}

type Matcher struct {
	GrpcCode string `parquet:"name=grpc_code,type=BYTE_ARRAY,convertedtype=UTF8" json:"grpc_code" diff:"grpc_code"`
	HttpCode string `parquet:"name=http_code,type=BYTE_ARRAY,convertedtype=UTF8" json:"http_code" diff:"http_code"`
}

type TargetHealthDescription struct {
	HealthCheckPort string             `parquet:"name=health_check_port,type=BYTE_ARRAY,convertedtype=UTF8" json:"health_check_port" diff:"health_check_port"`
	Target          *TargetDescription `parquet:"name=target" json:"target" diff:"target"`
	TargetHealth    *TargetHealth      `parquet:"name=target_health" json:"target_health" diff:"target_health"`
}

type TargetDescription struct {
	Id               string `parquet:"name=id,type=BYTE_ARRAY,convertedtype=UTF8" json:"id" diff:"id"`
	AvailabilityZone string `parquet:"name=availability_zone,type=BYTE_ARRAY,convertedtype=UTF8" json:"availability_zone" diff:"availability_zone"`
	Port             int32  `parquet:"name=port,type=INT32" json:"port" diff:"port"`
}

type TargetHealth struct {
	Description string `parquet:"name=description,type=BYTE_ARRAY,convertedtype=UTF8" json:"description" diff:"description"`
	Reason      string `parquet:"name=reason,type=BYTE_ARRAY,convertedtype=UTF8" json:"reason" diff:"reason"`
	State       string `parquet:"name=state,type=BYTE_ARRAY,convertedtype=UTF8" json:"state" diff:"state"`
}