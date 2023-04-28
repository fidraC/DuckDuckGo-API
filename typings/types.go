package typings

type Search struct {
	Query     string `json:"query"`
	Region    string `json:"region"`
	TimeRange string `json:"time_range"`
	Limit     int    `json:"limit"`
}

type Result struct {
	Title   string `json:"title"`
	Link    string `json:"link"`
	Snippet string `json:"snippet"`
}
