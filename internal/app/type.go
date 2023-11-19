package app

type Value struct {
	Artists   string `json:"artists"`
	Dates     string `json:"dates"`
	Locations string `json:"locations"`
	Relation  string `json:"relation"`
}

type PageData struct {
	Artist []Artist
}

type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}
type LOcDatRel struct {
	Location map[string][]Locations
	Date     map[string][]Dates
	Relation map[string][]Relations
}

type Global struct {
	Artist       PageData
	DonneRestant LOcDatRel
}
type Locations struct {
	Id       int      `json:"id"`
	Location []string `json:"locations"`
}
type Relations struct {
	Id        int                 `json:"id"`
	Relations map[string][]string `json:"datesLocations"`
}

type Dates struct {
	Id    int      `json:"id"`
	Dates []string `json:"dates"`
}

type Info struct {
	Artist   Artist
	Date     Dates
	Location Locations
	Relation Relations
}
