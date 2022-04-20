package meta

import "time"

type InventoryResults struct {
	InventoryUUID    string    `bson:"_id" dynamodbav:"_id"`
	Cloud            string    `bson:"cloud" dynamodbav:"cloud"`
	Service          string    `bson:"service" dynamodbav:"service"`
	Resource         string    `bson:"resource" dynamodbav:"resource"`
	AccountId        string    `bson:"account_id" dynamodbav:"account_id"`
	Region           string    `bson:"region" dynamodbav:"region"`
	FetchedResources int       `bson:"fetched_resources" dynamodbav:"fetched_resources"`
	FailedResources  int       `bson:"failed_resources" dynamodbav:"failed_resources"`
	HadErrors        bool      `bson:"had_errors" dynamodbav:"had_errors"`
	ReportTime       time.Time `bson:"report_time" dynamodbav:"report_time,unixtime"`
}
