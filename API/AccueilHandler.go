package API

import (
	"encoding/json"
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func AccueilHandler(w http.ResponseWriter, r *http.Request) {
	var artists []Artist
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Erreur lors de la requête HTTP :", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&artists)
	if err != nil {
		fmt.Println("Erreur lors de la lecture de la réponse JSON :", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Créer une copie de la liste d'artistes originale
	artistsOrig := make([]Artist, len(artists))
	copy(artistsOrig, artists)
	// Permuter la liste d'artistes aléatoirement
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(artists), func(i, j int) {
		artists[i], artists[j] = artists[j], artists[i]
	})
	for i, artist := range artists {
		artists[i].ID = artistsOrig[artist.ID-1].ID
	}
	tmpl, err := template.ParseFiles("front-end/home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Vérifie si un ID a été spécifié dans l'URL
	id := r.URL.Query().Get("id")
	if id != "" {
		for _, artist := range artists {
			if strconv.Itoa(artist.ID) == id {
				tmpl.Execute(w, artist)
				return
			}
		}
		http.NotFound(w, r)
		return
	}
	// Si aucun ID n'a été spécifié, affiche tous les artistes
	tmpl.Execute(w, artists)
}
