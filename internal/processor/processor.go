package processor

import (
	"context"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/sheacloud/cloud-inventory/internal/controller"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

type AwsJob struct {
	controller     controller.AwsController
	ctx            context.Context
	accountId      string
	region         string
	date           time.Time
	cfg            aws.Config
	storageManager *storage.StorageManager
}

func NewAwsJob(controller controller.AwsController, ctx context.Context, accountId, region string, date time.Time, cfg aws.Config, storageManager *storage.StorageManager) AwsJob {
	return AwsJob{
		controller:     controller,
		ctx:            ctx,
		accountId:      accountId,
		region:         region,
		date:           date,
		cfg:            cfg,
		storageManager: storageManager,
	}
}

func AwsJobWorker(jobChannel chan AwsJob, cancelChannel chan bool, jobWaitGroup, workerWaitGroup *sync.WaitGroup) {
InfiniteLoop:
	for {
		select {
		case <-cancelChannel:
			logrus.Info("stopping AwsJobWorker")
			break InfiniteLoop
		case job := <-jobChannel:
			errors := job.controller.Process(job.ctx, job.accountId, job.region, job.date, job.cfg, job.storageManager)
			if errors != nil {
				for datasource, err := range errors {
					logrus.WithFields(logrus.Fields{
						"account_id": job.accountId,
						"region":     job.region,
						"cloud":      "aws",
						"service":    job.controller.GetName(),
						"datasouce":  datasource,
						"error":      err,
					}).Error("Error processing datasource")
				}
			}
			jobWaitGroup.Done()
		}
	}

	workerWaitGroup.Done()
}

type AwsJobProcessor struct {
	cancelChannels  []chan bool
	jobChannel      chan AwsJob
	numWorkers      int
	jobWaitGroup    sync.WaitGroup
	jobBufferLength int
	workerWaitGroup sync.WaitGroup
}

func NewAwsJobProcessor(numWorkers, jobBufferLength int) *AwsJobProcessor {
	a := AwsJobProcessor{
		numWorkers:      numWorkers,
		jobBufferLength: jobBufferLength,
		jobChannel:      make(chan AwsJob, jobBufferLength),
	}
	a.cancelChannels = make([]chan bool, a.numWorkers)

	return &a
}

func (a *AwsJobProcessor) Start() {
	logrus.WithFields(logrus.Fields{
		"num_workers": a.numWorkers,
	}).Info("Starting AWS Job Processor")

	// spin up a worker
	for i := 0; i < a.numWorkers; i++ {
		a.cancelChannels[i] = make(chan bool)
		a.workerWaitGroup.Add(1)
		go AwsJobWorker(a.jobChannel, a.cancelChannels[i], &a.jobWaitGroup, &a.workerWaitGroup)
	}
}

func (a *AwsJobProcessor) AddJob(job AwsJob) {
	a.jobWaitGroup.Add(1)
	a.jobChannel <- job
}

func (a *AwsJobProcessor) Cancel() {
	logrus.Info("cancelling AWS Job Processor")
	for i := 0; i < a.numWorkers; i++ {
		a.cancelChannels[i] <- true
	}
	a.workerWaitGroup.Wait()
	for i := 0; i < a.numWorkers; i++ {
		close(a.cancelChannels[i])
	}
}

func (a *AwsJobProcessor) WaitForCompletion() {
	logrus.Info("Stopping AWS Job Processor, waiting for jobs to complete")
	//wait for jobs to be complete, then cancel workers and wait for them to be done
	a.jobWaitGroup.Wait()
	logrus.Info("Stopping AWS Job Processor, waiting for workers to finish stopping")
	a.Cancel()
	close(a.jobChannel)
}
