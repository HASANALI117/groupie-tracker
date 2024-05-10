package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/pkg/models"
)

func SuggestionsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	query := r.URL.Query().Get("query")
	if query == "" {
		BadRequestError(w, "Query parameter is missing")
		return
	}

	results, err := Suggest(query)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// return results in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(results)

}

type Suggestion struct {
	Name   string
	Type   string
	Artist models.Artist
}

func Suggest(query string) ([]byte, error) {
	var artists []models.Artist
	err := fetchData("artists", &artists)
	if err != nil {
		return nil, err
	}
	var allLocations models.AllLocations
	if fetchData("locations", &allLocations) != nil {
		return nil, err
	}

	lowerQuery := strings.ToLower(query)

	var artistResults []Suggestion
	var memberResults []Suggestion
	var locationResults []Suggestion
	var firstAlbumResults []Suggestion
	var creationDateResults []Suggestion

	for _, artist := range artists {
		if strings.HasPrefix(strings.ToLower(artist.Name), lowerQuery) {
			artistResults = append(artistResults, Suggestion{artist.Name, "artist/band", artist})
		}
		for _, member := range artist.Members {
			if strings.HasPrefix(strings.ToLower(member), lowerQuery) {
				memberResults = append(memberResults, Suggestion{member, "member", artist})
			}
		}
		locations := allLocations.Index[artist.ID-1]
		for _, location := range locations.Locations {
			if strings.Contains(strings.ToLower(location), lowerQuery) {
				locationResults = append(locationResults, Suggestion{location, "location", artist})
			}
		}
		if strings.HasPrefix(artist.FirstAlbum, query) {
			firstAlbumResults = append(firstAlbumResults, Suggestion{artist.FirstAlbum, "first album date", artist})
		}
		if strings.HasPrefix(strconv.Itoa(artist.CreationDate), query) {
			creationDateResults = append(creationDateResults, Suggestion{strconv.Itoa(artist.CreationDate), "creation date", artist})
		}
	}

	// combine all results
	var results []Suggestion
	results = append(results, artistResults...)
	results = append(results, memberResults...)
	results = append(results, locationResults...)
	results = append(results, firstAlbumResults...)
	results = append(results, creationDateResults...)
	// convert results to JSON
	resultsJSON, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}

	return resultsJSON, nil
}
