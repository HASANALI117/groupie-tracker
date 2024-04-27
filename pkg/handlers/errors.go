package handlers

import (
	"groupie-tracker/pkg/models"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, msg string) {
	ErrorData := models.Error{
		ErrorCode: 404,
		ErrorMsg:  msg,
		ErrorDsc:  "Sorry, we couldn't find the page you're looking for.",
	}
	templates.ExecuteTemplate(w, "error.html", ErrorData)
}
