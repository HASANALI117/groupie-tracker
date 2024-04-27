package handlers

import (
	"net/http"
	"text/template"

	"groupie-tracker/pkg/models"
)

var templates = template.Must(template.ParseGlob("web/templates/*.html"))

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, "Page Not Found")
		return
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var artists []models.Artist
	err := fetchData("artists", &artists)
	if err != nil {
		http.Error(w, "Error fetching artists", http.StatusInternalServerError)
		return
	}
	templates.ExecuteTemplate(w, "index.html", artists)
}
