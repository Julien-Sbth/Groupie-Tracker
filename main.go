package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type Artist struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	Albums         string      `json:"albums"`
	Tracks         string      `json:"tracks"`
	Self           string      `json:"self"`
	ImageURL       string      `json:"image"`
	CreationDate   int64       `json:"creationDate"`
	Members        []string    `json:"members"`
	FirstAlbum     string      `json:"FirstAlbum"`
	Locations      interface{} `json:"locations"`
	DatesLocations interface{} `json:"dateslocations"`
	ConcertDate    string      `json:"ConcertDate"`
	Dates          interface{} `json:"dates"`
	Latitude       float64     `json:"lat"`
	Longitude      float64     `json:"long"`
	Country        string      `json:"country"`
	City           string      `json:"city"`
	ConcertID      string      `json:"ConcertID"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/index/{id}", artistDetailsHandler).Methods("GET")
	r.HandleFunc("/date/", datesHandler).Methods("GET")
	r.HandleFunc("/date/", getLocations).Methods("GET")
	r.HandleFunc("/date/", GetRelation).Methods("GET")
	fs := http.FileServer(http.Dir("front-end/css"))
	cssHandler := http.StripPrefix("/css/", fs)
	r.PathPrefix("/css/").Handler(cssHandler)
	fs = http.FileServer(http.Dir("front-end"))
	fileHandler := http.StripPrefix("/front-end/", fs)
	r.PathPrefix("/front-end/").Handler(fileHandler)
	log.Fatal(http.ListenAndServe(":8080", r))
}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Erreur lors de la requête HTTP :", err)
		return
	}
	defer resp.Body.Close()
	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse JSON :", err)
		return
	}
	for i := range artists {
		artists[i].ID = i + 1
	}
	tmpl, err := template.ParseFiles("front-end/accueil.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, artists)
}
func getArtistWithDatesLocationsAndRelations(id string) (*Artist, error) {
	artistURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%s", id)
	datesLocationsURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%s", id)
	relationsURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relations/%s", id)

	artistResp, err := http.Get(artistURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get artist: %v", err)
	}
	defer artistResp.Body.Close()

	var artist Artist
	err = json.NewDecoder(artistResp.Body).Decode(&artist)
	if err != nil {
		return nil, fmt.Errorf("failed to decode artist response: %v", err)
	}

	datesLocationsResp, err := http.Get(datesLocationsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get dates and locations: %v", err)
	}
	defer datesLocationsResp.Body.Close()

	var datesLocations []DateLocation
	err = json.NewDecoder(datesLocationsResp.Body).Decode(&datesLocations)
	if err != nil {
		return nil, fmt.Errorf("failed to decode dates and locations response: %v", err)
	}

	relationsResp, err := http.Get(relationsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get relations: %v", err)
	}
	defer relationsResp.Body.Close()

	var relations interface{}
	err = json.NewDecoder(relationsResp.Body).Decode(&relations)
	if err != nil {
		return nil, fmt.Errorf("failed to decode relations response: %v", err)
	}

	for i := range relations {
		for j := range datesLocations {
			if relations[i].ConcertID == datesLocations[j].Location.ID {
				relations[i].Date = datesLocations[j].Date
				relations[i].Location = datesLocations[j].Location
			}
		}
	}

	artist.DatesLocations = relations
	artist.Dates = datesLocations

	return &artist, nil
}

func getLocations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing artist ID", http.StatusBadRequest)
		return
	}

	_, err := getArtistWithDatesLocationsAndRelations(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// call locationsHandler
	http.Redirect(w, r, fmt.Sprintf("/locations/%s", id), http.StatusSeeOther)
}
func datesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing artist ID", http.StatusBadRequest)
		return
	}

	_, err := getArtistWithDatesLocationsAndRelations(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// call locationsHandler
	http.Redirect(w, r, fmt.Sprintf("/locations/%s", id), http.StatusSeeOther)
}
func GetRelation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing artist ID", http.StatusBadRequest)
		return
	}

	_, err := getArtistWithDatesLocationsAndRelations(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// call locationsHandler
	http.Redirect(w, r, fmt.Sprintf("/locations/%s", id), http.StatusSeeOther)
}
func artistDetailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing artist ID", http.StatusBadRequest)
		return
	}

	artist, err := getArtistWithDatesLocationsAndRelations(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("front-end/artiste.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, artist)
}
