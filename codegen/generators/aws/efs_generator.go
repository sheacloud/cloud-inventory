package codegen

//go:generate go run ../../../cmd/model-generator/ -f ../../service-specs/aws/efs_spec.hcl -o ../../../pkg

//go:generate go fmt ../../../pkg/aws/efs/
