package codegen

//go:generate go run ../cmd/model-generator/ -f ./aws_iam_spec.hcl -o ../pkg

//go:generate go fmt ../pkg/aws/iam/
