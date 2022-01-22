package inventory

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/sheacloud/cloud-inventory/internal/indexedstorage"
	"github.com/sheacloud/cloud-inventory/pkg/examplecloud"
	"github.com/sirupsen/logrus"
)

type ExampleCloudFetchJob struct {
	Input              *examplecloud.ExampleCloudFetchInput
	IndexFileWaitGroup *sync.WaitGroup
	Function           func(context.Context, *examplecloud.ExampleCloudFetchInput) *examplecloud.ExampleCloudFetchOutput
}

func ProcessExampleCloudFetchJobs(ctx context.Context, jobs <-chan ExampleCloudFetchJob, results chan<- *examplecloud.ExampleCloudFetchOutput, waitGroup *sync.WaitGroup) {
	for job := range jobs {
		output := job.Function(ctx, job.Input)
		job.IndexFileWaitGroup.Done()
		results <- output
	}

	waitGroup.Done()
}

func FetchExampleCloudInventory(ctx context.Context, client examplecloud.ExampleClient, reportTime time.Time, fileManager *indexedstorage.IndexedFileManager, numWorkers int) {
	jobs := make(chan ExampleCloudFetchJob)
	results := make(chan *examplecloud.ExampleCloudFetchOutput)
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

			logrus.Infof("Fetched %d resources, failed to fetch %d resources", result.FetchedResources, result.FailedResources)
		}
		resultProcessorWaitGroup.Done()
	}()

	// start the worker routines
	for i := 0; i < numWorkers; i++ {
		workerWaitGroup.Add(1)
		go ProcessExampleCloudFetchJobs(ctx, jobs, results, workerWaitGroup)
	}

	resourceWaitGroup := sync.WaitGroup{}               // waits for all resources to be fetched, and corresponding index files to be written
	indexFileWaitGroups := map[string]*sync.WaitGroup{} // wait for all jobs for given index to be done

	reportDateString := reportTime.Format("2006-01-02")
	reportTimeMilli := reportTime.UTC().UnixMilli()

	for _, service := range ExampleCloudCatalog {
		for _, resource := range service.Resources {
			fileIndices := []string{"example_cloud", service.ServiceName, resource.ResourceName}
			fileIndex := strings.Join(fileIndices, "/")
			indexFile, err := fileManager.GetIndexedFile(fileIndices, reportDateString, reportTimeMilli, resource.ResourceModel)
			if err != nil {
				panic(err)
			}
			indexFileWaitGroups[fileIndex] = &sync.WaitGroup{}

			resourceWaitGroup.Add(1)

			for _, accountID := range []string{"A", "B", "C", "D", "E", "F"} {
				indexFileWaitGroups[fileIndex].Add(1)
				input := &examplecloud.ExampleCloudFetchInput{
					AccountID:  accountID,
					Client:     client,
					ReportTime: reportTime,
					OutputFile: indexFile,
				}
				jobs <- ExampleCloudFetchJob{
					Input:              input,
					IndexFileWaitGroup: indexFileWaitGroups[fileIndex],
					Function:           resource.FetchFunction,
				}
			}

			go func() {
				indexFileWaitGroups[fileIndex].Wait()
				logrus.Infof("Closing index file %s", fileIndex)
				err := indexFile.Close()
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
