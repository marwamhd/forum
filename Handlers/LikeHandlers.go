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

func AddLikePostHandler(w http.ResponseWriter, r *http.Request) {
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
	like := r.FormValue("like")

	if postID == "" || like == "" {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	pid, err := strconv.Atoi(postID)
	if err != nil {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	likenum, err := strconv.Atoi(like)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = use.DataBase.InsertLike(author, pid, likenum)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	likes, dislikes, err := use.DataBase.LikesDislikesTotal(postID)

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

//DidUserLikeComment

func DidUserLike(w http.ResponseWriter, r *http.Request) {
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

	var requestBody RequestBody
	err = json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	postID := requestBody.Pid

	if postID == 0 {
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	likedwhat, err := use.DataBase.WhatUserLiked(author, postID)
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

func LikedPost(w http.ResponseWriter, r *http.Request) {
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

	likedPost, errForLiked := use.DataBase.WhatUserLikedPosts(author)
	if errForLiked != nil {
		log.Println("Error in liked posts", errForLiked)
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
