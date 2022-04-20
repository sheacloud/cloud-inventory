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

type BaseConfig struct {
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
		serviceConfig := BaseConfig{
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

			filename := fmt.Sprintf("./terraform/autogen_aws_%s_dynamodb_tables.tf", service.ServiceName)
			err := os.WriteFile(filename, hclFile.Bytes(), 0755)
			if err != nil {
				panic(err)
			}
		}
	}
}
