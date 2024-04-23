package handlers

import (
	"fmt"
	"groupie-tracker/pkg/models"
	"log"
	"net/http"
	"path"
	"strconv"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	artistID, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	//panic("Simulated panic") // Uncomment this line to simulate a panic
	
	var artist models.Artist
	err = fetchData(fmt.Sprintf("artists/%d", artistID), &artist)
	if err != nil {
		log.Printf("Error fetching artist: %v", err) // Log the error
		http.Error(w, "Artist not found", http.StatusNotFound)
		return
	}
	templates.ExecuteTemplate(w, "info.html", artist)
}
