package terraformgen

import (
	"fmt"
	"os"
	"reflect"

	"github.com/fatih/structtag"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/sheacloud/cloud-inventory/internal/inventory"
)

func GetIonNameTag(tag reflect.StructTag) (string, error) {
	tags, err := structtag.Parse(string(tag))
	if err != nil {
		return "", err
	}

	ionTag, err := tags.Get("ion")
	if err != nil {
		return "", err
	}

	return ionTag.Name, nil
}

func ConvertStructToIonGlueTable(obj interface{}, cloud, service, resource string) (GlueCatalogTable, error) {
	tableName := fmt.Sprintf("ion_%s_%s_%s", cloud, service, resource)
	table := GlueCatalogTable{
		ResourceLabel: "aws_glue_catalog_table",
		NameLabel:     tableName,
		Name:          tableName,
		DatabaseName:  "replace-me",
		TableType:     "EXTERNAL_TABLE",
		Parameters: map[string]string{
			"EXTERNAL":                      "TRUE",
			"ion.encoding":                  "BINARY",
			"projection.report_date.type":   "date",
			"projection.report_date.range":  "NOW-3YEARS,NOW",
			"projection.report_date.format": "yyyy-MM-dd",
			"projection.enabled":            "true",
		},
		StorageDescriptor: StorageDescriptor{
			Location:     "replaceme",
			InputFormat:  "com.amazon.ionhiveserde.formats.IonInputFormat",
			OutputFormat: "com.amazon.ionhiveserde.formats.IonOutputFormat",
			SerDeInfo: SerDeInfo{
				Name:                 "ion",
				SerializationLibrary: "com.amazon.ionhiveserde.IonHiveSerDe",
			},
			Columns: []Columns{},
		},
		PartitionKeys: []PartitionKeys{
			{
				Name: "report_date",
				Type: "date",
			},
		},
	}

	objValue := reflect.ValueOf(obj)
	objValue = reflect.Indirect(objValue)
	objType := objValue.Type()

	//iterate over each field in obj
	for i := 0; i < objValue.NumField(); i++ {
		typeField := objType.Field(i)
		tag := typeField.Tag

		parquetFieldName, err := GetIonNameTag(tag)
		if err != nil {
			continue
		}

		typeString := ""

		if ParquetFieldIsTimestamp(tag) {
			typeString = "timestamp"
		} else {
			typeString = GetFieldTypeString(typeField.Type)
		}

		column := Columns{
			Name:       parquetFieldName,
			Type:       typeString,
			Parameters: map[string]string{},
		}
		table.StorageDescriptor.Columns = append(table.StorageDescriptor.Columns, column)
	}

	return table, nil
}

func GenerateIonGlueTerraform() {
	for _, service := range inventory.AwsCatalog {
		hclFile := hclwrite.NewEmptyFile()
		serviceConfig := BaseGlueConfig{
			Resources: []GlueCatalogTable{},
		}
		for _, resource := range service.Resources {
			table, err := ConvertStructToIonGlueTable(resource.ResourceModel, "aws", service.ServiceName, resource.ResourceName)
			if err != nil {
				panic(err)
			}
			serviceConfig.Resources = append(serviceConfig.Resources, table)
		}

		// encode the terraform into HCL
		gohcl.EncodeIntoBody(&serviceConfig, hclFile.Body())
		rootBody := hclFile.Body()

		// modify each block to use variables
		for _, resource := range service.Resources {
			resourceBlock := rootBody.FirstMatchingBlock("resource", []string{"aws_glue_catalog_table", fmt.Sprintf("ion_aws_%s_%s", service.ServiceName, resource.ResourceName)})
			resourceBlock.Body().SetAttributeTraversal("database_name", hcl.Traversal{
				hcl.TraverseRoot{
					Name: "var",
				},
				hcl.TraverseAttr{
					Name: "glue_database_name",
				},
			})

			//update storage location to use interpolation
			storageDescriptor := resourceBlock.Body().FirstMatchingBlock("storage_descriptor", []string{})
			// construct an interpolated string - see https://stackoverflow.com/questions/67945463/how-to-use-hcl-write-to-set-expressions-with for justification for this complexity
			locationTokens := hclwrite.Tokens{
				{
					Type:  hclsyntax.TokenOQuote,
					Bytes: []byte("\""),
				},
				{
					Type:  hclsyntax.TokenQuotedLit,
					Bytes: []byte("s3://"),
				},
				{
					Type:  hclsyntax.TokenTemplateInterp,
					Bytes: []byte("${"),
				},
				{
					Type:  hclsyntax.TokenIdent,
					Bytes: []byte("var"),
				},
				{
					Type:  hclsyntax.TokenDot,
					Bytes: []byte("."),
				},
				{
					Type:  hclsyntax.TokenIdent,
					Bytes: []byte("s3_bucket_name"),
				},
				{
					Type:  hclsyntax.TokenTemplateSeqEnd,
					Bytes: []byte("}"),
				},
				{
					Type:  hclsyntax.TokenQuotedLit,
					Bytes: []byte(fmt.Sprintf("/inventory/aws/%s/%s/", service.ServiceName, resource.ResourceName)),
				},
				{
					Type:  hclsyntax.TokenCQuote,
					Bytes: []byte("\""),
				},
			}
			storageDescriptor.Body().SetAttributeRaw("location", locationTokens)
		}

		filename := fmt.Sprintf("./terraform/glue/autogen_aws_%s_ion_glue_tables.tf", service.ServiceName)
		err := os.WriteFile(filename, hclFile.Bytes(), 0755)
		if err != nil {
			panic(err)
		}
	}
}
