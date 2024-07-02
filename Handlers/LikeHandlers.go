package Handlers

import (
	"encoding/json"
	"fmt"
	database "forum/Database"
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
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
		return
	}

	fmt.Printf("like: %v\n", like)

	pid, err := strconv.Atoi(postID)
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

//DidUserLikeComment

func DidUserLike(w http.ResponseWriter, r *http.Request) {
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
		ErrorHandler(w, r, http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
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
