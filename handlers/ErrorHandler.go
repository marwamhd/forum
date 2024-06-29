package handlers

import (
	"net/http"
	"text/template"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, statusCode int, errM string) {
	var errorMessage string
	switch statusCode {
	case http.StatusNotFound:
		//404
		errorMessage = "Page not found"
	case http.StatusBadRequest:
		//400
		errorMessage = "Bad request"
		if errM != "" {
			//400 with extra message
			errorMessage += ": " + errM
		}
	case http.StatusInternalServerError:
		//500
		errorMessage = "Internal server error"
	case http.StatusMethodNotAllowed:
		//405
		errorMessage = "Method not allowed"
	default:
		errorMessage = "Unexpected error"
	}
	data := struct {
		ErrorCode    int
		ErrorMessage string
	}{
		ErrorCode:    statusCode,
		ErrorMessage: errorMessage,
	}
	tmpl, err := template.ParseFiles("Templates/error.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	tmpl.Execute(w, data)
}
