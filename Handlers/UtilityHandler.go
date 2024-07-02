package Handlers

import (
	"net/http"

	"github.com/gofrs/uuid"
)

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
