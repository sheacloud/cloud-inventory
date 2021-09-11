package codegen

//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name iam -data-source-name roles -primary-resource-name Role -primary-resource-field RoleId -api-function ListRoles -primary-object-path Roles -library-path "github.com/aws/aws-sdk-go-v2/service/iam" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name iam -data-source-name policies -primary-resource-name Policy -primary-resource-field PolicyId -api-function ListPolicies -primary-object-path Policies -library-path "github.com/aws/aws-sdk-go-v2/service/iam" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name iam -data-source-name users -primary-resource-name User -primary-resource-field UserId -api-function ListUsers -primary-object-path Users -library-path "github.com/aws/aws-sdk-go-v2/service/iam" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name iam -data-source-name groups -primary-resource-name Group -primary-resource-field GroupId -api-function ListGroups -primary-object-path Groups -library-path "github.com/aws/aws-sdk-go-v2/service/iam" -base-output-path ../pkg
//go:generate go run ../cmd/model-generator/ -cloud-name aws -service-name iam -data-source-name instance_profiles -primary-resource-name InstanceProfile -primary-resource-field InstanceProfileId -api-function ListInstanceProfiles -primary-object-path InstanceProfiles -library-path "github.com/aws/aws-sdk-go-v2/service/iam" -base-output-path ../pkg

//go:generate go fmt ../pkg/aws/iam/
