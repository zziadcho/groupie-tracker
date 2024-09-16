package main

import (
	"01/groupie-tracker/common/functions"
	"log"
	"net/http"
	"path/filepath"
)

const PORT = ":8080"

func main() {
	http.HandleFunc("/", functions.MainHandler)
	staticDir := filepath.Join("common", "static") 
    fs := http.FileServer(http.Dir(staticDir))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))


}

	