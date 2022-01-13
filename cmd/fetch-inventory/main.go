package main

import (
	"context"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/sheacloud/cloud-inventory/internal/indexedstorage"
	"github.com/sheacloud/cloud-inventory/internal/inventory"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	s3Client := s3.NewFromConfig(cfg)

	fileManager := indexedstorage.NewIndexedFileManager("sheacloud-core-cloud-inventory", "test/", "parquet", s3Client, 4)

	accountIDs := []string{"306526781466", "166953723888", "055185845477", "023611290521", "261108431719", "536809099843"}
	regions := []string{"us-east-1", "us-west-2"}

	inventory.FetchAwsInventory(context.TODO(), accountIDs, regions, cfg, "gitlab-automation-role", time.Now(), fileManager, 32)
}
