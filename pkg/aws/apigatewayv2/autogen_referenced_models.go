//AUTOGENERATED CODE DO NOT EDIT
package apigatewayv2

import (
	"time"
)

type Cors struct {
	AllowCredentials bool     `bson:"allow_credentials,omitempty" dynamodbav:"allow_credentials" json:"allow_credentials,omitempty" diff:"allow_credentials"`
	AllowHeaders     []string `bson:"allow_headers,omitempty" dynamodbav:"allow_headers,omitempty" json:"allow_headers,omitempty" diff:"allow_headers"`
	AllowMethods     []string `bson:"allow_methods,omitempty" dynamodbav:"allow_methods,omitempty" json:"allow_methods,omitempty" diff:"allow_methods"`
	AllowOrigins     []string `bson:"allow_origins,omitempty" dynamodbav:"allow_origins,omitempty" json:"allow_origins,omitempty" diff:"allow_origins"`
	ExposeHeaders    []string `bson:"expose_headers,omitempty" dynamodbav:"expose_headers,omitempty" json:"expose_headers,omitempty" diff:"expose_headers"`
	MaxAge           int32    `bson:"max_age,omitempty" dynamodbav:"max_age,omitempty" json:"max_age,omitempty" diff:"max_age"`
}

type Stage struct {
	StageName                   string                    `bson:"stage_name,omitempty" dynamodbav:"stage_name,omitempty" json:"stage_name,omitempty" diff:"stage_name"`
	AccessLogSettings           *AccessLogSettings        `bson:"access_log_settings,omitempty" dynamodbav:"access_log_settings,omitempty" json:"access_log_settings,omitempty" diff:"access_log_settings"`
	ApiGatewayManaged           bool                      `bson:"api_gateway_managed,omitempty" dynamodbav:"api_gateway_managed" json:"api_gateway_managed,omitempty" diff:"api_gateway_managed"`
	AutoDeploy                  bool                      `bson:"auto_deploy,omitempty" dynamodbav:"auto_deploy" json:"auto_deploy,omitempty" diff:"auto_deploy"`
	ClientCertificateId         string                    `bson:"client_certificate_id,omitempty" dynamodbav:"client_certificate_id,omitempty" json:"client_certificate_id,omitempty" diff:"client_certificate_id"`
	CreatedDate                 *time.Time                `bson:"created_date,omitempty" dynamodbav:"created_date,unixtime,omitempty" json:"created_date,omitempty" diff:"created_date"`
	DefaultRouteSettings        *RouteSettings            `bson:"default_route_settings,omitempty" dynamodbav:"default_route_settings,omitempty" json:"default_route_settings,omitempty" diff:"default_route_settings"`
	DeploymentId                string                    `bson:"deployment_id,omitempty" dynamodbav:"deployment_id,omitempty" json:"deployment_id,omitempty" diff:"deployment_id"`
	Description                 string                    `bson:"description,omitempty" dynamodbav:"description,omitempty" json:"description,omitempty" diff:"description"`
	LastDeploymentStatusMessage string                    `bson:"last_deployment_status_message,omitempty" dynamodbav:"last_deployment_status_message,omitempty" json:"last_deployment_status_message,omitempty" diff:"last_deployment_status_message"`
	LastUpdatedDate             *time.Time                `bson:"last_updated_date,omitempty" dynamodbav:"last_updated_date,unixtime,omitempty" json:"last_updated_date,omitempty" diff:"last_updated_date"`
	RouteSettings               map[string]*RouteSettings `bson:"route_settings,omitempty" dynamodbav:"route_settings,omitempty" json:"route_settings,omitempty" diff:"route_settings"`
	StageVariables              map[string]string         `bson:"stage_variables,omitempty" dynamodbav:"stage_variables,omitempty" json:"stage_variables,omitempty" diff:"stage_variables"`
	Tags                        map[string]string         `bson:"tags,omitempty" dynamodbav:"tags,omitempty" json:"tags,omitempty" diff:"tags"`
}

type AccessLogSettings struct {
	DestinationArn string `bson:"destination_arn,omitempty" dynamodbav:"destination_arn,omitempty" json:"destination_arn,omitempty" diff:"destination_arn"`
	Format         string `bson:"format,omitempty" dynamodbav:"format,omitempty" json:"format,omitempty" diff:"format"`
}

type RouteSettings struct {
	DataTraceEnabled       bool    `bson:"data_trace_enabled,omitempty" dynamodbav:"data_trace_enabled" json:"data_trace_enabled,omitempty" diff:"data_trace_enabled"`
	DetailedMetricsEnabled bool    `bson:"detailed_metrics_enabled,omitempty" dynamodbav:"detailed_metrics_enabled" json:"detailed_metrics_enabled,omitempty" diff:"detailed_metrics_enabled"`
	LoggingLevel           string  `bson:"logging_level,omitempty" dynamodbav:"logging_level,omitempty" json:"logging_level,omitempty" diff:"logging_level"`
	ThrottlingBurstLimit   int32   `bson:"throttling_burst_limit,omitempty" dynamodbav:"throttling_burst_limit,omitempty" json:"throttling_burst_limit,omitempty" diff:"throttling_burst_limit"`
	ThrottlingRateLimit    float64 `bson:"throttling_rate_limit,omitempty" dynamodbav:"throttling_rate_limit,omitempty" json:"throttling_rate_limit,omitempty" diff:"throttling_rate_limit"`
}

type GetRouteOutput struct {
	ApiGatewayManaged                bool                             `bson:"api_gateway_managed,omitempty" dynamodbav:"api_gateway_managed" json:"api_gateway_managed,omitempty" diff:"api_gateway_managed"`
	ApiKeyRequired                   bool                             `bson:"api_key_required,omitempty" dynamodbav:"api_key_required" json:"api_key_required,omitempty" diff:"api_key_required"`
	AuthorizationScopes              []string                         `bson:"authorization_scopes,omitempty" dynamodbav:"authorization_scopes,omitempty" json:"authorization_scopes,omitempty" diff:"authorization_scopes"`
	AuthorizationType                string                           `bson:"authorization_type,omitempty" dynamodbav:"authorization_type,omitempty" json:"authorization_type,omitempty" diff:"authorization_type"`
	AuthorizerId                     string                           `bson:"authorizer_id,omitempty" dynamodbav:"authorizer_id,omitempty" json:"authorizer_id,omitempty" diff:"authorizer_id"`
	ModelSelectionExpression         string                           `bson:"model_selection_expression,omitempty" dynamodbav:"model_selection_expression,omitempty" json:"model_selection_expression,omitempty" diff:"model_selection_expression"`
	OperationName                    string                           `bson:"operation_name,omitempty" dynamodbav:"operation_name,omitempty" json:"operation_name,omitempty" diff:"operation_name"`
	RequestModels                    map[string]string                `bson:"request_models,omitempty" dynamodbav:"request_models,omitempty" json:"request_models,omitempty" diff:"request_models"`
	RequestParameters                map[string]*ParameterConstraints `bson:"request_parameters,omitempty" dynamodbav:"request_parameters,omitempty" json:"request_parameters,omitempty" diff:"request_parameters"`
	RouteId                          string                           `bson:"route_id,omitempty" dynamodbav:"route_id,omitempty" json:"route_id,omitempty" diff:"route_id"`
	RouteKey                         string                           `bson:"route_key,omitempty" dynamodbav:"route_key,omitempty" json:"route_key,omitempty" diff:"route_key"`
	RouteResponseSelectionExpression string                           `bson:"route_response_selection_expression,omitempty" dynamodbav:"route_response_selection_expression,omitempty" json:"route_response_selection_expression,omitempty" diff:"route_response_selection_expression"`
	Target                           string                           `bson:"target,omitempty" dynamodbav:"target,omitempty" json:"target,omitempty" diff:"target"`
}

type ParameterConstraints struct {
	Required bool `bson:"required,omitempty" dynamodbav:"required" json:"required,omitempty" diff:"required"`
}

type Integration struct {
	ApiGatewayManaged                      bool              `bson:"api_gateway_managed,omitempty" dynamodbav:"api_gateway_managed" json:"api_gateway_managed,omitempty" diff:"api_gateway_managed"`
	ConnectionId                           string            `bson:"connection_id,omitempty" dynamodbav:"connection_id,omitempty" json:"connection_id,omitempty" diff:"connection_id"`
	ConnectionType                         string            `bson:"connection_type,omitempty" dynamodbav:"connection_type,omitempty" json:"connection_type,omitempty" diff:"connection_type"`
	ContentHandlingStrategy                string            `bson:"content_handling_strategy,omitempty" dynamodbav:"content_handling_strategy,omitempty" json:"content_handling_strategy,omitempty" diff:"content_handling_strategy"`
	CredentialsArn                         string            `bson:"credentials_arn,omitempty" dynamodbav:"credentials_arn,omitempty" json:"credentials_arn,omitempty" diff:"credentials_arn"`
	Description                            string            `bson:"description,omitempty" dynamodbav:"description,omitempty" json:"description,omitempty" diff:"description"`
	IntegrationId                          string            `bson:"integration_id,omitempty" dynamodbav:"integration_id,omitempty" json:"integration_id,omitempty" diff:"integration_id"`
	IntegrationMethod                      string            `bson:"integration_method,omitempty" dynamodbav:"integration_method,omitempty" json:"integration_method,omitempty" diff:"integration_method"`
	IntegrationResponseSelectionExpression string            `bson:"integration_response_selection_expression,omitempty" dynamodbav:"integration_response_selection_expression,omitempty" json:"integration_response_selection_expression,omitempty" diff:"integration_response_selection_expression"`
	IntegrationSubtype                     string            `bson:"integration_subtype,omitempty" dynamodbav:"integration_subtype,omitempty" json:"integration_subtype,omitempty" diff:"integration_subtype"`
	IntegrationType                        string            `bson:"integration_type,omitempty" dynamodbav:"integration_type,omitempty" json:"integration_type,omitempty" diff:"integration_type"`
	IntegrationUri                         string            `bson:"integration_uri,omitempty" dynamodbav:"integration_uri,omitempty" json:"integration_uri,omitempty" diff:"integration_uri"`
	PassthroughBehavior                    string            `bson:"passthrough_behavior,omitempty" dynamodbav:"passthrough_behavior,omitempty" json:"passthrough_behavior,omitempty" diff:"passthrough_behavior"`
	PayloadFormatVersion                   string            `bson:"payload_format_version,omitempty" dynamodbav:"payload_format_version,omitempty" json:"payload_format_version,omitempty" diff:"payload_format_version"`
	RequestParameters                      map[string]string `bson:"request_parameters,omitempty" dynamodbav:"request_parameters,omitempty" json:"request_parameters,omitempty" diff:"request_parameters"`
	RequestTemplates                       map[string]string `bson:"request_templates,omitempty" dynamodbav:"request_templates,omitempty" json:"request_templates,omitempty" diff:"request_templates"`
	TemplateSelectionExpression            string            `bson:"template_selection_expression,omitempty" dynamodbav:"template_selection_expression,omitempty" json:"template_selection_expression,omitempty" diff:"template_selection_expression"`
	TimeoutInMillis                        int32             `bson:"timeout_in_millis,omitempty" dynamodbav:"timeout_in_millis,omitempty" json:"timeout_in_millis,omitempty" diff:"timeout_in_millis"`
	TlsConfig                              *TlsConfig        `bson:"tls_config,omitempty" dynamodbav:"tls_config,omitempty" json:"tls_config,omitempty" diff:"tls_config"`
}

type TlsConfig struct {
	ServerNameToVerify string `bson:"server_name_to_verify,omitempty" dynamodbav:"server_name_to_verify,omitempty" json:"server_name_to_verify,omitempty" diff:"server_name_to_verify"`
}

type Authorizer struct {
	Name                           string            `bson:"name,omitempty" dynamodbav:"name,omitempty" json:"name,omitempty" diff:"name"`
	AuthorizerCredentialsArn       string            `bson:"authorizer_credentials_arn,omitempty" dynamodbav:"authorizer_credentials_arn,omitempty" json:"authorizer_credentials_arn,omitempty" diff:"authorizer_credentials_arn"`
	AuthorizerId                   string            `bson:"authorizer_id,omitempty" dynamodbav:"authorizer_id,omitempty" json:"authorizer_id,omitempty" diff:"authorizer_id"`
	AuthorizerPayloadFormatVersion string            `bson:"authorizer_payload_format_version,omitempty" dynamodbav:"authorizer_payload_format_version,omitempty" json:"authorizer_payload_format_version,omitempty" diff:"authorizer_payload_format_version"`
	AuthorizerResultTtlInSeconds   int32             `bson:"authorizer_result_ttl_in_seconds,omitempty" dynamodbav:"authorizer_result_ttl_in_seconds,omitempty" json:"authorizer_result_ttl_in_seconds,omitempty" diff:"authorizer_result_ttl_in_seconds"`
	AuthorizerType                 string            `bson:"authorizer_type,omitempty" dynamodbav:"authorizer_type,omitempty" json:"authorizer_type,omitempty" diff:"authorizer_type"`
	AuthorizerUri                  string            `bson:"authorizer_uri,omitempty" dynamodbav:"authorizer_uri,omitempty" json:"authorizer_uri,omitempty" diff:"authorizer_uri"`
	EnableSimpleResponses          bool              `bson:"enable_simple_responses,omitempty" dynamodbav:"enable_simple_responses" json:"enable_simple_responses,omitempty" diff:"enable_simple_responses"`
	IdentitySource                 []string          `bson:"identity_source,omitempty" dynamodbav:"identity_source,omitempty" json:"identity_source,omitempty" diff:"identity_source"`
	IdentityValidationExpression   string            `bson:"identity_validation_expression,omitempty" dynamodbav:"identity_validation_expression,omitempty" json:"identity_validation_expression,omitempty" diff:"identity_validation_expression"`
	JwtConfiguration               *JWTConfiguration `bson:"jwt_configuration,omitempty" dynamodbav:"jwt_configuration,omitempty" json:"jwt_configuration,omitempty" diff:"jwt_configuration"`
}

type JWTConfiguration struct {
	Audience []string `bson:"audience,omitempty" dynamodbav:"audience,omitempty" json:"audience,omitempty" diff:"audience"`
	Issuer   string   `bson:"issuer,omitempty" dynamodbav:"issuer,omitempty" json:"issuer,omitempty" diff:"issuer"`
}
