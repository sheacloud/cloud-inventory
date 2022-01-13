package codegen

import (
	"errors"
	"go/types"
	"os"

	"github.com/sirupsen/logrus"
	"golang.org/x/tools/go/packages"
)

type AwsServiceConfig struct {
	Name                   string               `hcl:"name,label"`
	ServiceCapName         string               `hcl:"service_cap_name,attr"`
	SdkPath                string               `hcl:"sdk_path,attr"`
	Resources              []*AwsResourceConfig `hcl:"resource,block"`
	ExtraUtilizedFunctions []string             `hcl:"extra_utilized_functions,optional"`
	RegionOverride         string               `hcl:"region_override,optional"`
	TagObjectName          string               `hcl:"tag_object_name,optional"` // the name of the object in the SDK which represents a tag
}

func (c *AwsServiceConfig) HasRegionOverride() bool {
	return c.RegionOverride != ""
}

type AwsResourceConfig struct {
	Name                string              `hcl:"name,label"`
	FetchFunction       string              `hcl:"fetch_function,attr"`        // the API function used to fetch the resource
	ObjectName          string              `hcl:"object_name,attr"`           // the name of the object in the API
	ObjectUniqueId      string              `hcl:"object_unique_id,attr"`      // the name of a field in the object that uniquely identifies it
	ObjectResponseField string              `hcl:"object_response_field,attr"` // the name of the object in the API response
	ModelOnly           bool                `hcl:"model_only,attr"`            // whether or not to only generate data models, not fetching code
	Pagination          bool                `hcl:"pagination,attr"`            // whether or not the API supports pagination
	UsePostProcessing   bool                `hcl:"use_post_processing,attr"`   // whether or not to use post processing
	ConvertTags         bool                `hcl:"convert_tags,attr"`          // whether or not to convert tags to map[string]string
	TagFieldName        string              `hcl:"tag_field_name,optional"`    // the name of the field that contains the tags
	ExcludedFields      []string            `hcl:"excluded_fields,optional"`
	ExtraFields         []*ExtraFieldConfig `hcl:"extra_field,block"`
	Children            []*ChildConfig      `hcl:"child,block"`
}

type ExtraFieldConfig struct {
	Name string `hcl:"name"`
	Type string `hcl:"type"`
}

type ChildConfig struct {
	ObjectName     string              `hcl:"object_name"`
	NewFieldName   string              `hcl:"new_field_name"`
	FieldType      string              `hcl:"field_type"`
	ExcludedFields []string            `hcl:"excluded_fields,optional"`
	ExtraFields    []*ExtraFieldConfig `hcl:"extra_field,block"`
	Children       []*ChildConfig      `hcl:"child,block"`
}

var (
	defaultAwsResourceFields = []*FieldModel{
		{
			Name: "AccountId",
			Type: "string",
		},
		{
			Name: "Region",
			Type: "string",
		},
		{
			Name:                 "ReportTime",
			Type:                 "int64",
			IsConvertedTimeField: true,
		},
	}
)

func GenerateAwsServiceCode(template *AwsTemplate, outputBaseDirectory string) error {

	for _, service := range template.Services {
		logrus.Info("Generating code for service " + service.Name)
		makeDir(outputBaseDirectory + "services/" + service.Name)

		var resourceStructs []*StructModel
		var allReferencedStructs []*StructModel
		utilizedFunctions := service.ExtraUtilizedFunctions

		var servicePackages []*packages.Package

		servicePackagePath := service.SdkPath
		serviceTypesPackagePath := service.SdkPath + "/types"

		pkgs, err := packages.Load(&packages.Config{
			Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedImports | packages.NeedDeps,
		}, servicePackagePath, serviceTypesPackagePath)
		if err != nil {
			panic(err)
		}
		packageMap := map[string]*packages.Package{}

		servicePackages = pkgs
		for _, pkg := range pkgs {
			pkgPath := pkg.PkgPath
			packageMap[pkgPath] = pkg
		}

		for _, resource := range service.Resources {
			utilizedFunctions = append(utilizedFunctions, resource.FetchFunction)

			var foundObject types.Object
			for _, pkg := range servicePackages {
				foundObject = pkg.Types.Scope().Lookup(resource.ObjectName)
				if foundObject != nil {
					break
				}
			}
			if foundObject == nil {
				return errors.New("Could not find object " + resource.ObjectName + " in any of the loaded packages")
			}

			//object is a types.Object == types.TypeName, i.e. "type DescribeInstancesOutput struct {}..."
			object := foundObject.(*types.TypeName).Type().(*types.Named)

			structModel, err := ParseStruct(object, resource.ExcludedFields)
			if err != nil {
				return err
			}
			structModel.Fields = append(structModel.Fields, defaultAwsResourceFields...)

			// Add extra fields
			for _, extraField := range resource.ExtraFields {
				structModel.Fields = append(structModel.Fields, &FieldModel{
					Name: extraField.Name,
					Type: extraField.Type,
				})
			}

			// Add children
			PopulateChildStructs(structModel, resource.Children, servicePackages)

			// Create any time field conversions
			structModel.AddConvertedTimeFields()

			structModel.ConvertTagFields()

			// determine the field tags
			structModel.PopulateFieldTags(resource.ObjectUniqueId)

			referencedStructs := structModel.GetReferencedStructs()
			for _, s := range referencedStructs {
				s.AddConvertedTimeFields()
				s.PopulateFieldTags(resource.ObjectUniqueId)
			}

			resourceStructs = append(resourceStructs, structModel)
			allReferencedStructs = append(allReferencedStructs, referencedStructs...)

			logrus.WithFields(logrus.Fields{
				"service": service.Name,
			}).Info("Analyzed source code for " + resource.Name)
		}

		allReferencedStructs = DeduplicateStructs(allReferencedStructs)

		for i, resource := range service.Resources {
			template := AwsResourceTemplate{
				ServiceName:       service.Name,
				ResourceStruct:    resourceStructs[i],
				ResourceConfig:    resource,
				SdkClientName:     service.ServiceCapName,
				ShouldConvertTags: resource.ConvertTags,
				TagListFieldName:  resource.TagFieldName,
			}
			template.DetermineRequiredImports()

			outputPath := outputBaseDirectory + "services/" + service.Name + "/autogen_" + resource.Name + "_model.go"
			outputFile, err := os.Create(outputPath)
			if err != nil {
				panic(err)
			}
			resourceFileCode := template.GetResourceFileCode()
			outputFile.WriteString(resourceFileCode)
			outputFile.Close()

			if !resource.ModelOnly {
				outputPath = outputBaseDirectory + "services/" + service.Name + "/autogen_" + resource.Name + "_fetch.go"
				outputFile, err = os.Create(outputPath)
				if err != nil {
					panic(err)
				}
				fetchFileCode := template.GetFetchingFileCode()
				outputFile.WriteString(fetchFileCode)
				outputFile.Close()
			}

			logrus.WithFields(logrus.Fields{
				"service": service.Name,
			}).Info("Generated code for " + resource.Name)

		}

		referencedTemplate := AwsReferencedResourceTemplate{
			ServiceName:       service.Name,
			ReferencedStructs: allReferencedStructs,
		}
		referencedTemplate.DetermineRequiredImports()

		outputPath := outputBaseDirectory + "services/" + service.Name + "/" + "autogen_referenced_models.go"
		outputFile, err := os.Create(outputPath)
		if err != nil {
			panic(err)
		}
		referencedResourceFileCode := referencedTemplate.GetReferencedResourceFileCode()
		outputFile.WriteString(referencedResourceFileCode)
		outputFile.Close()

		// generate client interface code
		serviceTemplate := AwsServiceTemplate{
			ServiceName:       service.Name,
			TagObjectName:     service.TagObjectName,
			SdkPath:           service.SdkPath,
			SdkClientName:     service.ServiceCapName,
			UtilizedFunctions: utilizedFunctions,
		}

		outputPath = outputBaseDirectory + "interfaces/" + "autogen_" + service.Name + "_client.go"
		outputFile, err = os.Create(outputPath)
		if err != nil {
			panic(err)
		}
		clientFileCode := serviceTemplate.GetServiceClientInterfaceFileCode()
		outputFile.WriteString(clientFileCode)
		outputFile.Close()

		// generate the helper code
		if service.TagObjectName != "" {
			outputPath = outputBaseDirectory + "services/" + service.Name + "/autogen_helper.go"
			outputFile, err = os.Create(outputPath)
			if err != nil {
				panic(err)
			}
			helperFileCode := serviceTemplate.GetHelpersFileCode()
			outputFile.WriteString(helperFileCode)
			outputFile.Close()
		}
	}

	// generate client interface code
	outputPath := outputBaseDirectory + "interfaces/autogen_client.go"
	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	clientInterfaceCode := template.GetClientInterfaceFileCode()
	outputFile.WriteString(clientInterfaceCode)

	// generate client code
	outputPath = outputBaseDirectory + "autogen_client.go"
	outputFile, err = os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	clientCode := template.GetClientFileCode()
	outputFile.WriteString(clientCode)

	// generate catalog code
	outputPath = "./internal/inventory/autogen_aws_catalog.go"
	outputFile, err = os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	catalogCode := template.GetCatalogFileCode()
	outputFile.WriteString(catalogCode)

	return nil
}

func PopulateChildStructs(structModel *StructModel, children []*ChildConfig, servicePackages []*packages.Package) {
	for _, child := range children {
		var foundObject types.Object
		for _, pkg := range servicePackages {
			foundObject = pkg.Types.Scope().Lookup(child.ObjectName)
			if foundObject != nil {
				break
			}
		}
		if foundObject == nil {
			panic("Could not find object " + child.ObjectName + " in any of the loaded packages")
		}

		//object is a types.Object == types.TypeName, i.e. "type DescribeInstancesOutput struct {}..."
		object := foundObject.(*types.TypeName).Type().(*types.Named)

		childStructModel, err := ParseStruct(object, child.ExcludedFields)
		if err != nil {
			panic(err)
		}

		// Add extra fields
		for _, extraField := range child.ExtraFields {
			childStructModel.Fields = append(childStructModel.Fields, &FieldModel{
				Name: extraField.Name,
				Type: extraField.Type,
			})
		}

		var childType string
		switch child.FieldType {
		case "literal":
			childType = "*" + childStructModel.Name
		case "list":
			childType = "[]*" + childStructModel.Name
		}

		structModel.Fields = append(structModel.Fields, &FieldModel{
			Name:              child.NewFieldName,
			Type:              childType,
			ReferencedStructs: []*StructModel{childStructModel},
		})

		// recusively populate child structs
		PopulateChildStructs(childStructModel, child.Children, servicePackages)
	}
}
