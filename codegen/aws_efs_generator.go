package codegen

//go:generate go run ../cmd/model-generator/ -f ./aws_efs_spec.hcl -o ../pkg

//go:generate go fmt ../pkg/aws/efs/
