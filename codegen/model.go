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
	Name             string
	Fields           []*Field
	TagFieldsMapping map[string]string
}

type Field struct {
	Name string
	Type string
	Tags string
}

type DatasourceFile struct {
	ServiceName         string
	DataSourceName      string
	PrimaryResourceName string
	ApiFunction         string
	Models              []*Model
	PrimaryObjectPath   []string
	PrimaryModel        *Model
	Paginate            bool
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

var datasourceTemplate string = `// AUTOGENERATED, DO NOT EDIT
package {{.ServiceName}}
import (
	"context"
	"fmt"
	"sync"
	"time"
	
	"github.com/aws/aws-sdk-go-v2/service/{{.ServiceName}}"
	"github.com/jinzhu/copier"
	"github.com/sheacloud/cloud-inventory/internal/storage"
	"github.com/sirupsen/logrus"
)

var custom{{.PrimaryModel.Name}}PostprocessingFuncs []func(x *{{.PrimaryModel.Name}}) = []func(x *{{.PrimaryModel.Name}}){}
var custom{{.PrimaryModel.Name}}FuncsLock sync.Mutex

func registerCustom{{.PrimaryModel.Name}}PostprocessingFunc(f func(x *{{.PrimaryModel.Name}})) {
	custom{{.PrimaryModel.Name}}FuncsLock.Lock()
	defer custom{{.PrimaryModel.Name}}FuncsLock.Unlock()

	custom{{.PrimaryModel.Name}}PostprocessingFuncs = append(custom{{.PrimaryModel.Name}}PostprocessingFuncs, f)
}

func init() {
	Controller.RegisterDataSource("{{.DataSourceName}}", {{.PrimaryResourceName}}DataSource)
}

{{range $index, $element := .Models}}
type {{$element.Name}} struct {
{{range $subindex, $subelement := $element.Fields}}	{{$subelement.Name}} {{$subelement.Type}} {{$subelement.Tags}}
{{end}}}	
{{end}}
{{$pathLength := len .PrimaryObjectPath}}

func {{.PrimaryResourceName}}DataSource(ctx context.Context, client *{{.ServiceName}}.Client, reportTime time.Time, storageConfig storage.StorageContextConfig, storageManager *storage.StorageManager) error {
	storageContextSet, err := storageManager.GetStorageContextSet(storageConfig, new({{.PrimaryResourceName}}Model))
	if err != nil {
		return err
	}
	defer storageContextSet.Close(ctx)

	{{if .Paginate}}
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
	{{else}}
	params := &{{.ServiceName}}.{{.ApiFunction}}Input{}
	
	result, err := client.{{.ApiFunction}}(ctx, params)
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

	results := []*{{.ServiceName}}.{{.ApiFunction}}Output{result}
	for _, output := range results {
	{{end}}
		{{range $index, $element := .PrimaryObjectPath}}{{tab $index}}for _, var{{$index}} := range {{if eq $index 0}}output{{else}}var{{sub $index 1}}{{end}}.{{$element}} {
		{{end}}
		{{tab $pathLength}}model := new({{.PrimaryResourceName}}Model)
		{{tab $pathLength}}copier.Copy(&model, &var{{sub $pathLength 1}})

		{{range $listField, $mapField := .PrimaryModel.TagFieldsMapping}}{{tab $pathLength}}model.{{$mapField}} = GetTagMap(var{{sub $pathLength 1}}.{{$mapField}}){{end}}
		{{tab $pathLength}}model.AccountId = storageConfig.AccountId
		{{tab $pathLength}}model.Region = storageConfig.Region
		{{tab $pathLength}}model.ReportTime = reportTime.UTC().UnixMilli()

		{{tab $pathLength}}for _, f := range custom{{.PrimaryModel.Name}}PostprocessingFuncs {
		{{tab $pathLength}}	f(model)
		{{tab $pathLength}}}

		{{tab $pathLength}}errors := storageContextSet.Store(ctx, model)
		{{tab $pathLength}}for storageContext, err := range errors {
		{{tab $pathLength}}	storage.LogContextError(storageContext, fmt.Sprintf("Error storing {{.PrimaryResourceName}}Model: %v", err))
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

var mapRegex = regexp.MustCompile(`map\[(.*)\](.*)`)

var TagFieldNames []string = []string{"Tags", "TagSet"}

func isTagFieldName(s string) bool {
	for _, tagFieldName := range TagFieldNames {
		if s == tagFieldName {
			return true
		}
	}

	return false
}

func isTimeField(s string) bool {
	return s == "time.Time" || s == "*time.Time"
}

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func TypeToParquetType(s string) string {
	switch s {
	case "string":
		return "BYTE_ARRAY"
	case "int32":
		return "INT32"
	case "bool":
		return "BOOLEAN"
	case "int64":
		return "INT64"
	case "float32":
		return "FLOAT"
	case "float64":
		return "DOUBLE"
	default:
		return ""
	}
}

func TypeToParquetConvertedType(s string) string {
	switch s {
	case "string":
		return "UTF8"
	default:
		return ""
	}
}

func (f *Field) PopulateTags() {
	parquetName := ToSnakeCase(f.Name)

	var extraParquetTags string

	if strings.HasPrefix(f.Type, "[]") {
		extraParquetTags = ",type=LIST"
		listType := f.Type[2:]

		listTypeString := TypeToParquetType(listType)
		listConvertedTypeString := TypeToParquetConvertedType(listType)

		if listTypeString != "" {
			extraParquetTags = ",type=MAP,convertedtype=LIST"
			extraParquetTags += ",valuetype=" + listTypeString

			if listConvertedTypeString != "" {
				extraParquetTags += ",valueconvertedtype=" + listConvertedTypeString
			}
		}

	} else if strings.HasPrefix(f.Type, "map") {
		groups := mapRegex.FindStringSubmatch(f.Type)
		if len(groups) != 3 {
			fmt.Println(groups)
			panic("map type not matched by regex: " + f.Type)
		}

		keyTypeString := TypeToParquetType(groups[1])
		keyConvertedTypeString := TypeToParquetConvertedType(groups[1])

		valueTypeString := TypeToParquetType(groups[2])
		valueConvertedTypeString := TypeToParquetConvertedType(groups[2])

		extraParquetTags = fmt.Sprintf(",type=MAP,keytype=%s,valuetype=%s", keyTypeString, valueTypeString)
		if keyConvertedTypeString != "" {
			extraParquetTags += fmt.Sprintf(",keyconvertedtype=%s", keyConvertedTypeString)
		}
		if valueConvertedTypeString != "" {
			extraParquetTags += fmt.Sprintf(",valueconvertedtype=%s", valueConvertedTypeString)
		}
	} else {
		typeString := TypeToParquetType(f.Type)
		convertedTypeString := TypeToParquetConvertedType(f.Type)

		extraParquetTags = ""
		if typeString != "" {
			extraParquetTags += fmt.Sprintf(",type=%s", typeString)
		}
		if convertedTypeString != "" {
			extraParquetTags += fmt.Sprintf(",convertedtype=%s", convertedTypeString)
		}
	}

	f.Tags = fmt.Sprintf("`parquet:\"name=%s%s\"`", parquetName, extraParquetTags)
}

func (f *Field) AppendToTags(s string) {
	if f.Tags == "" {
		f.Tags = fmt.Sprintf("`%s`", s)
	} else {
		f.Tags = strings.TrimSuffix(f.Tags, "`")
		f.Tags += s + "`"
	}
}

func getTypeName(t types.Type, primaryResourcename string) string {
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
			fieldType = getTypeName(k, primaryResourcename)
		case *types.Struct:
			if isTimeField(v.Obj().Type().String()) {
				fieldType = "*" + v.Obj().Type().String()
			} else {
				// always use pointers to structs since they may be optional in some cases
				fieldType = "*" + v.Obj().Name() + primaryResourcename + "Model"
			}
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
		fieldType += getTypeName(elemType, primaryResourcename)
	}

	return fieldType
}

func getFieldModels(t types.Type, primaryResourcename string) ([]*Model, error) {
	switch v := t.(type) {
	default:
		return nil, nil
	case *types.Slice:
		return getFieldModels(v.Elem(), primaryResourcename)
	case *types.Named:
		switch v.Underlying().(type) {
		default:
			return nil, nil
		case *types.Struct:
			if isTimeField(v.Obj().Type().String()) {
				return nil, nil
			}
			return NewModelsFromObject(v.Obj(), primaryResourcename)
		}
	}
}

func NewModelsFromObject(object types.Object, primaryResourcename string) ([]*Model, error) {
	var modelName string
	// if the model is the primary resource, just append Model to the name
	// if it's a sub-model of the primary resource, append the primary resource name as well to make it unique in the package
	if object.Name() == primaryResourcename {
		modelName = object.Name() + "Model"
	} else {
		modelName = object.Name() + primaryResourcename + "Model"
	}
	model := Model{
		Fields:           []*Field{},
		Name:             modelName,
		TagFieldsMapping: map[string]string{},
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

	// iterate over each of the structs fields, get the name of the type (i.e. string, map[x]y, or *StructZ)
	for i := 0; i < underlyingStruct.NumFields(); i++ {
		structField := underlyingStruct.Field(i)
		if !structField.Exported() {
			continue
		}

		structFieldType := structField.Type()

		// dereference any pointers
		pointedType, ok := structFieldType.(*types.Pointer)
		if ok {
			structFieldType = pointedType.Elem()
		}

		//get the name of the field type (i.e. string, map[x]y, etc)
		fieldType := getTypeName(structFieldType, primaryResourcename)
		//recursively generate any necessary sub-models if this field is another struct
		models, err := getFieldModels(structFieldType, primaryResourcename)
		if err != nil {
			return nil, err
		}
		fieldModels = append(fieldModels, models...)

		// create a modified field for any Tags fields to change it from a list of tags to a key-value map of tags
		fieldName := structField.Name()
		if isTagFieldName(fieldName) {
			// add List to this fields name, create a new field with a map type to store the map of the tags

			mapField := Field{
				Name: fieldName,
				Type: "map[string]string",
			}
			mapField.PopulateTags()

			model.TagFieldsMapping[fieldName+"List"] = fieldName

			model.Fields = append(model.Fields, &mapField)
		} else if isTimeField(fieldType) {
			field := Field{
				Name: fieldName,
				Type: fieldType,
			}

			milliField := Field{
				Name: fieldName + "Milli",
				Type: "int64",
				Tags: fmt.Sprintf("`parquet:\"name=%s, type=INT64, convertedtype=TIMESTAMP_MILLIS\"`", ToSnakeCase(fieldName)),
			}
			model.Fields = append(model.Fields, &field, &milliField)
		} else {
			field := Field{
				Name: fieldName,
				Type: fieldType,
			}

			field.PopulateTags()

			model.Fields = append(model.Fields, &field)
		}
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

func (m *Model) UpdatePrimaryResourceFieldTag(primaryResourceField string) bool {
	for _, field := range m.Fields {
		if field.Name == primaryResourceField {
			field.AppendToTags(" inventory_primary_key:\"true\"")
			return true
		}
	}
	return false
}
