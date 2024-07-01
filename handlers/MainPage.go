package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"pl/database"
	use "pl/database"
	"pl/helpers"
	"strconv"
	"strings"
	"text/template"

	"github.com/gofrs/uuid"
)

type content struct {
	Authlevel     int
	U_id          int
	Posts         []use.Post
	FilteredPosts []use.Post
	LikedPosts    []use.Post
}

type RequestBody struct {
	Pid int `json:"pid"`
}

type CommentRequestBody struct {
	Pid int `json:"pid"`
	Cid int `json:"cid"`
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		helpers.HandleErrorPages(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))
		return
	}

	cook, cookieFound := r.Cookie("session_id")
	authlevel := 1
	var author int

	if cookieFound != nil {
		log.Println(cookieFound, "31")
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
			log.Println("error in getting author", err)
			return
		}
	}

	//is this cookie actually in the db? if yes, we enter as logged in, if not, we enter as not logged in and delete the cookie

	posts, errForPost := use.DataBase.GetPosts()
	if errForPost != nil {
		log.Println("error in getting posts", errForPost)
		return
	}
	com := "select * from posts"
	values := r.URL.Query()[("cat")]
	fmt.Println("val", values)
	if len(values) != 0 {
		str := strings.Join(values, " and ")
		com = "select * from posts where " + str
	}

	filteredPosts, errForFiltered := use.DataBase.GetFilteredPosts(com)
	if errForFiltered != nil {
		log.Println("error in filtering posts", errForFiltered)
		return
	}

	likedPost, errForLiked := use.DataBase.WhatUserLikedPosts(author)
	if errForLiked != nil {
		log.Println("error in liked posts", errForLiked)
		return
	}

	MainHtml, _ := template.ParseFiles("Templates/index.html")

	err := MainHtml.Execute(w, content{authlevel, author, posts, filteredPosts, likedPost})
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

func AddPostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.HandleErrorPages(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
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
	cats := []int{}

	for i := 1; i <= 4; i++ {
		found := false
		for _, v := range category {
			n, err := strconv.Atoi(v)
			if err != nil {
				helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
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
		helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	use.DataBase.InsertPost(author, title, content, cats)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.HandleErrorPages(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	cook, cookieFound := r.Cookie("session_id")
	if cookieFound != nil {
		log.Println(cookieFound, "31")
		OverWriteCookieValue(w, r, uuid.Nil)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("278")

	if cook.Value == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("284")

	activeSession, errForSes := use.DataBase.SessionExists(cook.Value)

	if !activeSession || errForSes != nil {
		OverWriteCookieValue(w, r, uuid.Nil)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("293")

	author, err := use.GetAuthor(cook.Value)
	if err != nil {
		log.Println("error in getting author", err)
		return
	}

	postID := r.FormValue("pid")
	content := r.FormValue("comment")

	if postID == "" || content == "" {
		helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	fmt.Println("309")

	pid, err := strconv.Atoi(postID)
	if err != nil {
		helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	fmt.Println("here")

	use.DataBase.InsertComment(author, pid, content)

	fmt.Println("Comment added.")

	posts, err := use.DataBase.GetFilteredPosts("select * from posts")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	// Create a JSON response struct
	response := CommentJsons{
		Success: true,
		Message: "Comment added successfully",
		Posts:   posts,
	}

	// Set content type and encode response as JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode the response struct directly
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

type jsonResponse struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Likes    int    `json:"likes"`
	Dislikes int    `json:"dislikes"`
}
type jsons struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Userl   int    `json:"userl"`
}

type CommentJsons struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Posts   []use.Post `json:"posts"`
}

func AddLikePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.HandleErrorPages(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	cook, cookieFound := r.Cookie("session_id")
	if cookieFound != nil {
		log.Println(cookieFound, "3231")
		OverWriteCookieValue(w, r, uuid.Nil)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("12278")

	if cook.Value == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("1212284")

	activeSession, errForSes := use.DataBase.SessionExists(cook.Value)

	if !activeSession || errForSes != nil {
		OverWriteCookieValue(w, r, uuid.Nil)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("2121293")

	author, err := use.GetAuthor(cook.Value)
	if err != nil {
		log.Println("error in getting author", err)
		return
	}

	postID := r.FormValue("pid")
	like := r.FormValue("like")
	fmt.Printf("author: %v\n", author)

	if postID == "" || like == "" {
		helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	fmt.Printf("like: %v\n", like)

	pid, err := strconv.Atoi(postID)
	if err != nil {
		helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	likenum, err := strconv.Atoi(like)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Println("here")

	fmt.Printf("pid: %v\n", pid)

	err = use.DataBase.InsertLike(author, pid, likenum)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Println("like added.")

	likes, dislikes, err := database.DataBase.LikesDislikesTotal(postID)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("likes: %v\n", likes)
	fmt.Printf("dislikes: %v\n", dislikes)

	// Create a JSON response struct
	response := jsonResponse{
		Success:  true,
		Message:  "like added successfully",
		Likes:    likes,
		Dislikes: dislikes,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func AddLikeCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.HandleErrorPages(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	cook, cookieFound := r.Cookie("session_id")
	if cookieFound != nil {
		log.Println(cookieFound, "3231")
		OverWriteCookieValue(w, r, uuid.Nil)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("12278")

	if cook.Value == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("1212284")

	activeSession, errForSes := use.DataBase.SessionExists(cook.Value)

	if !activeSession || errForSes != nil {
		OverWriteCookieValue(w, r, uuid.Nil)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("2121293")

	author, err := use.GetAuthor(cook.Value)
	if err != nil {
		log.Println("error in getting author", err)
		return
	}

	postID := r.FormValue("pid")
	commentID := r.FormValue("cid")
	like := r.FormValue("like" + commentID)
	fmt.Printf("author: %v\n", author)

	if commentID == "" || like == "" || postID == "" {
		helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	fmt.Printf("like: %v\n", like)

	pid, err := strconv.Atoi(postID)
	if err != nil {
		helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	cid, err := strconv.Atoi(commentID)
	if err != nil {
		helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	likenum, err := strconv.Atoi(like)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Println("here")

	fmt.Printf("cid: %v\n", cid)

	err = use.DataBase.InsertCommentLike(author, pid, cid, likenum)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Println("comment like added.")

	likes, dislikes, err := database.DataBase.CommentLikesDislikesTotal(postID, commentID)

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Printf("likes: %v\n", likes)
	fmt.Printf("dislikes: %v\n", dislikes)

	// Create a JSON response struct
	response := jsonResponse{
		Success:  true,
		Message:  "like added successfully",
		Likes:    likes,
		Dislikes: dislikes,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

//DidUserLikeComment

func DidUserLike(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post")
	if r.Method != "POST" {
		helpers.HandleErrorPages(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	cook, cookieFound := r.Cookie("session_id")
	if cookieFound != nil {
		log.Println(cookieFound, "3231")
		OverWriteCookieValue(w, r, uuid.Nil)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("sess")

	fmt.Println("12278")

	if cook.Value == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("1212284")

	activeSession, errForSes := use.DataBase.SessionExists(cook.Value)

	if !activeSession || errForSes != nil {
		OverWriteCookieValue(w, r, uuid.Nil)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("ass")

	fmt.Println("2121293a")

	author, err := use.GetAuthor(cook.Value)
	if err != nil {
		log.Println("error in getting author", err)
		return
	}

	var requestBody RequestBody
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	postID := requestBody.Pid
	fmt.Printf("author: %v\n", author)

	fmt.Printf("postID: %v\n", postID)

	if postID == 0 {
		helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	fmt.Printf("pid: %v\n", postID)

	likedwhat, err := database.DataBase.WhatUserLiked(author, postID)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	// Create a JSON response struct
	response := jsons{
		Success: true,
		Message: "like added successfully",
		Userl:   likedwhat,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}

func AddLikeCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		helpers.HandleErrorPages(w, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	cook, cookieFound := r.Cookie("session_id")
	if cookieFound != nil {
		log.Println(cookieFound, "3231")
		OverWriteCookieValue(w, r, uuid.Nil)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("sess")

	fmt.Println("12278")

	if cook.Value == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("1212284")

	activeSession, errForSes := use.DataBase.SessionExists(cook.Value)

	if !activeSession || errForSes != nil {
		OverWriteCookieValue(w, r, uuid.Nil)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	fmt.Println("ass")

	fmt.Println("2121293a")

	author, err := use.GetAuthor(cook.Value)
	if err != nil {
		log.Println("error in getting author", err)
		return
	}

	var requestBody CommentRequestBody
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	postID := requestBody.Pid
	commentID := requestBody.Cid
	fmt.Printf("author: %v\n", author)

	fmt.Printf("postID: %v\n", postID)
	fmt.Printf("commentPD: %v\n", commentID)

	if postID == 0 || commentID == 0 {
		helpers.HandleErrorPages(w, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	fmt.Printf("pid: %v\n", postID)

	likedwhat, err := database.DataBase.WhatUserLikedComment(author, postID, commentID)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	// Create a JSON response struct
	response := jsons{
		Success: true,
		Message: "like added successfully",
		Userl:   likedwhat,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}
