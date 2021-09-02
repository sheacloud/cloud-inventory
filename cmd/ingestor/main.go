package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sheacloud/aws-infra-warehouse/internal/parquetwriter"
	"github.com/sheacloud/aws-infra-warehouse/pkg/service/ec2"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	today := time.Now().UTC()

	parquetConfig := parquetwriter.ParquetConfig{
		Bucket:          "sheacloud-test-parquet",
		PathPrefix:      "parquet/",
		FileExtension:   "parquet",
		Api:             s3.NewFromConfig(cfg),
		NumParquetPages: 4,
		Year:            today.Year(),
		Month:           int(today.Month()),
		Day:             today.Day(),
	}

	errors := ec2.Controller.Process(context.TODO(), "000011112222", "us-east-1", cfg, parquetConfig)
	if errors != nil {
		fmt.Println(errors)
	}
}
