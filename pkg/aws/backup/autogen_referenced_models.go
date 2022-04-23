//AUTOGENERATED CODE DO NOT EDIT
package backup

type AdvancedBackupSetting struct {
	BackupOptions map[string]string `bson:"backup_options,omitempty" ion:"backup_options" dynamodbav:"backup_options,omitempty" parquet:"name=backup_options,type=MAP,keytype=BYTE_ARRAY,valuetype=BYTE_ARRAY,keyconvertedtype=UTF8,valueconvertedtype=UTF8" json:"backup_options,omitempty" diff:"backup_options"`
	ResourceType  string            `bson:"resource_type,omitempty" ion:"resource_type" dynamodbav:"resource_type,omitempty" parquet:"name=resource_type,type=BYTE_ARRAY,convertedtype=UTF8" json:"resource_type,omitempty" diff:"resource_type"`
}

type BackupSelectionsListMember struct {
	BackupPlanId     string `bson:"backup_plan_id,omitempty" ion:"backup_plan_id" dynamodbav:"backup_plan_id,omitempty" parquet:"name=backup_plan_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"backup_plan_id,omitempty" diff:"backup_plan_id"`
	CreationDate     int64  `bson:"creation_date,omitempty" ion:"creation_date" dynamodbav:"creation_date,omitempty" parquet:"name=creation_date,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"creation_date,omitempty" diff:"creation_date"`
	CreatorRequestId string `bson:"creator_request_id,omitempty" ion:"creator_request_id" dynamodbav:"creator_request_id,omitempty" parquet:"name=creator_request_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"creator_request_id,omitempty" diff:"creator_request_id"`
	IamRoleArn       string `bson:"iam_role_arn,omitempty" ion:"iam_role_arn" dynamodbav:"iam_role_arn,omitempty" parquet:"name=iam_role_arn,type=BYTE_ARRAY,convertedtype=UTF8" json:"iam_role_arn,omitempty" diff:"iam_role_arn"`
	SelectionId      string `bson:"selection_id,omitempty" ion:"selection_id" dynamodbav:"selection_id,omitempty" parquet:"name=selection_id,type=BYTE_ARRAY,convertedtype=UTF8" json:"selection_id,omitempty" diff:"selection_id"`
	SelectionName    string `bson:"selection_name,omitempty" ion:"selection_name" dynamodbav:"selection_name,omitempty" parquet:"name=selection_name,type=BYTE_ARRAY,convertedtype=UTF8" json:"selection_name,omitempty" diff:"selection_name"`
}
