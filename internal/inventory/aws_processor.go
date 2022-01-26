package inventory

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/sheacloud/cloud-inventory/internal/indexedstorage"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud"
	"github.com/sheacloud/cloud-inventory/pkg/awscloud/interfaces"
	"github.com/sheacloud/cloud-inventory/pkg/meta"
	"github.com/sirupsen/logrus"
)

type AwsFetchJob struct {
	Input              *awscloud.AwsFetchInput
	IndexFileWaitGroup *sync.WaitGroup
	Function           func(context.Context, *awscloud.AwsFetchInput) *awscloud.AwsFetchOutput
}

func stringInList(s string, list []string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}

	return false
}

func ProcessAwsFetchJobs(ctx context.Context, jobs <-chan AwsFetchJob, results chan<- *awscloud.AwsFetchOutput, waitGroup *sync.WaitGroup) {
	for job := range jobs {
		output := job.Function(ctx, job.Input)
		job.IndexFileWaitGroup.Done()
		results <- output
	}

	waitGroup.Done()
}

func FetchAwsInventory(ctx context.Context, accountIds []string, regions []string, baseAwsConfig aws.Config, useLocalCredentials bool, assumeRoleName string, reportTime time.Time, fileManager *indexedstorage.IndexedFileManager, numWorkers int) {
	jobs := make(chan AwsFetchJob)
	results := make(chan *awscloud.AwsFetchOutput)
	workerWaitGroup := &sync.WaitGroup{}

	// start the result processor
	resultProcessorWaitGroup := &sync.WaitGroup{}
	resultProcessorWaitGroup.Add(1)

	reportDateString := reportTime.Format("2006-01-02")
	reportTimeMilli := reportTime.UTC().UnixMilli()

	resultFileIndices := []string{"meta", "inventory_results"}
	resultIndexFile, err := fileManager.GetIndexedFile(resultFileIndices, reportDateString, reportTimeMilli, new(meta.InventoryResults))
	if err != nil {
		logrus.Errorf("error getting result index file: %v", err)
		panic(err)
	}
	go func() {
		for result := range results {
			if len(result.FetchingErrors) != 0 {
				for _, err := range result.FetchingErrors {
					logrus.Errorf("Fetching error: %v", err)
				}
			}

			err = resultIndexFile.Write(ctx, result.InventoryResults)
			if err != nil {
				logrus.Errorf("error writing result to index file: %v", err)
			}

			logrus.Infof("Processed account %s:%s for resource %s. Fetched %d resources, failed to fetch %d resources", result.AccountId, result.Region, result.ResourceName, result.InventoryResults.FetchedResources, result.InventoryResults.FailedResources)
		}
		err := resultIndexFile.Close()
		if err != nil {
			logrus.Errorf("error closing result index file: %v", err)
		}
		resultProcessorWaitGroup.Done()
	}()

	// start the worker routines
	for i := 0; i < numWorkers; i++ {
		workerWaitGroup.Add(1)
		go ProcessAwsFetchJobs(ctx, jobs, results, workerWaitGroup)
	}

	resourceWaitGroup := sync.WaitGroup{}               // waits for all resources to be fetched, and corresponding index files to be written
	indexFileWaitGroups := map[string]*sync.WaitGroup{} // wait for all jobs for given index to be done

	stsSvc := sts.NewFromConfig(baseAwsConfig)

	// create all the AWS clients for eacha account/region
	accountClients := map[string]map[string]interfaces.AwsClient{}

	if !useLocalCredentials {
		for _, accountId := range accountIds {
			accountClients[accountId] = map[string]interfaces.AwsClient{}
			accountCfg := baseAwsConfig.Copy()
			creds := stscreds.NewAssumeRoleProvider(stsSvc, fmt.Sprintf("arn:aws:iam::%s:role/%s", accountId, assumeRoleName))
			accountCfg.Credentials = aws.NewCredentialsCache(creds)
			for _, region := range regions {
				regionCfg := accountCfg.Copy()
				regionCfg.Region = region
				accountClients[accountId][region] = awscloud.NewAwsClient(regionCfg)
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
		accountClients[accountId] = map[string]interfaces.AwsClient{}
		for _, region := range regions {
			regionCfg := baseAwsConfig.Copy()
			regionCfg.Region = region
			accountClients[accountId][region] = awscloud.NewAwsClient(regionCfg)
		}
	}

	for _, service := range AwsCatalog {
		for _, resource := range service.Resources {
			fileIndices := []string{"aws", service.ServiceName, resource.ResourceName}
			fileIndex := strings.Join(fileIndices, "/")
			indexFile, err := fileManager.GetIndexedFile(fileIndices, reportDateString, reportTimeMilli, resource.ResourceModel)
			if err != nil {
				panic(err)
			}
			indexFileWaitGroups[fileIndex] = &sync.WaitGroup{}

			resourceWaitGroup.Add(1)

			for _, accountId := range accountIds {
				for _, region := range regions {
					// skip if there are region overrides in place
					if len(service.RegionOverrides) != 0 {
						if !stringInList(region, service.RegionOverrides) {
							continue
						}
					}
					indexFileWaitGroups[fileIndex].Add(1)
					input := &awscloud.AwsFetchInput{
						AccountId:       accountId,
						Region:          region,
						RegionalClients: accountClients[accountId],
						ReportTime:      reportTime,
						OutputFile:      indexFile,
					}
					jobs <- AwsFetchJob{
						Input:              input,
						IndexFileWaitGroup: indexFileWaitGroups[fileIndex],
						Function:           resource.FetchFunction,
					}
				}

			}

			go func() {
				indexFileWaitGroups[fileIndex].Wait()
				logrus.Infof("Writing manifest file for %s", fileIndex)
				err := indexFile.UpdateManifest(ctx)
				if err != nil {
					panic(err)
				}
				logrus.Infof("Closing index file %s", fileIndex)
				err = indexFile.Close()
				if err != nil {
					panic(err)
				}
				resourceWaitGroup.Done()
			}()

		}
	}

	// close the jobs channel, so workers will eventually exit
	close(jobs)
	logrus.Info("Waiting for all jobs to finish")

	// wait for all jobs to finish, along with all index files being closed
	resourceWaitGroup.Wait()
	logrus.Info("All resources fetched")

	// wait for all workers to exit, meaning all results have been written to the results channel
	workerWaitGroup.Wait()
	logrus.Info("All workers finished")

	// close the results channel, so result processor will exit
	close(results)
	// wait for the result processor to exit
	resultProcessorWaitGroup.Wait()
	logrus.Info("All results processed")
}
