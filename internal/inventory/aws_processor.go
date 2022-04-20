package inventory

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/sheacloud/cloud-inventory/internal/db"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
	"github.com/sirupsen/logrus"
)

type AwsFetchJob struct {
	Input     *localAws.AwsFetchInput
	Service   string
	Resource  string
	DAO       db.DAO
	WaitGroup *sync.WaitGroup
	Function  func(context.Context, db.DAO, *localAws.AwsFetchInput) (*localAws.AwsFetchOutputMetadata, error)
}

func stringInList(s string, list []string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}

	return false
}

func ProcessAwsFetchJobs(ctx context.Context, jobs <-chan AwsFetchJob, results chan<- *localAws.AwsFetchOutputMetadata, waitGroup *sync.WaitGroup) {
	for job := range jobs {
		output, err := job.Function(ctx, job.DAO, job.Input)
		if err != nil {
			logrus.Errorf("Error processing job: %v", err)
			job.WaitGroup.Done()
			continue
		}
		results <- output

		err = job.DAO.WriteInventoryResults(ctx, output.InventoryResults)
		if err != nil {
			logrus.Errorf("Error writing inventory results: %v", err)
		}

		job.WaitGroup.Done()
	}

	waitGroup.Done()
}

func FetchAwsInventory(ctx context.Context, accountIds []string, regions []string, baseAwsConfig aws.Config, useLocalCredentials bool, assumeRoleName string, reportTime time.Time, dao db.DAO, numWorkers int) {
	jobs := make(chan AwsFetchJob)
	results := make(chan *localAws.AwsFetchOutputMetadata)
	workerWaitGroup := &sync.WaitGroup{}

	// start the result processor
	resultProcessorWaitGroup := &sync.WaitGroup{}
	resultProcessorWaitGroup.Add(1)

	go func() {
		for result := range results {
			if len(result.FetchingErrors) != 0 {
				for _, err := range result.FetchingErrors {
					logrus.Errorf("Fetching error: %v", err)
				}
			}
			logrus.Infof("Processed account %s:%s for %s:%s:%s. Fetched %d resources, failed to fetch %d resources", result.InventoryResults.AccountId, result.InventoryResults.Region, result.InventoryResults.Cloud, result.InventoryResults.Service, result.InventoryResults.Resource, result.InventoryResults.FetchedResources, result.InventoryResults.FailedResources)
		}
		resultProcessorWaitGroup.Done()
	}()

	// start the worker routines
	for i := 0; i < numWorkers; i++ {
		workerWaitGroup.Add(1)
		go ProcessAwsFetchJobs(ctx, jobs, results, workerWaitGroup)
	}

	jobWaitGroup := &sync.WaitGroup{}

	stsSvc := sts.NewFromConfig(baseAwsConfig)

	// create all the AWS clients for eacha account/region
	accountClients := map[string]map[string]localAws.AwsClientInterface{}

	if !useLocalCredentials {
		for _, accountId := range accountIds {
			accountClients[accountId] = map[string]localAws.AwsClientInterface{}
			accountCfg := baseAwsConfig.Copy()
			creds := stscreds.NewAssumeRoleProvider(stsSvc, fmt.Sprintf("arn:aws:iam::%s:role/%s", accountId, assumeRoleName))
			accountCfg.Credentials = aws.NewCredentialsCache(creds)
			for _, region := range regions {
				regionCfg := accountCfg.Copy()
				regionCfg.Region = region
				accountClients[accountId][region] = localAws.NewAwsClient(regionCfg)
			}
		}
	} else {
		// get local account id from STS
		stsOutput, err := stsSvc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
		if err != nil {
			logrus.Errorf("error getting caller identity: %v", err)
			panic(err)
		}
		accountId := *stsOutput.Account
		accountIds = []string{accountId}
		accountClients[accountId] = map[string]localAws.AwsClientInterface{}
		for _, region := range regions {
			regionCfg := baseAwsConfig.Copy()
			regionCfg.Region = region
			accountClients[accountId][region] = localAws.NewAwsClient(regionCfg)
		}
	}

	for _, service := range AwsCatalog {
		for _, resource := range service.Resources {
			for _, accountId := range accountIds {
				for _, region := range regions {
					// skip if there are region overrides in place
					if len(service.RegionOverrides) != 0 {
						if !stringInList(region, service.RegionOverrides) {
							continue
						}
					}
					input := &localAws.AwsFetchInput{
						AccountId:       accountId,
						Region:          region,
						RegionalClients: accountClients[accountId],
						ReportTime:      reportTime,
					}
					jobWaitGroup.Add(1)
					jobs <- AwsFetchJob{
						Input:     input,
						Service:   service.ServiceName,
						Resource:  resource.ResourceName,
						Function:  resource.FetchFunction,
						DAO:       dao,
						WaitGroup: jobWaitGroup,
					}
				}

			}
		}
	}

	// close the jobs channel, so workers will eventually exit
	close(jobs)
	logrus.Info("Waiting for all jobs to finish")

	// wait for all jobs to finish, along with all index files being closed
	jobWaitGroup.Wait()
	logrus.Info("All jobs finished")

	// wait for all workers to exit, meaning all results have been written to the results channel
	workerWaitGroup.Wait()
	logrus.Info("All workers finished")

	// close the results channel, so result processor will exit
	close(results)
	// wait for the result processor to exit
	resultProcessorWaitGroup.Wait()
	logrus.Info("All results processed")

	// write metadata to database
	for _, service := range AwsCatalog {
		for _, resource := range service.Resources {
			dao.WriteIngestionTimestamp(ctx, &meta.IngestionTimestamp{
				Key:        fmt.Sprintf("aws:%s:%s", service.ServiceName, resource.ResourceName),
				ReportTime: reportTime,
			})
		}
	}
	logrus.Info("Wrote ingestion timestamps")
}
