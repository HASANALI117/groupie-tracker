package handlers

import (
	"groupie-tracker/pkg/models"
	"net/http"
)

func NotFoundError(w http.ResponseWriter, msg string) {
	ErrorData := models.Error{
		ErrorCode: 404,
		ErrorMsg:  msg,
		ErrorDsc:  "Sorry, we couldn't find the page you're looking for.",
	}
	templates.ExecuteTemplate(w, "error.html", ErrorData)
}

func BadRequestError(w http.ResponseWriter, msg string) {
	ErrorData := models.Error{
		ErrorCode: 400,
		ErrorMsg:  msg,
		ErrorDsc:  "Sorry, we couldn't process your request.",
	}
	templates.ExecuteTemplate(w, "error.html", ErrorData)
}

func InternalServerError(w http.ResponseWriter, msg string) {
	ErrorData := models.Error{
		ErrorCode: 500,
		ErrorMsg:  msg,
		ErrorDsc:  "Sorry, we encountered an error while processing your request.",
	}
	templates.ExecuteTemplate(w, "error.html", ErrorData)
}