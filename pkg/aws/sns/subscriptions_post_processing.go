package sns

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
	"github.com/sirupsen/logrus"
)

func PostProcessSubscription(ctx context.Context, params *localAws.AwsFetchInput, model *Subscription) error {
	awsClient := params.RegionalClients[params.Region]
	client := awsClient.SNS()

	if model.SubscriptionArn == "Deleted" {
		return nil
	}
	attributeResult, err := client.GetSubscriptionAttributes(ctx, &sns.GetSubscriptionAttributesInput{
		SubscriptionArn: aws.String(model.SubscriptionArn),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "sns",
			"data_source": "subscriptions",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"arn":         model.SubscriptionArn,
			"error":       err,
		}).Error("error calling GetSubscriptionAttributes")
		return err
	}

	for key, value := range attributeResult.Attributes {
		switch key {
		case "DeliveryPolicy":
			model.DeliveryPolicy = value
		case "ConfirmationWasAuthenticated":
			if value == "true" {
				model.ConfirmationWasAuthenticated = true
			}
		case "EffectiveDeliveryPolicy":
			model.EffectiveDeliveryPolicy = value
		case "FilterPolicy":
			model.FilterPolicy = value
		case "PendingConfirmation":
			if value == "true" {
				model.PendingConfirmation = true
			}
		case "RawMessageDelivery":
			if value == "true" {
				model.RawMessageDelivery = true
			}
		case "RedrivePolicy":
			model.RedrivePolicy = value
		case "SubscriptionRoleArn":
			model.SubscriptionRoleArn = value
		}
	}

	return nil
}
