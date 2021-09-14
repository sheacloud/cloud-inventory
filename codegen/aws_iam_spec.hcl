datasource "aws" "iam" "roles" {
    primary_resource_name = "Role"
    primary_resource_field = "RoleId"
    api_function = "ListRoles"
    primary_object_path = ["Roles"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/iam"
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
}

datasource "aws" "iam" "policies" {
    primary_resource_name = "Policy"
    primary_resource_field = "PolicyId"
    api_function = "ListPolicies"
    primary_object_path = ["Policies"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/iam"
    models_only = false
    paginate = true

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

datasource "aws" "iam" "users" {
    primary_resource_name = "User"
    primary_resource_field = "UserId"
    api_function = "ListUsers"
    primary_object_path = ["Users"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/iam"
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
}

datasource "aws" "iam" "groups" {
    primary_resource_name = "Group"
    primary_resource_field = "GroupId"
    api_function = "ListGroups"
    primary_object_path = ["Groups"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/iam"
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

datasource "aws" "iam" "instance_profiles" {
    primary_resource_name = "InstanceProfile"
    primary_resource_field = "InstanceProfileId"
    api_function = "ListInstanceProfiles"
    primary_object_path = ["InstanceProfiles"]
    library_path = "github.com/aws/aws-sdk-go-v2/service/iam"
    models_only = false
    paginate = true
}