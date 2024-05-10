package handlers

import (
    "log"
    "net/http"
    "strings"
    
    "groupie-tracker/pkg/models"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != "GET" {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    query := r.URL.Query().Get("query")
    if query == "" {
        BadRequestError(w, "Query parameter is missing")
        return
    }

    results, err := Search(query)
    if err != nil {
        log.Println(err)
        InternalServerError(w, "Failed to search")
        return
    }

    data := struct {
        Query   string
        Results []models.Artist
    }{
        Query:   query,
        Results: results,
    }

    templates.ExecuteTemplate(w, "search.html", data)
}

func Search(query string) ([]models.Artist, error) {
    var results []models.Artist
    var artists []models.Artist
	err := fetchData("artists", &artists)
    if err != nil {
        return nil, err
    }

    lowerQuery := strings.ToLower(query)
    for _, artist := range artists {
        if strings.Contains(strings.ToLower(artist.Name), lowerQuery) {
            results = append(results, artist)
        }
    }

    return results, nil
}