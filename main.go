package main

import (
	"01/groupie-tracker/common/functions"
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	fs := http.FileServer(http.Dir("static"))
	
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", functions.MainHandler)

	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

	