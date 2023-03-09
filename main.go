package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
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
	FirstAlbum     string      `json:"firstAlbum"`
	ConcertDate    string      `json:"concertDate"`
	Locations      interface{} `json:"locations"`
	Dates          interface{} `json:"dates"`
	Relations      interface{} `json:"relations"`
	FirstAlbumDate string
	Location       string
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
type Band struct {
	Name           string
	Members        []string
	Location       string
	FirstAlbumDate string
	Page           string
	Locations      string
	FirstAlbum     string
}

var result []Band

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/index/{id}", artistDetailsHandler).Methods("GET")
	r.HandleFunc("/lofi/{id}", Lofi).Methods("GET")
	r.HandleFunc("/search", searchHandler).Methods("GET")
	r.HandleFunc("/date/", datesHandler).Methods("GET")
	r.HandleFunc("/date/", getLocations).Methods("GET")
	r.HandleFunc("/date/", GetRelation).Methods("GET")

	fs := http.FileServer(http.Dir("front-end/css"))
	cssHandler := http.StripPrefix("/css/", fs)
	r.PathPrefix("/css/").Handler(cssHandler)
	fs = http.FileServer(http.Dir("front-end"))
	fileHandler := http.StripPrefix("/front-end/", fs)
	r.PathPrefix("/front-end/").Handler(fileHandler)
	fmt.Println("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.FormValue("search")
	var artists []Artist
	suggestions := autoCompleteSuggestions(artists, searchTerm)
	fmt.Println(suggestions) // Juste pour vérifier les suggestions dans la console

	// Retrieve artists from API
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Filter artists by search term
	var results []Artist
	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(searchTerm)) {
			results = append(results, artist)
		}
	}

	// If no artist matches the search term, display an error message
	if len(results) == 0 {
		fmt.Fprintf(w, "<head><style> body { font-family: 'Lato', sans-serif; text-align: center; background-color: #f2f2f2; } h1 { color: #3c3c3c; font-size: 3em; margin-bottom: 20px; } p { color: #3c3c3c; font-size: 2em; margin-top: 30px; margin-bottom: 10px; } </style></head>")
		fmt.Fprintf(w, "<h1>Search results for '%s'</h1>", searchTerm)
		fmt.Fprint(w, "<p>No artist found.</p>")
		return
	}

	// Display search results in HTML
	fmt.Fprintf(w, "<head><style> body { font-family: 'Lato', sans-serif; text-align: center; background-color: #f2f2f2; } h1 { color: #3c3c3c; font-size: 3em; margin-bottom: 20px; } ul { margin: 0; padding: 0; list-style: none; text-align: center; } ul li { font-size: 1.5em; margin-bottom: 10px; display: inline-block; } img { border-radius: 25px }</style></head>")
	fmt.Fprintf(w, "<h1>Search results for '%s'</h1>", searchTerm)
	fmt.Fprint(w, "<ul>\n")
	for _, artist := range results {
		fmt.Fprintf(w, "<li id=\"artist%d\">\n", artist.ID)
		fmt.Fprintf(w, "<div class=\"artiste-box\">\n")
		fmt.Fprintf(w, "<img src=\"%s\" alt=\"%s\" />\n", artist.ImageURL, artist.Name)
		fmt.Fprintf(w, "<div class=\"info\">\n")
		fmt.Fprintf(w, "<h2>%s Membre</h2>\n", artist.Name)
		fmt.Fprintf(w, "<p>Albums : %s</p>\n", artist.Albums)
		fmt.Fprintf(w, "<p>Pistes : %s</p>\n", artist.Tracks)
		fmt.Fprintf(w, "<p><a href=\"/index/%d\">Page de %s</a></p>\n", artist.ID, artist.Name)
		fmt.Fprintf(w, "</div>\n") // end div.info
		fmt.Fprintf(w, "</div>\n") // end div.artiste-box
		fmt.Fprintf(w, "</li>\n")
	}
}
func autoCompleteSuggestions(artists []Artist, search string) []string {
	var suggestions []string
	for _, artist := range artists {
		if strings.HasPrefix(strings.ToLower(artist.Name), strings.ToLower(search)) {
			suggestions = append(suggestions, artist.Name+" - artist/band")
		}
		for _, member := range artist.Members {
			if strings.HasPrefix(strings.ToLower(member), strings.ToLower(search)) {
				suggestions = append(suggestions, "Membre - "+member)
			}
		}
		if strings.HasPrefix(strings.ToLower(artist.Location), strings.ToLower(search)) {
			suggestions = append(suggestions, artist.Location+" - location")
		}
		if strings.HasPrefix(strings.ToLower(artist.FirstAlbumDate), strings.ToLower(search)) {
			suggestions = append(suggestions, artist.FirstAlbumDate+" - first album date")
			if strings.HasPrefix(strings.ToLower(strconv.FormatInt(artist.CreationDate, 10)), strings.ToLower(search)) {
				suggestions = append(suggestions, strconv.FormatInt(artist.CreationDate, 10)+" - creation date")
			}
		}
	}
	return suggestions
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
	relationsURL := fmt.Sprintf("https://groupietrackers.herokuapp.com/api/relation/%s", id)

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
func Lofi(w http.ResponseWriter, r *http.Request) {
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

	tmpl, err := template.ParseFiles("front-end/lofi.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tmpl.Execute(w, artist)
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
