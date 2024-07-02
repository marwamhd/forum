package Handlers

import (
	"encoding/json"
	"fmt"
	use "forum/Database"
	"html"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gofrs/uuid"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
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
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	fmt.Println("309")

	pid, err := strconv.Atoi(postID)
	if err != nil {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	fmt.Println("here")

	content = sanitizeInput(content)

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

func AddLikeCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
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
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	fmt.Printf("like: %v\n", like)

	pid, err := strconv.Atoi(postID)
	if err != nil {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	cid, err := strconv.Atoi(commentID)
	if err != nil {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
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

	likes, dislikes, err := use.DataBase.CommentLikesDislikesTotal(postID, commentID)

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

func DidUserLikeComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post")
	if r.Method != "POST" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
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
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	fmt.Printf("pid: %v\n", postID)

	likedwhat, err := use.DataBase.WhatUserLikedComment(author, postID, commentID)
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

func LikedPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post")
	if r.Method != "POST" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
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

	likedPost, errForLiked := use.DataBase.WhatUserLikedPosts(author)
	if errForLiked != nil {
		log.Println("error in liked posts", errForLiked)
		return
	}

	// Create a JSON response struct
	response := CommentJsons{
		Success: true,
		Message: "posts successfully retrieved",
		Posts:   likedPost,
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}