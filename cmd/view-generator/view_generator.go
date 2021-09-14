package main

import (
	"context"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/aws/aws-sdk-go-v2/service/athena/types"
	"github.com/fatih/structtag"
	"github.com/sheacloud/cloud-inventory/internal/catalog"
	"github.com/sirupsen/logrus"
)

var (
	dailySchema = `
-- AUTOGENERATED BY VIEW-GENERATOR, DO NOT MODIFY
CREATE OR REPLACE VIEW daily_%s_%s AS 
WITH
	filter AS (
	SELECT
		*
	, "max"(report_time) OVER (PARTITION BY report_date) oldest_report_time
	FROM
		%s_%s
) 
SELECT *
FROM
	filter
WHERE (report_time = oldest_report_time)`

	distinctSchema = `
-- AUTOGENERATED BY VIEW-GENERATOR, DO NOT MODIFY
CREATE OR REPLACE VIEW distinct_%s_%s AS 
WITH
	filter AS (
	SELECT
		*
	, "max"(report_time) OVER (PARTITION BY %s) oldest_report_time
	FROM
		%s_%s
) 
SELECT *
FROM
	filter
WHERE (report_time = oldest_report_time)`

	currentSchema = `
-- AUTOGENERATED BY VIEW-GENERATOR, DO NOT MODIFY
CREATE OR REPLACE VIEW current_%s_%s AS 
WITH
	filter AS (
	SELECT
		*
	, "max"(report_time) OVER (PARTITION BY report_date) oldest_report_time
	FROM
		%s_%s
	WHERE (report_date = current_date)
) 
SELECT *
FROM
	filter
WHERE (report_time = oldest_report_time)`

	dailyDistinctSchema = `
-- AUTOGENERATED BY VIEW-GENERATOR, DO NOT MODIFY
CREATE OR REPLACE VIEW daily_distinct_%s_%s AS 
WITH
	filter AS (
	SELECT
		*
	, "max"(report_time) OVER (PARTITION BY %s, report_date) oldest_report_time
	FROM
		%s_%s
) 
SELECT *
FROM
	filter
WHERE (report_time = oldest_report_time)`
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}
	athenaClient := athena.NewFromConfig(cfg)

	for cloud, serviceMapping := range catalog.DatasourceModels {
		viewDirectory := fmt.Sprintf("./view_queries/%s/", cloud)

		if _, err := os.Stat(viewDirectory); os.IsNotExist(err) {
			err := os.Mkdir(viewDirectory, 0755)
			if err != nil {
				panic(err)
			}
		}

		for service, datasourceMapping := range serviceMapping {
			serviceDirectory := fmt.Sprintf("%s%s/", viewDirectory, service)
			if _, err := os.Stat(serviceDirectory); os.IsNotExist(err) {
				err := os.Mkdir(serviceDirectory, 0755)
				if err != nil {
					panic(err)
				}
			}
			for datasource, model := range datasourceMapping {
				// get the primary key of the model
				modelValue := reflect.ValueOf(model)
				modelValue = reflect.Indirect(modelValue)
				modelType := modelValue.Type()

				var modelPrimaryKey string
				for i := 0; i < modelValue.NumField(); i++ {
					typeField := modelType.Field(i)
					tag := typeField.Tag

					tags, err := structtag.Parse(string(tag))
					if err != nil {
						continue
					}

					primaryKeyTag, err := tags.Get("inventory_primary_key")
					if err != nil {
						continue
					}
					if primaryKeyTag.Name != "true" {
						continue
					}

					parquetTag, err := tags.Get("parquet")
					if err != nil {
						panic(err)
					}
					modelPrimaryKey = strings.Split(parquetTag.Name, "=")[1]
				}

				if modelPrimaryKey == "" {
					logrus.WithFields(logrus.Fields{
						"cloud":      cloud,
						"service":    service,
						"datasource": datasource,
					}).Warning("No primary key found for datasource")
					continue
				}

				dailyFilename := fmt.Sprintf("%s%s_daily.sql", serviceDirectory, datasource)
				dailyFile, err := os.Create(dailyFilename)
				if err != nil {
					panic(err)
				}
				dailyQuery := fmt.Sprintf(dailySchema, service, datasource, service, datasource)
				dailyFile.WriteString(dailyQuery)
				dailyFile.Close()

				distinctFilename := fmt.Sprintf("%s%s_distinct.sql", serviceDirectory, datasource)
				distinctFile, err := os.Create(distinctFilename)
				if err != nil {
					panic(err)
				}
				distinctQuery := fmt.Sprintf(distinctSchema, service, datasource, modelPrimaryKey, service, datasource)
				distinctFile.WriteString(distinctQuery)
				distinctFile.Close()

				currentFilename := fmt.Sprintf("%s%s_current.sql", serviceDirectory, datasource)
				currentFile, err := os.Create(currentFilename)
				if err != nil {
					panic(err)
				}
				currentQuery := fmt.Sprintf(currentSchema, service, datasource, service, datasource)
				currentFile.WriteString(currentQuery)
				currentFile.Close()

				dailyDistinctFilename := fmt.Sprintf("%s%s_daily_distinct.sql", serviceDirectory, datasource)
				dailyDistinctFile, err := os.Create(dailyDistinctFilename)
				if err != nil {
					panic(err)
				}
				dailyDistinctQuery := fmt.Sprintf(dailyDistinctSchema, service, datasource, modelPrimaryKey, service, datasource)
				dailyDistinctFile.WriteString(dailyDistinctQuery)
				dailyDistinctFile.Close()

				_, err = athenaClient.StartQueryExecution(context.TODO(), &athena.StartQueryExecutionInput{
					QueryString: aws.String(dailyQuery),
					QueryExecutionContext: &types.QueryExecutionContext{
						Catalog:  aws.String("AwsGlueCatalog"),
						Database: aws.String("cloud-inventory"),
					},
					WorkGroup: aws.String("primary"),
				})
				if err != nil {
					panic(err)
				}

				_, err = athenaClient.StartQueryExecution(context.TODO(), &athena.StartQueryExecutionInput{
					QueryString: aws.String(distinctQuery),
					QueryExecutionContext: &types.QueryExecutionContext{
						Catalog:  aws.String("AwsGlueCatalog"),
						Database: aws.String("cloud-inventory"),
					},
					WorkGroup: aws.String("primary"),
				})
				if err != nil {
					panic(err)
				}

				_, err = athenaClient.StartQueryExecution(context.TODO(), &athena.StartQueryExecutionInput{
					QueryString: aws.String(currentQuery),
					QueryExecutionContext: &types.QueryExecutionContext{
						Catalog:  aws.String("AwsGlueCatalog"),
						Database: aws.String("cloud-inventory"),
					},
					WorkGroup: aws.String("primary"),
				})
				if err != nil {
					panic(err)
				}

				_, err = athenaClient.StartQueryExecution(context.TODO(), &athena.StartQueryExecutionInput{
					QueryString: aws.String(dailyDistinctQuery),
					QueryExecutionContext: &types.QueryExecutionContext{
						Catalog:  aws.String("AwsGlueCatalog"),
						Database: aws.String("cloud-inventory"),
					},
					WorkGroup: aws.String("primary"),
				})
				if err != nil {
					panic(err)
				}

				time.Sleep(time.Millisecond * 500)
			}
		}
	}
}
