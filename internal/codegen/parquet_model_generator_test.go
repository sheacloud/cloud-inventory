package codegen

import (
	"fmt"
	"go/types"
	"reflect"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/tools/go/packages"
)

func init() {
	pkg, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedImports | packages.NeedDeps,
	}, "github.com/sheacloud/cloud-inventory/internal/codegen")
	if err != nil {
		panic(err)
	}

	exampleStruct1Object = pkg[0].Types.Scope().Lookup("ExampleStruct1").(*types.TypeName).Type().(*types.Named)
}

var exampleStruct1Object *types.Named

func TestNewParquetModelStruct(t *testing.T) {
	type args struct {
		object            *types.Named
		primaryObjectName string
		excludedFields    []string
		fieldConversions  []*ParquetModelFieldConversion
	}
	tests := []struct {
		name    string
		args    args
		want    *ParquetModelStruct
		wantErr bool
	}{
		{
			name: "TestExampleStruct1",
			args: args{
				object:            exampleStruct1Object,
				primaryObjectName: "ExampleStruct1",
				excludedFields:    []string{},
				fieldConversions: []*ParquetModelFieldConversion{
					{
						SourceFieldName: "TestConversion",
						TargetField: ParquetModelField{
							Name: "TargetConversionField",
							Type: "int64",
						},
						ConversionFunctionName: "TestFunc",
					},
				},
			},
			want: &ParquetModelStruct{
				Name: "ExampleStruct1Model",
				Fields: []*ParquetModelField{
					{
						Name: "FieldOne",
						Type: "string",
						Tags: "`parquet:\"name=field_one,type=BYTE_ARRAY,convertedtype=UTF8\"`",
					},
					{
						Name: "FieldTwo",
						Type: "int64",
						Tags: "`parquet:\"name=field_two,type=INT64\"`",
					},
					{
						Name: "ExampleTime",
						Type: "*time.Time",
					},
					{
						Name: "ExampleTimePointer",
						Type: "*time.Time",
					},
					{
						Name: "ExampleList",
						Type: "[]int32",
						Tags: "`parquet:\"name=example_list,type=MAP,convertedtype=LIST,valuetype=INT32\"`",
					},
					{
						Name: "NestedStruct",
						Type: "*ExampleNestedStruct1ExampleStruct1Model",
						Tags: "`parquet:\"name=nested_struct\"`",
						ReferencedModels: []*ParquetModelStruct{
							{
								Name: "ExampleNestedStruct1ExampleStruct1Model",
								Fields: []*ParquetModelField{
									{
										Name: "FieldOne",
										Type: "string",
										Tags: "`parquet:\"name=field_one,type=BYTE_ARRAY,convertedtype=UTF8\"`",
									},
								},
								FieldConversions: []*ParquetModelFieldConversion{
									{
										SourceFieldName: "TestConversion",
										TargetField: ParquetModelField{
											Name: "TargetConversionField",
											Type: "int64",
										},
										ConversionFunctionName: "TestFunc",
									},
								},
							},
						},
					},
					{
						Name: "NestedPointerStruct",
						Type: "*ExampleNestedStruct2ExampleStruct1Model",
						Tags: "`parquet:\"name=nested_pointer_struct\"`",
						ReferencedModels: []*ParquetModelStruct{
							{
								Name: "ExampleNestedStruct2ExampleStruct1Model",
								Fields: []*ParquetModelField{
									{
										Name: "FieldOne",
										Type: "string",
										Tags: "`parquet:\"name=field_one,type=BYTE_ARRAY,convertedtype=UTF8\"`",
									},
								},
								FieldConversions: []*ParquetModelFieldConversion{
									{
										SourceFieldName: "TestConversion",
										TargetField: ParquetModelField{
											Name: "TargetConversionField",
											Type: "int64",
										},
										ConversionFunctionName: "TestFunc",
									},
								},
							},
						},
					},
					{
						Name: "ExampleMap",
						Type: "map[string]*ExampleNestedStruct3ExampleStruct1Model",
						Tags: "`parquet:\"name=example_map,type=MAP,keytype=BYTE_ARRAY,keyconvertedtype=UTF8\"`",
						ReferencedModels: []*ParquetModelStruct{
							{
								Name: "ExampleNestedStruct3ExampleStruct1Model",
								Fields: []*ParquetModelField{
									{
										Name: "FieldOne",
										Type: "string",
										Tags: "`parquet:\"name=field_one,type=BYTE_ARRAY,convertedtype=UTF8\"`",
									},
								},
								FieldConversions: []*ParquetModelFieldConversion{
									{
										SourceFieldName: "TestConversion",
										TargetField: ParquetModelField{
											Name: "TargetConversionField",
											Type: "int64",
										},
										ConversionFunctionName: "TestFunc",
									},
								},
							},
						},
					},
					{
						Name: "TestConversion",
						Type: "int32",
					},
					{
						Name:        "ExampleTimeMilli",
						Type:        "int64",
						Tags:        "`parquet:\"name=example_time_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS\"`",
						IsTimeField: true,
					},
					{
						Name:        "ExampleTimePointerMilli",
						Type:        "int64",
						Tags:        "`parquet:\"name=example_time_pointer_milli,type=INT64,convertedtype=TIMESTAMP_MILLIS\"`",
						IsTimeField: true,
					},
					{
						Name: "TargetConversionField",
						Type: "int64",
						Tags: "`parquet:\"name=target_conversion_field,type=INT64\"`",
					},
				},
				FieldConversions: []*ParquetModelFieldConversion{
					{
						SourceFieldName: "TestConversion",
						TargetField: ParquetModelField{
							Name: "TargetConversionField",
							Type: "int64",
						},
						ConversionFunctionName: "TestFunc",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewParquetModelStruct(tt.args.object, tt.args.primaryObjectName, tt.args.excludedFields, tt.args.fieldConversions)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewParquetModelStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got.PopulateFieldTags(tt.args.primaryObjectName)
			if !reflect.DeepEqual(got, tt.want) {
				spew.Dump(got)
				fmt.Print("\n\n\n")
				spew.Dump(tt.want)
				for i := 0; i < len(got.Fields); i++ {
					if !reflect.DeepEqual(got.Fields[i], tt.want.Fields[i]) {
						t.Log("---------------------------")
						t.Log(got.Fields[i])
						t.Log(tt.want.Fields[i])
					}
				}
				t.Errorf("NewParquetModelStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParquetModelStruct_StructSourceCode(t *testing.T) {
	tests := []struct {
		name string
		p    *ParquetModelStruct
		want string
	}{
		{
			name: "Basic",
			p: &ParquetModelStruct{
				Name: "ExampleModel",
				Fields: []*ParquetModelField{
					{
						Name: "FieldOne",
						Type: "string",
						Tags: "`parquet:\"name=field_one\"`",
					},
					{
						Name: "FieldTwo",
						Type: "int64",
						Tags: "`parquet:\"name=field_two\"`",
					},
				},
			},
			want: `type ExampleModel struct {
	FieldOne string ` + "`parquet:\"name=field_one\"`" + `
	FieldTwo int64 ` + "`parquet:\"name=field_two\"`" + `
}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.StructSourceCode(); got != tt.want {
				t.Errorf("ParquetModelStruct.StructSourceCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
