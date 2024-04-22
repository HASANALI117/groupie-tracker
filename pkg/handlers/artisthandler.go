package handlers

import (
	"fmt"
	"groupie-tracker/pkg/models"
	"net/http"
	"strconv"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	artistID, err := strconv.Atoi(r.URL.Path[8:])
	if err != nil {
		http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		return
	}

	var artist models.Artist
	err = fetchData(fmt.Sprintf("artists/%d", artistID), &artist)
	if err != nil {
		http.Error(w, "Error fetching artist", http.StatusInternalServerError)
		return
	}
	templates.ExecuteTemplate(w, "info.html", artist)
}
