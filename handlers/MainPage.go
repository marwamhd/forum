package handlers

import (
	"fmt"
	"log"
	"net/http"
	use "pl/database"
	"pl/helpers"
	"text/template"

	"github.com/gofrs/uuid"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		helpers.HandleErrorPages(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	fmt.Printf("r.Cookies(): %v\n", r.Cookies())

	cook, cookieFound := r.Cookie("session_id")
	authlevel := 1

	if cookieFound != nil {
		log.Println(cookieFound)
		cookie := http.Cookie{
			Name:     "session_id",
			Value:    "",
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteLaxMode,
		}

		http.SetCookie(w, &cookie)
		authlevel = 0
	} else if cook.Value == "" {
		authlevel = 0
	}
	MainHtml, _ := template.ParseFiles("Templates/index.html")

	err := MainHtml.Execute(w, authlevel)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.HandleErrorPages(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	email := r.FormValue("em")
	username := r.FormValue("us")
	password := r.FormValue("ps")

	fmt.Printf("email: %v\n", email)
	fmt.Printf("username: %v\n", username)
	fmt.Printf("password: %v\n", password)

	//we have to check for the each of the data and make sure all of them are valid, if not, throw an error

	if email == "" || username == "" || password == "" {
		helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	//we check if the email is unique

	exists, err := use.DataBase.EmailExists(email)
	if err != nil {
		log.Println("error in executing email exists", err)
		return
	}

	if exists {
		log.Println("Email already exists")
		helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	use.DataBase.InsertUser(email, username, password)

	authlevel, sid, err := use.DataBase.Login(email, password, r)
	if err != nil && (err.Error() == "invalid credentials" || err.Error() == "user already has an active session") {
		helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest)+" "+err.Error())
		return
	} else if err != nil {
		helpers.HandleErrorPages(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)+" "+err.Error())
		return
	}
	if authlevel == 1 {
		fmt.Println("authorized")
	}

	SetCookie(w, sid)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.HandleErrorPages(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	email := r.FormValue("em")
	password := r.FormValue("ps")

	//we have to check for the each of the data and make sure all of them are valid, if not, throw an error
	if email == "" || password == "" {
		helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	authlevel, sid, err := use.DataBase.Login(email, password, r)
	if err != nil && (err.Error() == "invalid credentials" || err.Error() == "user already has an active session") {
		helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest)+" "+err.Error())
		return
	} else if err != nil {
		helpers.HandleErrorPages(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)+" "+err.Error())
		return
	}
	if authlevel == 1 {
		fmt.Println("authorized")
	}

	SetCookie(w, sid)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func SetCookie(w http.ResponseWriter, sid uuid.UUID) {
	cookie := http.Cookie{
		Name:     "session_id",
		Value:    sid.String(),
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
