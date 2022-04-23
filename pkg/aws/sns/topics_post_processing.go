package sns

import (
	"context"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sirupsen/logrus"
)

func PostProcessTopic(ctx context.Context, params *localAws.AwsFetchInput, model *Topic) error {
	awsClient := params.RegionalClients[params.Region]
	client := awsClient.SNS()

	attributeResult, err := client.GetTopicAttributes(ctx, &sns.GetTopicAttributesInput{
		TopicArn: aws.String(model.TopicArn),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "sns",
			"data_source": "topics",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling GetTopicAttributes")
		return err
	}

	for key, value := range attributeResult.Attributes {
		switch key {
		case "DeliveryPolicy":
			model.DeliveryPolicy = value
		case "DisplayName":
			model.DisplayName = value
		case "Owner":
			model.Owner = value
		case "Policy":
			model.Policy = value
		case "SubscriptionsConfirmed":
			model.SubscriptionsConfirmed, err = strconv.Atoi(value)
			if err != nil {
				continue
			}
		case "SubscriptionsDeleted":
			model.SubscriptionsDeleted, err = strconv.Atoi(value)
			if err != nil {
				continue
			}
		case "SubscriptionsPending":
			model.SubscriptionsConfirmed, err = strconv.Atoi(value)
			if err != nil {
				continue
			}
		case "EffectiveDeliveryPolicy":
			model.EffectiveDeliveryPolicy = value
		case "KmsMasterKeyId":
			model.KmsMasterKeyId = value
		case "FifoTopic":
			model.FifoTopic = true
		case "ContentBasedDeduplication":
			if value == "true" {
				model.ContentBasedDeduplication = true
			}
		}
	}

	tagsResult, err := client.ListTagsForResource(ctx, &sns.ListTagsForResourceInput{
		ResourceArn: aws.String(model.TopicArn),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "sns",
			"data_source": "topics",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling ListTagsForResource")
		return err
	}

	model.Tags = ConvertTags(tagsResult.Tags)

	return nil
}
