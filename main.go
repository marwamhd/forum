package main

import (
	"fmt"
	"forum/Handlers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("the server is starting at https://localhost:5050/")
	StartServer()
}

func StartServer() {
	http.HandleFunc("/", Handlers.MainHandler)
	http.HandleFunc("/signup", Handlers.SignUpHandler)
	http.HandleFunc("/login", Handlers.LoginHandler)
	http.HandleFunc("/logout", Handlers.LogoutHandler)
	http.HandleFunc("/addpost", Handlers.AddPostHandler)
	http.HandleFunc("/addcomment", Handlers.AddCommentHandler)
	http.HandleFunc("/addlike", Handlers.AddLikePostHandler)
	http.HandleFunc("/diduserlike", Handlers.DidUserLike)
	http.HandleFunc("/addCommentlike", Handlers.AddLikeCommentHandler)
	http.HandleFunc("/diduserlikecomment", Handlers.DidUserLikeComment)
	http.HandleFunc("/likedpost", Handlers.LikedPost)

	http.Handle("/Static/", http.StripPrefix("/Static/", http.FileServer(http.Dir("Static"))))
	// err := http.ListenAndServeTLS("0.0.0.0:5050", "Security/server.pem", "Security/server.key", nil)
	// log.Fatal(err)

	err := http.ListenAndServe("localhost:1010", nil)
	log.Fatal(err)
}
