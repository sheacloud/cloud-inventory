aws_service "s3" {
    service_cap_name = "S3"
    sdk_path = "github.com/aws/aws-sdk-go-v2/service/s3"
    extra_utilized_functions = ["GetBucketLocation", "GetBucketReplication", "GetBucketAcl", "GetBucketCors", "GetBucketEncryption", "ListBucketIntelligentTieringConfigurations", "ListBucketInventoryConfigurations", "GetBucketLifecycleConfiguration", "GetBucketLogging", "GetBucketPolicy", "GetBucketPolicyStatus", "GetBucketTagging", "GetBucketVersioning"]
    tag_object_name = "Tag"
    region_override = "us-east-1"

    resource "buckets" {
        fetch_function = "ListBuckets"
        object_name = "Bucket"
        object_plural_name = "Buckets"
        object_unique_id = "Name"
        object_response_field = "Buckets"
        model_only = false
        pagination = false
        use_post_processing = true
        excluded_fields = []
        convert_tags = false

        // GetBucketReplicationConfiguration
        child {
            object_name = "ReplicationConfiguration"
            new_field_name = "ReplicationConfiguration"
            field_type = "literal"
        }

        //GetBucketAcl
        child {
            object_name = "Grant"
            new_field_name = "AclGrants"
            field_type = "list"
        }

        // GetBucketCors
        child {
            object_name = "CORSRule"
            new_field_name = "CorsRules"
            field_type = "list"
        }

        // GetBucketEncryption
        child {
            object_name = "ServerSideEncryptionConfiguration"
            new_field_name = "ServerSideEncryptionConfiguration"
            field_type = "literal"
        }

        // GetBucketIntelligentTieringConfiguration
        child {
            object_name = "IntelligentTieringConfiguration"
            new_field_name = "IntelligentTieringConfigurations"
            field_type = "list"
        }

        // GetBucketInventoryConfiguration
        child {
            object_name = "InventoryConfiguration"
            new_field_name = "InventoryConfigurations"
            field_type = "list"
            excluded_fields = ["SSES3"]
        }

        // GetBucketLifecycleConfiguration
        child {
            object_name = "LifecycleRule"
            new_field_name = "LifecycleRules"
            field_type = "list"
        }

        // GetBucketLocation

        // GetBucketLogging
        child {
            object_name = "LoggingEnabled"
            new_field_name = "Logging"
            field_type = "literal"
        }

        // GetBucketPolicy
        extra_field {
            name = "Policy"
            type = "string"
        }

        // GetBucketPolicyStatus
        extra_field {
            name = "IsPublic"
            type = "bool"
        }

        // GetBucketTagging
        extra_field {
            name = "Tags"
            type = "map[string]string"
        }

        // GetBucketVersioning
        extra_field {
            name = "VersioningStatus"
            type = "string"
        }
        // GetBucketVersioning
        extra_field {
            name = "MFADeleteStatus"
            type = "string"
        }
    }
}