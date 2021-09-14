package codegen

//go:generate go run ../cmd/model-generator/ -f ./aws_rds_spec.hcl -o ../pkg

//go:generate go fmt ../pkg/aws/rds/
