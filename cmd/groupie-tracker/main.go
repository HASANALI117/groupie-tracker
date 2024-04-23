package main

import (
	"fmt"
	"groupie-tracker/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	http.Handle("/web/static/", http.StripPrefix("/web/static/", http.FileServer(http.Dir("web/static"))))
	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/artist/", handlers.ArtistHandler)
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
