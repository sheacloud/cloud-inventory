package main

import "github.com/sheacloud/cloud-inventory/internal/terraformgen"

func main() {
	terraformgen.GenerateDynamoDBTerraform()
	// terraformgen.GenerateIonGlueTerraform()
	terraformgen.GenerateParquetGlueTerraform()
}
