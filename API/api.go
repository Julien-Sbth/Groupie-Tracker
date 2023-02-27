package API

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// Handle data for artist or relation

// handleLocations renvoie les informations sur les lieux de concert

// handleDates renvoie les informations sur les dates de concert

// Artist contient les informations sur un artiste
type Artist struct {
	Name       string   `json:"name"`
	Image      string   `json:"image"`
	YearActive int      `json:"yearActive"`
	FirstAlbum string   `json:"firstAlbum"`
	Members    []string `json:"members"`
	ID         string   `json:"id"`
}

// Location contient les informations sur un lieu de concert
type Location struct {
	Name        string  `json:"name"`
	Country     string  `json:"country"`
	City        string  `json:"city"`
	Venue       string  `json:"venue"`
	Date        string  `json:"date"`
	SoldOut     bool    `json:"soldOut"`
	TicketPrice float32 `json:"ticketPrice"`
}

// Date contient les informations sur une date de concert
type Date struct {
	Name       string `json:"name"`
	Date       string `json:"date"`
	SoldOut    bool   `json:"soldOut"`
	TicketLeft int    `json:"ticketLeft"`
}

// Relation contient les informations sur la relation entre un artiste, un lieu de concert et une date de concert
func HandleArtists(w http.ResponseWriter, r *http.Request) {
	var artists []Artist
	for _, relation := range relation {
		artists = append(artists, relation.Artist)
	}

	jsonData, err := json.Marshal(artists)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonData)
}
func HandleDates(w http.ResponseWriter, r *http.Request) {
	var dates []Date
	for _, relation := range relation {
		dates = append(dates, relation.Date)
	}

	jsonData, err := json.Marshal(dates)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

type Relation struct {
	Artist   Artist   `json:"artist"`
	Location Location `json:"location"`
	Date     Date     `json:"date"`
	Id       int      `json:"id"`
}

var relation = []Relation{
	{
		Artist: Artist{
			Name:       "Daft Punk",
			Image:      "band1.png",
			YearActive: 2021,
			FirstAlbum: "Album 1",
			Members:    []string{"test", "test"},
		},
		Location: Location{
			Name:    "Location 1",
			Country: "France",
			City:    "City 1",
			Venue:   "Venue 1",
			Date:    "2022-04-01",
		},
		Date: Date{
			Name: "Concert 1",
			Date: "2022-04-01",
		},
		Id: 1,
	},
}

/*
func HandleLocations(w http.ResponseWriter, r *http.Request) {
	var locations []Location
	for _, relation := range relation {
		locations = append(locations, relation.Location)
	}

	jsonData, err := json.Marshal(locations)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
*/

type Band struct {
	Name           string
	Members        []string
	Location       string
	FirstAlbumDate string
	Page           string
}

var data = []Band{
	{
		Name:           "2PAC",
		Members:        []string{"Tupac Shakur"},
		Location:       "United States, US",
		FirstAlbumDate: "1963",
		Page:           "/shakur",
	},
	{
		Name:           "Daft Punk",
		Members:        []string{"Daft Punk"},
		Location:       "France, Fr",
		FirstAlbumDate: "1973",
		Page:           "/daft",
	},
	{
		Name:           "Notorious",
		Members:        []string{"Notorious Big"},
		Location:       "United States, US",
		FirstAlbumDate: "1973",
		Page:           "/notorious",
	},
	{
		Name:           "Warren G",
		Members:        []string{"Warren G"},
		Location:       "United States, US",
		FirstAlbumDate: "1973",
		Page:           "/warren",
	},
	{
		Name:           "Lamar",
		Members:        []string{"Lamar"},
		Location:       "United States, US",
		FirstAlbumDate: "1973",
		Page:           "/lamar",
	},
	{
		Name:           "NWA",
		Members:        []string{"Eazy-E", "Dr Dre", "Ice Cube", "John Deacon"},
		Location:       "London, UK",
		FirstAlbumDate: "1973",
		Page:           "/nwa",
	},
	{
		Name:           "Halliday",
		Members:        []string{"Johnny Halliday"},
		Location:       "London, UK",
		FirstAlbumDate: "1973",
		Page:           "/halliday",
	},
	{
		Name:           "Dion",
		Members:        []string{"Celine Dion"},
		Location:       "London, UK",
		FirstAlbumDate: "1973",
		Page:           "/dion",
	},
	{
		Name:           "Soprano",
		Members:        []string{"Soprano"},
		Location:       "London, UK",
		FirstAlbumDate: "1973",
		Page:           "/soprano",
	},
	{
		Name:           "Kaaris",
		Members:        []string{"Kaaris"},
		Location:       "London, UK",
		FirstAlbumDate: "1973",
		Page:           "/kaaris",
	},
	{
		Name:           "Bigflo et Oli",
		Members:        []string{"Bigflo et Oli"},
		Location:       "London, UK",
		FirstAlbumDate: "1973",
		Page:           "/big",
	},
	{
		Name:           "SCH",
		Members:        []string{"SCH"},
		Location:       "London, UK",
		FirstAlbumDate: "1973",
		Page:           "/sch",
	},
	{
		Name:           "MAES",
		Members:        []string{"MAES"},
		Location:       "London, UK",
		FirstAlbumDate: "1973",
		Page:           "/maes",
	},
	{
		Name:           "Jul",
		Members:        []string{"Jul"},
		Location:       "London, UK",
		FirstAlbumDate: "1973",
		Page:           "/jul",
	},
	{
		Name:           "Naps",
		Members:        []string{"Naps"},
		Location:       "London, UK",
		FirstAlbumDate: "1973",
		Page:           "/naps",
	},
	{
		Name:           "Snoop",
		Members:        []string{"Snoop Dog"},
		Location:       "London, UK",
		FirstAlbumDate: "1973",
		Page:           "/snoop",
	},
	{
		Name:           "Eminem",
		Members:        []string{"Eminem"},
		Location:       "London, UK",
		FirstAlbumDate: "1973",
		Page:           "/eminem",
	},
	{
		Name:           "XXXtentation",
		Members:        []string{"XXXtentation"},
		Location:       "London, UK",
		FirstAlbumDate: "1973",
		Page:           "/tentation",
	},
}

func Search(query string) []Band {
	var results []Band
	for _, band := range data {
		if strings.Contains(strings.ToLower(band.Name), strings.ToLower(query)) {
			results = append(results, band)
		}
	}
	return results
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	// Parse the search term from the form
	searchTerm := r.FormValue("search")

	// Search the bands for matching results
	var results []Band
	for _, b := range data {
		if strings.Contains(strings.ToLower(b.Name), strings.ToLower(searchTerm)) {
			results = append(results, b)
			continue
		}
		for _, member := range b.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(searchTerm)) {
				results = append(results, b)
				break
			}
		}
		if strings.Contains(strings.ToLower(b.Location), strings.ToLower(searchTerm)) {
			results = append(results, b)
			continue
		}
		if strings.Contains(strings.ToLower(b.FirstAlbumDate), strings.ToLower(searchTerm)) {
			results = append(results, b)
			continue
		}
	}

	// Shuffle the results randomly
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(results), func(i, j int) {
		results[i], results[j] = results[j], results[i]
	})

	// Display the search results in the HTML
	fmt.Fprintf(w, "<head><style> body { font-family: 'Lato', sans-serif; text-align: center; background-color: #f2f2f2; } h1 { color: #3c3c3c; font-size: 3em; margin-bottom: 20px; } h2 { color: #3c3c3c; font-size: 2em; margin-top: 30px; margin-bottom: 10px; } ul { margin: 0; padding: 0; list-style: none; text-align: center; } ul li { font-size: 1.5em; margin-bottom: 10px; display: inline-block; } </style></head>")
	fmt.Fprintf(w, "<h1>Search results for '%s'</h1>", searchTerm)
	if len(results) == 0 {
		fmt.Fprint(w, "<p>No results found.</p>")
	} else {
		for _, b := range results {
			fmt.Fprintf(w, "<h2><a href=\"%s\">%s</a></h2>\n", b.Page, b.Name)
			fmt.Fprint(w, "<ul>\n")
			fmt.Fprintf(w, "<li>Members: %s</li>\n", strings.Join(b.Members, ", "))
			fmt.Fprint(w, "<br>\n")
			fmt.Fprintf(w, "<li>Location: %s</li>\n", b.Location)
			fmt.Fprint(w, "<br>\n")
			fmt.Fprintf(w, "<li>First album date: %s</li>\n", b.FirstAlbumDate)
			fmt.Fprint(w, "</ul>\n")
		}
	}
}
