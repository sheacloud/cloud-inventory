package main

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/athena"
	"github.com/sheacloud/cloud-inventory/internal/inventory"
	"github.com/sheacloud/cloud-inventory/internal/viewgen"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	logrusLevels = map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
	}
)

func initOptions() {
	viper.SetEnvPrefix("cloud_inventory")
	viper.AutomaticEnv()

	viper.BindEnv("log_level")
	viper.SetDefault("log_level", "info")

	viper.BindEnv("log_caller")
	viper.SetDefault("log_caller", false)

	viper.BindEnv("glue_database_name")

	viper.BindEnv("athena_workgroup_name")
}

func initLogging() {
	logrus.SetLevel(logrusLevels[viper.GetString("log_level")])
	logrus.SetReportCaller(viper.GetBool("log_caller"))
}

func validateOptions() {
	if viper.GetString("glue_database_name") == "" {
		panic("glue_database_name is required")
	}

	if viper.GetString("athena_workgroup_name") == "" {
		panic("athena_workgroup_name is required")
	}
}

func init() {
	initOptions()
	initLogging()
	validateOptions()
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		panic(err)
	}
	athenaClient := athena.NewFromConfig(cfg)

	viewProcessor := viewgen.NewViewJobProcessor(5, 1000)
	viewProcessor.Start()

	errorChan := make(chan error)
	numErrors := 0
	go func() {
		for err := range errorChan {
			numErrors++
			fmt.Println(err)
		}
	}()

	database := viper.GetString("glue_database_name")
	workgroup := viper.GetString("athena_workgroup_name")

	for _, service := range inventory.AwsCatalog {

		for _, resource := range service.Resources {
			views, err := viewgen.GetModelViews("aws", service.ServiceName, resource.ResourceName, resource.ResourceModel)
			if err != nil {
				panic(err)
			}

			for _, view := range views {
				viewProcessor.AddJob(viewgen.NewViewJob(context.TODO(), view.Name, view.View, "AwsGlueCatalog", database, workgroup, athenaClient, errorChan))
			}

		}
	}

	files, err := os.ReadDir("./custom_views/")
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		fileData, err := os.ReadFile("./custom_views/" + file.Name())
		if err != nil {
			panic(err)
		}
		viewProcessor.AddJob(viewgen.NewViewJob(context.TODO(), file.Name(), string(fileData), "AwsGlueCatalog", database, workgroup, athenaClient, errorChan))

	}

	viewProcessor.WaitForCompletion()
	close(errorChan)

	if numErrors > 0 {
		panic("Errors creating views")
	}
}
