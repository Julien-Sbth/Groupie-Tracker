package backend

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type NWA struct {
	ID       int      `json:"id"` // changed "ID" to "Id"
	Artist   Artist   `json:"artist"`
	Location Location `json:"location"`
	Date     Date     `json:"date"`
}

var nwa = []NWA{
	{
		ID: 7, // added the ID field and set it to 1
		Artist: Artist{
			Name:       "heheheha",
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
}

func HandleInfoJSONNWA(w http.ResponseWriter, r *http.Request) {
	// Set the content type of the response to application/json
	w.Header().Set("Content-Type", "application/json")

	// Loop through the relations to find the one with ID 1
	for _, rel := range nwa {
		if rel.ID == 7 {
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
func HandleInfoNWA(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("front-end/InfoGroup/infosNWA.html")
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
