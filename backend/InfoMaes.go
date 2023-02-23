package backend

import (
	"html/template"
	"net/http"
)

func HandleInfoMaes(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("front-end/InfoGroup/InfoMaes.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := "Hello, World!" // définir une variable factice pour 'data'
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
