package API

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	searchTerm := r.FormValue("search")
	var artists []Artist

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
		fmt.Fprintf(w, "<select name=\"artist\" id=\"artist\">\n")
		fmt.Fprintf(w, "{{if gt (len .Members) 2}}\n")
		fmt.Fprintf(w, "<option value=\"%s - groupe\">%s - groupe</option>\n", artist.Name, artist.Name)
		fmt.Fprintf(w, "{{else}}\n")
		fmt.Fprintf(w, "<option value=\"%s - membre\">%s - artiste</option>\n", artist.Name, artist.Name)
		fmt.Fprintf(w, "{{end}}\n")
		fmt.Fprintf(w, "</select>\n")
		fmt.Fprintf(w, "<p><a href=\"/index/%d\">Page de %s</a></p>\n", artist.ID, artist.Name)
		fmt.Fprintf(w, "</div>\n") // end div.info
		fmt.Fprintf(w, "</div>\n") // end div.artiste-box
		fmt.Fprintf(w, "</li>\n")
	}
}
