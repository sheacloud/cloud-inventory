package terraformgen

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/sheacloud/cloud-inventory/internal/codegen"
	"github.com/sheacloud/cloud-inventory/internal/inventory"
)

var (
	FieldKindToType = map[reflect.Kind]string{
		reflect.Bool:    "boolean",
		reflect.String:  "string",
		reflect.Int:     "int",
		reflect.Int32:   "int",
		reflect.Int64:   "bigint",
		reflect.Float32: "float",
		reflect.Float64: "double",
	}
)

type BaseDynamoConfig struct {
	Resources []DynamoDBTable `hcl:"resource,block"`
}

type DynamoDBTable struct {
	ResourceLabel          string                   `hcl:"resource_label,label"`
	NameLabel              string                   `hcl:"name_label,label"`
	Name                   string                   `hcl:"name"`
	BillingMode            string                   `hcl:"billing_mode"`
	HashKey                string                   `hcl:"hash_key"`
	Attributes             []DynamoDBTableAttribute `hcl:"attribute,block"`
	GlobalSecondaryIndexes []GlobalSecondaryIndex   `hcl:"global_secondary_index,block"`
}

type DynamoDBTableAttribute struct {
	Name string `hcl:"name"`
	Type string `hcl:"type"`
}

type GlobalSecondaryIndex struct {
	Name           string `hcl:"name"`
	HashKey        string `hcl:"hash_key"`
	RangeKey       string `hcl:"range_key"`
	ProjectionType string `hcl:"projection_type"`
}

func GenerateDynamoDBTerraform() {
	for _, service := range inventory.AwsCatalog {
		hclFile := hclwrite.NewEmptyFile()
		serviceConfig := BaseDynamoConfig{
			Resources: []DynamoDBTable{},
		}
		for _, resource := range service.Resources {
			resourceUniqueField := codegen.ToSnakeCase(resource.UniqueIdField)
			table := DynamoDBTable{
				ResourceLabel: "aws_dynamodb_table",
				NameLabel:     fmt.Sprintf("aws_%s_%s", service.ServiceName, resource.ResourceName),
				Name:          fmt.Sprintf("cloud-inventory-aws-%s-%s", service.ServiceName, strings.Replace(resource.ResourceName, "_", "-", -1)),
				BillingMode:   "PAY_PER_REQUEST",
				HashKey:       "_id",
				Attributes: []DynamoDBTableAttribute{
					{
						Name: "_id",
						Type: "S",
					},
					{
						Name: "report_time",
						Type: "N",
					},
					{
						Name: resourceUniqueField,
						Type: "S",
					},
				},
				GlobalSecondaryIndexes: []GlobalSecondaryIndex{
					{
						Name:           "report_time_id_index",
						HashKey:        "report_time",
						RangeKey:       resourceUniqueField,
						ProjectionType: "ALL",
					},
					{
						Name:           "id_report_time_index",
						HashKey:        resourceUniqueField,
						RangeKey:       "report_time",
						ProjectionType: "ALL",
					},
				},
			}
			serviceConfig.Resources = append(serviceConfig.Resources, table)
			gohcl.EncodeIntoBody(&serviceConfig, hclFile.Body())

			filename := fmt.Sprintf("./terraform/dynamodb/autogen_aws_%s_dynamodb_tables.tf", service.ServiceName)
			err := os.WriteFile(filename, hclFile.Bytes(), 0755)
			if err != nil {
				panic(err)
			}
		}
	}
}

type BaseGlueConfig struct {
	Resources []GlueCatalogTable `hcl:"resource,block"`
}

type GlueCatalogTable struct {
	ResourceLabel     string            `hcl:"resource_label,label"`
	NameLabel         string            `hcl:"name_label,label"`
	Name              string            `hcl:"name"`
	DatabaseName      string            `hcl:"database_name"`
	TableType         string            `hcl:"table_type"`
	Parameters        map[string]string `hcl:"parameters"`
	StorageDescriptor StorageDescriptor `hcl:"storage_descriptor,block"`
	PartitionKeys     []PartitionKeys   `hcl:"partition_keys,block"`
}

type StorageDescriptor struct {
	Location     string    `hcl:"location"`
	InputFormat  string    `hcl:"input_format"`
	OutputFormat string    `hcl:"output_format"`
	SerDeInfo    SerDeInfo `hcl:"ser_de_info,block"`
	Columns      []Columns `hcl:"columns,block"`
}

type SerDeInfo struct {
	Name                 string            `hcl:"name"`
	SerializationLibrary string            `hcl:"serialization_library"`
	Parameters           map[string]string `hcl:"parameters"`
}

type Columns struct {
	Name       string            `hcl:"name"`
	Type       string            `hcl:"type"`
	Comment    string            `hcl:"comment"`
	Parameters map[string]string `hcl:"parameters"`
}

type PartitionKeys struct {
	Name string `hcl:"name"`
	Type string `hcl:"type"`
}

func GetFieldTypeString(fieldType reflect.Type) string {

	if fieldType.Kind() == reflect.Ptr {
		fieldType = fieldType.Elem()
	}

	if fieldType.Kind() == reflect.Struct {
		//for each field of the struct, call GetFieldTypeString
		parquetString := "struct<"
		for i := 0; i < fieldType.NumField(); i++ {
			subfieldType := fieldType.Field(i)

			tag := subfieldType.Tag

			parquetFieldName, err := GetParquetNameTag(tag)
			if err != nil {
				continue
			}

			var typeString string
			if ParquetFieldIsTimestamp(tag) {
				typeString = "timestamp"
			} else {
				typeString = GetFieldTypeString(subfieldType.Type)
			}

			parquetString += fmt.Sprintf("%s:%s", parquetFieldName, typeString)
			if i != fieldType.NumField()-1 {
				parquetString += ","
			}
		}
		parquetString += ">"

		return parquetString
	} else if fieldType.Kind() == reflect.Array || fieldType.Kind() == reflect.Slice {
		parquetString := "array<"
		arrayType := fieldType.Elem()
		parquetString += GetFieldTypeString(arrayType)
		parquetString += ">"
		return parquetString
	} else if fieldType.Kind() == reflect.Map {
		parquetString := "map<"
		keyType := fieldType.Key()
		elemType := fieldType.Elem()
		keyString := GetFieldTypeString(keyType)
		elemString := GetFieldTypeString(elemType)
		parquetString += keyString
		parquetString += ","
		parquetString += elemString
		parquetString += ">"
		return parquetString
	} else {
		return FieldKindToType[fieldType.Kind()]
	}
}
