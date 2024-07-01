package main

import (
	"fmt"
	"log"
	"net/http"
	"pl/handlers"
)

func main() {
	fmt.Println("the server is starting at https://localhost:5050/")
	StartServer()
}

func StartServer() {

	http.HandleFunc("/", handlers.MainHandler)
	http.HandleFunc("/signup", handlers.SignUpHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/logout", handlers.LogoutHandler)
	http.HandleFunc("/addpost", handlers.AddPostHandler)
	http.HandleFunc("/addcomment", handlers.AddCommentHandler)
	http.HandleFunc("/addlike", handlers.AddLikePostHandler)
	http.HandleFunc("/diduserlike", handlers.DidUserLike)
	http.HandleFunc("/addCommentlike", handlers.AddLikeCommentHandler)
	http.HandleFunc("/diduserlikecomment", handlers.DidUserLikeComment)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	err := http.ListenAndServeTLS("0.0.0.0:5050", "server.pem", "server.key", nil)

	log.Fatal(err)
}
