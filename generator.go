package main

//go:generate go run ./cmd/autogen-code/main.go
//go:generate swag init -g ./cmd/cloud-inventory-api/main.go
//go:generate go fmt ./...
//go:generate go run ./cmd/autogen-terraform/main.go
