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
	Label               string              `hcl:"name,label"`
	FetchFunction       string              `hcl:"fetch_function,attr"`           // the API function used to fetch the resource
	ObjectSourceName    string              `hcl:"object_source_name,attr"`       // the name of the object in the API
	ObjectSingularName  string              `hcl:"object_singular_name,optional"` // the singular name of the object
	ObjectPluralName    string              `hcl:"object_plural_name"`            // the plural name of the object
	ObjectUniqueId      string              `hcl:"object_unique_id,attr"`         // the name of a field in the object that uniquely identifies it
	ObjectResponseField string              `hcl:"object_response_field,attr"`    // the name of the object in the API response
	ModelOnly           bool                `hcl:"model_only,attr"`               // whether or not to only generate data models, not fetching code
	Pagination          bool                `hcl:"pagination,attr"`               // whether or not the API supports pagination
	UsePostProcessing   bool                `hcl:"use_post_processing,attr"`      // whether or not to use post processing
	ConvertTags         bool                `hcl:"convert_tags,attr"`             // whether or not to convert tags to map[string]string
	TagFieldName        string              `hcl:"tag_field_name,optional"`       // the name of the field that contains the tags
	DisplayFields       []string            `hcl:"display_fields,optional"`       // the fields to display when listing resources
	ExcludedFields      []string            `hcl:"excluded_fields,optional"`
	ExtraFields         []*ExtraFieldConfig `hcl:"extra_field,block"`
	Children            []*ChildConfig      `hcl:"child,block"`
}

func (c AwsResourceConfig) ObjectUniqueIdSnakeCase() string {
	return ToSnakeCase(c.ObjectUniqueId)
}

func (c AwsResourceConfig) ObjectSingularSnakeName() string {
	return ToSnakeCase((c.ObjectSingularName))
}

func (c AwsResourceConfig) ObjectPluralSnakeName() string {
	return ToSnakeCase(c.ObjectPluralName)
}

type ExtraFieldConfig struct {
	Name string `hcl:"name"`
	Type string `hcl:"type"`
}

type ChildConfig struct {
	ObjectSourceName string              `hcl:"object_source_name"`
	NewFieldName     string              `hcl:"new_field_name"`
	FieldType        string              `hcl:"field_type"`
	ExcludedFields   []string            `hcl:"excluded_fields,optional"`
	ExtraFields      []*ExtraFieldConfig `hcl:"extra_field,block"`
	Children         []*ChildConfig      `hcl:"child,block"`
}

var (
	defaultAwsResourceFields = []*FieldModel{
		{
			Name:        "AccountId",
			Type:        "string",
			IsTimeField: false,
		},
		{
			Name:        "Region",
			Type:        "string",
			IsTimeField: false,
		},
		{
			Name:        "ReportTime",
			Type:        "int64",
			IsTimeField: true,
		},
		{
			Name:        "InventoryUUID",
			Type:        "string",
			IsTimeField: false,
		},
	}
)

func GenerateAwsServiceCode(template *AwsTemplate) error {

	apiRoutesBaseDirectory := "./internal/api/routes/aws/"
	baseDirectory := "./pkg/aws/"
	interfacesDirectory := "./pkg/aws/interfaces/"
	mongoDaoBaseDirectory := "./internal/db/mongo/"
	dynamodbDaoBaseDirectory := "./internal/db/dynamodb/"
	s3ionDaoBaseDirectory := "./internal/db/s3ion/"
	s3parquetDaoBaseDirectory := "./internal/db/s3parquet/"
	multiDaoBaseDirectory := "./internal/db/multi/"
	inventoryBaseDirectory := "./internal/inventory/"

	for _, service := range template.Services {
		logrus.Info("Generating code for service " + service.Name)

		serviceDirectory := baseDirectory + service.Name + "/"
		makeDir(serviceDirectory)

		serviceApiRoutesDirectory := apiRoutesBaseDirectory + service.Name + "/"
		makeDir(serviceApiRoutesDirectory)

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
			if resource.ObjectSingularName == "" {
				resource.ObjectSingularName = resource.ObjectSourceName
			}
			if !stringInList(resource.FetchFunction, utilizedFunctions) {
				utilizedFunctions = append(utilizedFunctions, resource.FetchFunction)
			}

			var foundObject types.Object
			for _, pkg := range servicePackages {
				foundObject = pkg.Types.Scope().Lookup(resource.ObjectSourceName)
				if foundObject != nil {
					break
				}
			}
			if foundObject == nil {
				return errors.New("Could not find object " + resource.ObjectSourceName + " in any of the loaded packages")
			}

			//object is a types.Object == types.TypeName, i.e. "type DescribeInstancesOutput struct {}..."
			object := foundObject.(*types.TypeName).Type().(*types.Named)

			structModel, err := ParseStruct(object, resource.ExcludedFields, &resource.ObjectSingularName)
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

			structModel.ConvertTagFields()

			// determine the field tags
			structModel.PopulateFieldTags(resource.ObjectUniqueId)

			referencedStructs := structModel.GetReferencedStructs()
			for _, s := range referencedStructs {
				s.PopulateFieldTags(resource.ObjectUniqueId)
			}

			resourceStructs = append(resourceStructs, structModel)
			allReferencedStructs = append(allReferencedStructs, referencedStructs...)

			logrus.WithFields(logrus.Fields{
				"service": service.Name,
			}).Info("Analyzed source code for " + resource.ObjectPluralName)
		}

		allReferencedStructs = DeduplicateStructs(allReferencedStructs)

		for i, resource := range service.Resources {
			template := AwsResourceTemplate{
				ServiceName:       service.Name,
				ServiceCapName:    service.ServiceCapName,
				ResourceStruct:    resourceStructs[i],
				ResourceConfig:    resource,
				SdkClientName:     service.ServiceCapName,
				ShouldConvertTags: resource.ConvertTags,
				TagListFieldName:  resource.TagFieldName,
			}
			template.DetermineRequiredImports()

			// generate the resource model
			outputPath := serviceDirectory + "autogen_" + resource.ObjectPluralSnakeName() + "_model.go"
			outputFile, err := os.Create(outputPath)
			if err != nil {
				panic(err)
			}
			resourceFileCode := template.GetResourceFileCode()
			outputFile.WriteString(resourceFileCode)
			outputFile.Close()

			// generate the resource fetching code
			if !resource.ModelOnly {
				outputPath = serviceDirectory + "autogen_" + resource.ObjectPluralSnakeName() + "_fetch.go"
				outputFile, err = os.Create(outputPath)
				if err != nil {
					panic(err)
				}
				fetchFileCode := template.GetAwsFetchingFileCode()
				outputFile.WriteString(fetchFileCode)
				outputFile.Close()
			}

			// generate the API route code
			outputPath = serviceApiRoutesDirectory + "autogen_" + resource.ObjectPluralSnakeName() + "_route.go"
			outputFile, err = os.Create(outputPath)
			if err != nil {
				panic(err)
			}
			routeFileCode := template.GetAwsApiRouteCode()
			outputFile.WriteString(routeFileCode)
			outputFile.Close()

			logrus.WithFields(logrus.Fields{
				"service": service.Name,
			}).Info("Generated code for " + resource.ObjectPluralName)
		}

		// generate the service sub-models
		referencedTemplate := AwsReferencedResourceTemplate{
			ServiceName:       service.Name,
			ReferencedStructs: allReferencedStructs,
		}
		referencedTemplate.DetermineRequiredImports()

		outputPath := serviceDirectory + "autogen_referenced_models.go"
		outputFile, err := os.Create(outputPath)
		if err != nil {
			panic(err)
		}
		referencedResourceFileCode := referencedTemplate.GetAwsReferencedResourceFileCode()
		outputFile.WriteString(referencedResourceFileCode)
		outputFile.Close()

		// generate client interface code
		serviceTemplate := AwsServiceTemplate{
			ServiceName:       service.Name,
			TagObjectName:     service.TagObjectName,
			SdkPath:           service.SdkPath,
			SdkClientName:     service.ServiceCapName,
			UtilizedFunctions: utilizedFunctions,
			ServiceConfig:     service,
		}

		outputPath = interfacesDirectory + "autogen_" + service.Name + "_client_interface.go"
		outputFile, err = os.Create(outputPath)
		if err != nil {
			panic(err)
		}
		clientFileCode := serviceTemplate.GetAwsServiceClientInterfaceFileCode()
		outputFile.WriteString(clientFileCode)
		outputFile.Close()

		// generate the helper code
		if service.TagObjectName != "" {
			outputPath = serviceDirectory + "autogen_helper.go"
			outputFile, err = os.Create(outputPath)
			if err != nil {
				panic(err)
			}
			helperFileCode := serviceTemplate.GetAwsHelpersFileCode()
			outputFile.WriteString(helperFileCode)
			outputFile.Close()
		}

		// generate the metadata route code
		outputPath = serviceApiRoutesDirectory + "autogen_metadata_route.go"
		outputFile, err = os.Create(outputPath)
		if err != nil {
			panic(err)
		}
		metadataRouteFileCode := serviceTemplate.GetServiceMetadataRouteFileCode()
		outputFile.WriteString(metadataRouteFileCode)
		outputFile.Close()

		// generate the inventory code
		outputPath = inventoryBaseDirectory + "autogen_aws_" + service.Name + "_inventory.go"
		outputFile, err = os.Create(outputPath)
		if err != nil {
			panic(err)
		}
		serviceInventoryFileCode := serviceTemplate.GetAwsServiceInventoryFileCode()
		outputFile.WriteString(serviceInventoryFileCode)
		outputFile.Close()
	}

	// generate aws client interface code
	outputPath := baseDirectory + "autogen_client_interface.go"
	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	clientInterfaceCode := template.GetClientInterfaceFileCode()
	outputFile.WriteString(clientInterfaceCode)

	// generate aws client code
	outputPath = baseDirectory + "autogen_client.go"
	outputFile, err = os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	clientCode := template.GetAwsClientFileCode()
	outputFile.WriteString(clientCode)

	// generate catalog code
	outputPath = inventoryBaseDirectory + "autogen_aws_catalog.go"
	outputFile, err = os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	catalogCode := template.GetAwsCatalogFileCode()
	outputFile.WriteString(catalogCode)

	// generate the implemented resources markdown
	outputPath = "./IMPLEMENTED_RESOURCES.md"
	outputFile, err = os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	implementedResourcesCode := template.GetImplementedResourcesCode()
	outputFile.WriteString(implementedResourcesCode)

	// generate aws router code
	outputPath = apiRoutesBaseDirectory + "autogen_router.go"
	outputFile, err = os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	awsRouterCode := template.GetAwsRouterFileCode()
	outputFile.WriteString(awsRouterCode)
	outputFile.Close()

	// generate aws dao code
	outputPath = "./internal/db/autogen_dao.go"
	outputFile, err = os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	awsDaoCode := template.GetDAOFileCode()
	outputFile.WriteString(awsDaoCode)
	outputFile.Close()

	// generate mongo dao code
	outputPath = mongoDaoBaseDirectory + "autogen_dao.go"
	outputFile, err = os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	mongoDaoCode := template.GetMongoDAOFileCode()
	outputFile.WriteString(mongoDaoCode)
	outputFile.Close()

	// generate dynamodb dao code
	outputPath = dynamodbDaoBaseDirectory + "autogen_dao.go"
	outputFile, err = os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	dynamodbDaoCode := template.GetDynamoDBDAOFileCode()
	outputFile.WriteString(dynamodbDaoCode)
	outputFile.Close()

	// generate s3 ion dao code
	outputPath = s3ionDaoBaseDirectory + "autogen_dao.go"
	outputFile, err = os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	s3IonDaoCode := template.GetS3IonDAOFileCode()
	outputFile.WriteString(s3IonDaoCode)
	outputFile.Close()

	// generate s3 parquet dao code
	outputPath = s3parquetDaoBaseDirectory + "autogen_dao.go"
	outputFile, err = os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	s3ParquetDaoCode := template.GetS3ParquetDAOFileCode()
	outputFile.WriteString(s3ParquetDaoCode)
	outputFile.Close()

	// generate multi dao code
	outputPath = multiDaoBaseDirectory + "autogen_dao.go"
	outputFile, err = os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	multiDaoCode := template.GetMultiDAOFileCode()
	outputFile.WriteString(multiDaoCode)
	outputFile.Close()

	return nil
}

func PopulateChildStructs(structModel *StructModel, children []*ChildConfig, servicePackages []*packages.Package) {
	for _, child := range children {
		var foundObject types.Object
		for _, pkg := range servicePackages {
			foundObject = pkg.Types.Scope().Lookup(child.ObjectSourceName)
			if foundObject != nil {
				break
			}
		}
		if foundObject == nil {
			panic("Could not find object " + child.ObjectSourceName + " in any of the loaded packages")
		}

		//object is a types.Object == types.TypeName, i.e. "type DescribeInstancesOutput struct {}..."
		object := foundObject.(*types.TypeName).Type().(*types.Named)

		childStructModel, err := ParseStruct(object, child.ExcludedFields, nil)
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
