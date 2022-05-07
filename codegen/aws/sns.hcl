aws_service "sns" {
  service_cap_name         = "SNS"
  sdk_path                 = "github.com/aws/aws-sdk-go-v2/service/sns"
  extra_utilized_functions = ["GetSubscriptionAttributes", "GetTopicAttributes", "ListTagsForResource"]
  tag_object_name          = "Tag"

  resource "topics" {
    fetch_function        = "ListTopics"
    object_source_name    = "Topic"
    object_plural_name    = "Topics"
    object_unique_id      = "TopicArn"
    object_response_field = "Topics"
    model_only            = false
    pagination            = true
    use_post_processing   = true
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["TopicArn"]

    extra_field {
      name = "DeliveryPolicy"
      type = "string"
    }

    extra_field {
      name = "DisplayName"
      type = "string"
    }

    extra_field {
      name = "Owner"
      type = "string"
    }

    extra_field {
      name = "Policy"
      type = "string"
    }

    extra_field {
      name = "SubscriptionsConfirmed"
      type = "int"
    }

    extra_field {
      name = "SubscriptionsDeleted"
      type = "int"
    }

    extra_field {
      name = "SubscriptionsPending"
      type = "int"
    }

    extra_field {
      name = "EffectiveDeliveryPolicy"
      type = "string"
    }

    extra_field {
      name = "KmsMasterKeyId"
      type = "string"
    }

    extra_field {
      name = "FifoTopic"
      type = "bool"
    }

    extra_field {
      name = "ContentBasedDeduplication"
      type = "bool"
    }

    extra_field {
      name = "Tags"
      type = "map[string]string"
    }
  }

  resource "subscriptions" {
    fetch_function        = "ListSubscriptions"
    object_source_name    = "Subscription"
    object_plural_name    = "Subscriptions"
    object_unique_id      = "SubscriptionArn"
    object_response_field = "Subscriptions"
    model_only            = false
    pagination            = true
    use_post_processing   = true
    excluded_fields       = []
    convert_tags          = false
    display_fields        = ["TopicArn", "Endpoint"]

    extra_field {
      name = "ConfirmationWasAuthenticated"
      type = "bool"
    }

    extra_field {
      name = "DeliveryPolicy"
      type = "string"
    }

    extra_field {
      name = "EffectiveDeliveryPolicy"
      type = "string"
    }

    extra_field {
      name = "FilterPolicy"
      type = "string"
    }

    extra_field {
      name = "PendingConfirmation"
      type = "bool"
    }

    extra_field {
      name = "RawMessageDelivery"
      type = "bool"
    }

    extra_field {
      name = "RedrivePolicy"
      type = "string"
    }

    extra_field {
      name = "SubscriptionRoleArn"
      type = "string"
    }
  }
}
