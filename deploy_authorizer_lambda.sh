#!/bin/bash
set -x

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./build/cloud-inventory-authorizer ./lambda/cloud-inventory-authorizer
cd ./build/
zip cloud-inventory-authorizer.zip cloud-inventory-authorizer
cd ../

aws lambda update-function-code --function-name cloud-inventory-authorizer --zip-file fileb://./build/cloud-inventory-authorizer.zip