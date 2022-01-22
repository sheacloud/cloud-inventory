aws_service "iam" {
    service_cap_name = "IAM"
    sdk_path = "github.com/aws/aws-sdk-go-v2/service/iam"
    extra_utilized_functions = ["ListAttachedGroupPolicies", "ListGroupPolicies", "GetGroup", "ListAttachedRolePolicies", "ListRolePolicies", "ListRoleTags", "ListAttachedUserPolicies", "ListUserPolicies", "ListUserTags", "GetLoginProfile", "ListAccessKeys", "ListGroupsForUser"]
    tag_object_name = "Tag"
    region_override = "us-east-1"

    resource "groups" {
        fetch_function = "ListGroups"
        object_name = "Group"
        object_plural_name = "Groups"
        object_unique_id = "GroupId"
        object_response_field = "Groups"
        model_only = false
        pagination = true
        use_post_processing = true
        excluded_fields = []
        convert_tags = false

        child {
            object_name = "AttachedPolicy"
            new_field_name = "AttachedPolicies"
            field_type = "list"
        }

        extra_field {
            name = "InlinePolicies"
            type = "[]string"
        }

        extra_field {
            name = "UserIds"
            type = "[]string"
        }
    }

    resource "policies" {
        fetch_function = "ListPolicies"
        object_name = "Policy"
        object_plural_name = "Policies"
        object_unique_id = "PolicyId"
        object_response_field = "Policies"
        model_only = true
        pagination = true
        use_post_processing = true
        excluded_fields = []
        convert_tags = false
    }

    resource "roles" {
        fetch_function = "ListRoles"
        object_name = "Role"
        object_plural_name = "Roles"
        object_unique_id = "RoleId"
        object_response_field = "Roles"
        model_only = false
        pagination = true
        use_post_processing = true
        excluded_fields = ["AssumeRolePolicyDocument"]
        convert_tags = true
        tag_field_name = "Tags"

        child {
            object_name = "AttachedPolicy"
            new_field_name = "AttachedPolicies"
            field_type = "list"
        }

        extra_field {
            name = "InlinePolicies"
            type = "[]string"
        }
    }

    resource "users" {
        fetch_function = "ListUsers"
        object_name = "User"
        object_plural_name = "Users"
        object_unique_id = "UserId"
        object_response_field = "Users"
        model_only = false
        pagination = true
        use_post_processing = true
        excluded_fields = []
        convert_tags = true
        tag_field_name = "Tags"

        child {
            object_name = "AccessKeyMetadata"
            new_field_name = "AccessKeys"
            field_type = "list"
            excluded_fields = ["UserName"]
        }

        child {
            object_name = "LoginProfile"
            new_field_name = "LoginProfile"
            field_type = "literal"
            excluded_fields = ["UserName"]
        }

        child {
            object_name = "AttachedPolicy"
            new_field_name = "AttachedPolicies"
            field_type = "list"
        }

        extra_field {
            name = "InlinePolicies"
            type = "[]string"
        }

        extra_field {
            name = "GroupIds"
            type = "[]string"
        }
    }
}