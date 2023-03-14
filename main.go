package main

import (
	"API/API"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", API.HomeHandler).Methods("GET")
	r.HandleFunc("/lofi/{id}", API.Lofi).Methods("GET")
	r.HandleFunc("/home", API.AccueilHandler).Methods("GET")
	r.HandleFunc("/search", API.SearchHandler).Methods("GET")
	r.HandleFunc("/index/{id}", API.ArtistDetailsHandler).Methods("GET")

	fs := http.FileServer(http.Dir("front-end/css"))
	cssHandler := http.StripPrefix("/css/", fs)
	r.PathPrefix("/css/").Handler(cssHandler)
	fs = http.FileServer(http.Dir("front-end"))
	fileHandler := http.StripPrefix("/front-end/", fs)
	r.PathPrefix("/front-end/").Handler(fileHandler)
	fmt.Println("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
