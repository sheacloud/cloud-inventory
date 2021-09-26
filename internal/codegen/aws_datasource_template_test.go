package codegen

import "testing"

func TestAwsDataSourceFile_SourceCode(t *testing.T) {
	tests := []struct {
		name string
		d    *AwsDataSourceFile
	}{
		{
			name: "BasicPaginate",
			d: &AwsDataSourceFile{
				ServiceName:       "service",
				DataSourceName:    "resource",
				PrimaryObjectName: "ExampleObject",
				ApiFunction:       "ListExampleObjects",
				Paginate:          true,
				ModelsOnly:        false,
				PrimaryObjectPath: []string{"ExampleObjects"},
				PrimaryModel: &ParquetModelStruct{
					Name: "ExampleObject",
					Fields: []*ParquetModelField{
						{
							Name: "FieldOne",
							Type: "string",
							Tags: "exampleTags",
						},
						{
							Name: "OldField",
							Type: "int32",
							Tags: "exampleTags2",
						},
						{
							Name: "NestedField",
							Type: "*ExampleNestedObject",
							Tags: "exampleTags3",
						},
						{
							Name: "NewField",
							Type: "int64",
							Tags: "exampleTags4",
						},
					},
					FieldConversions: []*ParquetModelFieldConversion{
						{
							SourceFieldName: "OldField",
							TargetField: ParquetModelField{
								Name: "NewField",
								Type: "int64",
								Tags: "exampleTags4",
							},
							ConversionFunctionName: "ConvertInt",
						},
					},
				},
				Models: []*ParquetModelStruct{
					{
						Name: "ExampleObject",
						Fields: []*ParquetModelField{
							{
								Name: "FieldOne",
								Type: "string",
								Tags: "exampleTags",
							},
							{
								Name: "OldField",
								Type: "int32",
								Tags: "exampleTags2",
							},
							{
								Name: "NestedField",
								Type: "*ExampleNestedObject",
								Tags: "exampleTags3",
							},
							{
								Name: "NewField",
								Type: "int64",
								Tags: "exampleTags4",
							},
						},
						FieldConversions: []*ParquetModelFieldConversion{
							{
								SourceFieldName: "OldField",
								TargetField: ParquetModelField{
									Name: "NewField",
									Type: "int64",
									Tags: "exampleTags4",
								},
								ConversionFunctionName: "ConvertInt",
							},
						},
					},
					{
						Name: "ExampleNestedObject",
						Fields: []*ParquetModelField{
							{
								Name: "FieldOne",
								Type: "string",
								Tags: "exampleTags5",
							},
						},
						FieldConversions: []*ParquetModelFieldConversion{},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.d.SourceCode()
		})
	}
}
