package routes

type Diff struct {
	Type string      `json:"type"`
	Path []string    `json:"path"`
	From interface{} `json:"from"`
	To   interface{} `json:"to"`
}

type AwsMetadata struct {
	Services []string `json:"services"`
}

type AwsResourceMetadata struct {
	DateTimes     []string `json:"datetimes"`
	IdField       string   `json:"id_field"`
	DisplayFields []string `json:"display_fields"`
}

type AwsServiceMetadata struct {
	Resources []string `json:"resources"`
}
