package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/pkg/handlers"
	"groupie-tracker/pkg/middleware"
)

func main() {
	http.Handle("/web/static/", http.StripPrefix("/web/static/", http.FileServer(http.Dir("web/static"))))
	http.HandleFunc("/", middleware.RecoverPanic(handlers.MainHandler))
	http.HandleFunc("/artist/", middleware.RecoverPanic(handlers.ArtistHandler))
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
