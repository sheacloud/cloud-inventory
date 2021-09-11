package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/sheacloud/cloud-inventory/codegen"
	"golang.org/x/tools/go/packages"
)

var (
	cloudName            = flag.String("cloud-name", "", "the cloud for the data source (i.e. AWS, Azure)")
	serviceName          = flag.String("service-name", "", "name of the cloud service (i.e. ec2, s3)")
	dataSourceName       = flag.String("data-source-name", "", "name of the service datasource (i.e. instances, buckets)")
	primaryResourcename  = flag.String("primary-resource-name", "", "name of the primary resource in the datasource (i.e. Instance, Bucket)")
	primaryResourceField = flag.String("primary-resource-field", "", "the unique ID field of the primary resource (i.e. InstanceId, Arn)")
	apiFunction          = flag.String("api-function", "", "name of the API function in the Go SDK for the datasource (i.e. DescribeInstances)")
	primaryObjectPath    = flag.String("primary-object-path", "", "Path of objects to reach the primary object, comma separated")
	libraryPath          = flag.String("library-path", "", "path to the module of the Go SDK containing the datasource function")
	baseOutputPath       = flag.String("base-output-path", "./pkg", "the path to write the files to")
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
	servicePackagePath := *libraryPath
	serviceTypesPackagePath := *libraryPath + "/types"

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

	ec2Package := packageMap[serviceTypesPackagePath]

	//object is a types.Object == types.TypeName, i.e. "type DescribeInstancesOutput struct {}..."
	object := ec2Package.Types.Scope().Lookup(*primaryResourcename)

	models, err := codegen.NewModelsFromObject(object, *primaryResourcename)
	if err != nil {
		panic(err)
	}

	models = codegen.DeduplicateModels(models)

	primaryModel := models[0]
	primaryModel.UpdatePrimaryResourceFieldTag(*primaryResourceField)
	primaryModel.Fields = append(primaryModel.Fields, primaryFields...)

	datasourceFile := codegen.DatasourceFile{
		ServiceName:         *serviceName,
		DataSourceName:      *dataSourceName,
		PrimaryResourceName: *primaryResourcename,
		ApiFunction:         *apiFunction,
		Models:              models,
		PrimaryModel:        primaryModel,
		PrimaryObjectPath:   strings.Split(*primaryObjectPath, ","),
	}

	outputPath := *baseOutputPath
	for _, dir := range []string{*cloudName, *serviceName} {
		outputPath += "/" + dir
		makeDir(outputPath)
	}
	outputPath += "/" + *dataSourceName + "_model.go"

	outputFile, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}

	datasourceFile.Print(outputFile)

	outputFile.Close()

	fmt.Println("generated file " + outputPath)
}
