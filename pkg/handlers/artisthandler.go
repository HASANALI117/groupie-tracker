package handlers

import (
	"fmt"
	"net/http"
	"path"
	"strconv"

	"groupie-tracker/pkg/models"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	artistID, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		// http.Error(w, "Invalid artist ID", http.StatusBadRequest)
		BadRequestError(w, "Invalid artist ID")
		return
	}

	// panic("Simulated panic") // Uncomment this line to simulate a panic

	var artist models.Artist
	var dates models.Dates
	var locations models.Locations
	var relation models.Relation

	fetchInfos := []models.FetchInfo{
		{Path: fmt.Sprintf("artists/%d", artistID), Data: &artist, ErrorMsg: "artist"},
		{Path: fmt.Sprintf("dates/%d", artistID), Data: &dates, ErrorMsg: "date"},
		{Path: fmt.Sprintf("locations/%d", artistID), Data: &locations, ErrorMsg: "location"},
		{Path: fmt.Sprintf("relation/%d", artistID), Data: &relation, ErrorMsg: "relation"},
	}

	for _, info := range fetchInfos {
		if fetchAndHandleData(info.Path, info.Data, info.ErrorMsg, w) != nil {
			return
		}

		// Check if artist is found
		if info.ErrorMsg == "artist" && info.Data.(*models.Artist).ID == 0 {
			NotFoundError(w, "Artist Not Found")
			return
		}
	}

	data := models.InfoTemplateData{
		Artist:    artist,
		Dates:     dates,
		Locations: locations,
		Relation:  relation,
	}

	templates.ExecuteTemplate(w, "info.html", data)

}
