package app

type Value struct {
	Artists   string `json:"artists"`
	Dates     string `json:"dates"`
	Locations string `json:"locations"`
	Relation  string `json:"relation"`
}

type PageData struct {
	Photo []Photo
}

type Photo struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}
