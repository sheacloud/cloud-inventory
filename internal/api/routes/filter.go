package routes

type AwsFilterableResource interface {
	GetAccountId() string
	GetRegion() string
}
