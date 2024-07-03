package Handlers

import (
	"html"
	"net/http"
	"regexp"
	"strings"

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

// Remove HTML tags, Convert \n to <br> tags, Replace special HTML characters with their escaped equivalents
func sanitizeInput(input string) string {

	input = removeHTMLTags(input)
	input = html.EscapeString(input)
	input = strings.ReplaceAll(input, "\r\n", "<br>")
	input = strings.TrimSpace(input)
	return input
}

func removeHTMLTags(input string) string {
	htmlTagRegex := regexp.MustCompile(`<[^>]*>`)
	sanitized := htmlTagRegex.ReplaceAllString(input, "")

	return sanitized
}
