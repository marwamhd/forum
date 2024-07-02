package Handlers

import (
	"log"
	"net/http"
	"text/template"
)

var MainHtml *template.Template
var ErrTmlp *template.Template
func init() {
	//we parse the templates once at the start so we no longer need to do them at each step.
	ErrTmlp, _ = template.ParseFiles("Templates/error.html")
	MainHtml, _ = template.ParseFiles("Templates/index.html")
	//if the main page wasn't parsed it means we couldn't find it. hence, internal server error.
	if MainHtml == nil {
		log.Fatal("Main page was not found.")
	}
}
// handles the error pages with status return
func HandleErrorPages(w http.ResponseWriter, statusCode int, message string) {
	//if we can't find the error html denoted by a null template, we simply makea temporary template and pass it in instead.
	if ErrTmlp == nil {
		tempTmlp := template.New("error template")
		tempTmlp, _ = tempTmlp.Parse("Internal server error page missing: {{.}}")
		w.WriteHeader(http.StatusInternalServerError)
		tempTmlp.Execute(w, "500")
		return
	}
	w.WriteHeader(statusCode)
	ErrTmlp.Execute(w, message)
}

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
