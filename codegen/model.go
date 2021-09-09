package codegen

import (
	"fmt"
	"go/types"
	"io"
	"regexp"
	"strings"
	"text/template"
)

type Model struct {
	Name   string
	Fields []Field
}

type Field struct {
	Name string
	Type string
	Tags string
}

type DatasourceFile struct {
	ServiceName       string
	DatasourceName    string
	DatasourceNameCap string
	ApiFunction       string
	Models            []*Model
	PrimaryObjectPath []string
}

var funcMap template.FuncMap = template.FuncMap{
	// The name "inc" is what the function will be called in the template text.
	"sub": func(i, j int) int {
		return i - j
	},
	"tab": func(i int) string {
		result := ""
		for k := 0; k < i; k++ {
			result += "	"
		}
		return result
	},
	"len": func(s []string) int {
		return len(s)
	},
}

var datasourceTemplate string = `// AUTOGENERED, DO NOT EDIT
package {{.ServiceName}}

import (
	"context"
	"fmt"
	"time"
	
	"github.com/aws/aws-sdk-go-v2/service/{{.ServiceName}}"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

func init() {
	Controller.RegisterDataSource("{{.DatasourceName}}", {{.DatasourceNameCap}}DataSource)
}

{{range $index, $element := .Models}}
type {{$element.Name}} struct {
{{range $subindex, $subelement := $element.Fields}}	{{$subelement.Name}} {{$subelement.Type}} {{$subelement.Tags}}
{{end}}
}	
{{end}}
{{$pathLength := len .PrimaryObjectPath}}

func {{.DatasourceNameCap}}DataSource(ctx context.Context, client *{{.ServiceName}}.Client, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new({{.DatasourceNameCap}}Model))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	paginator := {{.ServiceName}}.New{{.ApiFunction}}Paginator(client, &{{.ServiceName}}.{{.ApiFunction}}Input{})

	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"service":     storageConfig.Service,
				"data_source": storageConfig.DataSource,
				"account_id":  storageConfig.AccountId,
				"region":      storageConfig.Region,
				"cloud":       storageConfig.Cloud,
				"error":       err,
			}).Error("error calling {{.ApiFunction}}")
			return err
		}

		{{range $index, $element := .PrimaryObjectPath}}{{tab $index}}for _, var{{$index}} := range {{if eq $index 0}}output{{else}}var{{sub $index 1}}{{end}}.{{$element}} {
		{{end}}
		{{tab $pathLength}}model := new({{.DatasourceNameCap}}Model)
		{{tab $pathLength}}copier.Copy(&model, &var{{sub $pathLength 1}})

		{{tab $pathLength}}//model.Tags = GetTagMap(var{{sub $pathLength 1}}.Tags)
		{{tab $pathLength}}//model.LaunchTimeMillis = model.LaunchTime.UTC().UnixMilli()
		{{tab $pathLength}}model.AccountId = storageConfig.AccountId
		{{tab $pathLength}}model.Region = storageConfig.Region
		{{tab $pathLength}}model.ReportTime = reportTime.UTC().UnixMilli()

		{{tab $pathLength}}errors := storageContextSet.Store(ctx, model)
		{{tab $pathLength}}for storageContext, err := range errors {
		{{tab $pathLength}}	storage.LogContextError(storageContext, fmt.Sprintf("Error storing {{.DatasourceNameCap}}Model: %v", err))
		{{tab $pathLength}}}
		{{range $index, $element := .PrimaryObjectPath}}{{tab (sub (sub $pathLength 1) $index)}}}
		{{end}}
	}

	return nil
}
`

func (d DatasourceFile) Print(out io.Writer) {
	tmpl, err := template.New("test").Funcs(funcMap).Parse(datasourceTemplate)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(out, d)
	if err != nil {
		panic(err)
	}
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func (f *Field) PopulateTags() {
	parquetName := ToSnakeCase(f.Name)

	var extraParquetTags string
	switch f.Type {
	case "string":
		extraParquetTags = ",type=BYTE_ARRAY,convertedtype=UTF8"
	case "int32":
		extraParquetTags = ",type=INT32"
	case "bool":
		extraParquetTags = ",type=BOOLEAN"
	default:
		if strings.HasPrefix(f.Type, "[]") {
			extraParquetTags = ",type=LIST"
		}
	}

	f.Tags = fmt.Sprintf("`parquet:\"name=%s%s\"`", parquetName, extraParquetTags)
}

func getTypeName(t types.Type) string {
	var fieldType string
	switch v := t.(type) {
	default:
		fieldType = "unknown"
	case *types.Basic:
		fieldType = v.Name()
	case *types.Struct:
		fieldType = "struct"
	case *types.Named:
		switch k := v.Underlying().(type) {
		default:
			fieldType = getTypeName(k)
		case *types.Struct:
			fieldType = v.Obj().Name() + "Model"
		}
	case *types.Interface:
		fieldType = "interface"
	case *types.Signature:
		fieldType = "signature"
	case *types.Array:
		fieldType = fmt.Sprintf("[%v]", v.Len())
	case *types.Slice:
		fieldType = "[]"
		elemType := v.Elem()
		fieldType += getTypeName(elemType)
	}

	return fieldType
}

func getFieldModels(t types.Type) ([]*Model, error) {
	switch v := t.(type) {
	default:
		return nil, nil
	case *types.Slice:
		return getFieldModels(v.Elem())
	case *types.Named:
		switch v.Underlying().(type) {
		default:
			return nil, nil
		case *types.Struct:
			return NewModelsFromObject(v.Obj())
		}
	}
}

func NewModelsFromObject(object types.Object) ([]*Model, error) {
	model := Model{
		Fields: []Field{},
		Name:   object.Name() + "Model",
	}

	//objectType is a types.Type == types.Named, i.e. "DescribeInstancesOutput" with all the objects methods
	objectType := object.Type()
	// fmt.Printf("%v: %v\n", objectType, reflect.TypeOf(objectType))

	//underlyingObject is a types.Type == types.Struct with all the fields
	underlyingObject := objectType.Underlying()
	// fmt.Printf("%v: %v\n", underlyingObject, reflect.TypeOf(underlyingObject))

	underlyingStruct, ok := underlyingObject.(*types.Struct)
	if !ok {
		panic("object not a struct")
	}

	fieldModels := []*Model{}

	for i := 0; i < underlyingStruct.NumFields(); i++ {
		structField := underlyingStruct.Field(i)

		if !structField.Exported() {
			continue
		}

		structFieldType := structField.Type()

		pointedType, ok := structFieldType.(*types.Pointer)
		if ok {
			structFieldType = pointedType.Elem()
		}

		fieldType := getTypeName(structFieldType)
		models, err := getFieldModels(structFieldType)
		if err != nil {
			return nil, err
		}
		fieldModels = append(fieldModels, models...)

		field := Field{
			Name: structField.Name(),
			Type: fieldType,
		}

		field.PopulateTags()

		model.Fields = append(model.Fields, field)
	}

	models := []*Model{&model}
	models = append(models, fieldModels...)

	return models, nil
}

func DeduplicateModels(models []*Model) []*Model {
	modelNames := map[string]bool{}

	dedupedModels := []*Model{}
	for _, model := range models {
		if !modelNames[model.Name] {
			modelNames[model.Name] = true
			dedupedModels = append(dedupedModels, model)
		}
	}

	return dedupedModels
}
