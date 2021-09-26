package main

import (
	"flag"
	"fmt"
	"go/types"
	"os"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/sheacloud/cloud-inventory/internal/codegen"
	"golang.org/x/tools/go/packages"
)

var (
	configFileName = flag.String("f", "", "the configuration filename")
	baseOutputPath = flag.String("o", "", "")
)

var primaryFields []*codegen.ParquetModelField = []*codegen.ParquetModelField{
	{
		Name: "AccountId",
		Type: "string",
	},
	{
		Name: "Region",
		Type: "string",
	},
	{
		Name:        "ReportTime",
		Type:        "int64",
		IsTimeField: true,
	},
}

func makeDir(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, 0755)
		if err != nil {
			panic(err)
		}
	}
}

type FileConfiguration struct {
	ServiceConfigurations []ServiceConfiguration `hcl:"service,block"`
}

type ServiceConfiguration struct {
	Cloud       string       `hcl:"cloud,label"`
	Service     string       `hcl:"service,label"`
	LibraryPath string       `hcl:"library_path"`
	DataSources []DataSource `hcl:"datasource,block"`
}

type DataSource struct {
	DataSource         string                 `hcl:"datasource,label"`
	PrimaryObjectName  string                 `hcl:"primary_object_name"`
	PrimaryObjectField string                 `hcl:"primary_object_field"`
	ApiFunction        string                 `hcl:"api_function"`
	PrimaryObjectPath  []string               `hcl:"primary_object_path"`
	Paginate           bool                   `hcl:"paginate"`
	Children           []DataSourceChild      `hcl:"child,block"`
	ExtraFields        []DataSourceExtraField `hcl:"extra_field,block"`
	ModelsOnly         bool                   `hcl:"models_only"`
	ExcludedFields     []string               `hcl:"excluded_fields,optional"`
	FieldConversions   []FieldConversion      `hcl:"field_conversion,block"`
}

type DataSourceChild struct {
	ResourceName   string                 `hcl:"resource_name"`
	ResourceField  string                 `hcl:"resource_field"`
	ResourceType   string                 `hcl:"resource_type"`
	ExcludedFields []string               `hcl:"excluded_fields,optional"`
	ExtraFields    []DataSourceExtraField `hcl:"extra_field,block"`
	Children       []DataSourceChild      `hcl:"child,block"`
}

type DataSourceExtraField struct {
	Name string `hcl:"name"`
	Type string `hcl:"type"`
}

type FieldConversion struct {
	SourceFieldName         string  `hcl:"source_field_name"`
	SourceFieldNameOverride *string `hcl:"source_field_name_override,optional"`
	TargetFieldName         string  `hcl:"target_field_name"`
	TargetFieldType         string  `hcl:"target_field_type"`
	ConversionFunctionName  string  `hcl:"conversion_function_name"`
}

func populateChild(parentModel *codegen.ParquetModelStruct, child DataSourceChild, apiPackage *packages.Package, primaryObjectName string) []*codegen.ParquetModelStruct {
	childObject := apiPackage.Types.Scope().Lookup(child.ResourceName).(*types.TypeName).Type().(*types.Named)
	childModel, err := codegen.NewParquetModelStruct(childObject, primaryObjectName, child.ExcludedFields, nil)
	if err != nil {
		panic(err)
	}

	for _, extraField := range child.ExtraFields {
		childModel.Fields = append(childModel.Fields, &codegen.ParquetModelField{
			Name: extraField.Name,
			Type: extraField.Type,
		})
	}

	var childType string
	switch child.ResourceType {
	case "literal":
		childType = "*" + childModel.Name
	case "list":
		childType = "[]*" + childModel.Name
	}

	parentModel.Fields = append(parentModel.Fields, &codegen.ParquetModelField{
		Name: child.ResourceField,
		Type: childType,
	})
	models := []*codegen.ParquetModelStruct{childModel}
	models = append(models, childModel.GetChildModels()...)

	for _, grandChild := range child.Children {
		grandChildModels := populateChild(childModel, grandChild, apiPackage, primaryObjectName)
		models = append(models, grandChildModels...)
	}

	return models
}

func main() {
	flag.Parse()

	var config FileConfiguration
	err := hclsimple.DecodeFile(*configFileName, nil, &config)
	if err != nil {
		panic(err)
	}

	for _, serviceConfig := range config.ServiceConfigurations {
		servicePackagePath := serviceConfig.LibraryPath
		serviceTypesPackagePath := serviceConfig.LibraryPath + "/types"

		pkgs, err := packages.Load(&packages.Config{
			Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedImports | packages.NeedDeps,
		}, servicePackagePath, serviceTypesPackagePath)
		if err != nil {
			panic(err)
		}
		packageMap := map[string]*packages.Package{}

		for _, pkg := range pkgs {
			pkgPath := pkg.PkgPath
			packageMap[pkgPath] = pkg
		}

		apiPackage := packageMap[serviceTypesPackagePath]

		for _, datasourceConfig := range serviceConfig.DataSources {

			//object is a types.Object == types.TypeName, i.e. "type DescribeInstancesOutput struct {}..."
			object := apiPackage.Types.Scope().Lookup(datasourceConfig.PrimaryObjectName).(*types.TypeName).Type().(*types.Named)

			fieldConversions := []*codegen.ParquetModelFieldConversion{}
			for _, conversion := range datasourceConfig.FieldConversions {
				fieldConversions = append(fieldConversions, &codegen.ParquetModelFieldConversion{
					SourceFieldName:         conversion.SourceFieldName,
					SourceFieldNameOverride: conversion.SourceFieldNameOverride,
					ConversionFunctionName:  conversion.ConversionFunctionName,
					TargetField: codegen.ParquetModelField{
						Name: conversion.TargetFieldName,
						Type: conversion.TargetFieldType,
					},
				})
			}
			primaryModel, err := codegen.NewParquetModelStruct(object, datasourceConfig.PrimaryObjectName, datasourceConfig.ExcludedFields, fieldConversions)
			if err != nil {
				panic(err)
			}

			models := []*codegen.ParquetModelStruct{primaryModel}
			models = append(models, primaryModel.GetChildModels()...)
			primaryModel.Fields = append(primaryModel.Fields, primaryFields...)

			for _, extraField := range datasourceConfig.ExtraFields {
				primaryModel.Fields = append(primaryModel.Fields, &codegen.ParquetModelField{
					Name: extraField.Name,
					Type: extraField.Type,
				})
			}

			for _, child := range datasourceConfig.Children {
				childModels := populateChild(primaryModel, child, apiPackage, datasourceConfig.PrimaryObjectName)
				models = append(models, childModels...)
			}

			models = codegen.DeduplicateModels(models)

			for _, model := range models {
				model.PopulateFieldTags(datasourceConfig.PrimaryObjectField)
			}

			datasourceFile := codegen.AwsDataSourceFile{
				ServiceName:       serviceConfig.Service,
				DataSourceName:    datasourceConfig.DataSource,
				PrimaryObjectName: datasourceConfig.PrimaryObjectName,
				ApiFunction:       datasourceConfig.ApiFunction,
				Models:            models,
				PrimaryModel:      primaryModel,
				PrimaryObjectPath: datasourceConfig.PrimaryObjectPath,
				Paginate:          datasourceConfig.Paginate,
				ModelsOnly:        datasourceConfig.ModelsOnly,
			}

			outputPath := *baseOutputPath
			for _, dir := range []string{serviceConfig.Cloud, serviceConfig.Service} {
				outputPath += "/" + dir
				makeDir(outputPath)
			}
			outputPath += "/" + datasourceConfig.DataSource + "_model.go"

			outputFile, err := os.Create(outputPath)
			if err != nil {
				panic(err)
			}

			outputFile.WriteString(datasourceFile.SourceCode())

			outputFile.Close()

			fmt.Println("generated file " + outputPath)
		}
	}

}
