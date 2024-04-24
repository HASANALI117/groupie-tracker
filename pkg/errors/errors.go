package errors

import (
	"groupie-tracker/pkg/handlers"
	"groupie-tracker/pkg/models"
) 

func notFound(msg string) {
	errData : models.error := {
		errCode: 404
		errMsg: msg
		errDsc: "Sorry, we couldn’t find the page you’re looking for."
	} 
	handlers.template.ExecuteTemplate(w, "error.html", errData)
}