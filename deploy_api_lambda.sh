#!/bin/bash
set -x

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/cloud-inventory-api ./lambda/cloud-inventory-api
cd ./build/
zip cloud-inventory-api.zip cloud-inventory-api
cd ../

aws lambda update-function-code --function-name cloud-inventory-api --zip-file fileb://./build/cloud-inventory-api.zip