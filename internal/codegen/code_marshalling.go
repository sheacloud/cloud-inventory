package codegen

import (
	_ "embed"
	"strings"
	"text/template"
)

var (
	funcMap = template.FuncMap{
		"toSnakeCase":   ToSnakeCase,
		"snakeToHyphen": SnakeToHyphen,
	}

	//go:embed templates/struct_source_code.tmpl
	structSourceCodeTemplateString string
	structSourceCodeTemplate       = template.Must(template.New("struct").Funcs(funcMap).Parse(structSourceCodeTemplateString))

	//go:embed templates/aws_resource_file.tmpl
	awsResourceFileTemplateString string
	awsResourceFileTemplate       = template.Must(template.New("awsResourceFile").Funcs(funcMap).Parse(awsResourceFileTemplateString))

	//go:embed templates/aws_referenced_resource_file.tmpl
	awsReferencedResourceFileTemplateString string
	awsReferencedResourceFileTemplate       = template.Must(template.New("awsReferencedResourceFile").Funcs(funcMap).Parse(awsReferencedResourceFileTemplateString))

	//go:embed templates/aws_fetching_file.tmpl
	awsFetchingFileTemplateString string
	awsFetchingFileTemplate       = template.Must(template.New("awsFetchingFile").Funcs(funcMap).Parse(awsFetchingFileTemplateString))

	//go:embed templates/aws_service_client_interface_file.tmpl
	awsServiceClientInterfaceFileTemplateString string
	awsServiceClientInterfaceFileTemplate       = template.Must(template.New("awsServiceClientInterfaceFile").Funcs(funcMap).Parse(awsServiceClientInterfaceFileTemplateString))

	//go:embed templates/aws_service_inventory_file.tmpl
	awsServiceInventoryFileTemplateString string
	awsServiceInventoryFileTemplate       = template.Must(template.New("awsServiceInventoryFile").Funcs(funcMap).Parse(awsServiceInventoryFileTemplateString))

	//go:embed templates/aws_client_interface_file.tmpl
	awsClientInterfaceFileTemplateString string
	awsClientInterfaceFileTemplate       = template.Must(template.New("awsClientInterfaceFile").Funcs(funcMap).Parse(awsClientInterfaceFileTemplateString))

	//go:embed templates/aws_client_file.tmpl
	awsClientFileTemplateString string
	awsClientFileTemplate       = template.Must(template.New("awsClientFile").Funcs(funcMap).Parse(awsClientFileTemplateString))

	//go:embed templates/aws_catalog_file.tmpl
	awsCatalogFileTemplateString string
	awsCatalogFileTemplate       = template.Must(template.New("awsCatalogFile").Funcs(funcMap).Parse(awsCatalogFileTemplateString))

	//go:embed templates/aws_helpers_file.tmpl
	awsHelpersFileTemplateString string
	awsHelpersFileTemplate       = template.Must(template.New("awsHelpersFile").Funcs(funcMap).Parse(awsHelpersFileTemplateString))

	//go:embed templates/implemented_resources.tmpl
	implementedResourcesTemplateString string
	implementedResourcesTemplate       = template.Must(template.New("implementedResources").Funcs(funcMap).Parse(implementedResourcesTemplateString))

	//go:embed templates/aws_api_route.tmpl
	awsApiRouteTemplateString string
	awsApiRouteTemplate       = template.Must(template.New("awsApiRoute").Funcs(funcMap).Parse(awsApiRouteTemplateString))

	//go:embed templates/aws_router.tmpl
	awsRouterTemplateString string
	awsRouterTemplate       = template.Must(template.New("awsRouter").Funcs(funcMap).Parse(awsRouterTemplateString))

	//go:embed templates/aws_service_metadata_route.tmpl
	awsServiceMetadataRouteTemplateString string
	awsServiceMetadataRouteTemplate       = template.Must(template.New("awsServiceMetadataRoute").Funcs(funcMap).Parse(awsServiceMetadataRouteTemplateString))

	//go:embed templates/dao.tmpl
	daoTemplateString string
	daoTemplate       = template.Must(template.New("awsDao").Funcs(funcMap).Parse(daoTemplateString))

	//go:embed templates/mongo_dao.tmpl
	mongoDaoTemplateString string
	mongoDaoTemplate       = template.Must(template.New("mongoDao").Funcs(funcMap).Parse(mongoDaoTemplateString))

	//go:embed templates/dynamodb_dao.tmpl
	dynamodbDaoTemplateString string
	dynamodbDaoTemplate       = template.Must(template.New("dynamodbDao").Funcs(funcMap).Parse(dynamodbDaoTemplateString))

	//go:embed templates/s3_ion_dao.tmpl
	s3IonDaoTemplateString string
	s3IonDaoTemplate       = template.Must(template.New("s3IonDao").Funcs(funcMap).Parse(s3IonDaoTemplateString))

	//go:embed templates/s3_parquet_dao.tmpl
	s3ParquetDaoTemplateString string
	s3ParquetDaoTemplate       = template.Must(template.New("s3ParquetDao").Funcs(funcMap).Parse(s3ParquetDaoTemplateString))

	//go:embed templates/multi_dao.tmpl
	multiDaoTemplateString string
	multiDaoTemplate       = template.Must(template.New("multiDao").Funcs(funcMap).Parse(multiDaoTemplateString))
)

type AwsTemplate struct {
	Services []*AwsServiceConfig `hcl:"aws_service,block"`
}

func (t *AwsTemplate) GetClientInterfaceFileCode() string {
	var buf strings.Builder
	err := awsClientInterfaceFileTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsTemplate) GetAwsClientFileCode() string {
	var buf strings.Builder
	err := awsClientFileTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsTemplate) GetAwsCatalogFileCode() string {
	var buf strings.Builder
	err := awsCatalogFileTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsTemplate) GetImplementedResourcesCode() string {
	var buf strings.Builder
	err := implementedResourcesTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsTemplate) GetAwsRouterFileCode() string {
	var buf strings.Builder
	err := awsRouterTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsTemplate) GetDAOFileCode() string {
	var buf strings.Builder
	err := daoTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsTemplate) GetMongoDAOFileCode() string {
	var buf strings.Builder
	err := mongoDaoTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsTemplate) GetDynamoDBDAOFileCode() string {
	var buf strings.Builder
	err := dynamodbDaoTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsTemplate) GetS3IonDAOFileCode() string {
	var buf strings.Builder
	err := s3IonDaoTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsTemplate) GetS3ParquetDAOFileCode() string {
	var buf strings.Builder
	err := s3ParquetDaoTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsTemplate) GetMultiDAOFileCode() string {
	var buf strings.Builder
	err := multiDaoTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

type AwsServiceTemplate struct {
	ServiceName       string
	TagObjectName     string
	SdkPath           string
	SdkClientName     string
	UtilizedFunctions []string
	ServiceConfig     *AwsServiceConfig
}

func (t *AwsServiceTemplate) GetAwsServiceClientInterfaceFileCode() string {
	var buf strings.Builder
	err := awsServiceClientInterfaceFileTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsServiceTemplate) GetAwsHelpersFileCode() string {
	var buf strings.Builder
	err := awsHelpersFileTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsServiceTemplate) GetServiceMetadataRouteFileCode() string {
	var buf strings.Builder
	err := awsServiceMetadataRouteTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsServiceTemplate) GetAwsServiceInventoryFileCode() string {
	var buf strings.Builder
	err := awsServiceInventoryFileTemplate.Execute(&buf, t.ServiceConfig)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

type AwsResourceTemplate struct {
	ServiceName       string
	ServiceCapName    string
	ResourceStruct    *StructModel
	ResourceConfig    *AwsResourceConfig
	RequiredImports   []string
	SdkClientName     string
	ShouldConvertTags bool
	TagListFieldName  string
}

func (t *AwsResourceTemplate) DetermineRequiredImports() {
	t.RequiredImports = t.ResourceStruct.GetRequiredImports()
	t.RequiredImports = Deduplicate(t.RequiredImports)
}

func (t *AwsResourceTemplate) GetResourceFileCode() string {
	var buf strings.Builder
	err := awsResourceFileTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsResourceTemplate) GetAwsFetchingFileCode() string {
	var buf strings.Builder
	err := awsFetchingFileTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsResourceTemplate) GetAwsApiRouteCode() string {
	var buf strings.Builder
	err := awsApiRouteTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

type AwsReferencedResourceTemplate struct {
	ServiceName       string
	ReferencedStructs []*StructModel
	RequiredImports   []string
}

func (t *AwsReferencedResourceTemplate) DetermineRequiredImports() {
	t.RequiredImports = []string{}
	for _, structModel := range t.ReferencedStructs {
		t.RequiredImports = append(t.RequiredImports, structModel.GetRequiredImports()...)
	}
	t.RequiredImports = Deduplicate(t.RequiredImports)
}

func (t *AwsReferencedResourceTemplate) GetAwsReferencedResourceFileCode() string {
	var buf strings.Builder
	err := awsReferencedResourceFileTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
