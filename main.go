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
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Albums       string      `json:"albums"`
	Tracks       string      `json:"tracks"`
	Self         string      `json:"self"`
	ImageURL     string      `json:"image"`
	CreationDate int64       `json:"creationDate"`
	Members      []string    `json:"members"`
	FirstAlbum   string      `json:"FirstAlbum"`
	Locations    interface{} `json:"locations"`
	ConcertDate  string      `json:"ConcertDate"`
	Dates        interface{} `json:"dates"`
	Relations    interface{} `json:"relations"`
	Latitude     float64     `json:"lat"`
	Longitude    float64     `json:"long"`
	Country      string      `json:"country"`
	City         string      `json:"city"`
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
func getArtistWithDatesAndLocations(id string) (*Artist, error) {
	artistURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/artists/%s", id)
	datesURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/dates/%s", id)
	locationsURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/locations/%s", id)
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

	datesResp, err := http.Get(datesURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get dates: %v", err)
	}
	defer datesResp.Body.Close()

	var dates interface{}
	err = json.NewDecoder(datesResp.Body).Decode(&dates)
	if err != nil {
		return nil, fmt.Errorf("failed to decode dates response: %v", err)
	}

	locationsResp, err := http.Get(locationsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get locations: %v", err)
	}
	defer locationsResp.Body.Close()

	var locations interface{}
	err = json.NewDecoder(locationsResp.Body).Decode(&locations)
	if err != nil {
		return nil, fmt.Errorf("failed to decode locations response: %v", err)
	}

	relationResp, err := http.Get(relationsURL)
	if err != nil {
		return nil, fmt.Errorf("failed to get dates: %v", err)
	}
	defer relationResp.Body.Close()

	var relations interface{}
	err = json.NewDecoder(relationResp.Body).Decode(&relations)
	if err != nil {
		return nil, fmt.Errorf("failed to decode dates response: %v", err)
	}
	artist.Dates = dates
	artist.Locations = locations
	artist.Relations = relations

	return &artist, nil
}
func getLocations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "Missing artist ID", http.StatusBadRequest)
		return
	}

	_, err := getArtistWithDatesAndLocations(id)
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

	_, err := getArtistWithDatesAndLocations(id)
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

	_, err := getArtistWithDatesAndLocations(id)
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

	artist, err := getArtistWithDatesAndLocations(id)
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
