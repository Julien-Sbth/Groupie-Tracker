package backend

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type Biggies struct {
	ID       int      `json:"id"` // changed "ID" to "Id"
	Artist   Artist   `json:"artist"`
	Location Location `json:"location"`
	Date     Date     `json:"date"`
}

var biggies = []Biggies{
	{
		ID: 1, // added the ID field and set it to 1
		Artist: Artist{
			Name:       "Band 1",
			Image:      "band1.png",
			YearActive: 2000,
			FirstAlbum: "Album 1",
			Members:    []string{"Member 1", "Member 2", "Member 3"},
		},
		Location: Location{
			Name:        "Location 1",
			Country:     "Country 1",
			City:        "City 1",
			Venue:       "Venue 1",
			Date:        "2022-04-01",
			SoldOut:     true,
			TicketPrice: 50.00,
		},
		Date: Date{
			Name:       "Concert 1",
			Concert:    "2022-04-01",
			SoldOut:    true,
			TicketLeft: 0,
		},
	},
	{
		ID: 1, // added the ID field and set it to 2
		Artist: Artist{
			Name:       "Band 2",
			Image:      "band2.png",
			YearActive: 2010,
			FirstAlbum: "Album 2",
			Members:    []string{"Member 4", "Member 5"},
		},
		Location: Location{
			Name:        "Location 2",
			Country:     "Country 2",
			City:        "City 2",
			Venue:       "Venue 2",
			Date:        "2022-04-02",
			SoldOut:     false,
			TicketPrice: 30.00,
		},
		Date: Date{
			Name:       "Concert 2",
			Concert:    "2022-04-02",
			SoldOut:    false,
			TicketLeft: 100,
		},
	},
}

func HandleInfoJSONBig(w http.ResponseWriter, r *http.Request) {
	// Set the content type of the response to application/json
	w.Header().Set("Content-Type", "application/json")

	// Loop through the relations to find the one with ID 1
	for _, rel := range biggies {
		if rel.ID == 1 {
			// Marshal the relation to JSON and write it to the response
			jsonBytes, err := json.Marshal(rel)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(jsonBytes)
			return
		}
	}

	// If no relation with ID 1 was found, return a 404 error
	http.NotFound(w, r)
}
func HandleInfoNotorious(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("front-end/InfoGroup/InfoNotoriousBig.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type Data struct {
		ArtistName string
		Members    []string
	}

	data := Data{
		ArtistName: "Notorius Big",
		Members:    []string{"Notorius Big"},
	}
	if err := tmpl.Execute(w, data); err != nil {
		return
	}
}
