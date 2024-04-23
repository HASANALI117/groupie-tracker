package main

import (
	"fmt"
	"groupie-tracker/pkg/handlers"
	"groupie-tracker/pkg/middleware"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", middleware.RecoverPanic(handlers.MainHandler))
	http.HandleFunc("/artist/", middleware.RecoverPanic(handlers.ArtistHandler))
	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
