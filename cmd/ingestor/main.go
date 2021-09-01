package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"gitlab.com/sheacloud-infrastructure/aws/data-warehouse/internal/parquetwriter"
	"gitlab.com/sheacloud-infrastructure/aws/data-warehouse/pkg/service/ec2"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	parquetConfig := parquetwriter.ParquetConfig{
		Bucket:          "sheacloud-test-parquet",
		PathPrefix:      "parquet/",
		FileExtension:   "parquet",
		Api:             s3.NewFromConfig(cfg),
		NumParquetPages: 4,
		Year:            2021,
		Month:           8,
		Day:             31,
	}

	errors := ec2.Controller.Process(context.TODO(), "000011112222", "us-east-1", cfg, parquetConfig)
	if errors != nil {
		fmt.Println(errors)
	}
}
