package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/sheacloud/cloud-inventory/codegen"
	"golang.org/x/tools/go/packages"
)

var (
	configFileName = flag.String("f", "", "the configuration filename")
	baseOutputPath = flag.String("o", "", "")
)

var primaryFields []*codegen.Field = []*codegen.Field{
	{
		Name: "AccountId",
		Type: "string",
		Tags: "`parquet:\"name=account_id, type=BYTE_ARRAY, convertedtype=UTF8\"`",
	},
	{
		Name: "Region",
		Type: "string",
		Tags: "`parquet:\"name=region, type=BYTE_ARRAY, convertedtype=UTF8\"`",
	},
	{
		Name: "ReportTime",
		Type: "int64",
		Tags: "`parquet:\"name=report_time, type=INT64, convertedtype=TIMESTAMP_MILLIS\"`",
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

func main() {
	flag.Parse()

	var config codegen.ServiceConfiguration
	err := hclsimple.DecodeFile(*configFileName, nil, &config)
	if err != nil {
		panic(err)
	}

	for _, datasourceConfig := range config.DataSources {
		servicePackagePath := datasourceConfig.LibraryPath
		serviceTypesPackagePath := datasourceConfig.LibraryPath + "/types"

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

		//object is a types.Object == types.TypeName, i.e. "type DescribeInstancesOutput struct {}..."
		object := apiPackage.Types.Scope().Lookup(datasourceConfig.PrimaryResourceName)

		models, err := codegen.NewModelsFromObject(object, datasourceConfig.PrimaryResourceName, datasourceConfig.ExcludedFields)
		if err != nil {
			panic(err)
		}

		models = codegen.DeduplicateModels(models)

		primaryModel := models[0]
		primaryModel.UpdatePrimaryResourceFieldTag(datasourceConfig.PrimaryResourceField)
		primaryModel.Fields = append(primaryModel.Fields, primaryFields...)

		for _, child := range datasourceConfig.Children {
			childObject := apiPackage.Types.Scope().Lookup(child.ResourceName)
			childModels, err := codegen.NewModelsFromObject(childObject, datasourceConfig.PrimaryResourceName, child.ExcludedFields)
			if err != nil {
				panic(err)
			}

			models = append(models, childModels...)

			var childType string
			switch child.ResourceType {
			case "literal":
				childType = "*" + childModels[0].Name
			case "list":
				childType = "[]*" + childModels[0].Name
			}

			field := &codegen.Field{
				Name: child.ResourceField,
				Type: childType,
			}
			field.PopulateTags()
			primaryModel.Fields = append(primaryModel.Fields, field)
		}

		models = codegen.DeduplicateModels(models)

		for _, extraField := range datasourceConfig.ExtraFields {
			field := &codegen.Field{
				Name: extraField.Name,
				Type: extraField.Type,
			}
			field.PopulateTags()
			primaryModel.Fields = append(primaryModel.Fields, field)
		}

		models = codegen.DeduplicateModels(models)

		datasourceFile := codegen.DatasourceFile{
			ServiceName:         datasourceConfig.Service,
			DataSourceName:      datasourceConfig.DataSource,
			PrimaryResourceName: datasourceConfig.PrimaryResourceName,
			ApiFunction:         datasourceConfig.ApiFunction,
			Models:              models,
			PrimaryModel:        primaryModel,
			PrimaryObjectPath:   datasourceConfig.PrimaryObjectPath,
			Paginate:            datasourceConfig.Paginate,
			ModelsOnly:          datasourceConfig.ModelsOnly,
		}

		outputPath := *baseOutputPath
		for _, dir := range []string{datasourceConfig.Cloud, datasourceConfig.Service} {
			outputPath += "/" + dir
			makeDir(outputPath)
		}
		outputPath += "/" + datasourceConfig.DataSource + "_model.go"

		outputFile, err := os.Create(outputPath)
		if err != nil {
			panic(err)
		}

		datasourceFile.Print(outputFile)

		outputFile.Close()

		fmt.Println("generated file " + outputPath)
	}

}
