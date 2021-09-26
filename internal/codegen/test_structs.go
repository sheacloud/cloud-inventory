package codegen

import "time"

type ExampleStruct1 struct {
	FieldOne            string
	FieldTwo            *int64
	ExampleTime         time.Time
	ExampleTimePointer  *time.Time
	ExampleList         []int32
	NestedStruct        ExampleNestedStruct1
	NestedPointerStruct *ExampleNestedStruct2
	ExampleMap          map[string]ExampleNestedStruct3
	TestConversion      int32
}

type ExampleNestedStruct1 struct {
	FieldOne string
}

type ExampleNestedStruct2 struct {
	FieldOne string
}

type ExampleNestedStruct3 struct {
	FieldOne string
}
