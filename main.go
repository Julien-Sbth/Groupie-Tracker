package main

import (
<<<<<<< HEAD
<<<<<<< HEAD
=======
	"API/API"
	"API/backend"
	"fmt"
>>>>>>> master
=======
	"encoding/json"
	"fmt"
	"html/template"
	"log"
>>>>>>> master
	"net/http"

	"github.com/gorilla/mux"
)

<<<<<<< HEAD
func main() {
<<<<<<< HEAD
	http.Handle("/front-end/", http.StripPrefix("/front-end", http.FileServer(http.Dir("front-end/css"))))
	fs := http.FileServer(http.Dir("./front-end"))
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
=======
	// Serve static files from the "front-end" directory
	http.Handle("/front-end/", http.StripPrefix("/front-end/", http.FileServer(http.Dir("front-end/css"))))
	HandleTest := http.FileServer(http.Dir("./front-end"))

	// Set up API endpoints
	http.Handle("/", HandleTest)
	http.HandleFunc("/search", API.HandleSearch)
	http.HandleFunc("/date", API.HandleDates)

	http.HandleFunc("/nwa", backend.HandleInfoNWA)
	http.HandleFunc("/infoJSONAW", backend.HandleInfoJSONNWA)

	http.HandleFunc("/infoJSONAPS", backend.HandleInfoJSONNaps)
	http.HandleFunc("/naps", backend.HandleInfoNaps)

	http.HandleFunc("/sch", backend.HandleInfoSCH)
	http.HandleFunc("/infoJSONSCH", backend.HandleInfoJSONSCH)

	http.HandleFunc("/jul", backend.HandleInfoJul)
	http.HandleFunc("/infoJSONJul", backend.HandleInfoJSONJul)

	http.HandleFunc("/dion", backend.HandleInfoDion)
	http.HandleFunc("/infoJSONDion", backend.HandleInfoJSONDion)

	http.HandleFunc("/maes", backend.HandleInfoMaes)
	http.HandleFunc("/infoJSONMaes", backend.HandleInfoJSONMaes)

	http.HandleFunc("/lamar", backend.HandleInfoLamar)
	http.HandleFunc("/infoJSONLamar", backend.HandleInfoJSONLamar)

	http.HandleFunc("/shakur", backend.HandleInfoShakur)
	http.HandleFunc("/infoJSONShakur", backend.HandleInfoJSONShakur)

	http.HandleFunc("/kaaris", backend.HandleInfoKaaris)
	http.HandleFunc("/infoJSONKaaris", backend.HandleInfoJSONKaaris)

	http.HandleFunc("/warren", backend.HandleInfoWarren)
	http.HandleFunc("/infoJSONWarren", backend.HandleInfoJSONWarren)

	http.HandleFunc("/infoJSONDaftPunk", backend.HandleInfoJSONDaftPunk)
	http.HandleFunc("/daft", backend.HandleInfoDaftPunk)

	http.HandleFunc("/infoJSONSnoopDog", backend.HandleInfoJSONSnoopDog)
	http.HandleFunc("/snoop", backend.HandleInfoSnoopDog)

	http.HandleFunc("/soprano", backend.HandleInfoSoprano)
	http.HandleFunc("/infoJSONSoprano", backend.HandleInfoJSONSoprano)

	http.HandleFunc("/halliday", backend.HandleInfoHalliday)
	http.HandleFunc("/infoJSONHalliday", backend.HandleInfoJSONHalliday)

	http.HandleFunc("/big", backend.HandleInfoBigflOli)
	http.HandleFunc("/infoJSONBigflolie", backend.HandleInfoJSONBigflolie)

	http.HandleFunc("/infoJSONEminem", backend.HandleInfoJSONEminem)
	http.HandleFunc("/eminem", backend.HandleInfoEminem)

	http.HandleFunc("/infoJSONBiggies", backend.HandleInfoJSONBig)
	http.HandleFunc("/notorious", backend.HandleInfoNotorious)

	http.HandleFunc("/tentation", backend.HandleInfoXXXTentation)
	http.HandleFunc("/infoJSONXXX", backend.HandleInfoJSONXXXTentation)

	// Start server
	fmt.Println("Server started on port :8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
>>>>>>> master
=======
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
>>>>>>> master
}
