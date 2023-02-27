package backend

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type DaftPunk struct {
	ID       int      `json:"id"` // changed "ID" to "Id"
	Artist   Artist   `json:"artist"`
	Location Location `json:"location"`
	Date     Date     `json:"date"`
}

var daftPunk = []DaftPunk{
	{
		ID: 4,
		Artist: Artist{
			Name:       "Daft Punk",
			Image:      "band1.png",
			YearActive: 2021,
			FirstAlbum: "Album 1",
			Members:    []string{"Thomas Bangalter", "Guy-Manuel de Homem-Christo"},
		},
		Location: Location{
			Name:    "Location 1",
			Country: "Country 1",
			City:    "City 1",
			Venue:   "Venue 1",
			Date:    "2022-04-01",
		},
		Date: Date{
			Name:    "Concert 1",
			Concert: "2022-04-01",
		},
	},
}

func HandleInfoJSONDaftPunk(w http.ResponseWriter, r *http.Request) {
	// Set the content type of the response to application/json
	w.Header().Set("Content-Type", "application/json")

	// Loop through the relations to find the one with ID 1
	for _, rel := range daftPunk {
		if rel.ID == 4 {
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
func HandleInfoDaftPunk(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("front-end/InfoGroup/InfoDaftPunk.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	type Data struct {
		ArtistName string
		Members    []string
	}

	data := Data{
		ArtistName: "Daft Punk",
		Members:    []string{"Daft Punk"},
	}
	if err := tmpl.Execute(w, data); err != nil {
		return
	}
}
