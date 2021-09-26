package codegen

import (
	"errors"
	"fmt"
	"go/types"
	"regexp"
	"strings"
	"text/template"
)

var (
	ErrNotAStruct           = errors.New("given objects underlying type is not a struct")
	ErrUnsupportedFieldType = errors.New("given field type is not supported")

	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
	mapRegex      = regexp.MustCompile(`map\[(.*)\](.*)`)

	structSourceCodeTemplateString = `type {{.Name}} struct {
{{range $index, $value := .Fields}}	{{$value.Name}} {{$value.Type}} {{$value.Tags}}
{{end}}}`
	structSourceCodeTemplate = template.Must(template.New("struct").Parse(structSourceCodeTemplateString))
)

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func stringInList(s string, list []string) bool {
	for _, s2 := range list {
		if s == s2 {
			return true
		}
	}
	return false
}

func isTimeField(s string) bool {
	return s == "time.Time" || s == "*time.Time"
}

func typeToParquetType(s string) string {
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

func typeToParquetConvertedType(s string) string {
	switch s {
	case "string":
		return "UTF8"
	default:
		return ""
	}
}

type ParquetModelStruct struct {
	Name             string               // the name of the resulting struct
	Fields           []*ParquetModelField // the fields of the resulting struct
	FieldConversions []*ParquetModelFieldConversion
}

func (p *ParquetModelStruct) StructSourceCode() string {
	var buf strings.Builder
	err := structSourceCodeTemplate.Execute(&buf, p)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (p *ParquetModelStruct) generateExtraFields() {
	newFields := []*ParquetModelField{}

	for _, field := range p.Fields {
		// add any field conversion target fields to the model
		for _, conversion := range p.FieldConversions {
			if conversion.SourceFieldName == field.Name {
				if conversion.SourceFieldNameOverride != nil {
					field.Name = *conversion.SourceFieldNameOverride
				}
				newFields = append(newFields, &ParquetModelField{
					Name: conversion.TargetField.Name,
					Type: conversion.TargetField.Type,
				})
			}
		}

		// add any time conversion fields
		if isTimeField(field.Type) {
			newField := &ParquetModelField{
				Name:        field.Name + "Milli",
				Type:        "int64",
				IsTimeField: true,
			}

			newFields = append(newFields, newField)
		}
	}

	p.Fields = append(p.Fields, newFields...)
}

func (p *ParquetModelStruct) PopulateFieldTags(primaryObjectField string) {
	for _, field := range p.Fields {
		// don't tag field conversion source fields
		isConvertedSourceField := false
		for _, conversion := range p.FieldConversions {
			if conversion.SourceFieldNameOverride != nil {
				if field.Name == *conversion.SourceFieldNameOverride {
					isConvertedSourceField = true
					break
				}
			} else {
				if field.Name == conversion.SourceFieldName {
					isConvertedSourceField = true
					break
				}
			}
		}
		if isConvertedSourceField {
			continue
		}

		// don't add tags for time fields
		if isTimeField(field.Type) {
			continue
		}

		tags := ""

		// populate parquet tags

		fieldParquetName := toSnakeCase(field.Name)
		parquetTags := "name=" + fieldParquetName

		if strings.HasPrefix(field.Type, "[]") {
			listType := field.Type[2:]

			listTypeString := typeToParquetType(listType)
			listConvertedTypeString := typeToParquetConvertedType(listType)
			parquetTags += ",type=MAP,convertedtype=LIST"

			if listTypeString != "" {
				parquetTags += ",valuetype=" + listTypeString

				if listConvertedTypeString != "" {
					parquetTags += ",valueconvertedtype=" + listConvertedTypeString
				}
			}
		} else if strings.HasPrefix(field.Type, "map") {
			groups := mapRegex.FindStringSubmatch(field.Type)
			if len(groups) != 3 {
				fmt.Println(groups)
				panic("map type not matched by regex: " + field.Type)
			}
			parquetTags += ",type=MAP"

			keyTypeString := typeToParquetType(groups[1])
			keyConvertedTypeString := typeToParquetConvertedType(groups[1])

			valueTypeString := typeToParquetType(groups[2])
			valueConvertedTypeString := typeToParquetConvertedType(groups[2])

			if keyTypeString != "" {
				parquetTags += ",keytype=" + keyTypeString
			}
			if valueTypeString != "" {
				parquetTags += ",valuetype=" + valueTypeString
			}
			if keyConvertedTypeString != "" {
				parquetTags += ",keyconvertedtype=" + keyConvertedTypeString
			}
			if valueConvertedTypeString != "" {
				parquetTags += ",valueconvertedtype=" + valueConvertedTypeString
			}
		} else if field.IsTimeField {
			parquetTags += ",type=INT64,convertedtype=TIMESTAMP_MILLIS"
		} else {
			typeString := typeToParquetType(field.Type)
			convertedTypeString := typeToParquetConvertedType(field.Type)

			if typeString != "" {
				parquetTags += ",type=" + typeString
			}
			if convertedTypeString != "" {
				parquetTags += ",convertedtype=" + convertedTypeString
			}
		}

		tags += "parquet:\"" + parquetTags + "\""

		// add primary key tag
		if field.Name == primaryObjectField {
			tags += " inventory_primary_key:\"true\""
		}

		field.Tags = "`" + tags + "`"
	}
}

func (p *ParquetModelStruct) GetChildModels() []*ParquetModelStruct {
	models := []*ParquetModelStruct{}
	for _, field := range p.Fields {
		models = append(models, field.ReferencedModels...)
		for _, fieldModel := range field.ReferencedModels {
			models = append(models, fieldModel.GetChildModels()...)
		}
	}

	return models
}

type ParquetModelField struct {
	Name             string                // the variable name of the field
	Type             string                // the type of the field
	Tags             string                // the struct tags of the field
	ReferencedModels []*ParquetModelStruct // the newly modeled struct that this field is a type of, if applicable
	IsTimeField      bool                  // represents if the field is a timestamp
}

type ParquetModelFieldConversion struct {
	SourceFieldName         string // the source for the conversion
	SourceFieldNameOverride *string
	TargetField             ParquetModelField // the destination for the conversion
	ConversionFunctionName  string            // the name of the function to call for the conversion
}

// returns a ParquetModelStruct based on the source code named object "object". The underlying type of object must be a struct currently
func NewParquetModelStruct(object *types.Named, primaryObjectName string, excludedFields []string, fieldConversions []*ParquetModelFieldConversion) (*ParquetModelStruct, error) {
	var modelName string
	// append the primaryObjectName to each models name to ensure uniqueness within a package
	if object.Obj().Name() == primaryObjectName {
		modelName = object.Obj().Name() + "Model"
	} else {
		modelName = object.Obj().Name() + primaryObjectName + "Model"
	}

	underlyingObject := object.Underlying()
	underlyingStruct, ok := underlyingObject.(*types.Struct)
	if !ok {
		return nil, ErrNotAStruct
	}

	model := ParquetModelStruct{
		Name:             modelName,
		Fields:           []*ParquetModelField{},
		FieldConversions: fieldConversions,
	}

	for i := 0; i < underlyingStruct.NumFields(); i++ {
		structField := underlyingStruct.Field(i)
		// we can't model non-exported fields
		if !structField.Exported() {
			continue
		}

		if stringInList(structField.Name(), excludedFields) {
			continue
		}

		field, err := NewParquetModelField(structField, primaryObjectName, excludedFields, fieldConversions)
		if err != nil {
			return nil, err
		}

		// for if field couldn't be generated, like if it's an interface type
		if field != nil {
			model.Fields = append(model.Fields, field)
		}

	}

	model.generateExtraFields()

	return &model, nil
}

func NewParquetModelField(field *types.Var, primaryObjectName string, excludedFields []string, fieldConversions []*ParquetModelFieldConversion) (*ParquetModelField, error) {
	fieldType := field.Type()
	// dereference any pointers
	pointedType, ok := fieldType.(*types.Pointer)
	if ok {
		fieldType = pointedType.Elem()
	}

	fieldName := field.Name()

	fieldTypeName, err := getTypeName(fieldType, primaryObjectName)
	if err != nil {
		return nil, err
	}
	if fieldTypeName == "" {
		return nil, nil
	}
	fieldTypeModels, err := getTypeModels(fieldType, primaryObjectName, excludedFields, fieldConversions)
	if err != nil {
		return nil, err
	}

	return &ParquetModelField{
		Name:             fieldName,
		Type:             fieldTypeName,
		ReferencedModels: fieldTypeModels,
	}, nil
}

func getTypeName(t types.Type, primaryObjectName string) (string, error) {
	switch v := t.(type) {
	default:
		return "", ErrUnsupportedFieldType
	case *types.Interface:
		return "", nil
	case *types.Basic:
		return v.Name(), nil
	case *types.Named:
		switch k := v.Underlying().(type) {
		default:
			return getTypeName(k, primaryObjectName)
		case *types.Struct:
			if isTimeField(v.Obj().Type().String()) {
				return "*" + v.Obj().Type().String(), nil
			} else {
				// always use pointers to structs since they may be optional in some cases
				return "*" + v.Obj().Name() + primaryObjectName + "Model", nil
			}
		}
	case *types.Array:
		return fmt.Sprintf("[%v]", v.Len()), nil
	case *types.Slice:
		elemTypeName, err := getTypeName(v.Elem(), primaryObjectName)
		if err != nil {
			return "", err
		}
		return "[]" + elemTypeName, nil
	case *types.Map:
		keyTypeName, err := getTypeName(v.Key(), primaryObjectName)
		if err != nil {
			return "", err
		}
		elemTypeName, err := getTypeName(v.Elem(), primaryObjectName)
		if err != nil {
			return "", err
		}

		return "map[" + keyTypeName + "]" + elemTypeName, nil
	}
}

func getTypeModels(t types.Type, primaryObjectName string, excludedFields []string, fieldConversions []*ParquetModelFieldConversion) ([]*ParquetModelStruct, error) {
	switch v := t.(type) {
	default:
		return nil, nil
	case *types.Slice:
		return getTypeModels(v.Elem(), primaryObjectName, excludedFields, fieldConversions)
	case *types.Map:
		keyModels, err := getTypeModels(v.Key(), primaryObjectName, excludedFields, fieldConversions)
		if err != nil {
			return nil, err
		}
		valueModels, err := getTypeModels(v.Elem(), primaryObjectName, excludedFields, fieldConversions)
		if err != nil {
			return nil, err
		}
		return append(keyModels, valueModels...), nil
	case *types.Named:
		switch v.Underlying().(type) {
		default:
			return nil, nil
		case *types.Struct:
			if isTimeField(v.Obj().Type().String()) {
				return nil, nil
			}
			model, err := NewParquetModelStruct(v, primaryObjectName, excludedFields, fieldConversions)
			if err != nil {
				return nil, err
			}
			return []*ParquetModelStruct{model}, nil
		}
	}
}

func DeduplicateModels(models []*ParquetModelStruct) []*ParquetModelStruct {
	modelNames := map[string]bool{}

	dedupedModels := []*ParquetModelStruct{}
	for _, model := range models {
		if !modelNames[model.Name] {
			modelNames[model.Name] = true
			dedupedModels = append(dedupedModels, model)
		}
	}

	return dedupedModels
}
