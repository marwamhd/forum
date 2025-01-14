package Handlers

import (
	"fmt"
	use "forum/Database"
	"log"
	"net/http"

	"github.com/gofrs/uuid"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	email := r.FormValue("em")
	username := r.FormValue("us")
	password := r.FormValue("ps")

	semail := SanitizeInput(email)
	susername := SanitizeInput(username)

	if semail != email || susername != username {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	//we have to check for the each of the data and make sure all of them are valid, if not, throw an error

	if email == "" || username == "" || password == "" {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	exists, err := use.DataBase.EmailExists(email)
	if err != nil {
		log.Println("Error in executing email exists", err)
		return
	}

	uexists, err := use.DataBase.UsernameExists(username)

	if err != nil {
		log.Println("Error in executing username exists", err)
		return
	}

	if exists || uexists {
		log.Println("Email or email already exists")
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	hashed, err := use.HashPassword(password)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	use.DataBase.InsertUser(email, username, hashed)
	LoginHandler(w, r)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	email := r.FormValue("em")
	password := r.FormValue("ps")

	//we have to check for the each of the data and make sure all of them are valid, if not, throw an error
	if email == "" || password == "" {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	_, sid, err := use.DataBase.Login(email, password, r)
	if err != nil && (err.Error() == "invalid credentials" || err.Error() == "user already has an active session" || err.Error() == "sql: no rows in result set") {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest)+" "+err.Error())
		return
	} else if err != nil {
		ErrorHandler(w, r, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)+" "+err.Error())
		return
	}

	OverWriteCookieValue(w, r, sid)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	OverWriteCookieValue(w, r, uuid.Nil)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
