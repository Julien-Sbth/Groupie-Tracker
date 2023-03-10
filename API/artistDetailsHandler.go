package API

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func ArtistDetailsHandler(w http.ResponseWriter, r *http.Request) {
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
