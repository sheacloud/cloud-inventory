package ecs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ecs"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

func init() {
	registerCustomClusterModelPostprocessingFunc(PostProcessClusterModel)
}

func PostProcessClusterModel(ctx context.Context, client *ecs.Client, cfg aws.Config, model *ClusterModel) {
	//populate services
	listServicesPaginator := ecs.NewListServicesPaginator(client, &ecs.ListServicesInput{
		MaxResults: aws.Int32(10),
		Cluster:    aws.String(model.ClusterName),
	})

	for listServicesPaginator.HasMorePages() {
		output, err := listServicesPaginator.NextPage(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":     "ecs",
				"data_source": "clusters",
				"account_id":  model.AccountId,
				"region":      model.Region,
				"cloud":       "aws",
				"error":       err,
			}).Error("error calling ListServices")
			return
		}

		if len(output.ServiceArns) > 0 {
			servicesOutput, err := client.DescribeServices(ctx, &ecs.DescribeServicesInput{
				Services: output.ServiceArns,
				Cluster:  aws.String(model.ClusterName),
			})
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"service":     "ecs",
					"data_source": "clusters",
					"account_id":  model.AccountId,
					"region":      model.Region,
					"cloud":       "aws",
					"error":       err,
				}).Error("error calling DescribeServices")
				return
			}
			for _, service := range servicesOutput.Services {
				serviceModel := new(ServiceClusterModel)
				copier.Copy(serviceModel, service)
				model.Services = append(model.Services, serviceModel)
			}
		}
	}

	//populate tasks
	listTasksPaginator := ecs.NewListTasksPaginator(client, &ecs.ListTasksInput{
		MaxResults: aws.Int32(10),
		Cluster:    aws.String(model.ClusterName),
	})

	for listTasksPaginator.HasMorePages() {
		output, err := listTasksPaginator.NextPage(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":     "ecs",
				"data_source": "clusters",
				"account_id":  model.AccountId,
				"region":      model.Region,
				"cloud":       "aws",
				"error":       err,
			}).Error("error calling ListTasks")
			return
		}

		if len(output.TaskArns) > 0 {
			tasksOutput, err := client.DescribeTasks(ctx, &ecs.DescribeTasksInput{
				Tasks:   output.TaskArns,
				Cluster: aws.String(model.ClusterName),
			})
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"service":     "ecs",
					"data_source": "clusters",
					"account_id":  model.AccountId,
					"region":      model.Region,
					"cloud":       "aws",
					"error":       err,
				}).Error("error calling DescribeTasks")
				return
			}
			for _, task := range tasksOutput.Tasks {
				taskModel := new(TaskClusterModel)
				copier.Copy(taskModel, task)
				model.Tasks = append(model.Tasks, taskModel)
			}
		}
	}
}
