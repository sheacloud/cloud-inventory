package viewgen

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/sirupsen/logrus"
)

type ViewJob struct {
	ctx          context.Context
	name         string
	view         string
	catalog      string
	database     string
	workgroup    string
	athenaClient *athena.Client
	errorChan    chan<- error
}

func NewViewJob(ctx context.Context, name, view, catalog, database, workgroup string, athenaClient *athena.Client, errorChan chan<- error) ViewJob {
	return ViewJob{
		ctx:          ctx,
		name:         name,
		view:         view,
		catalog:      catalog,
		database:     database,
		workgroup:    workgroup,
		athenaClient: athenaClient,
		errorChan:    errorChan,
	}
}

func ViewJobWorker(jobChannel chan ViewJob, cancelChannel chan bool, jobWaitGroup, workerWaitGroup *sync.WaitGroup) {
InfiniteLoop:
	for {
		select {
		case <-cancelChannel:
			logrus.Info("stopping ViewJobWorker")
			break InfiniteLoop
		case job := <-jobChannel:
			// process job
			query, err := job.athenaClient.StartQueryExecution(job.ctx, &athena.StartQueryExecutionInput{
				QueryString: aws.String(job.view),
				QueryExecutionContext: &types.QueryExecutionContext{
					Catalog:  aws.String(job.catalog),
					Database: aws.String(job.database),
				},
				WorkGroup: aws.String(job.workgroup),
			})
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"error": err,
				}).Error("error creating view " + job.name)
				job.errorChan <- err
				jobWaitGroup.Done()
				break
			}

			// wait for query to finish
		QueryLoop:
			for {
				time.Sleep(time.Millisecond * 500)
				queryStatus, err := job.athenaClient.GetQueryExecution(job.ctx, &athena.GetQueryExecutionInput{
					QueryExecutionId: query.QueryExecutionId,
				})
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"error": err,
					}).Error("error geting view query status: " + job.name)
					job.errorChan <- err
					break
				}

				switch queryStatus.QueryExecution.Status.State {
				case types.QueryExecutionStateSucceeded:
					break QueryLoop
				case types.QueryExecutionStateCancelled:
					logrus.Error("view query cancelled")
					job.errorChan <- errors.New("view query cancelled: " + job.name)
					break QueryLoop
				case types.QueryExecutionStateFailed:
					logrus.Error("view query failed")
					job.errorChan <- errors.New("view query failed: " + job.name)
					break QueryLoop
				}
			}

			jobWaitGroup.Done()
			logrus.Info("completed query for view " + job.name)
		}
	}

	workerWaitGroup.Done()
}

type ViewJobProcessor struct {
	cancelChannels  []chan bool
	jobChannel      chan ViewJob
	numWorkers      int
	jobWaitGroup    sync.WaitGroup
	jobBufferLength int
	workerWaitGroup sync.WaitGroup
}

func NewViewJobProcessor(numWorkers, jobBufferLength int) *ViewJobProcessor {
	a := ViewJobProcessor{
		numWorkers:      numWorkers,
		jobBufferLength: jobBufferLength,
		jobChannel:      make(chan ViewJob, jobBufferLength),
	}
	a.cancelChannels = make([]chan bool, a.numWorkers)

	return &a
}

func (a *ViewJobProcessor) Start() {
	logrus.WithFields(logrus.Fields{
		"num_workers": a.numWorkers,
	}).Info("Starting View Job Processor")

	// spin up a worker
	for i := 0; i < a.numWorkers; i++ {
		a.cancelChannels[i] = make(chan bool)
		a.workerWaitGroup.Add(1)
		go ViewJobWorker(a.jobChannel, a.cancelChannels[i], &a.jobWaitGroup, &a.workerWaitGroup)
	}
}

func (a *ViewJobProcessor) AddJob(job ViewJob) {
	a.jobWaitGroup.Add(1)
	a.jobChannel <- job
}

func (a *ViewJobProcessor) Cancel() {
	logrus.Info("cancelling View Job Processor")
	for i := 0; i < a.numWorkers; i++ {
		a.cancelChannels[i] <- true
	}
	a.workerWaitGroup.Wait()
	for i := 0; i < a.numWorkers; i++ {
		close(a.cancelChannels[i])
	}
}

func (a *ViewJobProcessor) WaitForCompletion() {
	logrus.Info("Stopping View Job Processor, waiting for jobs to complete")
	//wait for jobs to be complete, then cancel workers and wait for them to be done
	a.jobWaitGroup.Wait()
	logrus.Info("Stopping View Job Processor, waiting for workers to finish stopping")
	a.Cancel()
	close(a.jobChannel)
}
