//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_service_client_interface_file.tmpl
package interfaces

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/athena"
)

type AthenaClient interface {
	GetDataCatalog(ctx context.Context, params *athena.GetDataCatalogInput, optFns ...func(*athena.Options)) (*athena.GetDataCatalogOutput, error)
	GetDatabase(ctx context.Context, params *athena.GetDatabaseInput, optFns ...func(*athena.Options)) (*athena.GetDatabaseOutput, error)
	GetWorkGroup(ctx context.Context, params *athena.GetWorkGroupInput, optFns ...func(*athena.Options)) (*athena.GetWorkGroupOutput, error)
	ListTagsForResource(ctx context.Context, params *athena.ListTagsForResourceInput, optFns ...func(*athena.Options)) (*athena.ListTagsForResourceOutput, error)
	ListWorkGroups(ctx context.Context, params *athena.ListWorkGroupsInput, optFns ...func(*athena.Options)) (*athena.ListWorkGroupsOutput, error)
	ListDataCatalogs(ctx context.Context, params *athena.ListDataCatalogsInput, optFns ...func(*athena.Options)) (*athena.ListDataCatalogsOutput, error)
	ListDatabases(ctx context.Context, params *athena.ListDatabasesInput, optFns ...func(*athena.Options)) (*athena.ListDatabasesOutput, error)
}
