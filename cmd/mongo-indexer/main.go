package main

import (
	"context"
	"fmt"

	"github.com/sheacloud/cloud-inventory/internal/inventory"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	viper.BindEnv("mongo_uri")
}

func init() {
	initOptions()
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(viper.GetString("mongo_uri")))
	if err != nil {
		panic(err)
	}

	db := client.Database("cloud-inventory")

	for _, service := range inventory.AwsCatalog {
		for _, resource := range service.Resources {
			coll := db.Collection(fmt.Sprintf("aws.%s.%s", service.ServiceName, resource.ResourceName))
			indexes := coll.Indexes()
			indexModels := []mongo.IndexModel{
				{
					Keys: bson.D{
						{"report_time", -1},
					},
				},
				{
					Keys: bson.D{
						{resource.UniqueIdField, 1},
					},
				},
			}

			_, err := indexes.CreateMany(context.TODO(), indexModels)
			if err != nil {
				panic(err)
			}
			logrus.WithFields(logrus.Fields{
				"service":  service.ServiceName,
				"resource": resource.ResourceName,
			}).Info("created indexes")
		}
	}

	coll := db.Collection("meta.inventory_results")
	indexes := coll.Indexes()
	_, err = indexes.CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.D{
			{"report_time", -1},
		},
	})
	if err != nil {
		panic(err)
	}
	logrus.Info("created meta.inventory_results index")

	coll = db.Collection("meta.ingestion_timestamps")
	indexes = coll.Indexes()
	_, err = indexes.CreateOne(context.TODO(), mongo.IndexModel{
		Keys: bson.D{
			{"report_time", -1},
		},
	})
	if err != nil {
		panic(err)
	}
	logrus.Info("created meta.ingestion_timestamps index")
}
