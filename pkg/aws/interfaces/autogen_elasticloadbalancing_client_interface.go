//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_service_client_interface_file.tmpl
package interfaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
)

type ElasticLoadBalancingClient interface {
	DescribeLoadBalancers(ctx context.Context, params *elasticloadbalancing.DescribeLoadBalancersInput, optFns ...func(*elasticloadbalancing.Options)) (*elasticloadbalancing.DescribeLoadBalancersOutput, error)
}
