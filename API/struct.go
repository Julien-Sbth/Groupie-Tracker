package API

type Artist struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	Albums         string      `json:"albums"`
	Tracks         string      `json:"tracks"`
	Self           string      `json:"self"`
	ImageURL       string      `json:"image"`
	CreationDate   int64       `json:"creationDate"`
	Members        []string    `json:"members"`
	FirstAlbum     string      `json:"firstAlbum"`
	ConcertDate    string      `json:"concertDate"`
	Locations      interface{} `json:"locations"`
	Dates          interface{} `json:"dates"`
	Relations      interface{} `json:"relations"`
	FirstAlbumDate string
}

type SearchResult struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	FirstAlbum   string      `json:"firstAlbum"`
	CreationDate int64       `json:"creationDate"`
	Members      []string    `json:"members"`
	Locations    interface{} `json:"locations"`
	Image        string      `json:"image"`
}
