package Handlers

import (
	use "forum/Database"
	"log"
	"net/http"
	"strings"
	"text/template"

	"github.com/gofrs/uuid"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	cook, cookieFound := r.Cookie("session_id")
	authlevel := 1
	var author int

	if cookieFound != nil {
		OverWriteCookieValue(w, r, uuid.Nil)
		authlevel = 0
	} else if cook.Value == "" {
		authlevel = 0
	} else {
		activeSession, errForSes := use.DataBase.SessionExists(cook.Value)

		if !activeSession || errForSes != nil {
			authlevel = 0
			OverWriteCookieValue(w, r, uuid.Nil)
		}
		var err error

		author, err = use.GetAuthor(cook.Value)
		if err != nil {
			log.Println("Error in getting author", err)
			return
		}
	}

	//is this cookie actually in the db? if yes, we enter as logged in, if not, we enter as not logged in and delete the cookie

	posts, errForPost := use.DataBase.GetFilteredPosts("select * from posts")
	if errForPost != nil {
		log.Println("Error in getting posts", errForPost)
		return
	}
	com := "select * from posts"
	values := r.URL.Query()[("cat")]
	if len(values) != 0 {
		str := strings.Join(values, " and ")
		com = "select * from posts where " + str
	}

	filteredPosts, errForFiltered := use.DataBase.GetFilteredPosts(com)
	if errForFiltered != nil {
		log.Println("Error in filtering posts", errForFiltered)
		return
	}

	likedPost, errForLiked := use.DataBase.WhatUserLikedPosts(author)
	if errForLiked != nil {
		log.Println("Error in liked posts", errForLiked)
		return
	}

	MainHtml, eror := template.ParseFiles("Templates/index.html")
	if eror != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	err := MainHtml.Execute(w, content{authlevel, author, posts, filteredPosts, likedPost})
	if err != nil {
		log.Fatal("Error", err)
		return
	}

}
