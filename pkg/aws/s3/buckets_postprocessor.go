package s3

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
)

func init() {
	registerCustomBucketModelPostprocessingFunc(PostProcessBucketModel)
}

func PostProcessBucketModel(ctx context.Context, client *s3.Client, cfg aws.Config, model *BucketModel) {
	if model.CreationDate != nil {
		model.CreationDateMilli = model.CreationDate.UTC().UnixMilli()
	}

	var apiError smithy.APIError

	// real bucket location
	locationOutput, err := client.GetBucketLocation(ctx, &s3.GetBucketLocationInput{
		Bucket: aws.String(model.Name),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "s3",
			"data_source": "buckets",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling GetBucketLocation")
		return
	}
	model.Region = string(locationOutput.LocationConstraint)
	if model.Region == "" {
		model.Region = "us-east-1"
	}

	realLocationCfg := cfg.Copy()
	realLocationCfg.Region = model.Region
	realLocationClient := s3.NewFromConfig(realLocationCfg)

	// bucket replication
	replicationOutput, err := realLocationClient.GetBucketReplication(ctx, &s3.GetBucketReplicationInput{
		Bucket: aws.String(model.Name),
	})
	if errors.As(err, &apiError) && apiError.ErrorCode() == "ReplicationConfigurationNotFoundError" {
		// do nothing
	} else if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "s3",
			"data_source": "buckets",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling GetBucketReplication")
		return
	} else {
		replicationModel := new(ReplicationConfigurationBucketModel)
		copier.Copy(replicationModel, replicationOutput.ReplicationConfiguration)
		model.ReplicationConfiguration = replicationModel
	}

	// bucket ACL
	aclOutput, err := realLocationClient.GetBucketAcl(ctx, &s3.GetBucketAclInput{
		Bucket: aws.String(model.Name),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "s3",
			"data_source": "buckets",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling GetBucketAcl")
		return
	}
	for _, grant := range aclOutput.Grants {
		grantModel := new(GrantBucketModel)
		copier.Copy(grantModel, grant)
		model.AclGrants = append(model.AclGrants, grantModel)
	}

	// bucket CORS
	corsOutput, err := realLocationClient.GetBucketCors(ctx, &s3.GetBucketCorsInput{
		Bucket: aws.String(model.Name),
	})
	if errors.As(err, &apiError); apiError.ErrorCode() == "NoSuchCORSConfiguration" {
		// do nothing
	} else if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "s3",
			"data_source": "buckets",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling GetBucketCors")
		return
	} else {
		for _, rule := range corsOutput.CORSRules {
			ruleModel := new(CORSRuleBucketModel)
			copier.Copy(ruleModel, rule)
			model.CorsRules = append(model.CorsRules, ruleModel)
		}
	}

	// bucket encryption
	encryptionOutput, err := realLocationClient.GetBucketEncryption(ctx, &s3.GetBucketEncryptionInput{
		Bucket: aws.String(model.Name),
	})
	if errors.As(err, &apiError); apiError.ErrorCode() == "ServerSideEncryptionConfigurationNotFoundError" {
		// do nothing
	} else if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "s3",
			"data_source": "buckets",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling GetBucketEncryption")
		return
	} else {
		encryptionModel := new(ServerSideEncryptionConfigurationBucketModel)
		copier.Copy(encryptionModel, encryptionOutput.ServerSideEncryptionConfiguration)
		model.ServerSideEncryptionConfiguration = encryptionModel
	}

	// bucket intelligent tiering config
	tieringOutput, err := realLocationClient.ListBucketIntelligentTieringConfigurations(ctx, &s3.ListBucketIntelligentTieringConfigurationsInput{
		Bucket: aws.String(model.Name),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "s3",
			"data_source": "buckets",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling ListBucketIntelligentTieringConfigurations")
		return
	}
	for _, tieringConfig := range tieringOutput.IntelligentTieringConfigurationList {
		tieringModel := new(IntelligentTieringConfigurationBucketModel)
		copier.Copy(tieringModel, tieringConfig)
		model.IntelligentTieringConfigurations = append(model.IntelligentTieringConfigurations, tieringModel)
	}

	// bucket inventory config
	inventoryOutput, err := realLocationClient.ListBucketInventoryConfigurations(ctx, &s3.ListBucketInventoryConfigurationsInput{
		Bucket: aws.String(model.Name),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "s3",
			"data_source": "buckets",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling ListBucketInventoryConfigurations")
		return
	}
	for _, inventoryConfig := range inventoryOutput.InventoryConfigurationList {
		inventoryModel := new(InventoryConfigurationBucketModel)
		copier.Copy(inventoryModel, inventoryConfig)
		model.InventoryConfigurations = append(model.InventoryConfigurations, inventoryModel)
	}

	// bucket lifecycle configuration
	lifecycleOutput, err := realLocationClient.GetBucketLifecycleConfiguration(ctx, &s3.GetBucketLifecycleConfigurationInput{
		Bucket: aws.String(model.Name),
	})
	if errors.As(err, &apiError); apiError.ErrorCode() == "NoSuchLifecycleConfiguration" {
		// do nothing
	} else if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "s3",
			"data_source": "buckets",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling GetBucketLifecycleConfiguration")
		return
	} else {
		for _, rule := range lifecycleOutput.Rules {
			ruleModel := new(LifecycleRuleBucketModel)
			copier.Copy(ruleModel, rule)
			model.LifecycleRules = append(model.LifecycleRules, ruleModel)
		}
	}

	// bucket logging
	loggingOutput, err := realLocationClient.GetBucketLogging(ctx, &s3.GetBucketLoggingInput{
		Bucket: aws.String(model.Name),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "s3",
			"data_source": "buckets",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling GetBucketLogging")
		return
	}
	loggingModel := new(LoggingEnabledBucketModel)
	copier.Copy(loggingModel, loggingOutput.LoggingEnabled)
	model.Logging = loggingModel

	// bucket policy
	// TODO determine if we keep this field or not, can be large
	policyOutput, err := realLocationClient.GetBucketPolicy(ctx, &s3.GetBucketPolicyInput{
		Bucket: aws.String(model.Name),
	})
	if errors.As(err, &apiError); apiError.ErrorCode() == "NoSuchBucketPolicy" {
		// do nothing
	} else if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "s3",
			"data_source": "buckets",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling GetBucketPolicy")
		return
	} else {
		if policyOutput.Policy != nil {
			model.Policy = *policyOutput.Policy
		}
	}

	// bucket policy status
	policyStatusOutput, err := realLocationClient.GetBucketPolicyStatus(ctx, &s3.GetBucketPolicyStatusInput{
		Bucket: aws.String(model.Name),
	})
	if errors.As(err, &apiError); apiError.ErrorCode() == "NoSuchBucketPolicy" {
		// do nothing
	} else if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "s3",
			"data_source": "buckets",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling GetBucketPolicyStatus")
		return
	} else {
		if policyStatusOutput.PolicyStatus != nil {
			model.IsPublic = policyStatusOutput.PolicyStatus.IsPublic
		}
	}

	// bucket tagging
	model.Tags = make(map[string]string)
	taggingOutput, err := realLocationClient.GetBucketTagging(ctx, &s3.GetBucketTaggingInput{
		Bucket: aws.String(model.Name),
	})
	if errors.As(err, &apiError); apiError.ErrorCode() == "NoSuchTagSet" {
		// do nothing
	} else if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "s3",
			"data_source": "buckets",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling GetBucketTagging")
		return
	} else {
		for _, tag := range taggingOutput.TagSet {
			model.Tags[*tag.Key] = *tag.Value
		}
	}

	// bucket versioning
	versioningOutput, err := realLocationClient.GetBucketVersioning(ctx, &s3.GetBucketVersioningInput{
		Bucket: aws.String(model.Name),
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service":     "s3",
			"data_source": "buckets",
			"account_id":  model.AccountId,
			"region":      model.Region,
			"cloud":       "aws",
			"error":       err,
		}).Error("error calling GetBucketVersioning")
		return
	}
	model.VersioningStatus = string(versioningOutput.Status)
	model.MFADeleteStatus = string(versioningOutput.MFADelete)
}
