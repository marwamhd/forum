package Handlers

import (
	"fmt"
	use "forum/Database"
	"log"
	"net/http"
	"strconv"

	"github.com/gofrs/uuid"
)

func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	cook, cookieFound := r.Cookie("session_id")
	if cookieFound != nil {
		log.Println(cookieFound, "31")
		OverWriteCookieValue(w, r, uuid.Nil)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if cook.Value == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	activeSession, errForSes := use.DataBase.SessionExists(cook.Value)

	if !activeSession || errForSes != nil {
		OverWriteCookieValue(w, r, uuid.Nil)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	author, err := use.GetAuthor(cook.Value)
	if err != nil {
		log.Println("error in getting author", err)
		return
	}

	title := r.FormValue("title")
	content := r.FormValue("post")
	category := r.Form["category"]
	fmt.Printf("category: %v\n", len(category))

	title = sanitizeInput(title)
	content = sanitizeInput(content)
	cats := []int{}

	for i := 1; i <= 4; i++ {
		found := false
		for _, v := range category {
			n, err := strconv.Atoi(v)
			if err != nil {
				ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
				return
			}
			if n == i {
				found = true
				cats = append(cats, 1)
			}
		}
		if !found {
			cats = append(cats, 0)
		}
	}

	fmt.Printf("cats: %v\n", cats)

	if title == "" || content == "" {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	use.DataBase.InsertPost(author, title, content, cats)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
