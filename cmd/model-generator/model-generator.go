package main

import (
	"os"

	"github.com/sheacloud/cloud-inventory/codegen"
	"golang.org/x/tools/go/packages"
)

var primaryFields []codegen.Field = []codegen.Field{
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

func main() {
	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedImports | packages.NeedDeps,
	}, "github.com/aws/aws-sdk-go-v2/service/ec2", "github.com/aws/aws-sdk-go-v2/service/ec2/types")
	if err != nil {
		panic(err)
	}
	packageMap := map[string]*packages.Package{}

	for _, pkg := range pkgs {
		pkgPath := pkg.PkgPath
		packageMap[pkgPath] = pkg
	}

	ec2Package := packageMap["github.com/aws/aws-sdk-go-v2/service/ec2/types"]

	// fmt.Println(ec2Package.PkgPath)

	//object is a types.Object == types.TypeName, i.e. "type DescribeInstancesOutput struct {}..."
	object := ec2Package.Types.Scope().Lookup("Instance")
	// fmt.Printf("%v: %v\n", object, reflect.TypeOf(object))

	models, err := codegen.NewModelsFromObject(object)
	if err != nil {
		panic(err)
	}

	models = codegen.DeduplicateModels(models)

	primaryModel := models[0]
	primaryModel.Fields = append(primaryModel.Fields, primaryFields...)

	datasourceFile := codegen.DatasourceFile{
		ServiceName:       "ec2",
		DatasourceName:    "instances",
		DatasourceNameCap: "Instance",
		ApiFunction:       "DescribeInstances",
		Models:            models,
		PrimaryObjectPath: []string{"Reservations", "Instances"},
	}

	datasourceFile.Print(os.Stdout)

}
