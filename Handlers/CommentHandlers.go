package Handlers

import (
	"encoding/json"
	"fmt"
	use "forum/Database"
	"log"
	"net/http"
	"strconv"

	"github.com/gofrs/uuid"
)

func AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	cook, cookieFound := r.Cookie("session_id")
	if cookieFound != nil {
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
		log.Println("Error in getting author", err)
		return
	}

	postID := r.FormValue("pid")
	content := r.FormValue("comment")
	content = SanitizeInput(content)

	if postID == "" || content == "" {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	pid, err := strconv.Atoi(postID)
	if err != nil {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	use.DataBase.InsertComment(author, pid, content)

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
		log.Println("Error in getting author", err)
		return
	}

	postID := r.FormValue("pid")
	commentID := r.FormValue("cid")
	like := r.FormValue("like" + commentID)

	if commentID == "" || like == "" || postID == "" {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

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
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = use.DataBase.InsertCommentLike(author, pid, cid, likenum)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	likes, dislikes, err := use.DataBase.CommentLikesDislikesTotal(postID, commentID)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

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
	if r.Method != "POST" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed))
		return
	}

	cook, cookieFound := r.Cookie("session_id")
	if cookieFound != nil {
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
		log.Println("Error in getting author", err)
		return
	}

	var requestBody CommentRequestBody
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	postID := requestBody.Pid
	commentID := requestBody.Cid

	if postID == 0 || commentID == 0 {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	likedwhat, err := use.DataBase.WhatUserLikedComment(author, postID, commentID)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
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
