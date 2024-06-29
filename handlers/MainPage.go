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
		OverWriteCookieValue(w, r, uuid.Nil)
		authlevel = 0
	} else if cook.Value == "" {
		authlevel = 0
	}

	//is this cookie actually in the db? if yes, we enter as logged in, if not, we enter as not logged in and delete the cookie

	activeSession, errForSes := use.DataBase.SessionExists(cook.Value)

	if !activeSession || errForSes != nil {
		authlevel = 0
		OverWriteCookieValue(w, r, uuid.Nil)
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

	// authlevel, sid, err := use.DataBase.Login(email, password, r)
	// if err != nil && (err.Error() == "invalid credentials" || err.Error() == "user already has an active session") {
	// 	helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest)+" "+err.Error())
	// 	return
	// } else if err != nil {
	// 	helpers.HandleErrorPages(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)+" "+err.Error())
	// 	return
	// }
	// if authlevel == 1 {
	// 	fmt.Println("authorized")
	// }

	// OverWriteCookieValue(w, r, sid)

	LoginHandler(w, r)

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
	if err != nil && (err.Error() == "invalid credentials" || err.Error() == "user already has an active session" || err.Error() == "sql: no rows in result set") {
		helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest)+" "+err.Error())
		return
	} else if err != nil {
		helpers.HandleErrorPages(w, http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)+" "+err.Error())
		return
	}
	if authlevel == 1 {
		fmt.Println("authorized")
	}

	OverWriteCookieValue(w, r, sid)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	OverWriteCookieValue(w, r, uuid.Nil)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func OverWriteCookieValue(w http.ResponseWriter, r *http.Request, sid uuid.UUID) {
	var val string

	if sid != uuid.Nil {
		val = sid.String()
	}

	cookie := http.Cookie{
		Name:     "session_id",
		Value:    val,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
}

// func IsOnASession(r *http.Request) bool {
// 	cook, _ := r.Cookie("session_id")

// 	if cook.Value == "" {
// 		return false
// 	}

// }
