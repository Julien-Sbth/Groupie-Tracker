package main

import (
	"net/http"
)

func main() {
	http.Handle("/front-end/", http.StripPrefix("/front-end", http.FileServer(http.Dir("front-end/css"))))
	fs := http.FileServer(http.Dir("./front-end"))
	http.Handle("/", fs)
	http.ListenAndServe(":8080", nil)
}
