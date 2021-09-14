package elasticloadbalancingv2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

func init() {
	registerCustomTargetGroupModelPostprocessingFunc(PostProcessTargetGroupModel)
}

func PostProcessTargetGroupModel(ctx context.Context, client *elasticloadbalancingv2.Client, cfg aws.Config, model *TargetGroupModel) {
	result, err := client.DescribeTags(ctx, &elasticloadbalancingv2.DescribeTagsInput{
		ResourceArns: []string{model.TargetGroupArn},
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "elasticloadbalancingv2",
			"data_source": "target_groups",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling DescribeTags")
		return
	}
	model.Tags = map[string]string{}
	for _, descriptions := range result.TagDescriptions {
		for _, tag := range descriptions.Tags {
			model.Tags[*tag.Key] = *tag.Value
		}
	}

	targetResult, err := client.DescribeTargetHealth(ctx, &elasticloadbalancingv2.DescribeTargetHealthInput{
		TargetGroupArn: aws.String(model.TargetGroupArn),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "elasticloadbalancingv2",
			"data_source": "target_groups",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling DescribeTargetHealth")
		return
	}
	for _, target := range targetResult.TargetHealthDescriptions {
		targetModel := new(TargetHealthDescriptionTargetGroupModel)
		copier.Copy(targetModel, target)
		model.Targets = append(model.Targets, targetModel)
	}
}
