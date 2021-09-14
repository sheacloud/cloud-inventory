package codegen

//go:generate go run ../cmd/model-generator/ -f ./aws_ecs_spec.hcl -o ../pkg

//go:generate go fmt ../pkg/aws/ecs/
