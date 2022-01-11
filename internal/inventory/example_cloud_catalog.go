package inventory

import (
	"context"

	"github.com/sheacloud/cloud-inventory/pkg/examplecloud"
	"github.com/sheacloud/cloud-inventory/pkg/examplecloud/exampleservice"
)

type ExampleCloudCatalogResource struct {
	ResourceName  string
	ResourceModel interface{}
	FetchFunction func(context.Context, *examplecloud.ExampleCloudFetchInput) *examplecloud.ExampleCloudFetchOutput
}

type ExampleCloudCatalogService struct {
	ServiceName string
	Resources   []ExampleCloudCatalogResource
}

var (
	ExampleCloudCatalog = []ExampleCloudCatalogService{
		{
			ServiceName: "exampleservice",
			Resources: []ExampleCloudCatalogResource{
				{
					ResourceName:  "example_resource",
					ResourceModel: &exampleservice.ExampleResource{},
					FetchFunction: exampleservice.FetchExampleResource,
				},
			},
		},
	}
)
