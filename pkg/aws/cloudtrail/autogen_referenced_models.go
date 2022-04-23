//AUTOGENERATED CODE DO NOT EDIT
// This file is automatically generated from /internal/codegen/templates/aws_referenced_resource_file.tmpl
package cloudtrail

type GetTrailStatusOutput struct {
	IsLogging                          bool   `bson:"is_logging,omitempty" ion:"is_logging" dynamodbav:"is_logging" parquet:"name=is_logging,type=BOOLEAN" json:"is_logging,omitempty" diff:"is_logging"`
	LatestCloudWatchLogsDeliveryError  string `bson:"latest_cloud_watch_logs_delivery_error,omitempty" ion:"latest_cloud_watch_logs_delivery_error" dynamodbav:"latest_cloud_watch_logs_delivery_error,omitempty" parquet:"name=latest_cloud_watch_logs_delivery_error,type=BYTE_ARRAY,convertedtype=UTF8" json:"latest_cloud_watch_logs_delivery_error,omitempty" diff:"latest_cloud_watch_logs_delivery_error"`
	LatestCloudWatchLogsDeliveryTime   int64  `bson:"latest_cloud_watch_logs_delivery_time,omitempty" ion:"latest_cloud_watch_logs_delivery_time" dynamodbav:"latest_cloud_watch_logs_delivery_time,omitempty" parquet:"name=latest_cloud_watch_logs_delivery_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"latest_cloud_watch_logs_delivery_time,omitempty" diff:"latest_cloud_watch_logs_delivery_time"`
	LatestDeliveryAttemptSucceeded     string `bson:"latest_delivery_attempt_succeeded,omitempty" ion:"latest_delivery_attempt_succeeded" dynamodbav:"latest_delivery_attempt_succeeded,omitempty" parquet:"name=latest_delivery_attempt_succeeded,type=BYTE_ARRAY,convertedtype=UTF8" json:"latest_delivery_attempt_succeeded,omitempty" diff:"latest_delivery_attempt_succeeded"`
	LatestDeliveryAttemptTime          string `bson:"latest_delivery_attempt_time,omitempty" ion:"latest_delivery_attempt_time" dynamodbav:"latest_delivery_attempt_time,omitempty" parquet:"name=latest_delivery_attempt_time,type=BYTE_ARRAY,convertedtype=UTF8" json:"latest_delivery_attempt_time,omitempty" diff:"latest_delivery_attempt_time"`
	LatestDeliveryError                string `bson:"latest_delivery_error,omitempty" ion:"latest_delivery_error" dynamodbav:"latest_delivery_error,omitempty" parquet:"name=latest_delivery_error,type=BYTE_ARRAY,convertedtype=UTF8" json:"latest_delivery_error,omitempty" diff:"latest_delivery_error"`
	LatestDeliveryTime                 int64  `bson:"latest_delivery_time,omitempty" ion:"latest_delivery_time" dynamodbav:"latest_delivery_time,omitempty" parquet:"name=latest_delivery_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"latest_delivery_time,omitempty" diff:"latest_delivery_time"`
	LatestDigestDeliveryError          string `bson:"latest_digest_delivery_error,omitempty" ion:"latest_digest_delivery_error" dynamodbav:"latest_digest_delivery_error,omitempty" parquet:"name=latest_digest_delivery_error,type=BYTE_ARRAY,convertedtype=UTF8" json:"latest_digest_delivery_error,omitempty" diff:"latest_digest_delivery_error"`
	LatestDigestDeliveryTime           int64  `bson:"latest_digest_delivery_time,omitempty" ion:"latest_digest_delivery_time" dynamodbav:"latest_digest_delivery_time,omitempty" parquet:"name=latest_digest_delivery_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"latest_digest_delivery_time,omitempty" diff:"latest_digest_delivery_time"`
	LatestNotificationAttemptSucceeded string `bson:"latest_notification_attempt_succeeded,omitempty" ion:"latest_notification_attempt_succeeded" dynamodbav:"latest_notification_attempt_succeeded,omitempty" parquet:"name=latest_notification_attempt_succeeded,type=BYTE_ARRAY,convertedtype=UTF8" json:"latest_notification_attempt_succeeded,omitempty" diff:"latest_notification_attempt_succeeded"`
	LatestNotificationAttemptTime      string `bson:"latest_notification_attempt_time,omitempty" ion:"latest_notification_attempt_time" dynamodbav:"latest_notification_attempt_time,omitempty" parquet:"name=latest_notification_attempt_time,type=BYTE_ARRAY,convertedtype=UTF8" json:"latest_notification_attempt_time,omitempty" diff:"latest_notification_attempt_time"`
	LatestNotificationError            string `bson:"latest_notification_error,omitempty" ion:"latest_notification_error" dynamodbav:"latest_notification_error,omitempty" parquet:"name=latest_notification_error,type=BYTE_ARRAY,convertedtype=UTF8" json:"latest_notification_error,omitempty" diff:"latest_notification_error"`
	LatestNotificationTime             int64  `bson:"latest_notification_time,omitempty" ion:"latest_notification_time" dynamodbav:"latest_notification_time,omitempty" parquet:"name=latest_notification_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"latest_notification_time,omitempty" diff:"latest_notification_time"`
	StartLoggingTime                   int64  `bson:"start_logging_time,omitempty" ion:"start_logging_time" dynamodbav:"start_logging_time,omitempty" parquet:"name=start_logging_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"start_logging_time,omitempty" diff:"start_logging_time"`
	StopLoggingTime                    int64  `bson:"stop_logging_time,omitempty" ion:"stop_logging_time" dynamodbav:"stop_logging_time,omitempty" parquet:"name=stop_logging_time,type=INT64,convertedtype=TIMESTAMP_MILLIS" json:"stop_logging_time,omitempty" diff:"stop_logging_time"`
	TimeLoggingStarted                 string `bson:"time_logging_started,omitempty" ion:"time_logging_started" dynamodbav:"time_logging_started,omitempty" parquet:"name=time_logging_started,type=BYTE_ARRAY,convertedtype=UTF8" json:"time_logging_started,omitempty" diff:"time_logging_started"`
	TimeLoggingStopped                 string `bson:"time_logging_stopped,omitempty" ion:"time_logging_stopped" dynamodbav:"time_logging_stopped,omitempty" parquet:"name=time_logging_stopped,type=BYTE_ARRAY,convertedtype=UTF8" json:"time_logging_stopped,omitempty" diff:"time_logging_stopped"`
}
