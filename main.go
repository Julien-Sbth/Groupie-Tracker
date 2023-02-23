package main

import (
	"API/API"
	"API/backend"
	"fmt"
	"net/http"
)

func main() {
	// Serve static files from the "front-end" directory
	http.Handle("/front-end/", http.StripPrefix("/front-end/", http.FileServer(http.Dir("front-end/css"))))
	HandleTest := http.FileServer(http.Dir("./front-end"))

	// Set up API endpoints
	http.Handle("/", HandleTest)
	http.HandleFunc("/data", API.HandleData)
	http.HandleFunc("/dates", API.HandleDates)
	http.HandleFunc("/artist", API.HandleArtists)
	http.HandleFunc("/search", API.HandleSearch)
	http.HandleFunc("/location", API.HandleLocations)
	http.HandleFunc("/relation", API.HandleRelations)

	http.HandleFunc("/nwa", backend.HandleInfoNWA)
	http.HandleFunc("/sch", backend.HandleInfoSCH)
	http.HandleFunc("/jul", backend.HandleInfoJul)

	http.HandleFunc("/dion", backend.HandleInfoDion)
	http.HandleFunc("/maes", backend.HandleInfoMaes)
	http.HandleFunc("/naps", backend.HandleInfoNaps)
	http.HandleFunc("/lamar", backend.HandleInfoLamar)
	http.HandleFunc("/shakur", backend.HandleInfoShakur)
	http.HandleFunc("/kaaris", backend.HandleInfoKaaris)
	http.HandleFunc("/eminem", backend.HandleInfoEminem)
	http.HandleFunc("/warren", backend.HandleInfoWarren)
	http.HandleFunc("/daft", backend.HandleInfoDaftPunk)
	http.HandleFunc("/snoop", backend.HandleInfoSnoopDog)
	http.HandleFunc("/soprano", backend.HandleInfoSoprano)
	http.HandleFunc("/halliday", backend.HandleInfoHalliday)
	http.HandleFunc("/big", backend.HandleInfoBigflOli)
	http.HandleFunc("/notorious", backend.HandleInfoNotorious)
	http.HandleFunc("/tentation", backend.HandleInfoXXXtentation)

	// Start server
	fmt.Println("Server started on port :8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
