package helpers

import (
	"html/template"
	"log"
	"net/http"
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
