package s3

import (
	"context"
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
	"github.com/jinzhu/copier"
	localAws "github.com/sheacloud/cloud-inventory/pkg/aws"
)

func PostProcessBucket(ctx context.Context, params *localAws.AwsFetchInput, model *Bucket) error {
	var apiError smithy.APIError

	awsClient := params.RegionalClients[params.Region]
	client := awsClient.S3()

	// real bucket location
	locationOutput, err := client.GetBucketLocation(ctx, &s3.GetBucketLocationInput{
		Bucket: aws.String(model.Name),
	})
	if err != nil {
		return fmt.Errorf("error calling GetBucketLocation: %s", err)
	}
	model.Region = string(locationOutput.LocationConstraint)
	if model.Region == "" {
		model.Region = "us-east-1"
	}

	realLocationClient, ok := params.RegionalClients[model.Region]
	if !ok {
		return fmt.Errorf("no regional client for region %s", model.Region)
	}
	realLocationS3Client := realLocationClient.S3()

	// bucket replication
	replicationOutput, err := realLocationS3Client.GetBucketReplication(ctx, &s3.GetBucketReplicationInput{
		Bucket: aws.String(model.Name),
	})
	if errors.As(err, &apiError) && apiError.ErrorCode() == "ReplicationConfigurationNotFoundError" {
		// do nothing
	} else if err != nil {
		return fmt.Errorf("error calling GetBucketReplication: %w", err)
	} else {
		replicationModel := new(ReplicationConfiguration)
		copier.CopyWithOption(replicationModel, replicationOutput.ReplicationConfiguration, localAws.CopyOption)
		model.ReplicationConfiguration = replicationModel
	}

	// bucket ACL
	aclOutput, err := realLocationS3Client.GetBucketAcl(ctx, &s3.GetBucketAclInput{
		Bucket: aws.String(model.Name),
	})
	if err != nil {
		return fmt.Errorf("error calling GetBucketAcl: %w", err)
	}
	for _, grant := range aclOutput.Grants {
		grantModel := new(Grant)
		copier.CopyWithOption(grantModel, grant, localAws.CopyOption)
		model.AclGrants = append(model.AclGrants, grantModel)
	}

	// bucket CORS
	corsOutput, err := realLocationS3Client.GetBucketCors(ctx, &s3.GetBucketCorsInput{
		Bucket: aws.String(model.Name),
	})
	if errors.As(err, &apiError); apiError.ErrorCode() == "NoSuchCORSConfiguration" {
		// do nothing
	} else if err != nil {
		return fmt.Errorf("error calling GetBucketCors: %w", err)
	} else {
		for _, rule := range corsOutput.CORSRules {
			ruleModel := new(CORSRule)
			copier.CopyWithOption(ruleModel, rule, localAws.CopyOption)
			model.CorsRules = append(model.CorsRules, ruleModel)
		}
	}

	// bucket encryption
	encryptionOutput, err := realLocationS3Client.GetBucketEncryption(ctx, &s3.GetBucketEncryptionInput{
		Bucket: aws.String(model.Name),
	})
	if errors.As(err, &apiError); apiError.ErrorCode() == "ServerSideEncryptionConfigurationNotFoundError" {
		// do nothing
	} else if err != nil {
		return fmt.Errorf("error calling GetBucketEncryption: %s", err)
	} else {
		encryptionModel := new(ServerSideEncryptionConfiguration)
		copier.CopyWithOption(encryptionModel, encryptionOutput.ServerSideEncryptionConfiguration, localAws.CopyOption)
		model.ServerSideEncryptionConfiguration = encryptionModel
	}

	// bucket intelligent tiering config
	tieringOutput, err := realLocationS3Client.ListBucketIntelligentTieringConfigurations(ctx, &s3.ListBucketIntelligentTieringConfigurationsInput{
		Bucket: aws.String(model.Name),
	})
	if err != nil {
		return fmt.Errorf("error calling ListBucketIntelligentTieringConfigurations: %s", err)
	}
	for _, tieringConfig := range tieringOutput.IntelligentTieringConfigurationList {
		tieringModel := new(IntelligentTieringConfiguration)
		copier.CopyWithOption(tieringModel, tieringConfig, localAws.CopyOption)
		model.IntelligentTieringConfigurations = append(model.IntelligentTieringConfigurations, tieringModel)
	}

	// bucket inventory config
	inventoryOutput, err := realLocationS3Client.ListBucketInventoryConfigurations(ctx, &s3.ListBucketInventoryConfigurationsInput{
		Bucket: aws.String(model.Name),
	})
	if err != nil {
		return fmt.Errorf("error calling ListBucketInventoryConfigurations: %w", err)
	}
	for _, inventoryConfig := range inventoryOutput.InventoryConfigurationList {
		inventoryModel := new(InventoryConfiguration)
		copier.CopyWithOption(inventoryModel, inventoryConfig, localAws.CopyOption)
		model.InventoryConfigurations = append(model.InventoryConfigurations, inventoryModel)
	}

	// bucket lifecycle configuration
	lifecycleOutput, err := realLocationS3Client.GetBucketLifecycleConfiguration(ctx, &s3.GetBucketLifecycleConfigurationInput{
		Bucket: aws.String(model.Name),
	})
	if errors.As(err, &apiError); apiError.ErrorCode() == "NoSuchLifecycleConfiguration" {
		// do nothing
	} else if err != nil {
		return fmt.Errorf("error calling GetBucketLifecycleConfiguration: %w", err)
	} else {
		for _, rule := range lifecycleOutput.Rules {
			ruleModel := new(LifecycleRule)
			copier.CopyWithOption(ruleModel, rule, localAws.CopyOption)
			model.LifecycleRules = append(model.LifecycleRules, ruleModel)
		}
	}

	// bucket logging
	loggingOutput, err := realLocationS3Client.GetBucketLogging(ctx, &s3.GetBucketLoggingInput{
		Bucket: aws.String(model.Name),
	})
	if err != nil {
		return fmt.Errorf("error calling GetBucketLogging: %w", err)
	}
	loggingModel := new(LoggingEnabled)
	copier.CopyWithOption(loggingModel, loggingOutput.LoggingEnabled, localAws.CopyOption)
	model.Logging = loggingModel

	// bucket policy
	// TODO determine if we keep this field or not, can be large
	policyOutput, err := realLocationS3Client.GetBucketPolicy(ctx, &s3.GetBucketPolicyInput{
		Bucket: aws.String(model.Name),
	})
	if errors.As(err, &apiError); apiError.ErrorCode() == "NoSuchBucketPolicy" {
		// do nothing
	} else if err != nil {
		return fmt.Errorf("error calling GetBucketPolicy: %s", err)
	} else {
		if policyOutput.Policy != nil {
			model.Policy = *policyOutput.Policy
		}
	}

	// bucket policy status
	policyStatusOutput, err := realLocationS3Client.GetBucketPolicyStatus(ctx, &s3.GetBucketPolicyStatusInput{
		Bucket: aws.String(model.Name),
	})
	if errors.As(err, &apiError); apiError.ErrorCode() == "NoSuchBucketPolicy" {
		// do nothing
	} else if err != nil {
		return fmt.Errorf("error calling GetBucketPolicyStatus: %w", err)
	} else {
		if policyStatusOutput.PolicyStatus != nil {
			model.IsPublic = policyStatusOutput.PolicyStatus.IsPublic
		}
	}

	// bucket tagging
	model.Tags = make(map[string]string)
	taggingOutput, err := realLocationS3Client.GetBucketTagging(ctx, &s3.GetBucketTaggingInput{
		Bucket: aws.String(model.Name),
	})
	if errors.As(err, &apiError); apiError.ErrorCode() == "NoSuchTagSet" {
		// do nothing
	} else if err != nil {
		return fmt.Errorf("error calling GetBucketTagging: %w", err)
	} else {
		for _, tag := range taggingOutput.TagSet {
			model.Tags[*tag.Key] = *tag.Value
		}
	}

	// bucket versioning
	versioningOutput, err := realLocationS3Client.GetBucketVersioning(ctx, &s3.GetBucketVersioningInput{
		Bucket: aws.String(model.Name),
	})
	if err != nil {
		return fmt.Errorf("error calling GetBucketVersioning: %w", err)
	}
	model.VersioningStatus = string(versioningOutput.Status)
	model.MFADeleteStatus = string(versioningOutput.MFADelete)

	return nil
}
