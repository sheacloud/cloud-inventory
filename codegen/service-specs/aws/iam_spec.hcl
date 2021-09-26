service "aws" "iam" {
    library_path = "github.com/aws/aws-sdk-go-v2/service/iam"

    datasource "roles" {
        primary_object_name = "Role"
        primary_object_field = "RoleId"
        api_function = "ListRoles"
        primary_object_path = ["Roles"]
        models_only = false
        paginate = true

        // ListAttachedRolePolicies
        child {
            resource_name = "AttachedPolicy"
            resource_field = "AttachedPolicies"
            resource_type = "list"
        }

        // ListRolePolicies
        extra_field {
            name = "InlinePolicies"
            type = "[]string"
        }

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "policies" {
        primary_object_name = "Policy"
        primary_object_field = "PolicyId"
        api_function = "ListPolicies"
        primary_object_path = ["Policies"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }

        # // ListEntitiesForPolicy
        # child {
        #     resource_name = "PolicyGroup"
        #     resource_field = "Groups"
        #     resource_type = "list"
        # }

        # child {
        #     resource_name = "PolicyRole"
        #     resource_field = "Roles"
        #     resource_type = "list"
        # }

        # child {
        #     resource_name = "PolicyUser"
        #     resource_field = "Users"
        #     resource_type = "list"
        # }
    }

    datasource "users" {
        primary_object_name = "User"
        primary_object_field = "UserId"
        api_function = "ListUsers"
        primary_object_path = ["Users"]
        models_only = false
        paginate = true

        child {
            resource_name = "AccessKeyMetadata"
            resource_field = "AccessKeys"
            resource_type = "list"
            excluded_fields = ["UserName"]
        }

        child {
            resource_name = "LoginProfile"
            resource_field = "LoginProfile"
            resource_type = "literal"
            excluded_fields = ["UserName"]
        }

        // ListAttachedUserPolicies
        child {
            resource_name = "AttachedPolicy"
            resource_field = "AttachedPolicies"
            resource_type = "list"
        }

        // ListUserPolicies
        extra_field {
            name = "InlinePolicies"
            type = "[]string"
        }

        // ListGroupsForUser
        extra_field {
            name = "GroupIds"
            type = "[]string"
        }

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }

    datasource "groups" {
        primary_object_name = "Group"
        primary_object_field = "GroupId"
        api_function = "ListGroups"
        primary_object_path = ["Groups"]
        models_only = false
        paginate = true

        // ListAttachedGroupPolicies
        child {
            resource_name = "AttachedPolicy"
            resource_field = "AttachedPolicies"
            resource_type = "list"
        }

        // ListGroupPolicies
        extra_field {
            name = "InlinePolicies"
            type = "[]string"
        }

        // GetGroup
        extra_field {
            name = "UserIds"
            type = "[]string"
        }
    }

    datasource "instance_profiles" {
        primary_object_name = "InstanceProfile"
        primary_object_field = "InstanceProfileId"
        api_function = "ListInstanceProfiles"
        primary_object_path = ["InstanceProfiles"]
        models_only = false
        paginate = true

        field_conversion {
            source_field_name = "Tags"
            source_field_name_override = "TagsOld"
            target_field_name = "Tags"
            target_field_type = "map[string]string"
            conversion_function_name = "GetTagMap"
        }
    }
}