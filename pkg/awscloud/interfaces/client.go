package interfaces

type AwsClient interface {
	EC2() EC2Client
	CloudWatchLogs() CloudWatchLogsClient
}
