package backend

import (
	"html/template"
	"net/http"
)

func HandleInfoNWA(w http.ResponseWriter, r *http.Request) {
	// Définir les données
	type Person struct {
		Name    string
		Age     int
		Address string
	}
	p := Person{"JAHHHHHHHH", 4200, "salutation"}

	// Charger et analyser le modèle HTML
	tmpl, err := template.ParseFiles("front-end/InfoGroup/infosNWA.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Exécuter le modèle HTML avec les données
	if err := tmpl.Execute(w, p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
