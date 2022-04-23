aws_service "sqs" {
  service_cap_name         = "SQS"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/sqs"
  extra_utilized_functions = ["GetQueueAttributes", "ListQueueTags"]

  resource "queues" {
    fetch_function        = "ListQueues"
    object_source_name    = "GetQueueAttributesOutput"
    object_singular_name  = "Queue"
    object_plural_name    = "Queues"
    object_unique_id      = "QueueArn"
    object_response_field = "Queues"
    model_only            = true
    pagination            = true
    use_post_processing   = false
    excluded_fields       = ["Attributes", "ResultMetadata"]
    convert_tags          = false
    display_fields        = ["QueueUrl"]

    extra_field {
      name = "QueueURL"
      type = "string"
    }

    extra_field {
      name = "Policy"
      type = "string"
    }

    extra_field {
      name = "VisibilityTimeout"
      type = "int64"
    }

    extra_field {
      name = "MaximumMessageSize"
      type = "int64"
    }

    extra_field {
      name = "MessageRetentionPeriod"
      type = "int64"
    }

    extra_field {
      name = "ApproximateNumberOfMessages"
      type = "int64"
    }

    extra_field {
      name = "ApproximateNumberOfMessagesNotVisible"
      type = "int64"
    }

    extra_field {
      name = "CreatedTimestamp"
      type = "string"
    }

    extra_field {
      name = "LastModifiedTimestamp"
      type = "string"
    }

    extra_field {
      name = "QueueArn"
      type = "string"
    }

    extra_field {
      name = "ApproximateNumberOfMessagesDelayed"
      type = "int64"
    }

    extra_field {
      name = "DelaySeconds"
      type = "int64"
    }

    extra_field {
      name = "ReceiveMessageWaitTimeSeconds"
      type = "int64"
    }

    extra_field {
      name = "RedrivePolicy"
      type = "string"
    }

    extra_field {
      name = "FifoQueue"
      type = "bool"
    }

    extra_field {
      name = "ContentBasedDeduplication"
      type = "bool"
    }

    extra_field {
      name = "KmsMasterKeyId"
      type = "string"
    }

    extra_field {
      name = "KmsDataKeyReusePeriodSeconds"
      type = "int64"
    }

    extra_field {
      name = "DeduplicationScope"
      type = "string"
    }

    extra_field {
      name = "FifoThroughputLimit"
      type = "string"
    }

    extra_field {
      name = "Tags"
      type = "map[string]string"
    }
  }
}
