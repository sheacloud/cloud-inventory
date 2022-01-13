package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"path/filepath"

	"github.com/hashicorp/hcl/v2/hclsimple"
	"github.com/sheacloud/cloud-inventory/internal/codegen"
)

var (
	configDirName = flag.String("d", "", "the configuration directory")
)

func main() {
	flag.Parse()

	var config codegen.AwsTemplate

	combinedHCL := bytes.NewBuffer([]byte{})
	files, err := ioutil.ReadDir(*configDirName)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileBytes, err := ioutil.ReadFile(filepath.Join(*configDirName, file.Name()))
		if err != nil {
			panic(err)
		}
		combinedHCL.Write(fileBytes)
		combinedHCL.Write([]byte("\n"))
	}

	err = hclsimple.Decode("configuration.hcl", combinedHCL.Bytes(), nil, &config)
	if err != nil {
		panic(err)
	}

	codegen.GenerateAwsServiceCode(&config, "pkg/awscloud/")

}
