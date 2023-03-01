package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Artist struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Albums   string `json:"albums"`
	Tracks   string `json:"tracks"`
	Self     string `json:"self"`
	ImageURL string `json:"image"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", homeHandler).Methods("GET")

	// Créer un gestionnaire de fichiers pour le répertoire "front-end/css"
	fs := http.FileServer(http.Dir("front-end/css"))
	cssHandler := http.StripPrefix("/css/", fs)

	// Servir les fichiers CSS à partir du chemin d'accès "/css/"
	r.PathPrefix("/css/").Handler(cssHandler)

	// Créer un gestionnaire de fichiers pour le répertoire "front-end"
	fs = http.FileServer(http.Dir("front-end"))
	fileHandler := http.StripPrefix("/front-end/", fs)

	// Servir les fichiers statiques à partir du chemin d'accès "/front-end/"
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

	// Ajouter un ID unique pour chaque artiste
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
