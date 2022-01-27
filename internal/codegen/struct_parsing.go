package codegen

import (
	"fmt"
	"go/types"
	"strings"

	"github.com/sirupsen/logrus"
)

// StructModel represents a parsed source-code of a struct and all its fields
type StructModel struct {
	Name   string
	Fields []*FieldModel
}

// FieldModel represents a parsed source-code of a struct field
type FieldModel struct {
	Name                 string
	Type                 string
	Tags                 string
	ReferencedStructs    []*StructModel // Parsed structs referenced by this field, could be multiple if it's a map
	IsConvertedTimeField bool
	NoParquetModel       bool
}

// ParseStruct parses a struct and all its fields
func ParseStruct(object *types.Named, excludedFields []string) (*StructModel, error) {

	// validate that the underlying type is a struct
	underlyingObject := object.Underlying()
	underlyingStruct, ok := underlyingObject.(*types.Struct)
	if !ok {
		return nil, ErrNotAStruct
	}

	StructModel := &StructModel{
		Name:   object.Obj().Name(),
		Fields: []*FieldModel{},
	}

	for i := 0; i < underlyingStruct.NumFields(); i++ {
		underlyingField := underlyingStruct.Field(i)
		// skip fields that are not exported
		if !underlyingField.Exported() {
			continue
		}
		// skip fields that are in the excluded list
		if stringInList(underlyingField.Name(), excludedFields) {
			continue
		}
		if isUnmodelableField(underlyingField.Type().String()) {
			continue
		}

		field, err := ParseStructField(underlyingField, excludedFields)
		if err != nil {
			return nil, err
		}

		if field != nil {
			StructModel.Fields = append(StructModel.Fields, field)
		}
	}

	return StructModel, nil
}

// ParseStructField parses a single struct field
func ParseStructField(field *types.Var, excludedFields []string) (*FieldModel, error) {
	fieldType := field.Type()

	pointedType, ok := fieldType.(*types.Pointer)
	if ok {
		fieldType = pointedType.Elem()
	}

	fieldName := field.Name()

	fieldTypeName, err := getTypeName(fieldType)
	if err != nil {
		return nil, err
	}
	// in the case that the field type can't be determined, like if it's an interface, do nothing
	if fieldTypeName == "" {
		return nil, nil
	}

	fieldTypeModels, err := parseReferencedStructs(fieldType, excludedFields)
	if err != nil {
		return nil, err
	}

	return &FieldModel{
		Name:              fieldName,
		Type:              fieldTypeName,
		ReferencedStructs: fieldTypeModels,
	}, nil

}

func (p *StructModel) MarshalToSourceCode() string {
	var buf strings.Builder
	err := structSourceCodeTemplate.Execute(&buf, p)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

func (p *StructModel) GetReferencedStructs() []*StructModel {
	structs := []*StructModel{}
	for _, field := range p.Fields {
		structs = append(structs, field.ReferencedStructs...)
		for _, fieldModel := range field.ReferencedStructs {
			structs = append(structs, fieldModel.GetReferencedStructs()...)
		}
	}

	return structs
}

func (p *StructModel) GetRequiredImports() []string {
	imports := []string{}
	for _, field := range p.Fields {
		if isTimeField(field.Type) {
			imports = append(imports, "time")
		}
	}
	return imports
}

func (p *StructModel) AddConvertedTimeFields() {
	fieldsToAdd := []*FieldModel{}
	for _, field := range p.Fields {
		if isTimeField(field.Type) {
			fieldsToAdd = append(fieldsToAdd, &FieldModel{
				Name:                 field.Name + "Milli",
				Type:                 "int64",
				IsConvertedTimeField: true,
			})
		}
	}
	p.Fields = append(p.Fields, fieldsToAdd...)
}

func (p *StructModel) ConvertTagFields() {
	for _, field := range p.Fields {
		if isTagField(field.Name) && strings.HasPrefix(field.Type, "[]") {
			field.Name = "Tags"
			field.Type = "map[string]string"
		}
	}
}

func (p *StructModel) PopulateFieldTags(primaryObjectField string) {
	for _, field := range p.Fields {

		if field.NoParquetModel {
			continue
		}

		// don't add tags for time fields
		if isTimeField(field.Type) {
			field.Tags = "`json:\"-\"`"
			continue
		}

		tags := ""

		// populate parquet tags

		fieldSnakeCaseName := toSnakeCase(field.Name)
		fieldSnakeCaseName = strings.TrimSuffix(fieldSnakeCaseName, "_milli")

		parquetTags := "name=" + fieldSnakeCaseName

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
		} else if field.IsConvertedTimeField {
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

		// add json tags
		tags += " json:\"" + fieldSnakeCaseName + "\""

		// add diff tags
		if field.Name == "ReportTime" {
			tags += " diff:\"report_time,immutable\""
		} else if field.Name == primaryObjectField {
			tags += " diff:\"" + fieldSnakeCaseName + ",identifier\""
		} else {
			tags += " diff:\"" + fieldSnakeCaseName + "\""
		}

		field.Tags = "`" + tags + "`"
	}
}

// getTypeName returns the "name" of the type as it should be written in the generated code
// for example, if the type is a pointer to a struct MyStruct, it will return "*MyStruct"
// this dereferences any pointers to basic types, and makes and reference to structs a pointer
func getTypeName(t types.Type) (string, error) {
	switch v := t.(type) {
	default:
		logrus.Errorf("unexpected type %T", v)
		return "", ErrUnsupportedFieldType
	case *types.Pointer:
		return getTypeName(v.Elem())
	case *types.Interface:
		return "", nil
	case *types.Basic:
		return v.Name(), nil
	case *types.Named:
		switch k := v.Underlying().(type) {
		default:
			return getTypeName(k)
		case *types.Struct:
			if isTimeField(v.Obj().Type().String()) {
				return "*" + v.Obj().Type().String(), nil
			} else {
				// always use pointers to structs since they may be optional in some cases
				return "*" + v.Obj().Name(), nil
			}
		}
	case *types.Array:
		return fmt.Sprintf("[%v]", v.Len()), nil
	case *types.Slice:
		elemTypeName, err := getTypeName(v.Elem())
		if err != nil {
			return "", err
		}
		return "[]" + elemTypeName, nil
	case *types.Map:
		keyTypeName, err := getTypeName(v.Key())
		if err != nil {
			return "", err
		}
		elemTypeName, err := getTypeName(v.Elem())
		if err != nil {
			return "", err
		}

		return "map[" + keyTypeName + "]" + elemTypeName, nil
	}
}

// parseReferencedStructs parses the referenced structs of a field, if any exist
// this returns a list since a Map type can reference multiple structs
func parseReferencedStructs(t types.Type, excludedFields []string) ([]*StructModel, error) {
	switch v := t.(type) {
	default:
		return nil, nil
	case *types.Pointer:
		return parseReferencedStructs(v.Elem(), excludedFields)
	case *types.Slice:
		return parseReferencedStructs(v.Elem(), excludedFields)
	case *types.Map:
		keyModels, err := parseReferencedStructs(v.Key(), excludedFields)
		if err != nil {
			return nil, err
		}
		valueModels, err := parseReferencedStructs(v.Elem(), excludedFields)
		if err != nil {
			fmt.Println(err)
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
			model, err := ParseStruct(v, excludedFields)
			if err != nil {
				return nil, err
			}
			return []*StructModel{model}, nil
		}
	}
}

func DeduplicateStructs(structs []*StructModel) []*StructModel {
	structNames := map[string]bool{}

	dedupedStructs := []*StructModel{}
	for _, s := range structs {
		if !structNames[s.Name] {
			structNames[s.Name] = true
			dedupedStructs = append(dedupedStructs, s)
		}
	}

	return dedupedStructs
}
