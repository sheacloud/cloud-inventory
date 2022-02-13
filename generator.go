package main

//go:generate go run ./cmd/autogen-code/main.go
//go:generate swag init -g ./internal/api/router.go
//go:generate go fmt ./...
//go:generate go run ./cmd/autogen-terraform/main.go
