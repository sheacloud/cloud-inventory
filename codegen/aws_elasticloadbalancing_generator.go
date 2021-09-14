package codegen

//go:generate go run ../cmd/model-generator/ -f ./aws_elasticloadbalancing_spec.hcl -o ../pkg

//go:generate go fmt ../pkg/aws/elasticloadbalancing/
