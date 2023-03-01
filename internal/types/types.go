package types

type Search struct {
	Query     string `json:"query"`
	Region    string `json:"region"`
	TimeRange string `json:"time_range"`
}
