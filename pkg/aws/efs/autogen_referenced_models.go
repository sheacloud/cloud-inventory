//AUTOGENERATED CODE DO NOT EDIT
package efs

import (
	"time"
)

type FileSystemSize struct {
	Value           int64      `bson:"value,omitempty" dynamodbav:"value,omitempty" json:"value,omitempty" diff:"value"`
	Timestamp       *time.Time `bson:"timestamp,omitempty" dynamodbav:"timestamp,unixtime,omitempty" json:"timestamp,omitempty" diff:"timestamp"`
	ValueInIA       int64      `bson:"value_in_ia,omitempty" dynamodbav:"value_in_ia,omitempty" json:"value_in_ia,omitempty" diff:"value_in_ia"`
	ValueInStandard int64      `bson:"value_in_standard,omitempty" dynamodbav:"value_in_standard,omitempty" json:"value_in_standard,omitempty" diff:"value_in_standard"`
}

type Tag struct {
	Key   string `bson:"key,omitempty" dynamodbav:"key,omitempty" json:"key,omitempty" diff:"key"`
	Value string `bson:"value,omitempty" dynamodbav:"value,omitempty" json:"value,omitempty" diff:"value"`
}