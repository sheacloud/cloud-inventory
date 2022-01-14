package codegen

import (
	_ "embed"
	"strings"
	"text/template"
)

var (
	//go:embed templates/struct_source_code.tmpl
	structSourceCodeTemplateString string
	structSourceCodeTemplate       = template.Must(template.New("struct").Parse(structSourceCodeTemplateString))

	//go:embed templates/resource_file.tmpl
	resourceFileTemplateString string
	resourceFileTemplate       = template.Must(template.New("resourceFile").Parse(resourceFileTemplateString))

	//go:embed templates/referenced_resource_file.tmpl
	referencedResourceFileTemplateString string
	referencedResourceFileTemplate       = template.Must(template.New("referencedResourceFile").Parse(referencedResourceFileTemplateString))

	//go:embed templates/fetching_file.tmpl
	fetchingFileTemplateString string
	fetchingFileTemplate       = template.Must(template.New("fetchingFile").Parse(fetchingFileTemplateString))

	//go:embed templates/service_client_interface_file.tmpl
	serviceClientInterfaceFileTemplateString string
	serviceClientInterfaceFileTemplate       = template.Must(template.New("serviceClientInterfaceFile").Parse(serviceClientInterfaceFileTemplateString))

	//go:embed templates/client_interface_file.tmpl
	clientInterfaceFileTemplateString string
	clientInterfaceFileTemplate       = template.Must(template.New("clientInterfaceFile").Parse(clientInterfaceFileTemplateString))

	//go:embed templates/client_file.tmpl
	clientFileTemplateString string
	clientFileTemplate       = template.Must(template.New("clientFile").Parse(clientFileTemplateString))

	//go:embed templates/catalog_file.tmpl
	catalogFileTemplateString string
	catalogFileTemplate       = template.Must(template.New("catalogFile").Parse(catalogFileTemplateString))

	//go:embed templates/helpers_file.tmpl
	helpersFileTemplateString string
	helpersFileTemplate       = template.Must(template.New("helpersFile").Parse(helpersFileTemplateString))

	//go:embed templates/implemented_resources.tmpl
	implementedResourcesTemplateString string
	implementedResourcesTemplate       = template.Must(template.New("implementedResources").Parse(implementedResourcesTemplateString))
)

type AwsTemplate struct {
	Services []*AwsServiceConfig `hcl:"aws_service,block"`
}

func (t *AwsTemplate) GetClientInterfaceFileCode() string {
	var buf strings.Builder
	err := clientInterfaceFileTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsTemplate) GetClientFileCode() string {
	var buf strings.Builder
	err := clientFileTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsTemplate) GetCatalogFileCode() string {
	var buf strings.Builder
	err := catalogFileTemplate.Execute(&buf, t)
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

type AwsServiceTemplate struct {
	ServiceName       string
	TagObjectName     string
	SdkPath           string
	SdkClientName     string
	UtilizedFunctions []string
}

func (t *AwsServiceTemplate) GetServiceClientInterfaceFileCode() string {
	var buf strings.Builder
	err := serviceClientInterfaceFileTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsServiceTemplate) GetHelpersFileCode() string {
	var buf strings.Builder
	err := helpersFileTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

type AwsResourceTemplate struct {
	ServiceName       string
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
	err := resourceFileTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (t *AwsResourceTemplate) GetFetchingFileCode() string {
	var buf strings.Builder
	err := fetchingFileTemplate.Execute(&buf, t)
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

func (t *AwsReferencedResourceTemplate) GetReferencedResourceFileCode() string {
	var buf strings.Builder
	err := referencedResourceFileTemplate.Execute(&buf, t)
	if err != nil {
		panic(err)
	}
	return buf.String()
}
