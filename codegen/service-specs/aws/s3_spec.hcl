service "aws" "s3" {
    library_path = "github.com/aws/aws-sdk-go-v2/service/s3"

    datasource "buckets" {
        primary_object_name = "Bucket"
        primary_object_field = "Name"
        api_function = "ListBuckets"
        primary_object_path = ["Buckets"]
        models_only = false
        paginate = false

        // GetBucketReplicationConfiguration
        child {
            resource_name = "ReplicationConfiguration"
            resource_field = "ReplicationConfiguration"
            resource_type = "literal"
        }

        //GetBucketAcl
        child {
            resource_name = "Grant"
            resource_field = "AclGrants"
            resource_type = "list"
        }

        // GetBucketCors
        child {
            resource_name = "CORSRule"
            resource_field = "CorsRules"
            resource_type = "list"
        }

        // GetBucketEncryption
        child {
            resource_name = "ServerSideEncryptionConfiguration"
            resource_field = "ServerSideEncryptionConfiguration"
            resource_type = "literal"
        }

        // GetBucketIntelligentTieringConfiguration
        child {
            resource_name = "IntelligentTieringConfiguration"
            resource_field = "IntelligentTieringConfigurations"
            resource_type = "list"
        }

        // GetBucketInventoryConfiguration
        child {
            resource_name = "InventoryConfiguration"
            resource_field = "InventoryConfigurations"
            resource_type = "list"
            excluded_fields = ["SSES3"]
        }

        // GetBucketLifecycleConfiguration
        child {
            resource_name = "LifecycleRule"
            resource_field = "LifecycleRules"
            resource_type = "list"
        }

        // GetBucketLocation

        // GetBucketLogging
        child {
            resource_name = "LoggingEnabled"
            resource_field = "Logging"
            resource_type = "literal"
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