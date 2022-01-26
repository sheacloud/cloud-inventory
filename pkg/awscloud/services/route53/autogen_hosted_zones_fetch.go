//AUTOGENERATED CODE DO NOT EDIT
package route53

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

func FetchHostedZone(ctx context.Context, params *awscloud.AwsFetchInput) *awscloud.AwsFetchOutput {
	fetchingErrors := []error{}
	var fetchedResources int
	var failedResources int
	inventoryResults := &meta.InventoryResults{
		Cloud:      "aws",
		Service:    "route53",
		Resource:   "hosted_zones",
		AccountId:  params.AccountId,
		Region:     params.Region,
		ReportTime: params.ReportTime.UTC().UnixMilli(),
	}

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.Route53()

	paginator := route53.NewListHostedZonesPaginator(client, &route53.ListHostedZonesInput{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)

		if err != nil {
			fetchingErrors = append(fetchingErrors, fmt.Errorf("error calling ListHostedZones in %s/%s: %w", params.AccountId, params.Region, err))
			break
		}

		for _, object := range output.HostedZones {

			model := new(HostedZone)
			copier.Copy(&model, &object)

			model.AccountId = params.AccountId
			model.Region = params.Region
			model.ReportTime = params.ReportTime.UTC().UnixMilli()

			if err = PostProcessHostedZone(ctx, params, model); err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error post-processing HostedZone %s %s/%s: %w", model.Id, params.AccountId, params.Region, err))
				failedResources++
			}

			err = params.OutputFile.Write(ctx, model)
			if err != nil {
				fetchingErrors = append(fetchingErrors, fmt.Errorf("error storing HostedZone model in %s/%s: %w", params.AccountId, params.Region, err))
			}
			fetchedResources++
		}

	}

	inventoryResults.FetchedResources = fetchedResources
	inventoryResults.FailedResources = failedResources
	inventoryResults.HadErrors = len(fetchingErrors) > 0

	return &awscloud.AwsFetchOutput{
		FetchingErrors:   fetchingErrors,
		InventoryResults: inventoryResults,
		ResourceName:     "hosted_zones",
		AccountId:        params.AccountId,
		Region:           params.Region,
	}
}
