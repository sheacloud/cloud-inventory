package codegen

//go:generate go run ../cmd/model-generator/ -f ./aws_s3_spec.hcl -o ../pkg

//go:generate go fmt ../pkg/aws/s3/
